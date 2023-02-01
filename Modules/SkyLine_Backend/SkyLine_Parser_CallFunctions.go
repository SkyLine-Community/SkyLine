package SkyLine

import "fmt"

func (parser *Parser) ParserErrors() []string          { return parser.Errors }
func (parser *Parser) PeekTokenIs(typ Token_Type) bool { return parser.PeekToken.Token_Type == typ }

func (parser *Parser) CurrentTokenIs(typ Token_Type) bool {
	return parser.CurrentToken.Token_Type == typ
}

func (parser *Parser) NextToken() {
	parser.CurrentToken = parser.PeekToken
	parser.PeekToken = parser.Lex.NextToken()
}

func (parser *Parser) PeekError(typ Token_Type) {
	msg := ErrorSymBolMap[CODE_EXPECT_PEEK_ERROR_DURING_CALL_TO_PEEK](fmt.Sprint(typ), string(parser.PeekToken.Token_Type))
	//msg := fmt.Sprintf("expected next token to be %s, got %s instead", typ, parser.PeekToken.Token_Type)
	parser.Errors = append(parser.Errors, msg)
}

func (parser *Parser) ExpectPeek(typ Token_Type) bool {
	if parser.PeekTokenIs(typ) {
		parser.NextToken()
		return true
	}
	parser.PeekError(typ)
	return false
}

func (parser *Parser) ParseProgram() *Program {
	program := &Program{
		Statements: []Statement{},
	}

	for !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.NextToken()
	}

	return program
}
