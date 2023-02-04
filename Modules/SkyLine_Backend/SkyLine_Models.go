package SkyLine

///////////////////////////////////////////////////////////////////////////
//// 						Tokenizer types                          ////

type Token_Type string

type Token struct {
	Token_Type Token_Type
	Literal    string
}

const (
	FOR             = "FOR"      // For loop 			 | Not implemeneted
	LTEQ            = "<="       // LT or equal to     | Implemented
	GTEQ            = ">="       // GT or equal to     | Implemented
	ILLEGAL         = "ILLEGAL"  // Illegal character  | Implemented
	EOF             = "EOF"      // End Of File        | Implemented
	IDENT           = "IDENT"    // Identifier         | Implimented
	INT             = "INT"      // TYPE integer       | Implemented
	FLOAT           = "FLOAT"    // TYPE float         | Implemented
	STRING          = "STRING"   // TYPE string        | Implemented
	UNSIGNED        = "UNSIGNED" // TYPE unsigned int  | Implemented
	CONST           = "CONST"    // Constant
	REGEXP          = "REGEXP"   // Open | REGEX data type
	FUNCTION        = "FUNCTION" // Function           | Implemented
	LET             = "LET"      // let statement      | Implemented
	TRUE            = "TRUE"     // boolean type true  | Implemented
	FALSE           = "FALSE"    // boolean type false | Implemented
	IF              = "IF"       // If statement       | Implemented
	ELSE            = "ELSE"     // Else statement     | Implemented
	RETURN          = "RETURN"   // return statement   | Implemented
	MACRO           = "MACRO"    // Macro statement    | Implemented but bugged
	CARRIER         = "CARRIER"  // Carry statement    | Implemented
	IMPORT          = "IMPORT"   // Import statement   | Implemented
	SWITCH          = "SWITCH"   // Switch statement   | Not implemented
	CASE            = "CASE"     // Case statement 	 | Not implemented
	DEFAULT         = "DEFAULT"  // Default keyword 	 | Not implemented
	WHILE           = "WHILE"    // Declare while loop | Implemented
	BREAK           = "BREAK"    // Break statement    | Not implemented
	ASTERISK_EQUALS = "*="       // Multiply equals
	BACKTICK        = "`"        // Backtick
	BANG            = "!"        // Boolean operator   | Implemented
	ASSIGN          = "="        // General assignment | Implemented
	PLUS            = "+"        // General operator   | Implemented
	MINUS           = "-"        // General operator   | Implemented
	MODULO          = "%"        // General operator   | Not Implemented
	ASTARISK        = "*"        // General operator   | Implemented
	POWEROF         = "**"       // General operator   | NOT Implemented
	SLASH           = "/"        // General operator   | Implemented
	LT              = "<"        // Boolean operator   | Implemented
	GT              = ">"        // Boolean operator   | Implemented
	EQ              = "=="       // Boolean operator   | Implemented
	MINUS_EQUALS    = "-="       // Minus equals
	MINUS_MINUS     = "--"       // Minus minux | Implemented
	OROR            = "||"       // Condition or or    | Not implemented
	ANDAND          = "&&"       // Boolean operator   | Not implemented
	NEQ             = "!="       // Boolean operator   | Implemented
	DIVEQ           = "/="       // Division operator  | Not Implemented
	NOTCONTAIN      = "!~"       // Boolean operator | Does not contain
	PERIOD          = "."        // Literally just a period
	PLUS_EQUALS     = "+="       // Plus equals
	PLUS_PLUS       = "++"       // Plus Plus
	QUESTION        = "?"        // Question que
	DOTDOT          = ".."       // Range
	CONTAINS        = "~="       // Contains
	COMMA           = ","        // Seperation         | Implemented
	SEMICOLON       = ";"        // SemiColon          | Implemented
	COLON           = ":"        // Colon              | Implemented
	LPAREN          = "("        // Args start         | Implemented
	RPAREN          = ")"        // Args end           | Implemented
	LINE            = "|"        // Line con           | Implemented
	LBRACE          = "{"        // Open  f            | Implemented
	RBRACE          = "}"        // Close f            | Implemented
	LBRACKET        = "["        // Open               | Implemented
	RBRACKET        = "]"        // Close              | Implemented

)

var keywords = map[string]Token_Type{
	"Func":     FUNCTION, // Function
	"function": FUNCTION, // Function
	"let":      LET,      // Variable declaration
	"cause":    LET,      // Variable declaration
	"allow":    LET,      // Variable declaration
	"true":     TRUE,     // Boolean true
	"false":    FALSE,    // Boolean false
	"if":       IF,       // Conditional start
	"else":     ELSE,     // Conditional alternative
	"return":   RETURN,   // Return decl
	"ret":      RETURN,   // Return decl
	"macro":    MACRO,    // Macro statement
	"carry":    CARRIER,  // Carry / exec
	"import":   IMPORT,   // Import statement
	"require":  IMPORT,   // Import statement
	"include":  IMPORT,   // Import statement
	"switch":   SWITCH,   // Switch statement
	"sw":       SWITCH,   // Switch statement
	"case":     CASE,     // Case statement
	"cs":       CASE,     // Case statement
	"default":  DEFAULT,  // Switch alternative
	"df":       DEFAULT,  // Switch alternative
	"while":    WHILE,    // While statement
	"break":    BREAK,    // Break statement
	"for":      FOR,      // For statement
}

///////////////////////////////////////////////////////////////////////////
////                         Lexer Types                               ////

type Lexer interface {
	NT() Token
}

type LexerStructure struct {
	CharInput string
	POS       int
	RPOS      int
	Char      byte
}

/////////////////////////////////////////////////////////////////////////////
////  						Object Types                                ////

type Type_Object string

// Data types
const (
	IntegerType     = "Integer"
	FloatType       = "Float"
	BooleanType     = "Boolean"
	NilType         = "Nil"
	ReturnValueType = "ReturnValue"
	ErrorType       = "Error"
	FunctionType    = "Function"
	StringType      = "String"
	BuiltinType     = "Builtin"
	ArrayType       = "Array"
	HashType        = "Hash"
	QuoteType       = "Quote"
	MacroType       = "Macro"
	CarrierType     = "Carrier"
	ImportingType   = "Importing"
	UnsignedType    = "Unsigned"
)

type Switch struct {
	Token   Token
	Value   Expression
	Choices []*Case
}

type Case struct {
	Token Token
	Def   bool
	Expr  []Expression
	Block *BlockStatement
}

type Object interface {
	Type_Object() Type_Object
	Inspect() string
}

type HashKey struct {
	Type_Object Type_Object
	Value       uint64
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value int64
}

type Unsigned struct {
	Value uint64
}

type Float struct {
	Value float64
}

type Boolean_Object struct {
	Value bool
}

type String struct {
	Value string
}

type Nil struct{}

type CarrierValue struct {
	Value Object
}

type ImporterValue struct {
	Value Object
}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        Environment
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []Object
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Quote struct {
	Node
}

type Macro struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        Environment
}

//////////////////////////////////////////////////////////////////////////////////////
/// 									ABSTRACT SYNTAX TREE MODELS				/////

var U UserInterpretData

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	SN()
}

type Expression interface {
	Node
	EN()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token Token
	Name  *Ident
	Value Expression
}

type AssignmentStatement struct {
	Token    Token
	Name     *Ident
	Value    Expression
	Operator string
}

type Ident struct {
	Token Token
	Value string
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

// Expression node
type Carrier struct {
	Token        Token
	CarrierValue Expression
}

// Expression node for importing files
type Import struct {
	Token       Token
	ImportValue Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

type UnsignedLiteral struct {
	Token Token
	Value uint64
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

type FloatLiteral struct {
	Token Token
	Value float64
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean_AST struct {
	Token Token
	Value bool
}

// Note

type Note struct {
	Token Token
	Value []string
}

type ConditionalExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type ForLoop struct {
	Token     Token
	Condition Expression
	Body      *BlockStatement
}

type WhileCondition struct {
	Token     Token
	Condition Expression
	Body      *BlockStatement
}

type BlockStatement struct {
	Token      Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

type CallExpression struct {
	Token     Token      // the '(' token
	Function  Expression // Ident or FunctionLiteral
	Arguments []Expression
}

type StringLiteral struct {
	Token Token
	Value string
}

type ArrayLiteral struct {
	Token    Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token Token // the '[' token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token Token // the '{' token
	Pairs map[Expression]Expression
}

type MacroLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

//////////////////////////////////////////////////////////////////////////////////////
///                       OBJECT ENVIRONMENT MODELS AND STRUCTURES              /////

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object) Object
}

type Environment_of_environment struct {
	Store map[string]Object
	Outer Environment
}

//////////////////////////////////////////////////////////////////////////////////////
///                       			Evaluation Models                           /////

var (
	NilValue   = &Nil{}
	TrueValue  = &Boolean_Object{Value: true}
	FalseValue = &Boolean_Object{Value: false}
)

const (
	FuncNameQuote   = "quote"
	FuncNameUnquote = "unquote"
)

var FileCurrent FileCurrentWithinParserEnvironment

//////////////////////////////////////////////////////////////////////////////////////
/// 									PARSER MODELS							 ///

const (
	_ int = iota
	LOWEST
	ASSIGNMENT   // =
	EQUALS       // ==
	LESSGREATER  // > or <
	SUM          // +
	PRODUCT      // *
	PREFIX       // -X or !X
	CALL         // myFunc(X)
	INDEX        // array[index]
	GTEQS        // greater than or equal to
	LTEQS        // Less than or equal to
	POWER        // Power of
	MOD          // Modulos
	DOT_DOT      // .. | DOTDOTs
	REGEXP_MATCH // !~ ~=
	TERNARY      // ? or :
	COND         // OR || AND

)

var Precedences = map[Token_Type]int{
	EQ:              EQUALS,
	GTEQ:            GTEQS,
	LTEQ:            LTEQS,
	NEQ:             EQUALS,
	LT:              LESSGREATER,
	GT:              LESSGREATER,
	PLUS:            SUM,
	PLUS_EQUALS:     SUM,
	MINUS:           SUM,
	MINUS_EQUALS:    SUM,
	SLASH:           PRODUCT,
	ASSIGN:          ASSIGNMENT,
	POWEROF:         POWER,
	QUESTION:        TERNARY,
	ASTARISK:        PRODUCT,
	ASTERISK_EQUALS: PRODUCT,
	DIVEQ:           PRODUCT,
	LPAREN:          CALL,
	ANDAND:          COND,
	OROR:            COND,
	PERIOD:          CALL,
	NOTCONTAIN:      REGEXP_MATCH,
	CONTAINS:        REGEXP_MATCH,
	DOTDOT:          DOT_DOT,
	LBRACKET:        INDEX,
}

type (
	PrefixParseFn  func() Expression
	InfixParseFn   func(Expression) Expression
	PostfixParseFn func() Expression
)

type Parser struct {
	Lex            Lexer
	Errors         []string
	CurrentToken   Token
	PeekToken      Token
	PrefixParseFns map[Token_Type]PrefixParseFn
	InfixParseFns  map[Token_Type]InfixParseFn
}
