package v1

import "github.com/deeptest-com/deeptest-next/internal/pkg/consts"

type TcNlpReq struct {
	Statement       string                   `json:"statement"`
	CurrInstruction consts.TcInstructionType `json:"currInstruction"`
	CurrStep        string                   `json:"currStep"`
}
type LlmNlpResp struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role      string      `json:"role"`
			Content   string      `json:"content"`
			ToolCalls interface{} `json:"tool_calls"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
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
