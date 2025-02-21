grammar MongoCommand;

// Parser Rules
parse
    : command (SEMICOLON EOF | EOF)
    ;

command
    : (DB DOT collection DOT operation arguments)
    | (DB DOT GET_COLLECTION LPAREN STRING RPAREN DOT operation arguments)
    ;

collection
    : IDENTIFIER
    ;

operation
    : 'find'
    | 'findOne'
    | 'insertOne'
    | 'insertMany'
    | 'updateOne'
    | 'updateMany'
    | 'deleteOne'
    | 'deleteMany'
    | 'aggregate'
    | 'count'
    ;

arguments
    : LPAREN argument (COMMA argument)* RPAREN
    | LPAREN RPAREN
    ;

argument
    : document
    | array
    ;

document
    : LCURLY (pair (COMMA pair)*)? RCURLY
    ;

pair
    : key=STRING COLON value
    | key=IDENTIFIER COLON value
    | operatorKey COLON value
    ;

operatorKey
    : DOLLAR_SET
    | DOLLAR_UNSET
    | DOLLAR_INC
    | DOLLAR_PUSH
    | DOLLAR_PULL
    | DOLLAR_IN
    | DOLLAR_NIN
    | DOLLAR_GT
    | DOLLAR_GTE
    | DOLLAR_LT
    | DOLLAR_LTE
    | DOLLAR_EQ
    | DOLLAR_NE
    | DOLLAR_EXISTS
    | DOLLAR_TYPE
    | DOLLAR_OR
    | DOLLAR_AND
    | DOLLAR_NOT
    | DOLLAR_MATCH
    | DOLLAR_GROUP
    | DOLLAR_PROJECT
    ;

value
    : simpleValue
    | document
    | array
    ;

simpleValue
    : STRING
    | NUMBER
    | BOOLEAN
    | NULL
    | operatorKey
    ;

array
    : LBRACK (value (COMMA value)*)? RBRACK
    ;

// Lexer Rules
DB: 'db';
GET_COLLECTION: 'getCollection';
DOT: '.';
COMMA: ',';
COLON: ':';
SEMICOLON: ';';
LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';
LCURLY: '{';
RCURLY: '}';

// MongoDB Operators
DOLLAR_SET: '$set';
DOLLAR_UNSET: '$unset';
DOLLAR_INC: '$inc';
DOLLAR_PUSH: '$push';
DOLLAR_PULL: '$pull';
DOLLAR_IN: '$in';
DOLLAR_NIN: '$nin';
DOLLAR_GT: '$gt';
DOLLAR_GTE: '$gte';
DOLLAR_LT: '$lt';
DOLLAR_LTE: '$lte';
DOLLAR_EQ: '$eq';
DOLLAR_NE: '$ne';
DOLLAR_EXISTS: '$exists';
DOLLAR_TYPE: '$type';
DOLLAR_OR: '$or';
DOLLAR_AND: '$and';
DOLLAR_NOT: '$not';
DOLLAR_MATCH: '$match';
DOLLAR_GROUP: '$group';
DOLLAR_PROJECT: '$project';

BOOLEAN: 'true' | 'false';
NULL: 'null';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
STRING: '"' (~["\r\n] | '\\"')* '"' | '\'' (~['\r\n] | '\\\'')* '\'';
NUMBER: '-'? [0-9]+ ('.' [0-9]+)?;

// Skip whitespace and allow multiline
WS: [ \t\r\n]+ -> skip;
SINGLE_LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;
MULTI_LINE_COMMENT: '/*' .*? '*/' -> skip;