# Push-Swap

## Description

Push-Swap is a sorting project written in Go.

The goal is to sort a list of integers in ascending order using two stacks, A and B, and a limited set of operations.

The program prints the operations needed to sort the numbers. The checker program verifies whether the operations produce the correct result.

## Project Structure

text
push-swap-project/
├── go.mod
├── piles/
│   ├── piles.go
│   └── sort.go
└── cmd/
    ├── push-swap/
    │   └── main.go
    └── checker/
        └── main.go


## Programs

### push-swap

Reads integer arguments, checks them, sorts the numbers and prints the operations.

### checker

Reads the same integer arguments and receives operations through standard input. It checks whether stack A is sorted and stack B is empty.

## Allowed Operations

- `sa`, `sb`, `ss`: swap the first two elements
- `pa`, `pb`: push an element between stacks
- `ra`, `rb`, `rr`: rotate stacks
- `rra`, `rrb`, `rrr`: reverse rotate stacks

## How to Run

Sort numbers:

bash
go run ./cmd/push-swap 5 2 9 1 7


Check the result:

bash
go run ./cmd/push-swap 5 2 9 1 7 | go run ./cmd/checker 5 2 9 1 7


Expected result:

text
OK


## Validation

Run the following commands:

bash
go test ./...
go vet ./...
go build ./...


## Error Handling

The program rejects invalid integers and duplicate numbers. The checker also rejects unknown instructions.

