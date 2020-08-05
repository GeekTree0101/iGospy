package worker

import (
	"fmt"

	"github.com/GeekTree0101/iSpygo/model"
)

// ParserError is error for parser
type ParserError struct {
	desc string
}

// Error is parser error constructor
func (pe *ParserError) Error() string {
	return fmt.Sprintf("%s", pe.desc)
}

// Parser converts pasted raw-swift-code to spy object
type Parser struct {
}

// NewParser is Parser constructor
func NewParser() Parser {
	return Parser{}
}

// Processing make a node-trees
func (p *Parser) Processing(str string) ([]model.Node, error) {
	// TODO:
	return nil, &ParserError{
		desc: "unknown error",
	}
}
