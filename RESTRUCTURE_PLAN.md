# Go Learning Workspace Restructure Plan

## Current Issues
- Inconsistent naming conventions (Arguments vs data-type)
- Mixed organization patterns
- Scattered related concepts
- Missing proper module structure

## Proposed Structure

```
go-learning/
├── 01-fundamentals/
│   ├── basics/
│   │   ├── hello-world/
│   │   ├── variables-types/
│   │   └── operators/
│   ├── control-flow/
│   │   ├── conditionals/
│   │   ├── loops/
│   │   └── switch/
│   └── functions/
│       ├── basic-functions/
│       ├── variadic/
│       └── closures/
├── 02-data-structures/
│   ├── arrays-slices/
│   ├── maps/
│   ├── structs/
│   └── interfaces/
├── 03-error-handling/
│   ├── basic-errors/
│   ├── custom-errors/
│   └── panic-recover/
├── 04-concurrency/
│   ├── goroutines/
│   ├── channels/
│   ├── select/
│   ├── sync/
│   └── patterns/
├── 05-packages-modules/
│   ├── creating-packages/
│   ├── imports/
│   └── go-modules/
├── 06-io-files/
│   ├── file-operations/
│   ├── csv-handling/
│   └── json-xml/
├── 07-web-development/
│   ├── http-server/
│   ├── templates/
│   ├── json-api/
│   └── middleware/
├── 08-testing/
│   ├── unit-tests/
│   ├── benchmarks/
│   └── table-tests/
├── 09-design-patterns/
│   ├── solid-principles/
│   ├── creational/
│   ├── structural/
│   └── behavioral/
├── 10-projects/
│   ├── cli-tools/
│   ├── web-apps/
│   └── games/
├── docs/
│   ├── cheatsheets/
│   ├── best-practices/
│   └── resources/
└── scripts/
    ├── setup.sh
    └── run-examples.sh
```

## Migration Steps

### 1. Create New Structure
```bash
mkdir -p 01-fundamentals/{basics,control-flow,functions}
mkdir -p 02-data-structures/{arrays-slices,maps,structs,interfaces}
mkdir -p 03-error-handling/{basic-errors,custom-errors,panic-recover}
mkdir -p 04-concurrency/{goroutines,channels,select,sync,patterns}
mkdir -p 05-packages-modules/{creating-packages,imports,go-modules}
mkdir -p 06-io-files/{file-operations,csv-handling,json-xml}
mkdir -p 07-web-development/{http-server,templates,json-api,middleware}
mkdir -p 08-testing/{unit-tests,benchmarks,table-tests}
mkdir -p 09-design-patterns/{solid-principles,creational,structural,behavioral}
mkdir -p 10-projects/{cli-tools,web-apps,games}
mkdir -p docs/{cheatsheets,best-practices,resources}
mkdir -p scripts
```

### 2. File Migrations

#### Arguments → 01-fundamentals/basics/
- `arg.go` → `command-line-args.go`
- `arg-length.go` → `args-validation.go`
- `cmd-exists.go` → `command-validation.go`

#### data-type → 02-data-structures/
- `data-type-explained.go` → `structs/basic-struct.go`
- `explore_map.go` → `maps/map-operations.go`
- `interface_example1.go` → `interfaces/logger-interface.go`

#### sample_programs/goRoutines → 04-concurrency/
- `practice1-9/` → `goroutines/examples/`
- Add proper naming: `01-basic-goroutine.go`, `02-waitgroup.go`, etc.

#### designPatterns → 09-design-patterns/
- `singleResponsibilityPrinciple/` → `solid-principles/srp/`
- `OpenClosePrinciple/` → `solid-principles/ocp/`

#### Projects → 10-projects/
- `quiz/` → `cli-tools/quiz-app/`
- `gamerockpaper/` → `games/rock-paper-scissors/`
- `WebApplication/` → `web-apps/game-api/`

#### Documentation → docs/
- `Notes/` → `docs/cheatsheets/`
- Add `docs/best-practices/` for coding standards

### 3. Standardize File Naming
- Use kebab-case: `basic-goroutine.go`
- Descriptive names: `logger-interface.go` vs `interface_example1.go`
- Consistent prefixes for related files

### 4. Add Module Structure
Each major section should have:
- `README.md` with learning objectives
- `go.mod` for standalone examples
- `examples/` subdirectory for multiple files
- `exercises/` for practice problems

### 5. Improve Code Quality
- Add proper comments and documentation
- Consistent error handling patterns
- Remove unused variables and imports
- Add unit tests for each example

## Benefits
- Clear learning progression (01-10)
- Related concepts grouped together
- Consistent naming conventions
- Better discoverability
- Easier maintenance
- Professional structure