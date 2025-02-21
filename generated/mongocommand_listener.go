// Code generated from MongoCommand.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MongoCommand

import "github.com/antlr4-go/antlr/v4"

// MongoCommandListener is a complete listener for a parse tree produced by MongoCommandParser.
type MongoCommandListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterOperatorKey is called when entering the operatorKey production.
	EnterOperatorKey(c *OperatorKeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterSimpleValue is called when entering the simpleValue production.
	EnterSimpleValue(c *SimpleValueContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitOperatorKey is called when exiting the operatorKey production.
	ExitOperatorKey(c *OperatorKeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitSimpleValue is called when exiting the simpleValue production.
	ExitSimpleValue(c *SimpleValueContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
