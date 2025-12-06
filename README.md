RealImage Challenge 2016 — Territory Authorization Engine (Go)

This repository contains a clean, production-style implementation of the RealImage 2016 backend task.
The goal is to determine whether a distributor is authorized to exhibit a film in a given region, based on hierarchical include/exclude rules.

The solution uses Go with a focus on clean architecture, modular components, and fully testable business logic.

1. Problem Overview

A distributor receives permission through:

Regions to INCLUDE

Regions to EXCLUDE

A query asks:

Is DISTRIBUTOR_X authorized for REGION_Y?

Regions follow the hierarchy:

CITY → STATE → COUNTRY

Example
INCLUDE: INDIA
EXCLUDE: KARNATAKA-INDIA


Queries:

D1 KARNATAKA-INDIA → NO
D1 TAMILNADU-INDIA → YES

2. Features

Parse regions in the format CITY-STATE-COUNTRY

In-memory geographical repository

Distributor permission model (include + exclude)

Deterministic authorization engine

Clean separation of modules

CLI interface with sample input

Unit tests (parser, permissions, checker)

3. Architecture

The system is divided into independent modules:

geo/ – region model, parsing, hierarchy

permissions/ – include/exclude evaluation engine

distributor/ – distributor registry

cmd/service/ – CLI entrypoint

Architecture Diagram

4. Folder Structure
realimage-challenge-2016/
│
├── cmd/
│   └── service/
│       └── main.go
│
├── internal/
│   ├── geo/
│   │   ├── model.go
│   │   ├── parser.go
│   │   └── repository.go
│   │
│   ├── permissions/
│   │   ├── model.go
│   │   ├── service.go
│   │   └── checker.go
│   │
│   └── distributor/
│       ├── model.go
│       └── store.go
│
├── examples/
│   ├── permissions.txt
│   └── queries.txt
│
├── data/
│   └── cities_sample.csv
│
├── docs/
│   ├── architecture.md
│   └── architecture.png
│
├── script/
│   └── run.sh
│
├── test/
│   ├── parser_test.go
│   ├── permissions_test.go
│   └── checker_test.go
│
├── go.mod
└── README.md

5. How the Engine Works
Step 1 — Load Permissions

Each distributor may define:

INCLUDE: <Region>
EXCLUDE: <Region>


Internally stored as:

[]Region includes

[]Region excludes

Step 2 — Parse Query

Format:

<DISTRIBUTOR> <REGION>


Example:

D1 KARNATAKA-INDIA

Step 3 — Evaluate

Check if region matches the include list

Check if it matches the exclude list

Apply hierarchical evaluation:

city

state

country

Output:

YES / NO

6. Running the Project
Build
go build -o auth ./cmd/service

Run
./auth examples/permissions.txt examples/queries.txt

Using the script (macOS/Linux)
sh script/run.sh

7. Running Tests
go test ./...


Test coverage includes:

Region parser

Permission logic

Authorization checker

8. Why This Implementation Is Strong

Clean, readable Go code

Modular architecture

Deterministic and testable business logic

Zero unnecessary dependencies

Simple to extend and modify

Easy for reviewers to evaluate

Includes documentation and architecture diagram

9. Sample Input/Output
permissions.txt
D1
INCLUDE: INDIA
EXCLUDE: KARNATAKA-INDIA

queries.txt
D1 KARNATAKA-INDIA
D1 TAMILNADU-INDIA

Output
NO
YES

10. License

Submitted as part of the RealImage Challenge 2016.
May be used or extended for educational and interview purposes.