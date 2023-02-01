package SkyLine

func (lex *LexerStructure) NextToken() Token {
	lex.ConsumeWhiteSpace()
	if lex.Char == '#' {
		lex.ConsumeComment()
	}

	var tok Token
	switch lex.Char {
	case '=':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: EQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(ASSIGN, lex.Char)
		}
	case '!':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: NEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(BANG, lex.Char)
		}
	case ';':
		tok = newToken(SEMICOLON, lex.Char)
	case ':':
		tok = newToken(COLON, lex.Char)
	case '(':
		tok = newToken(LPAREN, lex.Char)
	case '|':
		tok = newToken(LINE, lex.Char)
	case ')':
		tok = newToken(RPAREN, lex.Char)
	case ',':
		tok = newToken(COMMA, lex.Char)
	case '+':
		tok = newToken(PLUS, lex.Char)
	case '-':
		tok = newToken(MINUS, lex.Char)
	case '*':
		tok = newToken(ASTARISK, lex.Char)
	case '/':
		tok = newToken(SLASH, lex.Char)
	case '<':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: LTEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(LT, lex.Char)
		}
	case '>':
		if lex.Peek() == '=' {
			cha := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: GTEQ,
				Literal:    string(cha) + string(lex.Char),
			}
		} else {
			tok = newToken(GT, lex.Char)
		}
	case '{':
		tok = newToken(LBRACE, lex.Char)
	case '}':
		tok = newToken(RBRACE, lex.Char)
	case '[':
		tok = newToken(LBRACKET, lex.Char)
	case ']':
		tok = newToken(RBRACKET, lex.Char)
	case '"':
		tok.Token_Type = STRING
		tok.Literal = lex.ReadString()
	case 0:
		tok.Literal = ""
		tok.Token_Type = EOF
	default:
		if CharIsDigit(lex.Char) {
			return lex.ReadIntToken()
		}

		if CharIsLetter(lex.Char) {
			tok.Literal = lex.ReadIdentifier()
			tok.Token_Type = LookupIdentifier(tok.Literal)
			return tok
		}

		tok = newToken(ILLEGAL, lex.Char)
	}
	lex.ReadChar()
	return tok
}

/*


var LexerMap = map[byte]func(lex *LexerStructure) Token{
	'=': func(lex *LexerStructure) Token {
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			return Token{
				Token_Type: EQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			return newToken(ASSIGN, lex.Char)
		}
	}, // Assignment
	'!': func(lex *LexerStructure) Token {
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			return Token{
				Token_Type: NEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			return newToken(BANG, lex.Char)
		}
	}, // BANG
	':': func(lex *LexerStructure) Token {
		return newToken(COLON, lex.Char)
	}, // Colon BEGIN, KEY, SEPERATOR
	';': func(lex *LexerStructure) Token {
		return newToken(SEMICOLON, lex.Char)
	}, // Semicolon END
	'(': func(lex *LexerStructure) Token {
		return newToken(LPAREN, lex.Char)
	}, // L-Argument paren
	')': func(lex *LexerStructure) Token {
		return newToken(RPAREN, lex.Char)
	}, // R-Argument paren
	'{': func(lex *LexerStructure) Token {
		return newToken(LBRACE, lex.Char)
	}, // L-StartBrick brace
	'}': func(lex *LexerStructure) Token {
		return newToken(RBRACE, lex.Char)
	}, // R-EndBrick brace
	'[': func(lex *LexerStructure) Token {
		return newToken(LBRACKET, lex.Char)
	}, // L-StartArray Bracket
	']': func(lex *LexerStructure) Token {
		return newToken(RBRACKET, lex.Char)
	}, // R-EndArray Bracket
	'|': func(lex *LexerStructure) Token {
		return newToken(LINE, lex.Char)
	}, // Bar
	',': func(lex *LexerStructure) Token {
		return newToken(COMMA, lex.Char)
	}, // Seperation
	'+': func(lex *LexerStructure) Token {
		return newToken(PLUS, lex.Char)
	}, // Addition
	'/': func(lex *LexerStructure) Token {
		return newToken(SLASH, lex.Char)
	}, // Division
	'*': func(lex *LexerStructure) Token {
		return newToken(ASTARISK, lex.Char)
	}, // Multiplication
	'-': func(lex *LexerStructure) Token {
		return newToken(MINUS, lex.Char)
	}, // Subtract
	'<': func(lex *LexerStructure) Token {
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			return Token{
				Token_Type: LTEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			return newToken(LT, lex.Char)
		}
	}, // Less than
	'>': func(lex *LexerStructure) Token {
		if lex.Peek() == '=' {
			cha := lex.Char
			lex.ReadChar()
			return Token{
				Token_Type: GTEQ,
				Literal:    string(cha) + string(lex.Char),
			}
		} else {
			return newToken(GT, lex.Char)
		}
	}, // Greater than
	'"': func(lex *LexerStructure) Token {
		return Token{
			Token_Type: STRING,
			Literal:    lex.ReadString(),
		}
	},
	0: func(lex *LexerStructure) Token {
		return Token{
			Token_Type: EOF,
			Literal:    "",
		}
	}, // EOF

}


*/
