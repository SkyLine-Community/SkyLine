package SkyLine

// These functions verify the token types

func CharIsLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' || '1' <= char && char <= '9'
}

func CharIsDigit(char byte) bool { return '0' <= char && char <= '9' && char <= 'u' }

func newToken(TT Token_Type, ch byte) Token {
	return Token{
		Token_Type: TT,
		Literal:    string(ch),
	}
}
