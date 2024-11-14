package domain

import "github.com/deeptest-com/deeptest-next/internal/pkg/consts"

type InstructionDef []InstructionItem

type InstructionItem struct {
	Name  consts.TcInstructionType `json:"name"`
	Steps []string                 `json:"steps"`
}

type NlpResult struct {
	Instruction     consts.TcInstructionType `json:"instruct"`
	NextInstruction consts.TcInstructionType `json:"nextInstruction"`
	CurrStep        string                   `json:"currStep"`
	NextStep        string                   `json:"nextStep"`
}
