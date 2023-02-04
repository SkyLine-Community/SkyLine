package SkyLine

// All statement nodes
func (ls *LetStatement) SN()        {} // Statement Node | Allow condition
func (rs *ReturnStatement) SN()     {} // Statement Node | Allow Return
func (es *ExpressionStatement) SN() {} // Statement Node | Allow Expression
func (cs *Carrier) SN()             {} // Statement Node | Allow Carrier / Importer
func (importvalue *Import) SN()     {} // Statement Node | Allow import
func (For *ForLoop) SN()            {} // Statement Node | Allow For loops
