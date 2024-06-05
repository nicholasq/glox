package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ExprType struct {
	name   string
	fields []string
}

func main() {
	exprTypes := exprBuilder([]string{
		"Binary     :  left Expr, operator Token, right Expr",
		"Grouping   :  expression Expr",
		"Literal    :  value interface{}",
		"Unary      :  operator Token, right Expr",
	})

	defineAst("./", "Expr", exprTypes)

	fmtCmd := exec.Command("make", "format")

	_, err := fmtCmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

func exprBuilder(input []string) []ExprType {
	var expressions []ExprType

	for _, line := range input {
		sections := strings.Split(line, ":")
		typeName := strings.TrimSpace(sections[0])
		fields := strings.Split(sections[1], ",")
		for i, field := range fields {
			fields[i] = strings.TrimSpace(field)
		}
		expr := ExprType{typeName, fields}
		expressions = append(expressions, expr)
	}

	return expressions
}

func defineAst(outputDir string, baseName string, types []ExprType) {
	file, err := os.Create(outputDir + "/" + baseName + ".go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "package main\n\n")
	fmt.Fprintf(writer, "type %s interface {\n", baseName)
	fmt.Fprintf(writer, "Visit(v ExpressionVisitor) interface{}\n")
	fmt.Fprintf(writer, "}\n\n")

	defineVisitor(writer, baseName, types)

	for _, exprType := range types {
		defineTypes(writer, exprType.name, exprType.fields)
	}

	for _, exprType := range types {
		defineAccept(writer, exprType.name)
	}

	writer.Flush()
}

func defineVisitor(writer *bufio.Writer, baseName string, types []ExprType) {
	fmt.Fprintf(writer, "type ExpressionVisitor interface {\n")
	for _, exprType := range types {
		typeName := exprType.name
		fmt.Fprintf(writer, "  Visit%s%s(%s %s) interface{}\n", typeName, baseName, strings.ToLower(baseName), typeName)
	}
	fmt.Fprintf(writer, "}\n\n")
}

func defineTypes(writer *bufio.Writer, typeName string, fields []string) {
	fmt.Fprintf(writer, "type %s struct {\n", typeName)
	for _, field := range fields {
		fmt.Fprintf(writer, "  %s\n", field)
	}
	fmt.Fprintf(writer, "}\n\n")
}

func defineAccept(writer *bufio.Writer, baseName string) {
	fmt.Fprintf(writer, "func (expr %s) Accept(v ExpressionVisitor) interface{} {\n", baseName)
	fmt.Fprintf(writer, "    return v.Visit%sExpr(%s)\n", baseName, "expr")
	fmt.Fprintf(writer, "}\n\n")
}
