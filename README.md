# go-cards 🃏

A Go project for learning the language through building a playing card deck simulator.

## About

This project is a hands-on way to explore Go fundamentals — custom types, receiver functions, slices, loops, and package structure — using a familiar real-world concept: a deck of playing cards.

Built while following [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/) on Udemy.

## What It Does (So Far)

- Defines a custom `deck` type (a slice of strings)
- Generates a full 52-card deck across 4 suits
- Prints all cards to the console

## Project Structure

```
go-cards/
├── main.go      # Entry point
└── deck.go      # deck type and card logic
```

## Running the Project

Make sure you have [Go installed](https://go.dev/dl/), then:

```bash
go run main.go deck.go
```

## Planned Features

- Deal a hand of cards
- Shuffle the deck
- Save/load a deck to/from a file

## What I'm Learning

- Custom types in Go
- Receiver functions
- Working with slices and `range`
- Splitting code across multiple files in a package