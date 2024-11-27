package v1

import "github.com/deeptest-com/deeptest-next/internal/pkg/consts"

type TcNlpReq struct {
	Statement       string                   `json:"statement"`
	CurrInstruction consts.TcInstructionType `json:"currInstruction"`
	CurrStep        string                   `json:"currStep"`
}

type TcNlpResp struct {
	Category        consts.TcInstructionCategory `json:"category,omitempty"`
	CurrInstruction consts.TcInstructionType     `json:"currInstruction,omitempty"`
	CurrStep        string                       `json:"currStep,omitempty"`

	NextInstruction consts.TcInstructionType `json:"nextInstruction,omitempty"`
	NextStep        string                   `json:"nextStep,omitempty"`

	Slots []TcNlpSlot `json:"slots,omitempty"`
}
type TcNlpResult struct {
	Instruction consts.TcInstructionType `json:"instruction,omitempty"`
	Slots       []TcNlpSlot              `json:"slots,omitempty"`
}

type TcNlpSlot struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type TcCacheReq struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

type TcCacheResp struct {
}
