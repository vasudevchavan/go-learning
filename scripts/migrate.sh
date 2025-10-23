#!/bin/bash

# Go Learning Workspace Migration Script
set -e

echo "🚀 Starting Go Learning Workspace Restructure..."

# Create new directory structure
echo "📁 Creating new directory structure..."
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

# Migrate files with improved naming
echo "📦 Migrating files..."

# Arguments → Fundamentals
cp Arguments/arg.go 01-fundamentals/basics/command-line-args.go
cp Arguments/arg-legth.go 01-fundamentals/basics/args-validation.go
cp Arguments/cmd-exists.go 01-fundamentals/basics/command-validation.go

# Data types → Data structures
mkdir -p 02-data-structures/basics
cp data-type/data-type-explained.go 02-data-structures/structs/basic-struct.go
cp data-type/explore_map.go 02-data-structures/maps/map-operations.go
cp data-type/explore_strings.go 02-data-structures/basics/string-operations.go
cp data-type/explores_struct.go 02-data-structures/structs/advanced-struct.go
cp data-type/interface_example1.go 02-data-structures/interfaces/logger-interface.go
cp data-type/interfaces.go 02-data-structures/interfaces/basic-interface.go

# Error handling
cp error-handling/input-val-error.go 03-error-handling/basic-errors/input-validation.go

# Concurrency
for i in {1..9}; do
    if [ -d "sample_programs/goRoutines/practice$i" ]; then
        cp -r sample_programs/goRoutines/practice$i 04-concurrency/goroutines/example-$(printf "%02d" $i)
    fi
done
cp -r sample_programs/selectstatement 04-concurrency/select/

# I/O and Files
cp Read-input/dup-readline.go 06-io-files/file-operations/read-lines.go
cp quiz/problems.csv 06-io-files/csv-handling/

# Web Development
cp -r sample_programs/WebApplication 07-web-development/http-server/game-api

# Testing
cp testing/main.go 08-testing/unit-tests/basic-test.go
cp goOperator/diskBenchMark/main.go 08-testing/benchmarks/disk-benchmark.go

# Design Patterns
cp -r designPatterns/singleResponsibilityPrinciple 09-design-patterns/solid-principles/srp
cp -r designPatterns/OpenClosePrinciple 09-design-patterns/solid-principles/ocp
cp -r designPatterns/LiskovSubstitionPrinciple 09-design-patterns/solid-principles/lsp
cp -r designPatterns/builder 09-design-patterns/creational/builder

# Projects
cp -r quiz 10-projects/cli-tools/quiz-app
cp -r sample_programs/gamerockpaper 10-projects/games/rock-paper-scissors
cp -r sample_programs/exported 05-packages-modules/creating-packages/staff-example

# Documentation
cp -r Notes/* docs/cheatsheets/

echo "✅ Migration completed!"
echo "📖 Check RESTRUCTURE_PLAN.md for detailed information"
echo "🧹 You can now remove old directories after verifying the migration"