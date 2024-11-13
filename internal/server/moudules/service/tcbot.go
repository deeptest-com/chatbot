package service

import (
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
)

type TcbotService struct {
}

func (s *TcbotService) CreatePart(req v1.TcNlpReq) (ret v1.TcNlpResp, err error) {
	instruction := consts.TcInstructionUnknown
	currStep := ""
	nextStep := ""

	if req.Statement == "create part" || req.Statement == "create_part" { // TODO: check with llm
		instruction = consts.TcInstructionCreatePart
		currStep = "init"
		nextStep = "input_part_no"
	} else {
		instruction = req.CurrInstruction
		currStep = req.CurrStep

		if req.CurrStep == "input_part_no" {
			nextStep = "input_part_name"
		} else if req.CurrStep == "input_part_name" {
			nextStep = "show_part_form"
		} else if req.CurrStep == "show_part_form" {
			nextStep = "end"
		}
	}

	ret = v1.TcNlpResp{
		Category:        consts.TcCategoryInstruction,
		CurrInstruction: instruction,

		CurrStep: currStep,
		NextStep: nextStep,

		Params: nil,
	}

	return
}
