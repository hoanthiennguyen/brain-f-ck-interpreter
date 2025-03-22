package brainfuck

type OpCodeType int

const (
	OpIncr OpCodeType = iota
	OpDecr
	OpMoveRight
	OpMoveLeft
	OpInput
	OpOutput
	OpBeginLoop
	OpEndLoop
)

type Instruction struct {
	Op    OpCodeType
	Param int
	data  int
}

func (e OpCodeType) String() string {
	op := ""
	switch e {
	case OpIncr:
		op = "Incr"
	case OpDecr:
		op = "Decr"
	case OpMoveLeft:
		op = "MoveLeft"
	case OpMoveRight:
		op = "MoveRight"
	case OpBeginLoop:
		op = "BeginLoop"
	case OpEndLoop:
		op = "EndLoop"
	case OpOutput:
		op = "Output"
	}

	return op
}

func NewInstruction(op OpCodeType) *Instruction {
	return &Instruction{
		Op:    op,
		Param: 1,
	}
}

func (o OpCodeType) CanStack() bool {
	return o == OpIncr || o == OpDecr || o == OpMoveLeft || o == OpMoveRight
}

func NewInstructionWithParam(op OpCodeType, param int) *Instruction {
	return &Instruction{
		Op:    op,
		Param: param,
	}
}
