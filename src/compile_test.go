package brainfuck

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestCompile(t *testing.T) {
	type args struct {
		code string
	}

	tests := []struct {
		name string
		args args
		want []*Instruction
	}{
		{
			args: args{
				code: "++[--]-",
			},
			want: []*Instruction{
				NewInstructionWithParam(OpIncr, 2),
				NewInstructionWithParam(OpBeginLoop, 1),
				NewInstructionWithParam(OpDecr, 2),
				NewInstructionWithParam(OpEndLoop, 1),
				NewInstruction(OpDecr),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compileV2(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildMatchingBrackets(t *testing.T) {
	type args struct {
		ins []*Instruction
	}
	tests := []struct {
		name string
		args args
		want []*Instruction
	}{
		{
			args: args{
				ins: []*Instruction{
					NewInstructionWithParam(OpIncr, 2),
					NewInstructionWithParam(OpDecr, 2),
					NewInstructionWithParam(OpBeginLoop, 1),
					NewInstructionWithParam(OpDecr, 2),
					NewInstructionWithParam(OpEndLoop, 1),
					NewInstruction(OpDecr),
				},
			},
			want: []*Instruction{
				NewInstructionWithParam(OpIncr, 2),
				NewInstructionWithParam(OpDecr, 2),
				NewInstructionWithParam(OpBeginLoop, 4),
				NewInstructionWithParam(OpDecr, 2),
				NewInstructionWithParam(OpEndLoop, 2),
				NewInstruction(OpDecr),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildMatchingBrackets(tt.args.ins)
			got := tt.args.ins
			rawWant, _ := json.Marshal(tt.want)
			rawGot, _ := json.Marshal(got)
			fmt.Println(string(rawGot))

			if string(rawGot) != string(rawWant) {
				t.Errorf("BuildMatchingBrackets() = %s\n, want %s", rawGot, rawWant)
			}
		})
	}
}
