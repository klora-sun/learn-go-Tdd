# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
This is a Go web server that tracks player win counts. It follows Test-Driven Development (TDD) practices.

### Architecture
- `main.go`: Contains the HTTP server and playerServer handler function
- `main_test.go`: Contains tests for the player API endpoints
- `go.mod`: Go module definition

### API Endpoints
- `GET /players/{name}`: Returns the win count for a specific player
- `POST /players/{name}`: (Planned) Increment win count for a player

### Current Implementation
- Uses `strings.TrimPrefix()` to extract player names from URL paths
- Uses `switch` statement to handle different players
- Returns `http.NotFound()` for unknown players
- Hardcoded scores: Pepper=20, Floyd=40

## Common Commands

### Run Tests
```bash
go test -v
```

### Run Specific Test
```bash
go test -run TestGETPlayers -v
```

### Build and Run Server
```bash
go build -o server
./server
```

### Format Code
```bash
go fmt ./...
```

## Development Workflow
1. Write failing test (Red phase)
2. Write minimal code to pass test (Green phase)  
3. Refactor code while keeping tests passing (Refactor phase)

## Next Steps
- Implement POST endpoint for incrementing scores
- Add persistent data storage
- Set up actual HTTP server in main() function