package worker

import (
	"fmt"
	"strings"

	"github.com/GeekTree0101/iGospy/model"
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

// MakeUsecase build usecase with rawString
func (p *Parser) MakeUsecase(str string) (model.Usecase, error) {

	var usecase model.Usecase
	var title string = ""
	var contexts []string
	var syntex string = ""
	var syntexQueue []string
	var currentDepth int = 0

	for _, r := range str {

		char := string(r)

		switch char {
		case "{":
			trimmedSyntex := strings.TrimSpace(syntex)
			syntexQueue = append(syntexQueue, trimmedSyntex)
			syntex = ""
			currentDepth++

		case "}":
			if len(syntexQueue) == 0 {
				return usecase, &ParserError{
					desc: "queue is empty",
				}
			}

			lastSyntex := syntexQueue[len(syntexQueue)-1]
			splitSyntex := strings.Split(lastSyntex, " ")

			if len(splitSyntex) == 2 {

				name := splitSyntex[len(splitSyntex)-1]

				switch currentDepth {
				case 1:
					title = name

				case 2:
					contexts = append(contexts, name)

				default:
					break
				}
			}

			syntexQueue = syntexQueue[:len(syntexQueue)-1]
			currentDepth--

		default:
			switch char {
			case "\n":
				break
			case "\t":
				break
			default:
				if currentDepth < 2 {
					syntex += char
				}
			}
		}
	}

	if len(title) != 0 && len(contexts) > 0 {
		usecase = model.Usecase{
			Title:    title,
			Contexts: contexts,
		}

		return usecase, nil
	}

	return usecase, &ParserError{
		desc: "fail to generate usecase " + title + " " + strings.Join(contexts, ","),
	}
}
