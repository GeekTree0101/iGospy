package worker

import (
	"fmt"
	"strings"

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
func (p *Parser) Processing(str string) (model.Node, error) {

	var syntex string = ""
	var syntexQueue []string
	var root model.Node
	var children []model.Node
	var currentDepth int = 0

	// FIXME: processing needs depth inference logic

	for _, r := range strings.TrimSpace(str) {

		char := string(r)

		switch char {
		case "{":
			trimmedSyntex := strings.TrimSpace(syntex)
			syntexQueue = append(syntexQueue, trimmedSyntex)
			syntex = ""
			currentDepth++

		case "}":
			if len(syntexQueue) == 0 {
				return root, &ParserError{
					desc: "queue is empty",
				}
			}

			lastSyntex := syntexQueue[len(syntexQueue)-1]
			splitSyntex := strings.Split(lastSyntex, " ")

			if len(splitSyntex) != 2 {
				return root, &ParserError{
					desc: "invalid syntex structure",
				}
			}

			name := splitSyntex[len(splitSyntex)-1]

			switch currentDepth {
			case 1:
				root = model.Node{
					Type:     model.Title,
					Name:     name,
					Children: children,
				}

			case 2:
				var bros []model.Node
				var child []model.Node

				for _, c := range children {
					if c.Type == model.Behavior {
						child = append(child, c)
					} else {
						bros = append(bros, c)
					}
				}

				children = append(bros, model.Node{
					Type:     model.Context,
					Name:     name,
					Children: child,
				})

			case 3:
				children = append(children, model.Node{
					Type:     model.Behavior,
					Name:     name,
					Children: nil,
				})

			default:
				break
			}

			syntexQueue = syntexQueue[:len(syntexQueue)-1]
			currentDepth--

		default:
			if currentDepth < 3 && char != "\n" {
				syntex += char
			}
		}
	}

	return root, nil
}
