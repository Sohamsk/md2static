lexer grammar MarkdownLexer;

// Basic characters
WHITESPACE : [ \t]+ ;
NEWLINE : '\r'? '\n' ;
ESCAPE_CHAR : '\\' . ;

// Formatting markers
BOLD_MARKER : '**' | '__' ;
ITALIC_MARKER : '*' | '_' ;
STRIKETHROUGH_MARKER : '~~';
BOLD_AND_ITALIC_MARKER: '***' | '___';  
CODE_MARKER : '`' ;

// Block markers
BLOCKQUOTE_MARKER : '>' WHITESPACE? ;
UNORDERED_LIST_MARKER : ('-' | '*' | '+') WHITESPACE ;
ORDERED_LIST_MARKER : [0-9]+ '.' WHITESPACE ;
HORIZONTAL_RULE : ('---' | '___' | '***') NEWLINE ;

// Code blocks
FENCED_CODE_BLOCK
    : '```' .*? '```'
    | '~~~' .*? '~~~'
    ;
INDENTED_CODE_BLOCK : WHITESPACE WHITESPACE WHITESPACE WHITESPACE .*? NEWLINE ;

// Structure characters
HASH : '#' ;  // Used for headings
LBRACKET : '[' ;
RBRACKET : ']' ;
LPAREN : '(' ;
RPAREN : ')' ;
EXCLAMATION : '!' ;

// Text content (catch-all for normal text)
// TEXT : ~[ \t\r\n*_`\\[\]()#>~!/:.?=&%+\-]+ ;
TEXT : ~[ \r\n\\`*_#[\]()!>]+ ;

