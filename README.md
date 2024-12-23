# todo-app

![Cute Gopher](cute_gopher.png "Diligent Gopher")

This README provides instructions for setting up and managing the backend part of your application.

Please also check the README in the `frontend/` folder and the `scripts` section of `package.json`.

## Backend Setup

### Dependencies

Ensure all Go dependencies are installed:

```bash
go mod tidy
```

### Compile and run backend:

```bash
go run .
```

To build a binary if needed:

```bash
go build -o todo-app
```

### Project Structure

`api/`: Handlers and routing

## Alternative Setup

Option to start backend and frontend concurrently:

```bash
sh start_app.sh
```

Fire up the frontend by visiting [http://localhost:5173/](http://localhost:5173/)