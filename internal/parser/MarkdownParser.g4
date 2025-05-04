parser grammar MarkdownParser;

options { tokenVocab=MarkdownLexer; }

document
    : block* EOF
    ;

block
    : heading
    | paragraph
    | code
    ;

line
    : (WORD | WHITESPACE | NUMBER | PUNCTUATION | inline)+ NEWLINE*
    ;

inline
    : italic
    | bold
    | bold_and_italic
    ;

italic
    : ITALIC_MARKER data=italic_line ITALIC_MARKER
    ;

italic_line
    : ((WORD | WHITESPACE)+ | bold)+ 
    ;

bold
    : BOLD_MARKER data=bold_line BOLD_MARKER
    ;

bold_line
    : ((WORD | WHITESPACE)+ | italic)+
    ;

bold_and_italic
    : BOLD_AND_ITALIC_MARKER data=bold_line BOLD_AND_ITALIC_MARKER
    ;

code
    : CODE_MARKER data=code_text CODE_MARKER
    ;

code_text
    : (WORD | WHITESPACE | ESCAPE_CHAR)+
    ;

h1
    : HASH WHITESPACE data=line NEWLINE*
    ;

h2
    : H2 WHITESPACE data=line NEWLINE*
    ;

h3
    : H3 WHITESPACE data=line NEWLINE*
    ;

h4
    : H4 WHITESPACE data=line NEWLINE*
    ;

h5
    : H5 WHITESPACE data=line NEWLINE*
    ;

h6
    : H6 WHITESPACE data=line NEWLINE*
    ;

paragraph
    : line+
    ;

heading
    : h1 | h2 | h3 | h4 | h5 | h6
    ;
