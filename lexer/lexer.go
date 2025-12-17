package lexer


type Lexer struct {
    input        string
    position     int      // current position in input (points to current char)
    readPosition int      // current reading position in input (after current char)
    ch           byte     // current char under examination
    line         int      // current line number for error reporting
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{Type: ASSIGN, Literal: string(l.ch), Line: l.line}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.ch), Line: l.line}
	case '-':
		tok = Token{Type: MINUS, Literal: string(l.ch), Line: l.line}
	case ';':
		tok = Token{Type: SEMICOLON, Literal: string(l.ch), Line: l.line}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch), Line: l.line}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch), Line: l.line}
	case '{':
		tok = Token{Type: LBRACE, Literal: string(l.ch), Line: l.line}
	case '}':
		tok = Token{Type: RBRACE, Literal: string(l.ch), Line: l.line}
	case ',':
		tok = Token{Type: COMMA, Literal: string(l.ch), Line: l.line}
	case '#':
		tok = Token{Type: HASH, Literal: string(l.ch), Line: l.line}
	case '<':
		if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LSHIFT, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = Token{Type: LESS, Literal: string(l.ch), Line: l.line}
		}
	case '>':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: RSHIFT, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = Token{Type: GREATER, Literal: string(l.ch), Line: l.line}
		}
	case '/':
		if l.peekChar() == '/' {
			l.skipLineComment()
			return l.NextToken()
		} else if l.peekChar() == '*' {
			l.skipBlockComment()
			return l.NextToken()
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: l.line}
		}
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = INTEGER_LIT
			tok.Line = l.line
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: l.line}
		}
	}

	l.readChar()
	return tok
}


func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) skipLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	l.readChar() // رد کردن /
	l.readChar() // رد کردن *
	for !(l.ch == '*' && l.peekChar() == '/') && l.ch != 0 {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
	l.readChar() // رد کردن *
	l.readChar() // رد کردن /
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func LookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"int":     INT,
		"return":  RETURN,
		"include": INCLUDE,
		// "main" و "cout" طبق کدت معمولاً شناسه‌اند، اگر تی‌ای گفت کلمه کلیدی هستن اینجا اضافه‌شون کن.
	}
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}