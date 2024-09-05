package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"todo-app/models"
)

// ESClient wraps the Elasticsearch client
type ESClient struct {
	Client *elasticsearch.Client
}

func NewClient(addresses []string) (*ESClient, error) {
	// creates ES client with provided addresses
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ESClient{Client: client}, nil
}

func (es *ESClient) GetAll(ctx context.Context, index string) ([]models.Todo, error) {
	// create request
	var buf bytes.Buffer
	query := `{
		"query": {
			"match_all": {}
		}
	}`
	buf.WriteString(query)

	req := esapi.SearchRequest{
		Index: []string{index},
		Body: &buf,
	}

	// Do() executes request and returns response (or error)
	res, err := req.Do(ctx, es.Client)
	if err != nil {
		fmt.Printf("Error when executing request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError(){
		fmt.Printf("Elasticsearch error response: %v\n", res)
		return nil, errors.New(fmt.Sprintf("Error fetching Elastic documents: %s", res.String()))
	}

	// parse response into map
	var resMap map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		fmt.Printf("Error decoding resMap: %v", err)
		return nil, err
	}

	// get hits from response
	hits := resMap["hits"].(map[string]interface{})["hits"].([]interface{})
	var todos []models.Todo

	// iterate through kv hits
	for _, hit := range hits {
		// Each hit is a document
		doc := hit.(map[string]interface{})["_source"]

		// convert to Todo struct
		todoBytes, _ := json.Marshal(doc)

		var todo models.Todo
		if err := json.Unmarshal(todoBytes, &todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (es *ESClient) Insert(ctx context.Context, index string, doc interface{}) (string, error) {
	// convert doc to JSON first
	body, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	// create the request
	req := esapi.IndexRequest{
		Index: index,
		DocumentID: "",
		Body: bytes.NewReader(body),
		Refresh: "true",
	}

	res, err := req.Do(ctx, es.Client)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// check for errors in response
	if res.IsError(){
		return "", errors.New(fmt.Sprintf("Error indexing document: %s", res.String()))
	}

	// decode response
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		return "", err
	}

	// extract and return doc ID
	if id, ok := resMap["_id"].(string); ok {
		return id, nil
	}

	return "", errors.New("Failed to retrieve document ID from response")
}