package SkyLine

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (b *Builtin) Inspect() string                 { return "builtin function" }
func (s *String) Inspect() string                  { return s.Value }
func (q *Quote) Inspect() string                   { return fmt.Sprintf("%s(%s)", QuoteType, q.Node.String()) }
func (e *Error) Inspect() string                   { return e.Message }
func (cs *CarrierValue) Inspect() string           { return cs.Value.Inspect() }
func (importvalue *ImporterValue) Inspect() string { return importvalue.Value.Inspect() }
func (rv *ReturnValue) Inspect() string            { return rv.Value.Inspect() }
func (n *Nil) Inspect() string                     { return "nil" }
func (f *Float) Inspect() string                   { return strconv.FormatFloat(f.Value, 'f', -1, 64) }
func (i *Integer) Inspect() string                 { return strconv.FormatInt(i.Value, 10) }
func (ui *Unsigned) Inspect() string               { return strconv.FormatUint(ui.Value, 10) }
func (b *Boolean_Object) Inspect() string          { return strconv.FormatBool(b.Value) }
func (note *Note) Inspect() []string               { return note.Value }

func (f *Function) Inspect() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(f.Parameters))
	for _, parser := range f.Parameters {
		params = append(params, parser.String())
	}
	Out.WriteString(fmt.Sprint(f.Type_Object()) + "(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") {")
	Out.WriteString(f.Body.String())
	Out.WriteString("}")
	return Out.String()
}

func (a *Array) Inspect() string {
	if a == nil {
		return ""
	}
	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	var Out bytes.Buffer
	Out.WriteString("")
	Out.WriteString("[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")
	return Out.String()
}

func (h *Hash) Inspect() string {
	if h == nil {
		return ""
	}
	pairs := make([]string, 0, len(h.Pairs))
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+": "+pair.Value.Inspect())
	}
	var Out bytes.Buffer
	Out.WriteString("{")
	Out.WriteString(strings.Join(pairs, ", "))
	Out.WriteString("}")
	return Out.String()
}

func (m *Macro) Inspect() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(m.Parameters))
	for _, parser := range m.Parameters {
		params = append(params, parser.String())
	}
	Out.WriteString("macro(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") {\n")
	Out.WriteString(m.Body.String())
	Out.WriteString("\n}")
	return Out.String()
}
