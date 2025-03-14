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
}

func NewInstruction(op OpCodeType) *Instruction {
	return &Instruction{
		Op: op,
	}
}
