# CLAUDE.md

This file provides guidance to Claude (claude.ai/code, Cursor, Claude CLI) when working with code in the IDRM repository.

## Project Overview

**IDRM** (Intelligent Data Resource Management) is an intelligent data resource management platform built with Go-Zero microservices architecture. It provides comprehensive data lifecycle management including cataloging, access control, data quality, and intelligent analytics.

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

### Workflow Specifications
```bash
@sdd_doc/spec/workflow/phase1-specify.md              # Requirements (EARS)
@sdd_doc/spec/workflow/phase2-design.md               # Technical design
@sdd_doc/spec/workflow/ears-notation-guide.md         # EARS详解
```

---

## Development Workflow (5-Phase Mandatory)

**CRITICAL**: All development MUST follow the 5-phase workflow. No exceptions.

```
Phase 0: Context    → Understand specs and prepare environment
Phase 1: Specify    → Define business requirements (EARS notation)
Phase 2: Design     → Create technical solution
Phase 3: Tasks      → Break down into <50 line tasks
Phase 4: Implement  → Code, test, and verify
```

### Phase 0: Context
**Before starting any work:**
1. Read `sdd_doc/spec/core/project-charter.md`
2. Understand related architecture specs
3. Review coding standards
4. Confirm development environment is ready

### Phase 1: Specify (Requirements)
**Focus**: Business requirements ONLY, NO technical details

**Output**: `specs/features/{feature-name}/requirements.md`

**Must include:**
- User Stories (AS a/I WANT/SO THAT)
- Acceptance Criteria (EARS notation)
- Business Rules (NOT technical constraints)
- Data Considerations (NOT database schema)
- Open Questions

**EARS Format**:
```
WHEN [condition/event]
THE SYSTEM SHALL [expected behavior]
```

**Example**:
```markdown
WHEN 用户提交有效的分类创建请求
THE SYSTEM SHALL 保存分类并返回201状态码和分类ID

WHEN 用户提交的分类名称为空
THE SYSTEM SHALL 返回400错误和"名称不能为空"的错误信息
```

**Reference**: `@sdd_doc/spec/workflow/ears-notation-guide.md`

### Phase 2: Design (Technical Solution)
**Focus**: Technical implementation details

**Output**: `specs/features/{feature-name}/design.md`

**Must include:**
- Architecture adherence (Handler→Logic→Model)
- File structure and locations
- Interface definitions (Go interfaces)
- Sequence diagrams
- ORM choice (GORM vs SQLx) with justification
- Technical constraints
- Database schema

### Phase 3: Tasks (Breakdown)
**Output**: `specs/features/{feature-name}/tasks.md`

**Rules**:
- Each task ≤ 50 lines of code
- Clear dependencies
- Specific acceptance criteria
- Estimated lines

### Phase 4: Implement (Code & Test)
**Activities**:
1. Implement tasks sequentially
2. Write tests (>80% coverage)
3. Self review
4. Run quality gates

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

**Reference**: `@sdd_doc/spec/architecture/layered-architecture.md`

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

**Reference**: `@sdd_doc/spec/architecture/dual-orm-pattern.md`

---

## Technology Stack

### Core Stack
- **Language**: Go 1.21+
- **Framework**: Go-Zero v1.9+
- **Database**: MySQL 8.0
- **Cache**: Redis 7.0
- **Message Queue**: Kafka 3.0

### Development Tools
- **goctl**: Go-Zero code generation
- **golangci-lint**: Code linting
- **go test**: Testing framework
- **wire**: Dependency injection (optional)

### AI Tools
- **Kiro.dev**: Structured development, large features
- **Cursor**: Daily development and refactoring
- **Claude CLI**: Batch processing and CI/CD

**Reference**: `@sdd_doc/spec/core/tech-stack.md`

---

## Project Structure

```
idrm/
├── api/                          # API services
│   ├── doc/                      # API definitions (.api files)
│   ├── internal/
│   │   ├── handler/             # Handler layer
│   │   ├── logic/               # Logic layer
│   │   ├── svc/                 # Service context
│   │   └── types/               # Request/response types
│   └── etc/                     # Configuration
├── rpc/                          # RPC services
├── model/                        # Model layer (Dual ORM)
│   └── {module}/
│       ├── interface.go         # Model interface
│       ├── types.go             # Data structures
│       ├── factory.go           # ORM factory
│       ├── gorm_dao.go          # GORM implementation
│       └── sqlx_model.go        # SQLx implementation
├── common/                       # Shared utilities
└── sdd_doc/
    └── spec/                    # Project specifications (READ FIRST!)
        ├── core/                # Core specs
        ├── architecture/        # Architecture specs
        ├── coding-standards/    # Coding standards
        ├── workflow/           # Workflow guides
        └── quality/            # Quality standards
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
    
    "af_idrm/common/errorx"
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
- **Constants**: PascalCase or UPPER_SNAKE_CASE

**Comments**:
- All public functions MUST have comments
- Use Chinese for comments
- Format: `// FunctionName 功能描述`

**Function Size**:
- Handler: ≤ 30 lines
- Logic: ≤ 50 lines
- Model: ≤ 50 lines

**Error Handling**:
```go
// Always wrap errors with context
if err != nil {
    return fmt.Errorf("failed to create category: %w", err)
}
```

**Reference**: `@sdd_doc/spec/coding-standards/go-style-guide.md`

---

## Quality Gates

### Gate 1: Specify Phase
- [ ] User stories complete (AS/I WANT/SO THAT)
- [ ] EARS notation used
- [ ] Business rules clear
- [ ] Data considerations defined
- [ ] **NO technical implementation details**

### Gate 2: Design Phase
- [ ] Follows layered architecture
- [ ] File checklist complete
- [ ] Interfaces clearly defined
- [ ] Sequence diagrams included

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

**Reference**: `@sdd_doc/spec/quality/quality-gates.md`

---

## Common Development Commands

### Building and Running
```bash
# Run API service
cd api
go run xxx.go

# Run RPC service
cd rpc/resource_catalog
go run resource_catalog.go
```

### Code Generation
```bash
# Generate API code from .api file
goctl api go -api api/doc/api.api -dir api/

# Generate RPC code from .proto file
goctl rpc protoc rpc/resource_catalog/resource_catalog.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test -v ./api/internal/logic/category/...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Database
```bash
# Connect to MySQL
mysql -h localhost -u root -p idrm

# Run migrations (if using)
migrate -path ./migrations -database "mysql://user:pass@tcp(localhost:3306)/idrm" up
```

---

## API Design Standards

### RESTful Conventions
```
GET    /api/v1/categories       # List
GET    /api/v1/categories/:id   # Get one
POST   /api/v1/categories       # Create
PUT    /api/v1/categories/:id   # Update
DELETE /api/v1/categories/:id   # Delete
```

### Response Format
```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

### Error Response
```json
{
  "code": 40001,
  "message": "名称不能为空",
  "data": null
}
```

### Status Codes
- `200`: Success
- `201`: Created
- `400`: Bad request / Validation error
- `401`: Unauthorized
- `403`: Forbidden
- `404`: Not found
- `409`: Conflict
- `500`: Internal server error

**Reference**: `@sdd_doc/spec/architecture/api-design-guide.md`

---

## Error Handling

### Error Code Range
- `10000-19999`: System errors
- `20000-29999`: Business errors
- `30000-39999`: Validation errors
- `40000-49999`: Permission errors

### Error Wrapping
```go
// Always use %w to wrap errors
if err != nil {
    return fmt.Errorf("failed to query category: %w", err)
}
```

### Custom Errors
```go
// Use errorx package for business errors
return errorx.NewCodeError(40001, "分类名称已存在")
```

**Reference**: `@sdd_doc/spec/architecture/error-handling.md`

---

## Testing Standards

### Test Coverage
- **Minimum**: 80% for business logic
- **Unit tests**: All logic layer functions
- **Integration tests**: Critical flows
- **Table-driven tests**: Preferred pattern

### Test File Naming
```
category_logic.go      → category_logic_test.go
gorm_dao.go           → gorm_dao_test.go
```

### Test Structure
```go
func TestCreateCategory(t *testing.T) {
    tests := []struct {
        name    string
        input   *CreateCategoryRequest
        want    *Category
        wantErr bool
    }{
        {
            name: "valid input",
            input: &CreateCategoryRequest{Name: "test"},
            want: &Category{ID: 1, Name: "test"},
            wantErr: false,
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic
        })
    }
}
```

**Reference**: `@sdd_doc/spec/coding-standards/testing-standards.md`

---

## AI Collaboration Guidelines

### When to Use Which AI Tool

**Kiro.dev** - Use for:
- New large features (>500 lines)
- Complex multi-file changes
- Full 5-phase workflow
- Team collaboration

**Cursor** - Use for:
- Daily development (<500 lines)
- Quick refactoring
- Single feature implementation
- Code review assistance

**Claude CLI** - Use for:
- Batch processing
- Documentation generation
- Test generation
- CI/CD integration

### AI Prompt Template

When requesting AI assistance, structure your prompt:

```
@sdd_doc/spec/workflow/phase{X}-{name}.md
@sdd_doc/spec/architecture/{relevant-spec}.md

Context: [Describe what you're working on]

Current Phase: Phase {X}

Task: [Specific task]

Requirements:
- Follow layered architecture
- Use EARS notation (if Phase 1)
- Functions < 50 lines
- Complete Chinese comments
- >80% test coverage

Output: [What you expect]
```

### AI Code Review Checklist

Before accepting AI-generated code, verify:
- [ ] Follows layered architecture
- [ ] Function size limits respected
- [ ] Complete error handling
- [ ] Chinese comments present
- [ ] Tests included
- [ ] No hardcoded values
- [ ] Follows naming conventions

---

## Git Workflow

### Branch Naming
```
feature/category-management
fix/query-performance
docs/update-api-spec
refactor/model-layer
```

### Commit Convention
```bash
feat: 添加资源分类管理功能
fix: 修复分类查询bug
docs: 更新API文档
refactor: 重构model层
test: 添加单元测试
chore: 更新依赖版本
```

### PR Process
1. Self review
2. Run quality gates locally
3. Create PR with description
4. Wait for CI/CD checks
5. Peer review (≥1 person)
6. Tech lead approval
7. Merge to develop

---

## Configuration

### Environment Variables
```bash
# MySQL
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=idrm

# Redis
REDIS_HOST=127.0.0.1
REDIS_PORT=6379

# Service
SERVICE_PORT=8888
```

### Config Files
- API service: `api/etc/{service}.yaml`
- RPC service: `rpc/{service}/etc/{service}.yaml`

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
- **ALWAYS** run quality gates before committing

### DON'T ❌
- **NEVER** skip Phase 1 (Specify)
- **NEVER** include technical details in Phase 1
- **NEVER** put business logic in handlers
- **NEVER** put data access in logic layer
- **NEVER** ignore error returns
- **NEVER** commit without tests
- **NEVER** hardcode configuration
- **NEVER** exceed function size limits
- **NEVER** skip quality gate checks

---

## Quick Reference

### Key Specification Files
| Topic | File |
|-------|------|
| Project Mission | `sdd_doc/spec/core/project-charter.md` |
| Workflow | `sdd_doc/spec/core/workflow.md` |
| Architecture | `sdd_doc/spec/architecture/layered-architecture.md` |
| EARS Guide | `sdd_doc/spec/workflow/ears-notation-guide.md` |
| Quality Gates | `sdd_doc/spec/quality/quality-gates.md` |

### Contact
- Tech Lead: [Name]
- Architecture Questions: See `sdd_doc/spec/architecture/`
- Workflow Questions: See `sdd_doc/spec/workflow/`

---

**Remember: Specifications are law. Read them. Follow them. Update them when needed.**

**Version**: Spec v2.0  
**Last Updated**: 2025-12-29
