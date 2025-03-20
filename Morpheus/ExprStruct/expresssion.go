package expr

import token "NeoLang/Token"

type Expr struct {
	Left     *Expr
	Operator token.Type
	Right    *Expr
}

func NewExpr(left *Expr, operator token.Type, right *Expr) Expr {
	return Expr{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}
