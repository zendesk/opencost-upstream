package parser

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*BaseallocationfilterListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func TestParseSimpleFilter(t *testing.T) {
	str := ("label:app=X AND namespace:kubecost")
	input := antlr.NewInputStream(str)
	lexer := NewallocationfilterLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewallocationfilterParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// parser.BuildParseTrees = true
	p.BuildParseTrees = true
	filter := p.Filter()
	fmt.Println("walking")
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), filter)
	fmt.Println("done walking")
}
