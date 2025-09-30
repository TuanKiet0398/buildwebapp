# BuildWebApp

A simple Go web application demonstrating template rendering and basic routing.

## Features
- Home and About pages using HTML templates
- Modular handler and render packages
- Easy to extend for more routes and templates

## Project Structure
```
buildwebapp/
├── cmd/web/main.go         # Entry point for the web server
├── pkg/handlers/handlers.go # HTTP handlers for routes
├── pkg/render/render.go     # Template rendering logic
├── templates/
│   ├── home.page.tmpl      # Home page template
│   └── about.page.tmpl     # About page template
└── go.mod                  # Go module definition
```

## Getting Started

### Prerequisites
- Go 1.18 or newer

### Installation
1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd buildwebapp
   ```
2. Download dependencies:
   ```bash
   go mod tidy
   ```

### Running the App
1. Start the server:
   ```bash
   go run cmd/web/main.go
   ```
2. Visit [http://localhost:8080](http://localhost:8080) for the Home page.
3. Visit [http://localhost:8080/about](http://localhost:8080/about) for the About page.

## Customization
- Add more templates to the `templates/` folder.
- Add new handler functions in `pkg/handlers/handlers.go` and register them in `main.go`.

## License
MIT
