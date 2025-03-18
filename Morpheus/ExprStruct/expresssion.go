package expr

type Expr struct {
	Left     *Expr
	Operator any // SWITCH: Token
	Right    *Expr
}

func NewExpr(left *Expr, operator any, right *Expr) Expr { // SWITCH: Token
	return Expr{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}
