lexer grammar MarkdownLexer;

WHITESPACE: [ \t]+ ;
NEWLINE: '\r'? '\n' ;
ESCAPE_CHAR : '\\' . ;
LINE_BREAK: WHITESPACE WHITESPACE+ NEWLINE;

// headings
HASH: '#';
H2: HASH HASH;
H3: HASH HASH HASH; 
H4: HASH HASH HASH HASH;
H5: HASH HASH HASH HASH HASH;
H6: HASH HASH HASH HASH HASH HASH;

// setex
DASHES: '-'+;
EQUALS: '='+;

// inlines
BOLD_AND_ITALIC_MARKER: '***';
ITALIC_MARKER: '*';
BOLD_MARKER: '**';
CODE_MARKER: '`';

PUNCTUATION: [.,!?;:(){}'"] | '[' | ']' ;
NUMBER: [0-9]+ ;

WORD : ~[\r\n#* \t`\\]+;
