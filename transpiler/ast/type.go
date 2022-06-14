package ast

type ASTTypeKind int

const (
	ASTTypeKindUnknown ASTTypeKind = iota
	ASTTypeKindInt
	ASTTypeKindFloat
)

type ASTType struct {
	Kind ASTTypeKind
}

func NewASTType(kind ASTTypeKind) *ASTType {
	return &ASTType{
		Kind: kind,
	}
}

func (t *ASTType) String() string {
	switch t.Kind {
	case ASTTypeKindInt:
		return "integer"
	case ASTTypeKindFloat:
		return "floating"
	default:
		return "unknown"
	}
}
