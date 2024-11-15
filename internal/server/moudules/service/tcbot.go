package service

import (
	"encoding/json"
	"github.com/deeptest-com/deeptest-next"
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
	"github.com/deeptest-com/deeptest-next/internal/pkg/domain"
	_logUtils "github.com/deeptest-com/deeptest-next/pkg/libs/log"
	"strings"
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
	if instructionDef == nil {
		return
	}

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

	if nextStep == "" {
		nextInstruction = ""
	}

	return
}

func (s *TcbotService) NlpParse(req v1.TcNlpReq) (ret domain.NlpResult, err error) {
	statement := strings.ReplaceAll(req.Statement, " ", "_")

	switch statement { // TODO: parse by llm
	case "create_part":
		ret.Instruction = consts.TcInstructionCreatePart
		ret.CurrStep = "init"
	case "attach_material":
		ret.Instruction = consts.TcInstructionAttachMaterial
		ret.CurrStep = "init"
	case "attach_geometry":
		ret.Instruction = consts.TcInstructionAttachGeometry
		ret.CurrStep = "init"
	case "create_st":
		ret.Instruction = consts.TcInstructionCreateSt
		ret.CurrStep = "init"
	case "assign_project":
		ret.Instruction = consts.TcInstructionAssignProject
		ret.CurrStep = "init"
	case "check_data":
		ret.Instruction = consts.TcInstructionCheckData
		ret.CurrStep = "init"
	case "freeze_st":
		ret.Instruction = consts.TcInstructionFreezeSt
		ret.CurrStep = "init"
	case "submit_st":
		ret.Instruction = consts.TcInstructionSubmitSt
		ret.CurrStep = "init"
	case "track_st":
		ret.Instruction = consts.TcInstructionTrackSt
		ret.CurrStep = "init"
	}

	return
}

func (s *TcbotService) GetInstructionDef() *domain.InstructionDef {
	if s.InstructionDef == nil {
		instructionDef := domain.InstructionDef{}
		s.InstructionDef = &instructionDef

		bytes, err := deeptest.ReadResData("res/instruction-def.json")
		if err != nil {
			_logUtils.Info(err.Error())
			return nil
		}

		err = json.Unmarshal(bytes, s.InstructionDef)
		if err != nil {
			_logUtils.Info(err.Error())
			return nil
		}
	}

	return s.InstructionDef
}
