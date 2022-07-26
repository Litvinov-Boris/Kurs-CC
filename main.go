package main
// debug false

import (
	"awesomeProject/lolcode3"
	visualization "awesomeProject/vizualization"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	//"github.com/go-echarts/go-echarts/v2/charts"
	//"github.com/go-echarts/go-echarts/v2/opts"
)

func main()  {
	is, err := antlr.NewFileStream("./tests/test_func.txt")
	if err != nil {
		fmt.Printf("No input file provided")
	}

	lexer := parser.NewlolcodeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewlolcodeParser(stream)
	listener := parser.NewBaseListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	visualization.Graph(listener)
	parser.Print_decl()
	parser.Ret_funcs()
}
