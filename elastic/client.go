package api

import (
  "fmt"
  "context"
  "encoding/json"
  "strings"
  "github.com/elastic/go-elasticsearch/v7"
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
    index,                            // Index name
    strings.NewReader(string(data)),  // Document to be indexed
    c.Client.Index.WithContext(ctx),
  )
  if err != nil {
      return "", err // return immediately if error occurs
  }
  defer res.Body.Close() // postpone Close() if there is no error. Carry on.

  if res.IsError() {
    var id string

    // check if document implements an interface with a method named ID
    if docIdProvider, ok := document.(interface{ ID() string}); ok {
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
