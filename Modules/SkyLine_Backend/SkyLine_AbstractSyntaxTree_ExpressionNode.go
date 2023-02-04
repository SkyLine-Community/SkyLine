package SkyLine

// All expression nodes
func (i *Ident) EN()                  {} // Expression Node   | Identifier
func (ui *UnsignedLiteral) EN()       {} // Unsigned Literal  | Unsigned Literal
func (il *IntegerLiteral) EN()        {} // Expression Node   | Integer Literal
func (fl *FloatLiteral) EN()          {} // Expression Node   | Float Literal
func (pe *PrefixExpression) EN()      {} // Expression Node   | Prefix Expression
func (ie *InfixExpression) EN()       {} // Expression Node   | Infix Expression
func (b *Boolean_AST) EN()            {} // Expression Node   | Boolean
func (ie *ConditionalExpression) EN() {} // Expression Node   | Conditional Expression
func (bs *BlockStatement) EN()        {} // Expression Node   | Block
func (fl *FunctionLiteral) EN()       {} // Expression Node   | Function Literal
func (ce *CallExpression) EN()        {} // Expression Node   | Call
func (sl *StringLiteral) EN()         {} // Expression Node   | String
func (*ArrayLiteral) EN()             {} // Expression Node   | Array Literal
func (*IndexExpression) EN()          {} // Expression Node   | Index expression
func (*HashLiteral) EN()              {} // Expression Node   | Hash Literal
func (ml *MacroLiteral) EN()          {} // Expression Node   | Macro Literal
func (while *WhileCondition) EN()     {} // Expression Node   | While Condition
func (For *ForLoop) EN()              {} // Expression Node   | For Loop
func (Case *Case) EN()                {} // Expression Node   | Case Expression
func (Switch *Switch) EN()            {} // Expression Node   | Switch Expression
