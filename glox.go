package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var hadError = false
var hadRuntimeError = false

//var interpreter = Interpreter{}

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		//TestPrinter()
		runPrompt()
	}
}

func runFile(fileName string) {
	println("Running file: ", fileName)
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
	//scanner := MakeScanner(script)
	scanner := Scanner{
		Source:  script,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
	}

	for _, token := range scanner.ScanTokens() {
		fmt.Println(token)
	}

	// parser := Parser{Tokens: scanner.ScanTokens()}
	// statements := parser.parse()
	// if hadError {
	// 	return
	// }
	// resolver := Resolver{interpreter: &interpreter}
	// resolver.resolve(statements)
	// if hadError {
	// 	return
	// }

}

//func error(Line uint, message string) {
//	report(Line, "", message)
//}

func report(line uint, where string, message string) {
	fmt.Printf("[Line %d] Error %s: %s\n", line, where, message)
	hadError = true
}

//func errorToken(token Token, message string) {
//	if token.TokenType == EOF {
//		report(token.Line, " at end", message)
//	} else {
//		report(token.Line, " at '"+token.Lexeme+"'", message)
//	}
//}

//func runtimeError(err *RuntimeError) {
//	fmt.Printf("%s\n[Line %d]\n", err.Message, err.Token.Line)
//	hadError = true
//}
