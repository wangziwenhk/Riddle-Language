package main

import (
	"Riddle/src/gen"
	"Riddle/src/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

func main() {
	fileName := "E:/Riddle-Go/test/main.red"
	input, err := os.ReadFile(fileName)

	inputStr := string(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	is := antlr.NewInputStream(inputStr)
	lexer := parser.NewRiddleLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRiddleParser(stream)

	// 使用默认的诊断错误监听器
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	programContext := p.Program()

	visitor := gen.NewStmtVisitor()

	result := visitor.Visit(programContext)

	fmt.Println(result)
}
