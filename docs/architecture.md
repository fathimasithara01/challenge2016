RealImage Challenge 2016 — System Architecture

This document explains the architecture, design principles, data flow, and folder structure of the Regional Authorization Engine.
The goal is to determine whether a distributor is authorized to exhibit a film in a given region, based on hierarchical geo rules and include/exclude permissions.

1. High-Level Overview

The system processes:

Geo Hierarchy – Country → State → City

Distributor Permissions – Includes & Excludes

Queries – “Is distributor X allowed in region Y?”

The engine reads region data from CSV, loads authorization rules from examples, and validates queries using the permission logic.

2. Architecture Diagram
                +---------------------------+
                |      cmd/service          |
                |     (entrypoint CLI)      |
                +-------------+-------------+
                              |
                              v
                +---------------------------+
                |   internal/bootstrap      |
                |  (load CSV + init repos) |
                +-------------+-------------+
                              |
        +---------------------+----------------------+
        |                                            |
        v                                            v
+---------------+                           +----------------+
|  Geo Module   |                           | Permissions    |
| (parse/query) |                           |  Module        |
+-------+-------+                           +-------+--------+
        |                                           |
        v                                           v
+---------------+                           +----------------+
| Distributor   |                           |  Checker       |
|  Store        |                           | (authorize)    |
+-------+-------+                           +-------+--------+
                              |
                              v
                    +-----------------+
                    |  CLI Output     |
                    +-----------------+

3. Folder Structure (Explained)
realimage-challenge-2016/
│
├── cmd/
│   └── service/main.go          # CLI entrypoint
│
├── internal/
│   ├── geo/                     # Region parsing + repository
│   │   ├── model.go
│   │   ├── parser.go
│   │   └── repository.go
│   │
│   ├── permissions/             # Include/Exclude authorization logic
│   │   ├── model.go
│   │   ├── validator.go
│   │   ├── service.go
│   │   └── checker.go
│   │
│   └── distributor/             # Distributor info + storage
│       ├── model.go
│       └── store.go
│
├── data/world.csv               # Region definitions
│
├── examples/                    # Sample inputs
│   ├── permissions.txt
│   └── queries.txt
│
├── test/                        # Unit tests
│   ├── permissions_test.go
│   ├── parser_test.go
│   └── checker_test.go
│
├── docs/
│   └── architecture.md          # (this file)
│
├── go.mod
└── README.md

4. Modules & Responsibilities
4.1 Geo Module

Handles:

✔ Parsing "CITY-STATE-COUNTRY" formats
✔ Normalizing names
✔ Looking up parent/child hierarchy
✔ Repository for region data

Input: world.csv
Output: region struct with ID, State, Country.

4.2 Permissions Module

Handles authorization rules:

Include Rules

Distributor is allowed in a region and its descendants.

Exclude Rules

Distributor is specifically blocked in:

the region

or sub-region of that region

Service Responsibilities

✔ Load permissions
✔ Validate allowed/excluded syntax
✔ Expand rules hierarchically

4.3 Checker Module

This is the heart of the challenge.

Given:

distributor

city/state/country

It answers:

YES or NO

By evaluating:

Does any include rule match?

Does an exclusion override it?

Does parent region match?

Does closest rule take priority?

Example:

INCLUDE INDIA
EXCLUDE KARNATAKA-INDIA


Query:

BANGALORE-KARNATAKA-INDIA → NO
DELHI-DELHI-INDIA → YES

4.4 Distributor Module

Just a small layer:

distributor model (name, permissions)

in-memory store

lookup during authorization

5. Data Flow

Load Regions
world.csv → GeoRepo

Parse Permissions
Example file → PermissionService

Store Distributors
distributor → DistributorStore

Run Query
input: distributor + region string

Authorize
Checker validates using rules.

Output Result
YES / NO printed to CLI.

6. Design Principles Used

✔ Clean Architecture (independent modules)
✔ Single Responsibility (each folder has one purpose)
✔ Testability (parser, permissions, checker have unit tests)
✔ Readability (small, focused files)

7. Why This Architecture Works Well

Easy to extend new region data

Simple to add new distributors

All logic is isolated & testable

Professional folder layout (industry-standard)

HR + engineers can understand quickly

8. Future Improvements

Optional upgrades:

Convert to REST API

Add caching

Add logging middleware

Support multiple films

Add web UI to query