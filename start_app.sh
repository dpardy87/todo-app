#!/bin/bash

# check if ES container is already running
if [ $(docker ps -q -f name=elasticsearch | wc -l) -eq 1 ]; then
  echo "Elasticsearch container is already running."
else
  # start it
  docker run -d --name elasticsearch -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.13.4
  echo "Elasticsearch container started."
fi

# check if 'todos' index exists
if curl -s -o /dev/null -w "%{http_code}" "http://localhost:9200/todos" | grep -q 200; then
    echo "Index 'todos' already exists."
else
    # create 'todos' index
    curl -X PUT "localhost:9200/todos" -H 'Content-Type: application/json' -d'
    {
      "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 1
      }
    }
    '
    echo "\nIndex 'todos' created successfully."
fi

# get current dir
current_dir=$(basename "$PWD")

# check if the current directory is the root directory
# if so, switch to frontend/
if [ "$current_dir" = "todo-app" ]; then
    echo "Running from root, switching to frontend..."
    cd frontend
fi

# this package.json script starts the Go server and Vue frontend concurrently.
npm run dev
