package service

import (
	"encoding/json"
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
	"github.com/deeptest-com/deeptest-next/internal/pkg/domain"
	_logUtils "github.com/deeptest-com/deeptest-next/pkg/libs/log"
)

type TcbotService struct {
	InstructionDef *domain.InstructionDef
}

func (s *TcbotService) CreatePart(req v1.TcNlpReq) (ret v1.TcNlpResp, err error) {
	nlpResult, _ := s.NlpParse(req)

	if nlpResult.Instruction == "" { // llm not known, use the one in request
		nlpResult.Instruction = req.CurrInstruction
		nlpResult.CurrStep = req.CurrStep
	}

	nlpResult.NextInstruction, nlpResult.NextStep = s.GetNextStep(nlpResult.Instruction, nlpResult.CurrStep)

	ret = v1.TcNlpResp{
		Category:        consts.TcCategoryInstruction,
		CurrInstruction: nlpResult.Instruction,

		CurrStep: nlpResult.CurrStep,

		NextInstruction: nlpResult.NextInstruction,
		NextStep:        nlpResult.NextStep,

		Params: nil,
	}

	return
}

func (s *TcbotService) GetNextStep(instruction consts.TcInstructionType, step string) (
	nextInstruction consts.TcInstructionType, nextStep string) {

	instructionDef := s.GetInstructionDef()

	for _, instructionItem := range *instructionDef {
		if instructionItem.Name == instruction {
			nextInstruction = instruction // default is current instruction

			for stepIndex, stepItem := range instructionItem.Steps {
				if stepItem.Name != step {
					continue
				}

				if stepItem.NextInstruction != "" { // next step is set
					nextInstruction = stepItem.NextInstruction
					nextStep = stepItem.NextStep

				} else if stepIndex < len(instructionItem.Steps)-1 { // use next step
					nextStep = instructionItem.Steps[stepIndex+1].Name

				}

				break
			}

			break
		}
	}

	return
}

func (s *TcbotService) NlpParse(req v1.TcNlpReq) (ret domain.NlpResult, err error) {
	if req.Statement == "create part" || req.Statement == "create_part" { // TODO: parse by llm
		ret.Instruction = consts.TcInstructionCreatePart
		ret.CurrStep = "init"
	}

	return
}

func (s *TcbotService) GetInstructionDef() *domain.InstructionDef {
	if s.InstructionDef == nil {
		instructionDef := domain.InstructionDef{}
		s.InstructionDef = &instructionDef

		err := json.Unmarshal([]byte(consts.InstructionDef), s.InstructionDef)
		if err != nil {
			_logUtils.Info(err.Error())
		}
	}

	return s.InstructionDef
}
