
parser grammar MarkdownParser;

options { tokenVocab=MarkdownLexer; }

document
    : block* EOF
    ;

block
    : heading
    | paragraph
    | unorderedList
    | NEWLINE+
    ;

// paragraph and co.
paragraph
    : (line (linebreak | NEWLINE?))+
    | linebreak
    ;

line
    : (WORD | WHITESPACE | NUMBER | PUNCTUATION | DASHES | EQUALS | escape_char | inline)+
    ;

escape_char
    : ESCAPE_CHAR
    ;

linebreak
    : LINE_BREAK
    ;

inline
    : italic
    | bold
    | bold_and_italic
    | code
    ;

italic
    : ASTERISK data=italic_line ASTERISK
    ;

italic_line
    : ((WORD | WHITESPACE | ESCAPE_CHAR)+ | bold)+ 
    ;

bold
    : BOLD_MARKER data=bold_line BOLD_MARKER
    ;

bold_line
    : ((WORD | WHITESPACE | ESCAPE_CHAR)+ | italic)+
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

// headings and co.
h1
    : HASH WHITESPACE data=line (NEWLINE? | LINE_BREAK)
    | data=line (NEWLINE | LINE_BREAK) EQUALS NEWLINE?
    ;

h2
    : H2 WHITESPACE data=line (NEWLINE? | LINE_BREAK)
    | data=line (NEWLINE | LINE_BREAK) DASHES NEWLINE?
    ;

h3
    : H3 WHITESPACE data=line NEWLINE?
    ;

h4
    : H4 WHITESPACE data=line NEWLINE?
    ;

h5
    : H5 WHITESPACE data=line NEWLINE?
    ;

h6
    : H6 WHITESPACE data=line NEWLINE?
    ;


heading
    : h1 | h2 | h3 | h4 | h5 | h6
    ;

// unorderedlist and co.
unorderedList
    : unorderedListItem+
    ;

unorderedListItem
    : listMarker line (NEWLINE continuationLine)* NEWLINE?
    ;

listMarker
    : (DASH | ASTERISK | PLUS) WHITESPACE
    ;

continuationLine
    : WHITESPACE line
    ;

