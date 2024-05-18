# Frontend

This README provides instructions for setting up and managing the frontend part of your application.

## Project Setup

Install all necessary dependencies:

```bash
npm install
```

## Development

To compile and start the frontend with hot-reload while simultaneously running the Go backend, use the following command:

```bash
npm run dev
```

Please note that this will not hot-reload Go code.

Open your browser and navigate to http://localhost:5173 to view the application.

## Production Build

To compile and minify the frontend for production, which creates a /web directory, run:

```bash
npm run build
```

This command adjusts the build according to the configurations specified in `vite.config.js`.

To see the prod build in action, `cd ..` from `frontend/` and start the Go backend with `go run .`

