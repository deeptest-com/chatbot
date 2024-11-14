package domain

import "github.com/deeptest-com/deeptest-next/internal/pkg/consts"

type InstructionDef []InstructionItem

type InstructionItem struct {
	Name  consts.TcInstructionType `json:"name"`
	Steps []InstructionStep        `json:"steps"`
}

type InstructionStep struct {
	Name            string                   `json:"name"`
	NextInstruction consts.TcInstructionType `json:"nextInstruction"`
	NextStep        string                   `json:"nextStep"`
}

type NlpResult struct {
	Instruction     consts.TcInstructionType `json:"instruct"`
	CurrStep        string                   `json:"currStep"`
	NextInstruction consts.TcInstructionType `json:"nextInstruction"`
	NextStep        string                   `json:"nextStep"`
}
