# Microservices in Go - Cat Facts Service

A simple Go microservice that demonstrates service architecture patterns by fetching cat facts from an external API with logging middleware.

## 🚀 Quick Start

```bash
# Build the application
make build

# Run the application
make run
```

## 📁 Project Structure

```
microservicesInGo/
├── main.go          # Entry point of the application
├── service.go       # Core business logic for fetching cat facts
├── logging.go       # Logging middleware service
├── types.go         # Data structures/types
├── Makefile         # Build automation
├── go.mod           # Go module definition
└── README.md        # This file
```

## 📋 File Details

### `main.go` - Application Entry Point
```go
func main() {
    svc := NewCatFactService("https://catfact.ninja/fact")
    svc = NewLoggingService(svc)
    fact, err := svc.GetCatFact(context.TODO())
    // ... handle result
}
```
- **Purpose**: Orchestrates the application startup
- **Flow**: Creates services → Wraps with logging → Fetches cat fact → Displays result

### `service.go` - Core Business Logic
```go
type Service interface {
    GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct{ url string }
```
- **Purpose**: Defines the main service interface and implementation
- **Responsibilities**:
  - Makes HTTP GET request to cat facts API
  - Parses JSON response into Go struct
  - Returns structured data or error

### `logging.go` - Middleware Layer
```go
type LoggingService struct {
    Next Service
}
```
- **Purpose**: Adds logging capabilities without modifying core service
- **Pattern**: Decorator pattern - wraps another service
- **Features**:
  - Measures execution time
  - Logs performance metrics
  - Passes calls through to the wrapped service

### `types.go` - Data Structures
```go
type CatFact struct {
    Fact string `json:"fact"`
}
```
- **Purpose**: Defines data models used across the application
- **Features**: JSON tags for automatic parsing

### `Makefile` - Build Automation
```makefile
build:
    go build -o bin/microservicesInGo .

run:
    go run .
```
- **Purpose**: Simplifies common development tasks
- **Commands**:
  - `make build`: Compiles binary to `bin/` directory
  - `make run`: Runs the application directly

## 🔄 Application Flow

### Step-by-Step Execution

1. **Initialization** (`main.go`)
   ```
   Application starts → Creates CatFactService → Wraps with LoggingService
   ```

2. **Service Call** (`main.go`)
   ```
   Calls GetCatFact() → Goes through LoggingService first
   ```

3. **Logging Layer** (`logging.go`)
   ```
   Records start time → Calls wrapped service → Measures duration → Logs timing
   ```

4. **Core Service** (`service.go`)
   ```
   Makes HTTP GET request → Receives JSON response → Parses to CatFact struct
   ```

5. **Response Flow**
   ```
   CatFact struct ← Core Service ← Logging Service ← Main Function
   ```

6. **Output**
   ```
   Prints execution time + Cat fact to console
   ```

### Visual Flow Diagram

```
┌─────────────┐    ┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   main.go   │───▶│ LoggingService  │───▶│ CatFactService  │───▶│  External API   │
│             │    │   (Middleware)  │    │ (Core Service)  │    │ catfact.ninja   │
└─────────────┘    └─────────────────┘    └─────────────────┘    └─────────────────┘
       ▲                     │                       │                       │
       │                     ▼                       ▼                       ▼
   Displays              Logs timing           Makes HTTP GET          Returns JSON
   cat fact              information              request               response
```

## 🏗️ Architecture Patterns

### 1. **Service Interface Pattern**
- All services implement the same `Service` interface
- Enables interchangeable implementations
- Makes testing easier with mock services

### 2. **Decorator Pattern** (Middleware)
- `LoggingService` wraps another service
- Adds functionality without modifying original code
- Can chain multiple decorators (logging, metrics, caching, etc.)

### 3. **Dependency Injection**
- Services receive dependencies through constructors
- Makes code more testable and flexible
- Clear separation of concerns

## 🛠️ Development Commands

### Building
```bash
# Build binary
make build

# Manual build (alternative)
go build -o bin/microservicesInGo .
```

### Running
```bash
# Run with Makefile
make run

# Run directly with Go
go run .

# Run built binary
./bin/microservicesInGo
```

### Development
```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run tests (if any exist)
go test ./...

# Initialize Go module (already done)
go mod init microservicesInGo

# Update dependencies
go mod tidy
```

## 📊 Sample Output

```
GetCatFact took 752.8878ms
Cats have 30 vertebrae (humans have 33 vertebrae during early development; 26 after the sacral and coccygeal regions fuse)
```

- First line: Execution timing from `LoggingService`
- Second line: Random cat fact from the API

## 🧪 Testing the Service

You can test the service manually:

1. **Check if API is accessible**:
   ```bash
   curl https://catfact.ninja/fact
   ```

2. **Verify JSON structure**:
   ```json
   {
     "fact": "Some interesting cat fact here",
     "length": 42
   }
   ```

3. **Run the application**:
   ```bash
   make run
   ```

## 🚀 Next Steps for Learning

### Beginner Enhancements:
1. Add error handling for network failures
2. Add configuration for different API endpoints
3. Create unit tests for each service
4. Add more middleware (metrics, caching)

### Intermediate Enhancements:
1. Add HTTP server to expose REST endpoints
2. Add database storage for facts
3. Implement circuit breaker pattern
4. Add Docker containerization

### Advanced Enhancements:
1. Convert to gRPC service
2. Add distributed tracing
3. Implement service discovery
4. Add Kubernetes deployment files
