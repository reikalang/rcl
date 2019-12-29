grammar YAML;

// TODO: not working, just copied from RCL, not handling any indent
// start of parser
pair
    : k=ID ':' value
    | k=STRING ':' value
    ;

value
    : BOOL    #ValBool
    | INT     #ValInt // TODO: negative number
    | DOUBLE  #ValDouble
    | STRING  #ValString
    ;

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

NEWLINE
    : '\r'? '\n'
    ;
