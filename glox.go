package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/nicholasq/glox/interpreter"
	"github.com/nicholasq/glox/parser"
	"github.com/nicholasq/glox/scanner"
)

var hadError = false
var hadRuntimeError = false

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(fileName string) {
	if len(fileName) < 4 || fileName[len(fileName)-4:] != ".lox" {
		fmt.Println("File must be a .lox file.")
		os.Exit(1)
	}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	defer file.Close()
	bytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	strContents := string(bytes)
	run(strContents)

	if hadError {
		os.Exit(65)
	}

	if hadRuntimeError {
		os.Exit(70)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		if input == ".exit" {
			println("Goodbye!")
			break
		}
		run(input)
		hadError = false
	}
}

func run(script string) {
	scanner := scanner.New(script)
	tokens := scanner.ScanTokens()
	parser := parser.NewParser(&tokens)
	stmts, err := parser.Parse()

	if err != nil {
		// todo add error handling
		panic(err)
	}

	if hadError {
		return
	}
	interpreter := interpreter.Interpreter{}
	value := interpreter.Interpret(*stmts)
	fmt.Printf("%v", value)
}
