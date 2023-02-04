package SkyLine

func (lex *LexerStructure) ConsumeWhiteSpace() {
	for lex.Char == ' ' || lex.Char == '\t' || lex.Char == '\n' || lex.Char == '\r' {
		lex.ReadChar()
	}
}

func (lex *LexerStructure) ConsumeComment() {
	for lex.Char != '\n' && lex.Char != '\r' && lex.Char != '\t' {
		lex.ReadChar()
	}
	lex.ConsumeWhiteSpace()
}

func (lex *LexerStructure) ConsumeMultiLineComment() {
	Mult := false
	for !Mult {
		if lex.Char == '0' {
			Mult = true
		}
		if lex.Char == '*' && lex.Peek() == '/' {
			Mult = true
			lex.ReadChar()
		}
		lex.ReadChar()
	}
	lex.ConsumeWhiteSpace()
}
