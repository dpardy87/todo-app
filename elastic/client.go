package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"strings"
)

type ESClient struct {
	// struct for Golang Elastic client
	Client *elasticsearch.Client
}

func NewClient(addresses []string) *ESClient {
	// create new ES client
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	// instantiate client with config -> url(s)
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Errorf("error creating Elastic client: %s", err.Error())
	}

	return &ESClient{Client: client}
}

func (c *ESClient) Insert(ctx context.Context, index string, document interface{}) (string, error) {
	// insert new doc into specified index

	// serialize the Todo struct into JSON
	data, err := json.Marshal(document)
	if err != nil {
		return "", err
	}

	// insert it
	res, err := c.Client.Index(
		index,                           // Index name
		strings.NewReader(string(data)), // Document to be indexed
		c.Client.Index.WithContext(ctx),
	)
	if err != nil {
		return "", err // return immediately if error occurs
	}
	defer res.Body.Close() // postpone Close() if there is no error. Carry on.

	if res.IsError() {
		var id string

		// check if document implements an interface with a method named ID
		if docIdProvider, ok := document.(interface{ ID() string }); ok {
			id = docIdProvider.ID()
		}
		return "", fmt.Errorf("error indexing document ID=%s: %s", id, res.String())
	}

	// parse ES response to extract Elastic assigned _id
	var esResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return "", fmt.Errorf("error parsing response: %s", err)
	}

	// convert _id to string. ok var will return true if conversion is successful
	if _id, ok := esResponse["_id"].(string); ok {
		return _id, nil
	} else {
		// Handle the case where _id is not a string or conversion is not feasible
		return "", fmt.Errorf("Elasticsearch _id not found or is not a string")
	}
}

func (c *ESClient) Search(ctx context.Context) ([]interface{}, error) {
	// search Elastic, return results

	// will match all documents as is
	searchQuery := map[string]interface{}{
		"size": 500,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": {},
			},
		},
	}

	// convert searchQuery into JSON
	queryJSON, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("could not encode search query: %s", err)
	}

	// set up request, convert JSON into a Reader obj (non-writable)
	req := c.SearchRequest{
		Index: []string{"_all"},
		Body:  strings.NewReader(string(queryJSON)),
	}

	// execute request
	response, err := req.Do(ctx, c.Client)
	if err != nil {
		return nil, fmt.Errorf("elastic search failed: %s", err)
	}
	defer res.Body.Close() // postpone Close() if there is no error. Carry on.

	// decode response
	var results map[string]interface{}
	// pass pointer to results, so Decode can modify it directly
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("error parsing elastic response: %s", err)
	}

	// verify map kv structure
	hitsMap, ok := results["hits"].(map[string]interface{})
	if !ok {
		return []interface{}{}, fmt.Println("Error with structure returned: %s", ok)
	}

	// if no hits, return empty slice
	hitsSlice, ok := hitsMap["hits"].([]interface{})
	if !ok || len(hitsSlice) == 0 {
		return []interface{}, fmt.Println("Successful response but no results.")
	}

	return hitsSlice, nil

}
