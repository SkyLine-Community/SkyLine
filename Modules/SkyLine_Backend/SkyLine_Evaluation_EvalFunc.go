package SkyLine

import (
	"fmt"
	"os"
	"strconv"
)

func Eval(node Node, Env Environment) Object {
	switch node := node.(type) {
	case *Program:
		return evalProgram(node, Env)
	case *ExpressionStatement:
		return Eval(node.Expression, Env)
	case *Carrier:
		value := Eval(node.CarrierValue, Env)
		if isError(value) {
			return value
		}
		return &CarrierValue{
			Value: value,
		}
	case *ReturnStatement:
		value := Eval(node.ReturnValue, Env)
		if isError(value) {
			return value
		}
		return &ReturnValue{Value: value}
	case *Import:
		value := Eval(node.ImportValue, Env)
		if isError(value) {
			return value
		}
		return &ImporterValue{
			Value: value,
		}
	case *BlockStatement:
		return evalBlockStatement(node, Env)
	case *Switch:
		return EvalSwitch(node, Env)
	case *LetStatement:
		value := Eval(node.Value, Env)
		if isError(value) {
			return value
		}
		Env.Set(node.Name.Value, value)

	// Expressions
	case *UnsignedLiteral:
		return &Unsigned{Value: node.Value}
	case *IntegerLiteral:
		return &Integer{Value: node.Value}

	case *FloatLiteral:
		return &Float{Value: node.Value}

	case *Boolean_AST:
		return nativeBoolToBooleanObject(node.Value)

	case *PrefixExpression:
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *InfixExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *WhileCondition:
		return EvalWhileCondition(node, Env)
	case *ConditionalExpression:
		return evalIfExpression(node, Env)

	case *Ident:
		return evalIdent(node, Env)

	case *FunctionLiteral:
		return &Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        Env,
		}
	case *CallExpression:
		if node.Function.TokenLiteral() == FuncNameQuote {
			return quote(node.Arguments[0], Env)
		}

		function := Eval(node.Function, Env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, Env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

	case *StringLiteral:
		return &String{Value: node.Value}

	case *ArrayLiteral:
		elems := evalExpressions(node.Elements, Env)
		if len(elems) == 1 && isError(elems[0]) {
			return elems[0]
		}
		return &Array{Elements: elems}

	case *IndexExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, Env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *HashLiteral:
		return evalHashLiteral(node, Env)
	}

	return nil
}

func evalProgram(program *Program, Env Environment) Object {
	var result Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, Env)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *Boolean_Object {
	if input {
		return TrueValue
	}
	return FalseValue
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		// Eventually add support for prefix and infix seperate errors
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator))
	}
}

func evalBangOperatorExpression(right Object) Object {
	if right == NilValue || right == FalseValue {
		return TrueValue
	}
	return FalseValue
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	switch right := right.(type) {
	case *Integer:
		return &Integer{Value: -right.Value}
	case *Float:
		return &Float{Value: -right.Value}
	default:
		return newError("unknown operator: -%s", right.Type_Object())
	}
}

func evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type_Object() == IntegerType && right.Type_Object() == IntegerType:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type_Object() == FloatType || right.Type_Object() == FloatType:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type_Object() == StringType && right.Type_Object() == StringType:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type_Object() != right.Type_Object():
		return newError(ErrorSymBolMap[CODE_PARSE_TYPE_ERROR](string(left.Type_Object()), string(right.Type_Object()), operator))
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalFloatInfixExpression(operator string, left, right Object) Object {
	var leftVal, rightVal float64

	switch left := left.(type) {
	case *Integer:
		leftVal = float64(left.Value)
	case *Float:
		leftVal = left.Value
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}

	switch right := right.(type) {
	case *Integer:
		rightVal = float64(right.Value)
	case *Float:
		rightVal = right.Value
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}

	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalStringInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalBlockStatement(block *BlockStatement, Env Environment) Object {
	var result Object

	for _, stmt := range block.Statements {
		result = Eval(stmt, Env)
		if result == nil {
			continue
		}

		if rt := result.Type_Object(); rt == ReturnValueType || rt == ErrorType {
			return result
		}
	}

	return result
}

func EvalWhileCondition(while *WhileCondition, Env Environment) Object {
	condition := Eval(while.Condition, Env)
	if isError(condition) {
		return condition
	}
	for isTruthy(condition) {
		Eval(while.Body, Env)
	}
	return NilValue
}

func evalIfExpression(ie *ConditionalExpression, Env Environment) Object {
	condition := Eval(ie.Condition, Env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, Env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, Env)
	}
	return NilValue
}

func isTruthy(obj Object) bool {
	return obj != NilValue && obj != FalseValue
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	return obj != nil && obj.Type_Object() == ErrorType
}

func evalIdent(node *Ident, Env Environment) Object {
	if val, ok := Env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	} else if builtin, ok := FInfo[node.Value]; ok {
		return builtin
	}
	return newError(ErrorSymBolMap[CODE_PARSE_IDENTIFIER_ERROR](node.Value))
}

func evalExpressions(exprs []Expression, Env Environment) []Object {
	result := make([]Object, 0, len(exprs))

	for _, expr := range exprs {
		evaluated := Eval(expr, Env)
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func extendFunctionEnv(fn *Function, args []Object) Environment {
	Env := NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		Env.Set(param.Value, args[i])
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(ErrorSymBolMap[CODE_PARSE_FUNCTION_ARGUMENTS_NOT_ENOUGH_ERROR](fmt.Sprint(fn.Inspect()), fmt.Sprint(len(args)), fmt.Sprint(len(fn.Parameters))))
				os.Exit(0)
			}
		}()
	}
	return Env
}

func applyFunction(fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *Builtin:
		return fn.Fn(args...)
	default:
		// expecting crash if it is not a function type object | Macros are bugged
		defer func() {
			if x := recover(); x != nil {
				fmt.Print("\n\n Assuming this is a macro, there was something wrong. Macro does not exist")
				os.Exit(1)
			}
		}()
		return newError(ErrorSymBolMap[CODE_PREFIX_PARSE_FUNCTION_INVALID_OR_UNFOUND_WITHIN_PARSER_AND_INTERPRETRR](string(fn.Type_Object())))
	}
}

func unwrapReturnValue(obj Object) Object {
	if returnValue, ok := obj.(*ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalIndexExpression(left, index Object) Object {
	switch {
	case left.Type_Object() == ArrayType && index.Type_Object() == IntegerType:
		return evalArrayIndexExpression(left, index)
	case left.Type_Object() == HashType:
		return evalHashIndexExpression(left, index)
	default:
		return newError(ErrorSymBolMap[CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED](string(left.Type_Object())))
	}
}

func evalArrayIndexExpression(array, index Object) Object {
	arrObj := array.(*Array)
	idx := index.(*Integer).Value
	max := int64(len(arrObj.Elements) - 1)

	if idx < 0 || idx > max {
		return NilValue
	}

	return arrObj.Elements[idx]
}

func evalHashLiteral(node *HashLiteral, Env Environment) Object {
	pairs := make(map[HashKey]HashPair, len(node.Pairs))

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, Env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(Hashable)
		if !ok {
			return newError(ErrorSymBolMap[CODE_PARSE_HASHKEY_ERROR](string(key.Type_Object())))
		}

		value := Eval(valueNode, Env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &Hash{Pairs: pairs}
}

func evalHashIndexExpression(left, index Object) Object {
	key, ok := index.(Hashable)
	if !ok {
		return newError(ErrorSymBolMap[CODE_PARSE_HASHKEY_ERROR](string(index.Type_Object())))
	}

	hashObj := left.(*Hash)
	if pair, exists := hashObj.Pairs[key.HashKey()]; exists {
		return pair.Value
	}
	return NilValue
}

func DefineMacros(program *Program, Env Environment) {
	defs := make([]int, 0)
	stmts := program.Statements

	for pos, stmt := range stmts {
		if isMacroDefinition(stmt) {
			addMacro(stmt, Env)
			defs = append(defs, pos)
		}
	}

	for i := len(defs) - 1; i >= 0; i-- {
		pos := defs[i]
		program.Statements = append(stmts[:pos], stmts[pos+1:]...)
	}
}

func isMacroDefinition(node Statement) bool {
	letStmt, ok := node.(*LetStatement)
	if !ok {
		return false
	}

	_, ok = letStmt.Value.(*MacroLiteral)
	return ok
}

func addMacro(stmt Statement, Env Environment) {
	letStmt := stmt.(*LetStatement)
	macroLit := letStmt.Value.(*MacroLiteral)
	macro := &Macro{
		Parameters: macroLit.Parameters,
		Env:        Env,
		Body:       macroLit.Body,
	}
	Env.Set(letStmt.Name.Value, macro)
}

func ExpandMacros(program Node, Env Environment) Node {
	modifier := func(node Node) Node {
		call, ok := node.(*CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(call, Env)
		if !ok {
			return node
		}

		args := quoteArgs(call)
		evalEnv := extendMacroEnv(macro, args)

		quote, ok := Eval(macro.Body, evalEnv).(*Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return quote.Node
	}

	return Modify(program, modifier)
}

func isMacroCall(call *CallExpression, Env Environment) (macro *Macro, ok bool) {
	ident, ok := call.Function.(*Ident)
	if !ok {
		return nil, false
	}

	obj, ok := Env.Get(ident.Value)
	if !ok {
		return nil, false
	}

	macro, ok = obj.(*Macro)
	return macro, ok
}

func quoteArgs(call *CallExpression) []*Quote {
	args := make([]*Quote, 0, len(call.Arguments))
	for _, arg := range call.Arguments {
		args = append(args, &Quote{Node: arg})
	}
	return args
}

func extendMacroEnv(macro *Macro, args []*Quote) Environment {
	extended := NewEnclosedEnvironment(macro.Env)
	for i, param := range macro.Parameters {
		extended.Set(param.Value, args[i])
	}
	return extended
}

func quote(node Node, Env Environment) Object {
	node = evalUnquoteCalls(node, Env)
	return &Quote{Node: node}
}

func evalUnquoteCalls(quoted Node, Env Environment) Node {
	modifier := func(node Node) Node {
		call, ok := node.(*CallExpression)
		if !ok || call.Function.TokenLiteral() != FuncNameUnquote || len(call.Arguments) != 1 {
			return node
		}

		unquoted := Eval(call.Arguments[0], Env)
		return convertObjectToASTNode(unquoted)
	}
	return Modify(quoted, modifier)
}

func convertObjectToASTNode(obj Object) Node {
	switch obj := obj.(type) {
	case *Integer:
		base := 10
		t := Token{
			Token_Type: INT,
			Literal:    strconv.FormatInt(obj.Value, base),
		}
		return &IntegerLiteral{Token: t, Value: obj.Value}
	case *Boolean_Object:
		var t Token
		if obj.Value {
			t = Token{Token_Type: TRUE, Literal: "true"}
		} else {
			t = Token{Token_Type: FALSE, Literal: "false"}
		}
		return &Boolean_AST{Token: t, Value: obj.Value}
	case *Quote:
		return obj.Node
	default:
		return nil
	}
}

func EvalSwitch(Switch *Switch, Env Environment) Object {
	object := Eval(Switch.Value, Env)
	for _, options := range Switch.Choices {
		if options.Def {
			continue
		}
		for _, values := range options.Expr {
			output := Eval(values, Env)
			if object.Type_Object() == output.Type_Object() && (object.Inspect() == output.Inspect()) {
				Block := evalBlockStatement(options.Block, Env)
				return Block
			}
		}
	}
	for _, opt := range Switch.Choices {
		if opt.Def {
			output := evalBlockStatement(opt.Block, Env)
			return output
		}
	}
	return nil
}
