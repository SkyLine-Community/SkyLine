package SkyLine

func LookupIdentifier(ident string) Token_Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
