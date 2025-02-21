// Code generated from MongoCommand.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MongoCommand

import "github.com/antlr4-go/antlr/v4"

// BaseMongoCommandListener is a complete listener for a parse tree produced by MongoCommandParser.
type BaseMongoCommandListener struct{}

var _ MongoCommandListener = &BaseMongoCommandListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMongoCommandListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMongoCommandListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMongoCommandListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMongoCommandListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMongoCommand is called when production mongoCommand is entered.
func (s *BaseMongoCommandListener) EnterMongoCommand(ctx *MongoCommandContext) {}

// ExitMongoCommand is called when production mongoCommand is exited.
func (s *BaseMongoCommandListener) ExitMongoCommand(ctx *MongoCommandContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseMongoCommandListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseMongoCommandListener) ExitCommand(ctx *CommandContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseMongoCommandListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseMongoCommandListener) ExitCollection(ctx *CollectionContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseMongoCommandListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseMongoCommandListener) ExitOperation(ctx *OperationContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMongoCommandListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMongoCommandListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseMongoCommandListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseMongoCommandListener) ExitArgument(ctx *ArgumentContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseMongoCommandListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseMongoCommandListener) ExitDocument(ctx *DocumentContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMongoCommandListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMongoCommandListener) ExitPair(ctx *PairContext) {}

// EnterOperatorKey is called when production operatorKey is entered.
func (s *BaseMongoCommandListener) EnterOperatorKey(ctx *OperatorKeyContext) {}

// ExitOperatorKey is called when production operatorKey is exited.
func (s *BaseMongoCommandListener) ExitOperatorKey(ctx *OperatorKeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMongoCommandListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMongoCommandListener) ExitValue(ctx *ValueContext) {}

// EnterSimpleValue is called when production simpleValue is entered.
func (s *BaseMongoCommandListener) EnterSimpleValue(ctx *SimpleValueContext) {}

// ExitSimpleValue is called when production simpleValue is exited.
func (s *BaseMongoCommandListener) ExitSimpleValue(ctx *SimpleValueContext) {}

// EnterArray is called when production array is entered.
func (s *BaseMongoCommandListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseMongoCommandListener) ExitArray(ctx *ArrayContext) {}
