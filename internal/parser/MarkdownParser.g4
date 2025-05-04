parser grammar MarkdownParser;

options { tokenVocab=MarkdownLexer; }

/*
 * Parser Rules
 */

document
    : block* EOF
    ;

block
    : heading
    | paragraph
    | blockquote
    | unorderedList
    | orderedList
    | codeBlock
    | horizontalRule
    ;

h1
    : HASH WHITESPACE data=line NEWLINE*
    ;

h2
    : HASH HASH WHITESPACE data=line NEWLINE*
    ;

h3
    : HASH HASH HASH WHITESPACE data=line NEWLINE*
    ;

h4
    : HASH HASH HASH HASH WHITESPACE data=line NEWLINE*
    ;

h5
    : HASH HASH HASH HASH HASH WHITESPACE data=line NEWLINE*
    ;

h6
    : HASH HASH HASH HASH HASH HASH WHITESPACE data=line NEWLINE*
    ;

heading
    : h1
    | h2
    | h3
    | h4
    | h5
    | h6
    ;

paragraph
    : (line | WHITESPACE | inlineElement)+ NEWLINE? 
    ;

blockquote
    : BLOCKQUOTE_MARKER (line | inlineElement)+ (NEWLINE BLOCKQUOTE_MARKER (line | inlineElement)+)* NEWLINE?
    ;

unorderedList
    : listItem+ 
    ;

orderedList
    : orderedListItem+
    ;

listItem
    : UNORDERED_LIST_MARKER (line | inlineElement)+ (NEWLINE WHITESPACE UNORDERED_LIST_MARKER? (line | inlineElement)+)* 
    ;

orderedListItem
    : ORDERED_LIST_MARKER (line | inlineElement)+ (NEWLINE WHITESPACE ORDERED_LIST_MARKER? (line | inlineElement)+)*
    ;

codeBlock
    : FENCED_CODE_BLOCK
    | INDENTED_CODE_BLOCK
    ;

horizontalRule
    : HORIZONTAL_RULE
    ;

line
    : (inlineElement | WHITESPACE | TEXT)+ NEWLINE?
    ;

inlineElement
    : bold
    | italic
    | boldAndItalic
    | strikethrough
    | code
    ;

bold
    : BOLD_MARKER data=boldText BOLD_MARKER
    ;
    
boldAndItalic
    : BOLD_AND_ITALIC_MARKER data=italicText BOLD_AND_ITALIC_MARKER
    ;

boldText
    : (TEXT | WHITESPACE | ITALIC_MARKER | CODE_MARKER)+
    ;

italic
    : ITALIC_MARKER data=italicText ITALIC_MARKER
    ;

italicText
    : (TEXT | WHITESPACE | CODE_MARKER)+
    ;

strikethrough
    : STRIKETHROUGH_MARKER data=strikethroughText STRIKETHROUGH_MARKER
    ;

strikethroughText
    : (TEXT | WHITESPACE | ITALIC_MARKER | BOLD_MARKER | CODE_MARKER)+
    ;

code
    : CODE_MARKER data=codeText CODE_MARKER
    ;

codeText
    : (TEXT | WHITESPACE | ESCAPE_CHAR)+
    ;

