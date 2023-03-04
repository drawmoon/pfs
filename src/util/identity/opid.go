package identity

type OpId struct {
	Real   uint64
	IsDir  bool
	IsRoot bool
}

func ParseOpId(str string) (OpId, error) {
	return OpId{}, nil
}
