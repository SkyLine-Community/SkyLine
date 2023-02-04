package SkyLine

import (
	"bytes"
	"fmt"
	"strings"
)

func (For *ForLoop) String() string {
	var buf bytes.Buffer

	buf.WriteString("for (")
	buf.WriteString(For.Condition.String())
	buf.WriteString(")\n{")
	for _, x := range For.Body.Statements {
		buf.WriteString("\n")
		buf.WriteString(x.String())
	}
	buf.WriteString("\n}")
	return buf.String()
}

func (while *WhileCondition) String() string {
	var Out bytes.Buffer
	Out.WriteString("while")
	Out.WriteString(while.Condition.String())
	return Out.String()
}

func (ml *MacroLiteral) String() string {
	var Out bytes.Buffer

	params := make([]string, 0, len(ml.Parameters))
	for _, param := range ml.Parameters {
		params = append(params, param.String())
	}

	Out.WriteString(ml.TokenLiteral())
	Out.WriteString("(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") ")
	Out.WriteString(ml.Body.String())

	return Out.String()
}

func (hl *HashLiteral) String() string {
	if hl == nil {
		return ""
	}

	pairs := make([]string, len(hl.Pairs))
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+": "+value.String())
	}

	var Out bytes.Buffer
	Out.WriteString("{")
	Out.WriteString(strings.Join(pairs, ", "))
	Out.WriteString("}")
	return Out.String()
}

func (ie *IndexExpression) String() string {
	if ie == nil {
		return ""
	}

	var Out bytes.Buffer

	Out.WriteString("(")
	Out.WriteString(ie.Left.String())
	Out.WriteString("[")
	Out.WriteString(ie.Index.String())
	Out.WriteString("])")

	return Out.String()
}

func (al *ArrayLiteral) String() string {
	if al == nil {
		return ""
	}

	elements := make([]string, 0, len(al.Elements))
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	var Out bytes.Buffer

	Out.WriteString("[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")

	return Out.String()
}

func (sl *StringLiteral) String() string {
	return sl.TokenLiteral()
}

func (ce *CallExpression) String() string {
	var Out bytes.Buffer

	args := make([]string, 0, len(ce.Arguments))
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	Out.WriteString(ce.Function.String())
	Out.WriteString("(")
	Out.WriteString(strings.Join(args, ", "))
	Out.WriteString(")")

	return Out.String()
}

func (fl *FunctionLiteral) String() string {
	var Out bytes.Buffer

	params := make([]string, 0, len(fl.Parameters))
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}

	Out.WriteString(fl.TokenLiteral())
	Out.WriteString("(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") ")
	Out.WriteString(fl.Body.String())

	return Out.String()
}

func (bs *BlockStatement) String() string {
	var Out bytes.Buffer

	for _, s := range bs.Statements {
		Out.WriteString(s.String())
	}

	return Out.String()
}

func (ie *ConditionalExpression) String() string {
	var Out bytes.Buffer

	Out.WriteString("if")
	Out.WriteString(ie.Condition.String())
	Out.WriteString(" ")
	Out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		Out.WriteString("else ")
		Out.WriteString(ie.Alternative.String())
	}

	return Out.String()
}

func (b *Boolean_AST) String() string {
	return b.Token.Literal
}

func (ie *InfixExpression) String() string {
	var Out bytes.Buffer

	Out.WriteString("(")
	Out.WriteString(ie.Left.String())
	Out.WriteString(" " + ie.Operator + " ")
	Out.WriteString(ie.Right.String())
	Out.WriteString(")")

	return Out.String()
}

func (pe *PrefixExpression) String() string {
	var Out bytes.Buffer

	Out.WriteString("(")
	Out.WriteString(pe.Operator)
	Out.WriteString(pe.Right.String())
	Out.WriteString(")")

	return Out.String()
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (ul *UnsignedLiteral) String() string {
	return ul.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (importvalue *Import) String() string {
	var Out bytes.Buffer
	Out.WriteString("import (")
	Out.WriteString(importvalue.ImportValue.String())
	Out.WriteString(")")
	return Out.String()
}

func (note *Note) String() string {
	var Out bytes.Buffer
	Out.WriteString("!*")
	for _, l := range note.Value {
		Out.WriteString(fmt.Sprint(l))
	}
	Out.WriteString("\n*!")
	return Out.String()
}

func (cs *Carrier) String() string {
	var Out bytes.Buffer

	Out.WriteString(cs.TokenLiteral() + " ")

	if cs.CarrierValue != nil {
		Out.WriteString(cs.CarrierValue.String())
	}

	Out.WriteString(";")

	return Out.String()
}

func (rs *ReturnStatement) String() string {
	var Out bytes.Buffer

	Out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		Out.WriteString(rs.ReturnValue.String())
	}

	Out.WriteString(";")

	return Out.String()
}

func (i *Ident) String() string {
	return i.Value
}

func (ls *LetStatement) String() string {
	var Out bytes.Buffer

	Out.WriteString(ls.TokenLiteral() + " ")
	Out.WriteString(ls.Name.String())
	Out.WriteString(" = ")

	if ls.Value != nil {
		Out.WriteString(ls.Value.String())
	}

	Out.WriteString(";")

	return Out.String()
}

func (prog *Program) String() string {
	var Out bytes.Buffer

	for _, s := range prog.Statements {
		Out.WriteString(s.String())
	}

	return Out.String()
}

func (Switch *Switch) String() string {
	var Out bytes.Buffer
	Out.WriteString("\nswitch ( ")
	Out.WriteString(Switch.Value.String())
	Out.WriteString(")\n{\n")
	for _, bod := range Switch.Choices {
		if bod != nil {
			Out.WriteString(bod.String())
		}
	}
	Out.WriteString("}\n")
	return Out.String()
}

func (Case *Case) String() string {
	var Out bytes.Buffer
	if Case.Def {
		Out.WriteString("defualt")
	} else {
		Out.WriteString("case ")
		bod := []string{}
		for _, body := range Case.Expr {
			bod = append(bod, body.String())
		}
		Out.WriteString(strings.Join(bod, ","))
	}
	Out.WriteString(Case.Block.String())
	return Out.String()
}
