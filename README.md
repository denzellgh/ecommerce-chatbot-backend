# Makers Tech - Backend API

RESTful API service for inventory management and AI-powered chatbot recommendations built with Go, Chi router, and PostgreSQL.

> **Important**: For code review, please use the `main` branch.

## Tech Stack

- **Language**: [Go 1.21+](https://golang.org/)
- **Router**: [Chi v5](https://github.com/go-chi/chi)
- **Database**: [PostgreSQL 15+](https://www.postgresql.org/)
- **AI Integration**: [Ollama](https://ollama.ai/) with DeepSeek-R1:1.5b
- **Database Driver**: [lib/pq](https://github.com/lib/pq)
- **Containerization**: Docker & Docker Compose

## Features

- RESTful API for product inventory management
- AI-powered chatbot with natural language processing
- Smart recommendation system with user preferences
- Real-time inventory queries and stock management
- Session-based user preference storage
- CORS-enabled for frontend integration
- Health check endpoints for monitoring

## Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 15+
- Docker & Docker Compose (optional)
- Ollama with DeepSeek-R1:1.5b model

## AI Integration

### Ollama Setup

```bash
# Install Ollama
curl -fsSL https://ollama.ai/install.sh | sh

# Pull the required model
ollama pull deepseek-r1:1.5b

# Start Ollama service
ollama serve
```

### Chat Processing Flow

1. User sends message via POST `/api/chat`
2. System retrieves current inventory data
3. Message + inventory context sent to Ollama
4. AI processes request and generates response
5. Formatted response returned to user

## Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/new-feature`
3. Make changes and test locally
4. Run tests: `go test ./...`
5. Commit changes: `git commit -am 'Add new feature'`
6. Push to branch: `git push origin feature/new-feature`
7. Submit Pull Request to `dev` branch
