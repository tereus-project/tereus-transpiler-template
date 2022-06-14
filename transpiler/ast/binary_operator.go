package ast

type ASTBinaryOperatorKind int

const (
	ASTBinaryOperatorKindUnknown ASTBinaryOperatorKind = iota
	ASTBinaryOperatorKindAdd
	ASTBinaryOperatorKindSubtract
	ASTBinaryOperatorKindMultiply
	ASTBinaryOperatorKindDivide
)

type ASTBinaryOperator struct {
	Kind ASTBinaryOperatorKind
}

func NewASTBinaryOperator(kind ASTBinaryOperatorKind) *ASTBinaryOperator {
	return &ASTBinaryOperator{
		Kind: kind,
	}
}

func (o *ASTBinaryOperator) String() string {
	switch o.Kind {
	case ASTBinaryOperatorKindAdd:
		return "plus"
	case ASTBinaryOperatorKindSubtract:
		return "minus"
	case ASTBinaryOperatorKindMultiply:
		return "times"
	case ASTBinaryOperatorKindDivide:
		return "divided by"
	default:
		return "unknown"
	}
}
