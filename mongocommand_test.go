package examples

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/kmdrn7/mongo-parser/generated"
)

// ErrorListener to catch syntax errors
type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors int
}

// SyntaxError
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors++
	fmt.Printf("Error at line %d:%d - %s\n", line, column, msg)
}

func TestParseMongoQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
	}{
		{
			name: "parse deleteMany query without filters",
			query: `
				db.users.deleteMany({})
			`,
		},
		{
			name: "parse deleteMany query with filters",
			query: `
				db.users.deleteMany({exampleArg:true})
			`,
		},
		{
			name: "parse find multiline query with filters",
			query: `
				db.users.find({
					name: "John",
					age: { $gt: 21 }
				})
			`,
		},
		{
			name: "parse find multiline query with filters",
			query: `
				db.users.find({
					name: "John",
					age: { $gt: 21 }
				});
				db.users.find({ name: "Doe", age: { $gt: 21 } })
			`,
		},
		{
			name: "parse updateMany query with filters using getCollection",
			query: `
				db.getCollection("users").deleteMany({exampleArg:true});
				db.getCollection("users").deleteMany({exampleArg:false});
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create error listener
			errorListener := &ErrorListener{}
			// create the lexer and parser
			input := antlr.NewInputStream(tt.query)
			lexer := parser.NewMongoCommandLexer(input)
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(errorListener)
			// create the token stream
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewMongoCommandParser(stream)
			p.RemoveErrorListeners()
			p.AddErrorListener(errorListener)
			// parse the input
			tree := p.MongoCommand()
			if errorListener.errors > 0 {
				fmt.Printf("Found %d errors\n", errorListener.errors)
			} else {
				fmt.Println("Successfully parsed!")
				fmt.Printf("Parse tree: %v\n", tree.ToStringTree(p.GetRuleNames(), p))
			}
		})
	}
}
