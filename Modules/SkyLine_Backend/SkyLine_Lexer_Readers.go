package SkyLine

func (lex *LexerStructure) ReadIdentifier() string { return lex.read(CharIsLetter) }
func (lex *LexerStructure) ReadNumber() string     { return lex.read(CharIsDigit) }

func (lex *LexerStructure) ReadChar() {
	if lex.RPOS >= len(lex.CharInput) {
		lex.Char = 0
	} else {
		lex.Char = lex.CharInput[lex.RPOS]
	}
	lex.POS = lex.RPOS
	lex.RPOS++
}

func (lex *LexerStructure) ReadString() string {
	POS := lex.POS + 1
	for {
		lex.ReadChar()
		if lex.Char == '"' || lex.Char == 0 {
			break
		}
	}
	return lex.CharInput[POS:lex.POS]
}

func (lex *LexerStructure) read(checkFn func(byte) bool) string {
	position := lex.POS
	for checkFn(lex.Char) {
		lex.ReadChar()
	}
	return lex.CharInput[position:lex.POS]
}

func (lex *LexerStructure) ReadIntToken() Token {
	intPart := lex.ReadNumber()
	if lex.Char != '.' && lex.Char != 'u' {
		return Token{
			Token_Type: INT,
			Literal:    intPart,
		}
	}

	lex.ReadChar()
	fracPart := lex.ReadNumber()
	return Token{
		Token_Type: FLOAT,
		Literal:    intPart + "." + fracPart,
	}
}

func LexNew(input string) *LexerStructure {
	lex := &LexerStructure{CharInput: input}
	lex.ReadChar()
	return lex
}

func (lex *LexerStructure) Peek() byte {
	if lex.RPOS >= len(lex.CharInput) {
		return 0
	}
	return lex.CharInput[lex.RPOS]
}
