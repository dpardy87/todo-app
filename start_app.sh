#!/bin/bash

# check if ES container is already running
if [ $(docker ps -q -f name=elasticsearch | wc -l) -eq 1 ]; then
  echo "Elasticsearch container is already running."
else
  # start it
  docker run -d --name elasticsearch -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 \
  -e "discovery.type=single-node" \
  -e "xpack.security.enabled=false" \
  -e "xpack.security.http.ssl.enabled=false" \
  docker.elastic.co/elasticsearch/elasticsearch:8.13.1
  echo "Elasticsearch container started."
fi

# wait for Elasticsearch to be up and running
echo "Waiting for Elasticsearch to start"
until curl -s -o /dev/null -w "%{http_code}" "http://localhost:9200" | grep -q 200; do
  printf '.'
  sleep 1
done
echo "Elasticsearch is up and running."

# check if 'todos' index exists
if curl -s -o /dev/null -w "%{http_code}" "http://localhost:9200/todos" | grep -q 200; then
    echo "Index 'todos' already exists."
else
   # Create 'todos' index
    response=$(curl -s -X PUT "localhost:9200/todos" -H 'Content-Type: application/json' -d'
    {
      "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 1
      }
    }
    ')

    if [ -z "$response" ]; then
        echo "Failed to create 'todos' index. No response from server."
        exit 1
    else
        echo "\nIndex 'todos' created successfully."
    fi
fi

# get current dir
current_dir=$(basename "$PWD")

# check if the current directory is the root directory
# if so, switch to frontend/
if [ "$current_dir" = "todo-app" ]; then
    echo "Running from root, switching to frontend/ directory..."
    cd frontend
fi

# this package.json script starts the Go server and Vue frontend concurrently.
npm run dev
