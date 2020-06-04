grammar RCL;

// start of parser
rcl
    : pair+
    ;

pair
    : k=ID '=' value // bare key
    | k=STRING '=' value
    ;

obj
    : '{' pair (',' pair)* '}'
    | '{' '}'
    ;

array
    : '[' value (',' value)* ']'
    | '[' ']'
    ;

value
    : BOOL    #ValBool
    | INT     #ValInt // TODO: negative number, or simple arith expression, TODO: support _
    | DOUBLE  #ValDouble
    | STRING  #ValString
    | obj     #ValObject
    | array   #ValArray
    ;
// end of parser

// start of lexer
BOOL
    : 'true'
    | 'false'
    ;

INT
    : '0'
    | [1-9] DIGIT*
    ;

DOUBLE
    : DIGIT+ '.' DIGIT+
    ;

STRING
    : '"' (ESC | ~["\\])* '"'
    ;

ID
    : [a-z] ID_LETTER*
    ;

fragment DIGIT
    : [0-9]
    ;

fragment ESC
    : '\\' ["\\/bfnrt]
    ;

fragment ID_LETTER
    : [a-z]
    | [A-Z]
    | '_'
    ;

WS  : [ \t\r\n] -> skip;
BLOCK_COMMENT  : '/*' .*? '*/' -> skip;
SINGLE_COMMENT : '//' .*? '\n' -> skip;
// end of lexer