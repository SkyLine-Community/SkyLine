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
