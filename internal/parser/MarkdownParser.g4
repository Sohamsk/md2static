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

heading
    : HASH+ WHITESPACE line NEWLINE*
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
    : BOLD_MARKER text=boldText BOLD_MARKER
    ;
    
boldAndItalic
    : BOLD_AND_ITALIC_MARKER text=italicText BOLD_AND_ITALIC_MARKER
    ;

boldText
    : (TEXT | WHITESPACE | ITALIC_MARKER | CODE_MARKER)+
    ;

italic
    : ITALIC_MARKER text=italicText ITALIC_MARKER
    ;

italicText
    : (TEXT | WHITESPACE | CODE_MARKER)+
    ;

strikethrough
    : STRIKETHROUGH_MARKER text=strikethroughText STRIKETHROUGH_MARKER
    ;

strikethroughText
    : (TEXT | WHITESPACE | ITALIC_MARKER | BOLD_MARKER | CODE_MARKER)+
    ;

code
    : CODE_MARKER text=codeText CODE_MARKER
    ;

codeText
    : (TEXT | WHITESPACE | ESCAPE_CHAR)+
    ;
