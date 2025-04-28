// Code generated from MongoCommand.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MongoCommand

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MongoCommandParser struct {
	*antlr.BaseParser
}

var MongoCommandParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mongocommandParserInit() {
	staticData := &MongoCommandParserStaticData
	staticData.LiteralNames = []string{
		"", "'findOne'", "'insertOne'", "'insertMany'", "'updateOne'", "'updateMany'",
		"'deleteOne'", "'deleteMany'", "'replaceOne'", "'replaceAll'", "'aggregate'",
		"'count'", "'countDocuments'", "'estimatedDocumentCount'", "'distinct'",
		"'createIndex'", "'createIndexes'", "'dropIndex'", "'dropIndexes'",
		"'listIndexes'", "'reIndex'", "'drop'", "'rename'", "'stats'", "'bulkWrite'",
		"'findOneAndUpdate'", "'findOneAndDelete'", "'findOneAndReplace'", "'isCapped'",
		"'explain'", "'watch'", "'db'", "'getCollection'", "'.'", "','", "':'",
		"';'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'$set'", "'$unset'",
		"'$inc'", "'$push'", "'$pull'", "'$addToSet'", "'$pop'", "'$pullAll'",
		"'$each'", "'$position'", "'$slice'", "'$sort'", "'$mul'", "'$min'",
		"'$max'", "'$rename'", "'$currentDate'", "'$bit'", "'$replaceAll'",
		"'find'", "'$in'", "'$nin'", "'$gt'", "'$gte'", "'$lt'", "'$lte'", "'$eq'",
		"'$ne'", "'$exists'", "'$type'", "'$or'", "'$and'", "'$not'", "'$nor'",
		"'$mod'", "'$regex'", "'$trim'", "'$text'", "'$where'", "'$all'", "'$elemMatch'",
		"'$size'", "'$match'", "'$group'", "'$project'", "'$limit'", "'$skip'",
		"'$unwind'", "'$lookup'", "'$sample'", "'$count'", "'$facet'", "'$bucket'",
		"'$bucketAuto'", "'$sum'", "'$avg'", "'$first'", "'$last'", "'$stdDevPop'",
		"'$stdDevSamp'", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "DB", "GET_COLLECTION",
		"DOT", "COMMA", "COLON", "SEMICOLON", "LPAREN", "RPAREN", "LBRACK",
		"RBRACK", "LCURLY", "RCURLY", "DOLLAR_SET", "DOLLAR_UNSET", "DOLLAR_INC",
		"DOLLAR_PUSH", "DOLLAR_PULL", "DOLLAR_ADD_TO_SET", "DOLLAR_POP", "DOLLAR_PULL_ALL",
		"DOLLAR_EACH", "DOLLAR_POSITION", "DOLLAR_SLICE", "DOLLAR_SORT", "DOLLAR_MUL",
		"DOLLAR_MIN", "DOLLAR_MAX", "DOLLAR_RENAME", "DOLLAR_CURRENT_DATE",
		"DOLLAR_BIT", "DOLLAR_REPLACE_ALL", "DOLLAR_FIND", "DOLLAR_IN", "DOLLAR_NIN",
		"DOLLAR_GT", "DOLLAR_GTE", "DOLLAR_LT", "DOLLAR_LTE", "DOLLAR_EQ", "DOLLAR_NE",
		"DOLLAR_EXISTS", "DOLLAR_TYPE", "DOLLAR_OR", "DOLLAR_AND", "DOLLAR_NOT",
		"DOLLAR_NOR", "DOLLAR_MOD", "DOLLAR_REGEX", "DOLLAR_TRIM", "DOLLAR_TEXT",
		"DOLLAR_WHERE", "DOLLAR_ALL", "DOLLAR_ELEM_MATCH", "DOLLAR_SIZE", "DOLLAR_MATCH",
		"DOLLAR_GROUP", "DOLLAR_PROJECT", "DOLLAR_LIMIT", "DOLLAR_SKIP", "DOLLAR_UNWIND",
		"DOLLAR_LOOKUP", "DOLLAR_SAMPLE", "DOLLAR_COUNT", "DOLLAR_FACET", "DOLLAR_BUCKET",
		"DOLLAR_BUCKET_AUTO", "DOLLAR_SUM", "DOLLAR_AVG", "DOLLAR_FIRST", "DOLLAR_LAST",
		"DOLLAR_STDDEV_POP", "DOLLAR_STDDEV_SAMP", "BOOLEAN", "NULL", "IDENTIFIER",
		"STRING", "NUMBER", "WS", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"mongoCommand", "command", "collection", "operation", "arguments", "argument",
		"document", "pair", "operatorKey", "value", "simpleValue", "array",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 110, 123, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		45, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 55, 8,
		4, 10, 4, 12, 4, 58, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 64, 8, 4, 1, 5,
		1, 5, 1, 5, 3, 5, 69, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 75, 8, 6, 10,
		6, 12, 6, 78, 9, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 94, 8, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 3, 9, 101, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 108,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 114, 8, 11, 10, 11, 12, 11, 117,
		9, 11, 3, 11, 119, 8, 11, 1, 11, 1, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 0, 3, 1, 1, 36, 36, 2, 0, 1, 30, 62, 62, 1,
		0, 43, 102, 127, 0, 24, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0,
		0, 6, 48, 1, 0, 0, 0, 8, 63, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 70, 1,
		0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 95, 1, 0, 0, 0, 18, 100, 1, 0, 0, 0, 20,
		107, 1, 0, 0, 0, 22, 109, 1, 0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 7, 0,
		0, 0, 26, 1, 1, 0, 0, 0, 27, 28, 5, 31, 0, 0, 28, 29, 5, 33, 0, 0, 29,
		30, 3, 4, 2, 0, 30, 31, 5, 33, 0, 0, 31, 32, 3, 6, 3, 0, 32, 33, 3, 8,
		4, 0, 33, 45, 1, 0, 0, 0, 34, 35, 5, 31, 0, 0, 35, 36, 5, 33, 0, 0, 36,
		37, 5, 32, 0, 0, 37, 38, 5, 37, 0, 0, 38, 39, 5, 106, 0, 0, 39, 40, 5,
		38, 0, 0, 40, 41, 5, 33, 0, 0, 41, 42, 3, 6, 3, 0, 42, 43, 3, 8, 4, 0,
		43, 45, 1, 0, 0, 0, 44, 27, 1, 0, 0, 0, 44, 34, 1, 0, 0, 0, 45, 3, 1, 0,
		0, 0, 46, 47, 5, 105, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 7, 1, 0, 0, 49,
		7, 1, 0, 0, 0, 50, 51, 5, 37, 0, 0, 51, 56, 3, 10, 5, 0, 52, 53, 5, 34,
		0, 0, 53, 55, 3, 10, 5, 0, 54, 52, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56,
		54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0,
		0, 59, 60, 5, 38, 0, 0, 60, 64, 1, 0, 0, 0, 61, 62, 5, 37, 0, 0, 62, 64,
		5, 38, 0, 0, 63, 50, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 9, 1, 0, 0, 0,
		65, 69, 3, 12, 6, 0, 66, 69, 3, 22, 11, 0, 67, 69, 5, 105, 0, 0, 68, 65,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 11, 1, 0, 0, 0,
		70, 79, 5, 41, 0, 0, 71, 76, 3, 14, 7, 0, 72, 73, 5, 34, 0, 0, 73, 75,
		3, 14, 7, 0, 74, 72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		76, 77, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 71, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 5, 42, 0, 0, 82,
		13, 1, 0, 0, 0, 83, 84, 5, 106, 0, 0, 84, 85, 5, 35, 0, 0, 85, 94, 3, 18,
		9, 0, 86, 87, 5, 105, 0, 0, 87, 88, 5, 35, 0, 0, 88, 94, 3, 18, 9, 0, 89,
		90, 3, 16, 8, 0, 90, 91, 5, 35, 0, 0, 91, 92, 3, 18, 9, 0, 92, 94, 1, 0,
		0, 0, 93, 83, 1, 0, 0, 0, 93, 86, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 94, 15,
		1, 0, 0, 0, 95, 96, 7, 2, 0, 0, 96, 17, 1, 0, 0, 0, 97, 101, 3, 20, 10,
		0, 98, 101, 3, 12, 6, 0, 99, 101, 3, 22, 11, 0, 100, 97, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 19, 1, 0, 0, 0, 102, 108, 5,
		106, 0, 0, 103, 108, 5, 107, 0, 0, 104, 108, 5, 103, 0, 0, 105, 108, 5,
		104, 0, 0, 106, 108, 3, 16, 8, 0, 107, 102, 1, 0, 0, 0, 107, 103, 1, 0,
		0, 0, 107, 104, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0,
		108, 21, 1, 0, 0, 0, 109, 118, 5, 39, 0, 0, 110, 115, 3, 18, 9, 0, 111,
		112, 5, 34, 0, 0, 112, 114, 3, 18, 9, 0, 113, 111, 1, 0, 0, 0, 114, 117,
		1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 119, 1, 0,
		0, 0, 117, 115, 1, 0, 0, 0, 118, 110, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 121, 5, 40, 0, 0, 121, 23, 1, 0, 0, 0, 11, 44,
		56, 63, 68, 76, 79, 93, 100, 107, 115, 118,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MongoCommandParserInit initializes any static state used to implement MongoCommandParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMongoCommandParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MongoCommandParserInit() {
	staticData := &MongoCommandParserStaticData
	staticData.once.Do(mongocommandParserInit)
}

// NewMongoCommandParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMongoCommandParser(input antlr.TokenStream) *MongoCommandParser {
	MongoCommandParserInit()
	this := new(MongoCommandParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MongoCommandParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MongoCommand.g4"

	return this
}

// MongoCommandParser tokens.
const (
	MongoCommandParserEOF                 = antlr.TokenEOF
	MongoCommandParserT__0                = 1
	MongoCommandParserT__1                = 2
	MongoCommandParserT__2                = 3
	MongoCommandParserT__3                = 4
	MongoCommandParserT__4                = 5
	MongoCommandParserT__5                = 6
	MongoCommandParserT__6                = 7
	MongoCommandParserT__7                = 8
	MongoCommandParserT__8                = 9
	MongoCommandParserT__9                = 10
	MongoCommandParserT__10               = 11
	MongoCommandParserT__11               = 12
	MongoCommandParserT__12               = 13
	MongoCommandParserT__13               = 14
	MongoCommandParserT__14               = 15
	MongoCommandParserT__15               = 16
	MongoCommandParserT__16               = 17
	MongoCommandParserT__17               = 18
	MongoCommandParserT__18               = 19
	MongoCommandParserT__19               = 20
	MongoCommandParserT__20               = 21
	MongoCommandParserT__21               = 22
	MongoCommandParserT__22               = 23
	MongoCommandParserT__23               = 24
	MongoCommandParserT__24               = 25
	MongoCommandParserT__25               = 26
	MongoCommandParserT__26               = 27
	MongoCommandParserT__27               = 28
	MongoCommandParserT__28               = 29
	MongoCommandParserT__29               = 30
	MongoCommandParserDB                  = 31
	MongoCommandParserGET_COLLECTION      = 32
	MongoCommandParserDOT                 = 33
	MongoCommandParserCOMMA               = 34
	MongoCommandParserCOLON               = 35
	MongoCommandParserSEMICOLON           = 36
	MongoCommandParserLPAREN              = 37
	MongoCommandParserRPAREN              = 38
	MongoCommandParserLBRACK              = 39
	MongoCommandParserRBRACK              = 40
	MongoCommandParserLCURLY              = 41
	MongoCommandParserRCURLY              = 42
	MongoCommandParserDOLLAR_SET          = 43
	MongoCommandParserDOLLAR_UNSET        = 44
	MongoCommandParserDOLLAR_INC          = 45
	MongoCommandParserDOLLAR_PUSH         = 46
	MongoCommandParserDOLLAR_PULL         = 47
	MongoCommandParserDOLLAR_ADD_TO_SET   = 48
	MongoCommandParserDOLLAR_POP          = 49
	MongoCommandParserDOLLAR_PULL_ALL     = 50
	MongoCommandParserDOLLAR_EACH         = 51
	MongoCommandParserDOLLAR_POSITION     = 52
	MongoCommandParserDOLLAR_SLICE        = 53
	MongoCommandParserDOLLAR_SORT         = 54
	MongoCommandParserDOLLAR_MUL          = 55
	MongoCommandParserDOLLAR_MIN          = 56
	MongoCommandParserDOLLAR_MAX          = 57
	MongoCommandParserDOLLAR_RENAME       = 58
	MongoCommandParserDOLLAR_CURRENT_DATE = 59
	MongoCommandParserDOLLAR_BIT          = 60
	MongoCommandParserDOLLAR_REPLACE_ALL  = 61
	MongoCommandParserDOLLAR_FIND         = 62
	MongoCommandParserDOLLAR_IN           = 63
	MongoCommandParserDOLLAR_NIN          = 64
	MongoCommandParserDOLLAR_GT           = 65
	MongoCommandParserDOLLAR_GTE          = 66
	MongoCommandParserDOLLAR_LT           = 67
	MongoCommandParserDOLLAR_LTE          = 68
	MongoCommandParserDOLLAR_EQ           = 69
	MongoCommandParserDOLLAR_NE           = 70
	MongoCommandParserDOLLAR_EXISTS       = 71
	MongoCommandParserDOLLAR_TYPE         = 72
	MongoCommandParserDOLLAR_OR           = 73
	MongoCommandParserDOLLAR_AND          = 74
	MongoCommandParserDOLLAR_NOT          = 75
	MongoCommandParserDOLLAR_NOR          = 76
	MongoCommandParserDOLLAR_MOD          = 77
	MongoCommandParserDOLLAR_REGEX        = 78
	MongoCommandParserDOLLAR_TRIM         = 79
	MongoCommandParserDOLLAR_TEXT         = 80
	MongoCommandParserDOLLAR_WHERE        = 81
	MongoCommandParserDOLLAR_ALL          = 82
	MongoCommandParserDOLLAR_ELEM_MATCH   = 83
	MongoCommandParserDOLLAR_SIZE         = 84
	MongoCommandParserDOLLAR_MATCH        = 85
	MongoCommandParserDOLLAR_GROUP        = 86
	MongoCommandParserDOLLAR_PROJECT      = 87
	MongoCommandParserDOLLAR_LIMIT        = 88
	MongoCommandParserDOLLAR_SKIP         = 89
	MongoCommandParserDOLLAR_UNWIND       = 90
	MongoCommandParserDOLLAR_LOOKUP       = 91
	MongoCommandParserDOLLAR_SAMPLE       = 92
	MongoCommandParserDOLLAR_COUNT        = 93
	MongoCommandParserDOLLAR_FACET        = 94
	MongoCommandParserDOLLAR_BUCKET       = 95
	MongoCommandParserDOLLAR_BUCKET_AUTO  = 96
	MongoCommandParserDOLLAR_SUM          = 97
	MongoCommandParserDOLLAR_AVG          = 98
	MongoCommandParserDOLLAR_FIRST        = 99
	MongoCommandParserDOLLAR_LAST         = 100
	MongoCommandParserDOLLAR_STDDEV_POP   = 101
	MongoCommandParserDOLLAR_STDDEV_SAMP  = 102
	MongoCommandParserBOOLEAN             = 103
	MongoCommandParserNULL                = 104
	MongoCommandParserIDENTIFIER          = 105
	MongoCommandParserSTRING              = 106
	MongoCommandParserNUMBER              = 107
	MongoCommandParserWS                  = 108
	MongoCommandParserSINGLE_LINE_COMMENT = 109
	MongoCommandParserMULTI_LINE_COMMENT  = 110
)

// MongoCommandParser rules.
const (
	MongoCommandParserRULE_mongoCommand = 0
	MongoCommandParserRULE_command      = 1
	MongoCommandParserRULE_collection   = 2
	MongoCommandParserRULE_operation    = 3
	MongoCommandParserRULE_arguments    = 4
	MongoCommandParserRULE_argument     = 5
	MongoCommandParserRULE_document     = 6
	MongoCommandParserRULE_pair         = 7
	MongoCommandParserRULE_operatorKey  = 8
	MongoCommandParserRULE_value        = 9
	MongoCommandParserRULE_simpleValue  = 10
	MongoCommandParserRULE_array        = 11
)

// IMongoCommandContext is an interface to support dynamic dispatch.
type IMongoCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Command() ICommandContext
	SEMICOLON() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsMongoCommandContext differentiates from other interfaces.
	IsMongoCommandContext()
}

type MongoCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMongoCommandContext() *MongoCommandContext {
	var p = new(MongoCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_mongoCommand
	return p
}

func InitEmptyMongoCommandContext(p *MongoCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_mongoCommand
}

func (*MongoCommandContext) IsMongoCommandContext() {}

func NewMongoCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MongoCommandContext {
	var p = new(MongoCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_mongoCommand

	return p
}

func (s *MongoCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *MongoCommandContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *MongoCommandContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserSEMICOLON, 0)
}

func (s *MongoCommandContext) EOF() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserEOF, 0)
}

func (s *MongoCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MongoCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MongoCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterMongoCommand(s)
	}
}

func (s *MongoCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitMongoCommand(s)
	}
}

func (p *MongoCommandParser) MongoCommand() (localctx IMongoCommandContext) {
	localctx = NewMongoCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MongoCommandParserRULE_mongoCommand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Command()
	}
	{
		p.SetState(25)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoCommandParserEOF || _la == MongoCommandParserSEMICOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DB() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	Collection() ICollectionContext
	Operation() IOperationContext
	Arguments() IArgumentsContext
	GET_COLLECTION() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) DB() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDB, 0)
}

func (s *CommandContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MongoCommandParserDOT)
}

func (s *CommandContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOT, i)
}

func (s *CommandContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *CommandContext) Operation() IOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *CommandContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CommandContext) GET_COLLECTION() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserGET_COLLECTION, 0)
}

func (s *CommandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserLPAREN, 0)
}

func (s *CommandContext) STRING() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserSTRING, 0)
}

func (s *CommandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserRPAREN, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *MongoCommandParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MongoCommandParserRULE_command)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(MongoCommandParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.Match(MongoCommandParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Collection()
		}
		{
			p.SetState(30)
			p.Match(MongoCommandParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Operation()
		}
		{
			p.SetState(32)
			p.Arguments()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(MongoCommandParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(MongoCommandParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(MongoCommandParserGET_COLLECTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(MongoCommandParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(MongoCommandParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(MongoCommandParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(MongoCommandParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Operation()
		}
		{
			p.SetState(42)
			p.Arguments()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserIDENTIFIER, 0)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *MongoCommandParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MongoCommandParserRULE_collection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(MongoCommandParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOLLAR_FIND() antlr.TerminalNode

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_operation
	return p
}

func InitEmptyOperationContext(p *OperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_operation
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) DOLLAR_FIND() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_FIND, 0)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *MongoCommandParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MongoCommandParserRULE_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686020574871550) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserLPAREN, 0)
}

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserRPAREN, 0)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoCommandParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoCommandParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *MongoCommandParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MongoCommandParserRULE_arguments)
	var _la int

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(MongoCommandParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Argument()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MongoCommandParserCOMMA {
			{
				p.SetState(52)
				p.Match(MongoCommandParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(53)
				p.Argument()
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(59)
			p.Match(MongoCommandParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(MongoCommandParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(MongoCommandParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Document() IDocumentContext
	Array() IArrayContext
	IDENTIFIER() antlr.TerminalNode

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *ArgumentContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserIDENTIFIER, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *MongoCommandParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MongoCommandParserRULE_argument)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoCommandParserLCURLY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Document()
		}

	case MongoCommandParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Array()
		}

	case MongoCommandParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(MongoCommandParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllPair() []IPairContext
	Pair(i int) IPairContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserLCURLY, 0)
}

func (s *DocumentContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserRCURLY, 0)
}

func (s *DocumentContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *DocumentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoCommandParserCOMMA)
}

func (s *DocumentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoCommandParserCOMMA, i)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *MongoCommandParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MongoCommandParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(MongoCommandParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-43)) & ^0x3f) == 0 && ((int64(1)<<(_la-43))&-3458764513820540929) != 0 {
		{
			p.SetState(71)
			p.Pair()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MongoCommandParserCOMMA {
			{
				p.SetState(72)
				p.Match(MongoCommandParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.Pair()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(81)
		p.Match(MongoCommandParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// Getter signatures
	COLON() antlr.TerminalNode
	Value() IValueContext
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	OperatorKey() IOperatorKeyContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetKey() antlr.Token { return s.key }

func (s *PairContext) SetKey(v antlr.Token) { s.key = v }

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserCOLON, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserSTRING, 0)
}

func (s *PairContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserIDENTIFIER, 0)
}

func (s *PairContext) OperatorKey() IOperatorKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorKeyContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *MongoCommandParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MongoCommandParserRULE_pair)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoCommandParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)

			var _m = p.Match(MongoCommandParserSTRING)

			localctx.(*PairContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(MongoCommandParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Value()
		}

	case MongoCommandParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)

			var _m = p.Match(MongoCommandParserIDENTIFIER)

			localctx.(*PairContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(MongoCommandParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Value()
		}

	case MongoCommandParserDOLLAR_SET, MongoCommandParserDOLLAR_UNSET, MongoCommandParserDOLLAR_INC, MongoCommandParserDOLLAR_PUSH, MongoCommandParserDOLLAR_PULL, MongoCommandParserDOLLAR_ADD_TO_SET, MongoCommandParserDOLLAR_POP, MongoCommandParserDOLLAR_PULL_ALL, MongoCommandParserDOLLAR_EACH, MongoCommandParserDOLLAR_POSITION, MongoCommandParserDOLLAR_SLICE, MongoCommandParserDOLLAR_SORT, MongoCommandParserDOLLAR_MUL, MongoCommandParserDOLLAR_MIN, MongoCommandParserDOLLAR_MAX, MongoCommandParserDOLLAR_RENAME, MongoCommandParserDOLLAR_CURRENT_DATE, MongoCommandParserDOLLAR_BIT, MongoCommandParserDOLLAR_REPLACE_ALL, MongoCommandParserDOLLAR_FIND, MongoCommandParserDOLLAR_IN, MongoCommandParserDOLLAR_NIN, MongoCommandParserDOLLAR_GT, MongoCommandParserDOLLAR_GTE, MongoCommandParserDOLLAR_LT, MongoCommandParserDOLLAR_LTE, MongoCommandParserDOLLAR_EQ, MongoCommandParserDOLLAR_NE, MongoCommandParserDOLLAR_EXISTS, MongoCommandParserDOLLAR_TYPE, MongoCommandParserDOLLAR_OR, MongoCommandParserDOLLAR_AND, MongoCommandParserDOLLAR_NOT, MongoCommandParserDOLLAR_NOR, MongoCommandParserDOLLAR_MOD, MongoCommandParserDOLLAR_REGEX, MongoCommandParserDOLLAR_TRIM, MongoCommandParserDOLLAR_TEXT, MongoCommandParserDOLLAR_WHERE, MongoCommandParserDOLLAR_ALL, MongoCommandParserDOLLAR_ELEM_MATCH, MongoCommandParserDOLLAR_SIZE, MongoCommandParserDOLLAR_MATCH, MongoCommandParserDOLLAR_GROUP, MongoCommandParserDOLLAR_PROJECT, MongoCommandParserDOLLAR_LIMIT, MongoCommandParserDOLLAR_SKIP, MongoCommandParserDOLLAR_UNWIND, MongoCommandParserDOLLAR_LOOKUP, MongoCommandParserDOLLAR_SAMPLE, MongoCommandParserDOLLAR_COUNT, MongoCommandParserDOLLAR_FACET, MongoCommandParserDOLLAR_BUCKET, MongoCommandParserDOLLAR_BUCKET_AUTO, MongoCommandParserDOLLAR_SUM, MongoCommandParserDOLLAR_AVG, MongoCommandParserDOLLAR_FIRST, MongoCommandParserDOLLAR_LAST, MongoCommandParserDOLLAR_STDDEV_POP, MongoCommandParserDOLLAR_STDDEV_SAMP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.OperatorKey()
		}
		{
			p.SetState(90)
			p.Match(MongoCommandParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorKeyContext is an interface to support dynamic dispatch.
type IOperatorKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOLLAR_SET() antlr.TerminalNode
	DOLLAR_UNSET() antlr.TerminalNode
	DOLLAR_INC() antlr.TerminalNode
	DOLLAR_PUSH() antlr.TerminalNode
	DOLLAR_PULL() antlr.TerminalNode
	DOLLAR_ADD_TO_SET() antlr.TerminalNode
	DOLLAR_POP() antlr.TerminalNode
	DOLLAR_PULL_ALL() antlr.TerminalNode
	DOLLAR_EACH() antlr.TerminalNode
	DOLLAR_POSITION() antlr.TerminalNode
	DOLLAR_SLICE() antlr.TerminalNode
	DOLLAR_SORT() antlr.TerminalNode
	DOLLAR_MUL() antlr.TerminalNode
	DOLLAR_MIN() antlr.TerminalNode
	DOLLAR_MAX() antlr.TerminalNode
	DOLLAR_RENAME() antlr.TerminalNode
	DOLLAR_CURRENT_DATE() antlr.TerminalNode
	DOLLAR_BIT() antlr.TerminalNode
	DOLLAR_REPLACE_ALL() antlr.TerminalNode
	DOLLAR_FIND() antlr.TerminalNode
	DOLLAR_IN() antlr.TerminalNode
	DOLLAR_NIN() antlr.TerminalNode
	DOLLAR_GT() antlr.TerminalNode
	DOLLAR_GTE() antlr.TerminalNode
	DOLLAR_LT() antlr.TerminalNode
	DOLLAR_LTE() antlr.TerminalNode
	DOLLAR_EQ() antlr.TerminalNode
	DOLLAR_NE() antlr.TerminalNode
	DOLLAR_EXISTS() antlr.TerminalNode
	DOLLAR_TYPE() antlr.TerminalNode
	DOLLAR_OR() antlr.TerminalNode
	DOLLAR_AND() antlr.TerminalNode
	DOLLAR_NOT() antlr.TerminalNode
	DOLLAR_NOR() antlr.TerminalNode
	DOLLAR_MOD() antlr.TerminalNode
	DOLLAR_TRIM() antlr.TerminalNode
	DOLLAR_REGEX() antlr.TerminalNode
	DOLLAR_TEXT() antlr.TerminalNode
	DOLLAR_WHERE() antlr.TerminalNode
	DOLLAR_ALL() antlr.TerminalNode
	DOLLAR_ELEM_MATCH() antlr.TerminalNode
	DOLLAR_SIZE() antlr.TerminalNode
	DOLLAR_MATCH() antlr.TerminalNode
	DOLLAR_GROUP() antlr.TerminalNode
	DOLLAR_PROJECT() antlr.TerminalNode
	DOLLAR_LIMIT() antlr.TerminalNode
	DOLLAR_SKIP() antlr.TerminalNode
	DOLLAR_UNWIND() antlr.TerminalNode
	DOLLAR_LOOKUP() antlr.TerminalNode
	DOLLAR_SAMPLE() antlr.TerminalNode
	DOLLAR_COUNT() antlr.TerminalNode
	DOLLAR_FACET() antlr.TerminalNode
	DOLLAR_BUCKET() antlr.TerminalNode
	DOLLAR_BUCKET_AUTO() antlr.TerminalNode
	DOLLAR_SUM() antlr.TerminalNode
	DOLLAR_AVG() antlr.TerminalNode
	DOLLAR_FIRST() antlr.TerminalNode
	DOLLAR_LAST() antlr.TerminalNode
	DOLLAR_STDDEV_POP() antlr.TerminalNode
	DOLLAR_STDDEV_SAMP() antlr.TerminalNode

	// IsOperatorKeyContext differentiates from other interfaces.
	IsOperatorKeyContext()
}

type OperatorKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorKeyContext() *OperatorKeyContext {
	var p = new(OperatorKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_operatorKey
	return p
}

func InitEmptyOperatorKeyContext(p *OperatorKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_operatorKey
}

func (*OperatorKeyContext) IsOperatorKeyContext() {}

func NewOperatorKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorKeyContext {
	var p = new(OperatorKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_operatorKey

	return p
}

func (s *OperatorKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorKeyContext) DOLLAR_SET() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SET, 0)
}

func (s *OperatorKeyContext) DOLLAR_UNSET() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_UNSET, 0)
}

func (s *OperatorKeyContext) DOLLAR_INC() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_INC, 0)
}

func (s *OperatorKeyContext) DOLLAR_PUSH() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_PUSH, 0)
}

func (s *OperatorKeyContext) DOLLAR_PULL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_PULL, 0)
}

func (s *OperatorKeyContext) DOLLAR_ADD_TO_SET() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_ADD_TO_SET, 0)
}

func (s *OperatorKeyContext) DOLLAR_POP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_POP, 0)
}

func (s *OperatorKeyContext) DOLLAR_PULL_ALL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_PULL_ALL, 0)
}

func (s *OperatorKeyContext) DOLLAR_EACH() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_EACH, 0)
}

func (s *OperatorKeyContext) DOLLAR_POSITION() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_POSITION, 0)
}

func (s *OperatorKeyContext) DOLLAR_SLICE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SLICE, 0)
}

func (s *OperatorKeyContext) DOLLAR_SORT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SORT, 0)
}

func (s *OperatorKeyContext) DOLLAR_MUL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_MUL, 0)
}

func (s *OperatorKeyContext) DOLLAR_MIN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_MIN, 0)
}

func (s *OperatorKeyContext) DOLLAR_MAX() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_MAX, 0)
}

func (s *OperatorKeyContext) DOLLAR_RENAME() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_RENAME, 0)
}

func (s *OperatorKeyContext) DOLLAR_CURRENT_DATE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_CURRENT_DATE, 0)
}

func (s *OperatorKeyContext) DOLLAR_BIT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_BIT, 0)
}

func (s *OperatorKeyContext) DOLLAR_REPLACE_ALL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_REPLACE_ALL, 0)
}

func (s *OperatorKeyContext) DOLLAR_FIND() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_FIND, 0)
}

func (s *OperatorKeyContext) DOLLAR_IN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_IN, 0)
}

func (s *OperatorKeyContext) DOLLAR_NIN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_NIN, 0)
}

func (s *OperatorKeyContext) DOLLAR_GT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_GT, 0)
}

func (s *OperatorKeyContext) DOLLAR_GTE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_GTE, 0)
}

func (s *OperatorKeyContext) DOLLAR_LT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_LT, 0)
}

func (s *OperatorKeyContext) DOLLAR_LTE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_LTE, 0)
}

func (s *OperatorKeyContext) DOLLAR_EQ() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_EQ, 0)
}

func (s *OperatorKeyContext) DOLLAR_NE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_NE, 0)
}

func (s *OperatorKeyContext) DOLLAR_EXISTS() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_EXISTS, 0)
}

func (s *OperatorKeyContext) DOLLAR_TYPE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_TYPE, 0)
}

func (s *OperatorKeyContext) DOLLAR_OR() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_OR, 0)
}

func (s *OperatorKeyContext) DOLLAR_AND() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_AND, 0)
}

func (s *OperatorKeyContext) DOLLAR_NOT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_NOT, 0)
}

func (s *OperatorKeyContext) DOLLAR_NOR() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_NOR, 0)
}

func (s *OperatorKeyContext) DOLLAR_MOD() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_MOD, 0)
}

func (s *OperatorKeyContext) DOLLAR_TRIM() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_TRIM, 0)
}

func (s *OperatorKeyContext) DOLLAR_REGEX() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_REGEX, 0)
}

func (s *OperatorKeyContext) DOLLAR_TEXT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_TEXT, 0)
}

func (s *OperatorKeyContext) DOLLAR_WHERE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_WHERE, 0)
}

func (s *OperatorKeyContext) DOLLAR_ALL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_ALL, 0)
}

func (s *OperatorKeyContext) DOLLAR_ELEM_MATCH() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_ELEM_MATCH, 0)
}

func (s *OperatorKeyContext) DOLLAR_SIZE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SIZE, 0)
}

func (s *OperatorKeyContext) DOLLAR_MATCH() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_MATCH, 0)
}

func (s *OperatorKeyContext) DOLLAR_GROUP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_GROUP, 0)
}

func (s *OperatorKeyContext) DOLLAR_PROJECT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_PROJECT, 0)
}

func (s *OperatorKeyContext) DOLLAR_LIMIT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_LIMIT, 0)
}

func (s *OperatorKeyContext) DOLLAR_SKIP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SKIP, 0)
}

func (s *OperatorKeyContext) DOLLAR_UNWIND() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_UNWIND, 0)
}

func (s *OperatorKeyContext) DOLLAR_LOOKUP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_LOOKUP, 0)
}

func (s *OperatorKeyContext) DOLLAR_SAMPLE() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SAMPLE, 0)
}

func (s *OperatorKeyContext) DOLLAR_COUNT() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_COUNT, 0)
}

func (s *OperatorKeyContext) DOLLAR_FACET() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_FACET, 0)
}

func (s *OperatorKeyContext) DOLLAR_BUCKET() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_BUCKET, 0)
}

func (s *OperatorKeyContext) DOLLAR_BUCKET_AUTO() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_BUCKET_AUTO, 0)
}

func (s *OperatorKeyContext) DOLLAR_SUM() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_SUM, 0)
}

func (s *OperatorKeyContext) DOLLAR_AVG() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_AVG, 0)
}

func (s *OperatorKeyContext) DOLLAR_FIRST() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_FIRST, 0)
}

func (s *OperatorKeyContext) DOLLAR_LAST() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_LAST, 0)
}

func (s *OperatorKeyContext) DOLLAR_STDDEV_POP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_STDDEV_POP, 0)
}

func (s *OperatorKeyContext) DOLLAR_STDDEV_SAMP() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserDOLLAR_STDDEV_SAMP, 0)
}

func (s *OperatorKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterOperatorKey(s)
	}
}

func (s *OperatorKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitOperatorKey(s)
	}
}

func (p *MongoCommandParser) OperatorKey() (localctx IOperatorKeyContext) {
	localctx = NewOperatorKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MongoCommandParserRULE_operatorKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-43)) & ^0x3f) == 0 && ((int64(1)<<(_la-43))&1152921504606846975) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleValue() ISimpleValueContext
	Document() IDocumentContext
	Array() IArrayContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) SimpleValue() ISimpleValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleValueContext)
}

func (s *ValueContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *MongoCommandParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MongoCommandParserRULE_value)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoCommandParserDOLLAR_SET, MongoCommandParserDOLLAR_UNSET, MongoCommandParserDOLLAR_INC, MongoCommandParserDOLLAR_PUSH, MongoCommandParserDOLLAR_PULL, MongoCommandParserDOLLAR_ADD_TO_SET, MongoCommandParserDOLLAR_POP, MongoCommandParserDOLLAR_PULL_ALL, MongoCommandParserDOLLAR_EACH, MongoCommandParserDOLLAR_POSITION, MongoCommandParserDOLLAR_SLICE, MongoCommandParserDOLLAR_SORT, MongoCommandParserDOLLAR_MUL, MongoCommandParserDOLLAR_MIN, MongoCommandParserDOLLAR_MAX, MongoCommandParserDOLLAR_RENAME, MongoCommandParserDOLLAR_CURRENT_DATE, MongoCommandParserDOLLAR_BIT, MongoCommandParserDOLLAR_REPLACE_ALL, MongoCommandParserDOLLAR_FIND, MongoCommandParserDOLLAR_IN, MongoCommandParserDOLLAR_NIN, MongoCommandParserDOLLAR_GT, MongoCommandParserDOLLAR_GTE, MongoCommandParserDOLLAR_LT, MongoCommandParserDOLLAR_LTE, MongoCommandParserDOLLAR_EQ, MongoCommandParserDOLLAR_NE, MongoCommandParserDOLLAR_EXISTS, MongoCommandParserDOLLAR_TYPE, MongoCommandParserDOLLAR_OR, MongoCommandParserDOLLAR_AND, MongoCommandParserDOLLAR_NOT, MongoCommandParserDOLLAR_NOR, MongoCommandParserDOLLAR_MOD, MongoCommandParserDOLLAR_REGEX, MongoCommandParserDOLLAR_TRIM, MongoCommandParserDOLLAR_TEXT, MongoCommandParserDOLLAR_WHERE, MongoCommandParserDOLLAR_ALL, MongoCommandParserDOLLAR_ELEM_MATCH, MongoCommandParserDOLLAR_SIZE, MongoCommandParserDOLLAR_MATCH, MongoCommandParserDOLLAR_GROUP, MongoCommandParserDOLLAR_PROJECT, MongoCommandParserDOLLAR_LIMIT, MongoCommandParserDOLLAR_SKIP, MongoCommandParserDOLLAR_UNWIND, MongoCommandParserDOLLAR_LOOKUP, MongoCommandParserDOLLAR_SAMPLE, MongoCommandParserDOLLAR_COUNT, MongoCommandParserDOLLAR_FACET, MongoCommandParserDOLLAR_BUCKET, MongoCommandParserDOLLAR_BUCKET_AUTO, MongoCommandParserDOLLAR_SUM, MongoCommandParserDOLLAR_AVG, MongoCommandParserDOLLAR_FIRST, MongoCommandParserDOLLAR_LAST, MongoCommandParserDOLLAR_STDDEV_POP, MongoCommandParserDOLLAR_STDDEV_SAMP, MongoCommandParserBOOLEAN, MongoCommandParserNULL, MongoCommandParserSTRING, MongoCommandParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.SimpleValue()
		}

	case MongoCommandParserLCURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Document()
		}

	case MongoCommandParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Array()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleValueContext is an interface to support dynamic dispatch.
type ISimpleValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NULL() antlr.TerminalNode
	OperatorKey() IOperatorKeyContext

	// IsSimpleValueContext differentiates from other interfaces.
	IsSimpleValueContext()
}

type SimpleValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleValueContext() *SimpleValueContext {
	var p = new(SimpleValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_simpleValue
	return p
}

func InitEmptySimpleValueContext(p *SimpleValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_simpleValue
}

func (*SimpleValueContext) IsSimpleValueContext() {}

func NewSimpleValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleValueContext {
	var p = new(SimpleValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_simpleValue

	return p
}

func (s *SimpleValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserSTRING, 0)
}

func (s *SimpleValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserNUMBER, 0)
}

func (s *SimpleValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserBOOLEAN, 0)
}

func (s *SimpleValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserNULL, 0)
}

func (s *SimpleValueContext) OperatorKey() IOperatorKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorKeyContext)
}

func (s *SimpleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterSimpleValue(s)
	}
}

func (s *SimpleValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitSimpleValue(s)
	}
}

func (p *MongoCommandParser) SimpleValue() (localctx ISimpleValueContext) {
	localctx = NewSimpleValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MongoCommandParserRULE_simpleValue)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoCommandParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(MongoCommandParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoCommandParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(MongoCommandParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoCommandParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(MongoCommandParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoCommandParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Match(MongoCommandParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoCommandParserDOLLAR_SET, MongoCommandParserDOLLAR_UNSET, MongoCommandParserDOLLAR_INC, MongoCommandParserDOLLAR_PUSH, MongoCommandParserDOLLAR_PULL, MongoCommandParserDOLLAR_ADD_TO_SET, MongoCommandParserDOLLAR_POP, MongoCommandParserDOLLAR_PULL_ALL, MongoCommandParserDOLLAR_EACH, MongoCommandParserDOLLAR_POSITION, MongoCommandParserDOLLAR_SLICE, MongoCommandParserDOLLAR_SORT, MongoCommandParserDOLLAR_MUL, MongoCommandParserDOLLAR_MIN, MongoCommandParserDOLLAR_MAX, MongoCommandParserDOLLAR_RENAME, MongoCommandParserDOLLAR_CURRENT_DATE, MongoCommandParserDOLLAR_BIT, MongoCommandParserDOLLAR_REPLACE_ALL, MongoCommandParserDOLLAR_FIND, MongoCommandParserDOLLAR_IN, MongoCommandParserDOLLAR_NIN, MongoCommandParserDOLLAR_GT, MongoCommandParserDOLLAR_GTE, MongoCommandParserDOLLAR_LT, MongoCommandParserDOLLAR_LTE, MongoCommandParserDOLLAR_EQ, MongoCommandParserDOLLAR_NE, MongoCommandParserDOLLAR_EXISTS, MongoCommandParserDOLLAR_TYPE, MongoCommandParserDOLLAR_OR, MongoCommandParserDOLLAR_AND, MongoCommandParserDOLLAR_NOT, MongoCommandParserDOLLAR_NOR, MongoCommandParserDOLLAR_MOD, MongoCommandParserDOLLAR_REGEX, MongoCommandParserDOLLAR_TRIM, MongoCommandParserDOLLAR_TEXT, MongoCommandParserDOLLAR_WHERE, MongoCommandParserDOLLAR_ALL, MongoCommandParserDOLLAR_ELEM_MATCH, MongoCommandParserDOLLAR_SIZE, MongoCommandParserDOLLAR_MATCH, MongoCommandParserDOLLAR_GROUP, MongoCommandParserDOLLAR_PROJECT, MongoCommandParserDOLLAR_LIMIT, MongoCommandParserDOLLAR_SKIP, MongoCommandParserDOLLAR_UNWIND, MongoCommandParserDOLLAR_LOOKUP, MongoCommandParserDOLLAR_SAMPLE, MongoCommandParserDOLLAR_COUNT, MongoCommandParserDOLLAR_FACET, MongoCommandParserDOLLAR_BUCKET, MongoCommandParserDOLLAR_BUCKET_AUTO, MongoCommandParserDOLLAR_SUM, MongoCommandParserDOLLAR_AVG, MongoCommandParserDOLLAR_FIRST, MongoCommandParserDOLLAR_LAST, MongoCommandParserDOLLAR_STDDEV_POP, MongoCommandParserDOLLAR_STDDEV_SAMP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.OperatorKey()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoCommandParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoCommandParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserLBRACK, 0)
}

func (s *ArrayContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MongoCommandParserRBRACK, 0)
}

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoCommandParserCOMMA)
}

func (s *ArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoCommandParserCOMMA, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoCommandListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *MongoCommandParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MongoCommandParserRULE_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(MongoCommandParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6047313952768) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&15393162788863) != 0) {
		{
			p.SetState(110)
			p.Value()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MongoCommandParserCOMMA {
			{
				p.SetState(111)
				p.Match(MongoCommandParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(112)
				p.Value()
			}

			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(120)
		p.Match(MongoCommandParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
