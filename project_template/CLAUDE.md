# CLAUDE.md

This file provides guidance to Claude (claude.ai/code, Cursor, Claude CLI) when working with code in the **{{PROJECT_NAME}}** repository.

## Project Overview

**{{PROJECT_NAME}}** is a project built with Go-Zero microservices architecture, following IDRM development specifications.

**Version**: Spec v2.0  
**Architecture**: Go-Zero Microservices + Dual ORM  
**Development**: AI-Assisted with Spec-Driven approach

---

## Critical: Read Specifications First

Before any development work, you **MUST** read and understand the project specifications:

### Core Specifications
```bash
@sdd_doc/spec/core/project-charter.md    # Project mission & principles
@sdd_doc/spec/core/tech-stack.md          # Technology stack
@sdd_doc/spec/core/workflow.md            # 5-phase workflow (CRITICAL)
```

### Architecture Specifications
```bash
@sdd_doc/spec/architecture/layered-architecture.md    # Handler→Logic→Model
@sdd_doc/spec/architecture/dual-orm-pattern.md        # GORM + SQLx usage
@sdd_doc/spec/architecture/api-design-guide.md        # RESTful standards
@sdd_doc/spec/architecture/error-handling.md          # Error patterns
```

---

## Development Workflow (5-Phase Mandatory)

**CRITICAL**: All development MUST follow the 5-phase workflow. No exceptions.

```
Phase 0: Context    → Understand specs and prepare environment
   ⚠️ STOP - Wait for user confirmation
Phase 1: Specify    → Define business requirements (EARS notation)
   ⚠️ STOP - Wait for user confirmation
Phase 2: Design     → Create technical solution
   ⚠️ STOP - Wait for user confirmation
Phase 3: Tasks      → Break down into <50 line tasks
   ⚠️ STOP - Wait for user confirmation
Phase 4: Implement  → Code, test, and verify
```

### ⚠️ CRITICAL: Agent Behavior Rules

> **YOU MUST FOLLOW THESE RULES**

1. **ONE PHASE AT A TIME**: Execute only ONE phase per conversation turn
2. **WAIT FOR APPROVAL**: After completing a phase, STOP and ask for user confirmation
3. **NO AUTO-CONTINUE**: NEVER automatically proceed to the next phase
4. **EXPLICIT OUTPUT**: After each phase, list your outputs and ask "Continue to Phase X?"

---

## Architecture Principles

### Layered Architecture (MANDATORY)

```
HTTP Request → Handler → Logic → Model → Database
```

**Rules**:
1. **Handler Layer** (`api/internal/handler/`)
   - Parse and validate parameters ONLY
   - Call logic layer
   - Format response
   - NO business logic
   - Functions ≤ 30 lines

2. **Logic Layer** (`api/internal/logic/`)
   - Business logic implementation
   - Data transformation
   - Call model layer
   - Return business data
   - Functions ≤ 50 lines

3. **Model Layer** (`model/`)
   - Data access ONLY
   - Implement Model interface
   - Choose appropriate ORM
   - Return data models

### Dual ORM Pattern

**GORM** - Use for:
- Complex queries with joins
- Relationships and associations
- Batch operations
- Transactions

**SQLx** - Use for:
- Simple CRUD
- High performance needs
- Direct SQL control
- Lightweight operations

---

## Technology Stack

### Core Stack
- **Language**: Go 1.21+
- **Framework**: Go-Zero v1.9+
- **Database**: MySQL 8.0 / 国产数据库
- **Cache**: Redis 7.0
- **Message Queue**: Kafka 3.0

### Development Tools
- **goctl**: Go-Zero code generation
- **golangci-lint**: Code linting
- **go test**: Testing framework

---

## Project Structure

```
{{PROJECT_NAME}}/
├── api/                          # API services
│   ├── doc/                      # API definitions (.api files)
│   ├── internal/
│   │   ├── handler/             # Handler layer
│   │   ├── logic/               # Logic layer
│   │   ├── svc/                 # Service context
│   │   └── types/               # Request/response types
│   └── etc/                     # Configuration
├── rpc/                          # RPC services (optional)
├── job/                          # Job services (optional)
├── consumer/                     # Consumer services (optional)
├── model/                        # Model layer (Dual ORM)
│   └── {module}/
│       ├── interface.go         # Model interface
│       ├── types.go             # Data structures
│       ├── factory.go           # ORM factory
│       ├── gorm_dao.go          # GORM implementation
│       └── sqlx_model.go        # SQLx implementation
├── pkg/                          # Shared utilities
└── sdd_doc/
    └── spec/                    # Project specifications
```

---

## Coding Standards

### Go Style Guide

**File Organization**:
```go
package xxx

// 1. Imports (grouped: stdlib, external, internal)
import (
    "context"
    "fmt"
    
    "github.com/zeromicro/go-zero/core/logx"
    
    "{{MODULE_PATH}}/pkg/errorx"
)

// 2. Constants
const (
    MaxPageSize = 100
)

// 3. Types
type Service struct { ... }

// 4. Functions
func NewService() *Service { ... }
```

**Naming Conventions**:
- **Files**: lowercase_underscore (`category_logic.go`)
- **Packages**: lowercase, no underscore (`category`)
- **Types**: PascalCase (`CategoryModel`)
- **Functions**: camelCase (private), PascalCase (public)

**Comments**:
- All public functions MUST have comments
- Use Chinese for comments
- Format: `// FunctionName 功能描述`

**Function Size**:
- Handler: ≤ 30 lines
- Logic: ≤ 50 lines
- Model: ≤ 50 lines

---

## Quality Gates

### Gate 1: Specify Phase
- [ ] User stories complete (AS/I WANT/SO THAT)
- [ ] EARS notation used
- [ ] Business rules clear
- [ ] **NO technical implementation details**

### Gate 2: Design Phase
- [ ] Follows layered architecture
- [ ] File checklist complete
- [ ] Interfaces clearly defined

### Gate 3: Tasks Phase
- [ ] Each task < 50 lines
- [ ] Dependencies clear
- [ ] Acceptance criteria specific

### Gate 4: Implement Phase
```bash
# Build check
go build ./...

# Test check (>80% coverage)
go test -cover ./...

# Lint check
golangci-lint run
```

---

## Common Development Commands

### Building and Running
```bash
# Run API service
cd api
go run api.go

# Run RPC service (if included)
cd rpc
go run main.go
```

### Code Generation
```bash
# Generate API code from .api file
goctl api go -api api/doc/api.api -dir api/
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

---

## Important Reminders

### DO ✅
- **ALWAYS** read relevant specs before coding
- **ALWAYS** follow 5-phase workflow
- **ALWAYS** use EARS notation in Phase 1
- **ALWAYS** separate Handler/Logic/Model layers
- **ALWAYS** write tests (>80% coverage)
- **ALWAYS** use Chinese comments
- **ALWAYS** handle errors with context
- **ALWAYS** keep functions < 50 lines

### DON'T ❌
- **NEVER** skip Phase 1 (Specify)
- **NEVER** include technical details in Phase 1
- **NEVER** put business logic in handlers
- **NEVER** put data access in logic layer
- **NEVER** ignore error returns
- **NEVER** commit without tests

---

**Remember: Specifications are law. Read them. Follow them.**

**Version**: Spec v2.0  
**Last Updated**: 2025-12-30
