grammar MongoCommand;

// Parser Rules
mongoCommand
    : command (SEMICOLON | EOF)
    ;

command
    : (DB DOT collection DOT operation arguments)
    | (DB DOT GET_COLLECTION LPAREN STRING RPAREN DOT operation arguments)
    ;

collection
    : IDENTIFIER
    ;

operation
    : // CRUD Operations
      'find'
    | 'findOne'
    | 'insertOne'
    | 'insertMany'
    | 'updateOne'
    | 'updateMany'
    | 'deleteOne'
    | 'deleteMany'
    | 'replaceOne'
    | 'replaceAll'

    // Aggregation
    | 'aggregate'
    | 'count'
    | 'countDocuments'
    | 'estimatedDocumentCount'
    | 'distinct'

    // Index Management
    | 'createIndex'
    | 'createIndexes'
    | 'dropIndex'
    | 'dropIndexes'
    | 'listIndexes'
    | 'reIndex'

    // Collection Management
    | 'drop'
    | 'rename'
    | 'stats'

    // Bulk Operations
    | 'bulkWrite'

    // Additional Operations
    | 'findOneAndUpdate'
    | 'findOneAndDelete'
    | 'findOneAndReplace'
    | 'isCapped'
    | 'explain'
    | 'watch'
    ;

arguments
    : LPAREN argument (COMMA argument)* RPAREN
    | LPAREN RPAREN
    ;

argument
    : document
    | array
    | IDENTIFIER // Allow bare identifiers
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
    : // Update Operators
      DOLLAR_SET
    | DOLLAR_UNSET
    | DOLLAR_INC
    | DOLLAR_PUSH
    | DOLLAR_PULL
    | DOLLAR_ADD_TO_SET
    | DOLLAR_POP
    | DOLLAR_PULL_ALL
    | DOLLAR_EACH
    | DOLLAR_POSITION
    | DOLLAR_SLICE
    | DOLLAR_SORT
    | DOLLAR_MUL
    | DOLLAR_MIN
    | DOLLAR_MAX
    | DOLLAR_RENAME
    | DOLLAR_CURRENT_DATE
    | DOLLAR_BIT
    | DOLLAR_REPLACE_ALL

    // Query Operators
    | DOLLAR_FIND
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
    | DOLLAR_NOR
    | DOLLAR_MOD
    | DOLLAR_TRIM
    | DOLLAR_REGEX
    | DOLLAR_TEXT
    | DOLLAR_WHERE
    | DOLLAR_ALL
    | DOLLAR_ELEM_MATCH
    | DOLLAR_SIZE

    // Aggregation Pipeline Operators
    | DOLLAR_MATCH
    | DOLLAR_GROUP
    | DOLLAR_PROJECT
    | DOLLAR_LIMIT
    | DOLLAR_SKIP
    | DOLLAR_SORT
    | DOLLAR_UNWIND
    | DOLLAR_LOOKUP
    | DOLLAR_SAMPLE
    | DOLLAR_COUNT
    | DOLLAR_FACET
    | DOLLAR_BUCKET
    | DOLLAR_BUCKET_AUTO

    // Aggregation Accumulators
    | DOLLAR_SUM
    | DOLLAR_AVG
    | DOLLAR_FIRST
    | DOLLAR_LAST
    | DOLLAR_PUSH
    | DOLLAR_ADD_TO_SET
    | DOLLAR_STDDEV_POP
    | DOLLAR_STDDEV_SAMP
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
// Update Operators
DOLLAR_SET: '$set';
DOLLAR_UNSET: '$unset';
DOLLAR_INC: '$inc';
DOLLAR_PUSH: '$push';
DOLLAR_PULL: '$pull';
DOLLAR_ADD_TO_SET: '$addToSet';
DOLLAR_POP: '$pop';
DOLLAR_PULL_ALL: '$pullAll';
DOLLAR_EACH: '$each';
DOLLAR_POSITION: '$position';
DOLLAR_SLICE: '$slice';
DOLLAR_SORT: '$sort';
DOLLAR_MUL: '$mul';
DOLLAR_MIN: '$min';
DOLLAR_MAX: '$max';
DOLLAR_RENAME: '$rename';
DOLLAR_CURRENT_DATE: '$currentDate';
DOLLAR_BIT: '$bit';
DOLLAR_REPLACE_ALL: '$replaceAll';

// Query Operators
DOLLAR_FIND: 'find';
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
DOLLAR_NOR: '$nor';
DOLLAR_MOD: '$mod';
DOLLAR_REGEX: '$regex';
DOLLAR_TRIM: '$trim';
DOLLAR_TEXT: '$text';
DOLLAR_WHERE: '$where';
DOLLAR_ALL: '$all';
DOLLAR_ELEM_MATCH: '$elemMatch';
DOLLAR_SIZE: '$size';

// Aggregation Pipeline Operators
DOLLAR_MATCH: '$match';
DOLLAR_GROUP: '$group';
DOLLAR_PROJECT: '$project';
DOLLAR_LIMIT: '$limit';
DOLLAR_SKIP: '$skip';
DOLLAR_UNWIND: '$unwind';
DOLLAR_LOOKUP: '$lookup';
DOLLAR_SAMPLE: '$sample';
DOLLAR_COUNT: '$count';
DOLLAR_FACET: '$facet';
DOLLAR_BUCKET: '$bucket';
DOLLAR_BUCKET_AUTO: '$bucketAuto';

// Aggregation Accumulators
DOLLAR_SUM: '$sum';
DOLLAR_AVG: '$avg';
DOLLAR_FIRST: '$first';
DOLLAR_LAST: '$last';
DOLLAR_STDDEV_POP: '$stdDevPop';
DOLLAR_STDDEV_SAMP: '$stdDevSamp';

BOOLEAN: 'true' | 'false';
NULL: 'null';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
STRING: '"' (~["\r\n] | '\\"')* '"' | '\'' (~['\r\n] | '\\\'')* '\'';
NUMBER: '-'? [0-9]+ ('.' [0-9]+)?;

// Skip whitespace and allow multiline
WS: [ \t\r\n]+ -> skip;
SINGLE_LINE_COMMENT: '//' .*? '\r'? '\n' -> skip;
MULTI_LINE_COMMENT: '/*' .*? '*/' -> skip;