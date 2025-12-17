package lexer

type TokenType string

const (
	INT      TokenType = "INT"
	RETURN   TokenType = "RETURN"
	INCLUDE  TokenType = "INCLUDE"

	IDENTIFIER TokenType = "IDENTIFIER" // a , b, main, cout
	INTEGER_LIT TokenType = "INTEGER_LIT" 

	
	ASSIGN   TokenType = "ASSIGN"   // =
	PLUS     TokenType = "PLUS"     // +
	MINUS    TokenType = "MINUS"    // -
	LSHIFT   TokenType = "LSHIFT"   // << 
	RSHIFT   TokenType = "RSHIFT"   // >>
	
	LPAREN   TokenType = "LPAREN"   // (
	RPAREN   TokenType = "RPAREN"   // )
	LBRACE   TokenType = "LBRACE"   // {
	RBRACE   TokenType = "RBRACE"   // }
	SEMICOLON TokenType = "SEMICOLON" // ;
	COMMA     TokenType = "COMMA"     // ,
	HASH      TokenType = "HASH"      // #
	LESS      TokenType = "LESS"      // <
	GREATER   TokenType = "GREATER"   // >

	COMMENTLINE TokenType = "COMMENTLINE" //  //
    COMMENTPARAGRAPH TokenType = "COMMENTPARAGRAPH" // /* */


    EOF     TokenType = "EOF"     // FILE END
	ILLEGAL TokenType = "ILLEGAL" // UNKNOWN TOKEN
)

type Token struct {
    Type    TokenType
    Literal string
    Line    int
}