#!/bin/bash
# Description: This script facilitates the development process by ensuring that 
# 'npm run dev' is executed from the correct directory. If the script is run from the root directory 
# of the todo-app, it automatically switches to the frontend directory before executing the command.
# This setup starts both the Go server and the Vue frontend concurrently.

# Get the current working directory name
current_dir=$(basename "$PWD")

# Check if the current directory is the root directory of the todo-app.
# If so, switch to the frontend/ directory.
if [ "$current_dir" = "todo-app" ]; then
    echo "Running from root, switching to frontend..."
    cd frontend
fi

# Run the development script listed in frontend/package.json
# This script starts the Go server and Vue frontend concurrently.
npm run dev
