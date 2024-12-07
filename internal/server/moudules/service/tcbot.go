package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest-next"
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/pkg/config"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
	"github.com/deeptest-com/deeptest-next/internal/pkg/domain"
	_http "github.com/deeptest-com/deeptest-next/pkg/libs/http"
	_logUtils "github.com/deeptest-com/deeptest-next/pkg/libs/log"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type TcbotService struct {
	InstructionDef *domain.InstructionDef
}

func (s *TcbotService) Index(req v1.TcNlpReq, ctx iris.Context) (ret v1.TcNlpResp, err error) {
	nlpResult, _ := s.NlpParse(req)

	if nlpResult.Instruction == "" { // llm not known, use the one in request
		nlpResult.Instruction = req.CurrInstruction
		nlpResult.CurrStep = req.CurrStep
	}

	var slots []v1.TcNlpSlot

	isMock := config.CONFIG.System.IsMock
	isFound := false
	nlpResult.NextInstruction, nlpResult.NextStep, isFound = s.GetNextStep(nlpResult.Instruction, nlpResult.CurrStep)
	if !isFound {
		//slots = []v1.TcNlpSlot{{Name: "ERROR", Value: fmt.Sprintf("%s - %s NOT FOUND", nlpResult.Instruction, nlpResult.CurrStep)}}
		//goto RETURN
	}

	if nlpResult.Instruction == "" { // no value, need to parse
		if isMock {
			if strings.Index(req.Statement, "Back Supply status") > -1 {
				nlpResult.Instruction = consts.TcInstructionTrackSt
			} else {
				nlpResult.Instruction = consts.TcInstructionGreetings
			}

		} else {
			nlpResult.Instruction, slots, err = s.ChatCompletion("", req.Statement)
		}

	} else if nlpResult.Instruction == consts.TcInstructionConfirm { // parse
		if isMock {
			slots = append(slots, v1.TcNlpSlot{
				Name:  "result",
				Value: true,
			})
		} else {
			_, slots, err = s.ChatCompletion(nlpResult.Instruction.String(), req.Statement)
		}

	} else if nlpResult.CurrStep == "input_materials" { // parse
		if isMock {
			slots = append(slots, v1.TcNlpSlot{
				Name:  "1",
				Value: "PA6+30GF",
			}, v1.TcNlpSlot{
				Name:  "2",
				Value: "LASW3",
			})
		} else {
			_, slots, err = s.ChatCompletion(nlpResult.CurrStep, req.Statement)
		}

	} else if nlpResult.CurrStep == "input_geometry" { // parse
		if isMock {
			slots = append(slots, v1.TcNlpSlot{
				Name:  "1",
				Value: "BBA-1047285",
			})
			slots = append(slots, v1.TcNlpSlot{
				Name:  "2",
				Value: "BBA-1047286",
			})
		} else {
			_, slots, err = s.ChatCompletion(nlpResult.CurrStep, req.Statement)
		}
	}

	if err != nil {
		slots = []v1.TcNlpSlot{{Name: "ERROR", Value: err.Error()}}
	}

	//RETURN:

	ret = v1.TcNlpResp{
		Category:        consts.TcCategoryInstruction,
		CurrInstruction: nlpResult.Instruction,

		CurrStep: nlpResult.CurrStep,

		NextInstruction: nlpResult.NextInstruction,
		NextStep:        nlpResult.NextStep,

		Slots: slots,
	}

	return
}

func (s *TcbotService) GetNextStep(instruction consts.TcInstructionType, step string) (
	nextInstruction consts.TcInstructionType, nextStep string, isFound bool) {

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

				isFound = true

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

func (s *TcbotService) ChatCompletion(tmpl, content string) (
	instruction consts.TcInstructionType, slots []v1.TcNlpSlot, err error) {

	url := _http.AddSepIfNeeded(config.CONFIG.System.LLmUrl) + "v1/chat/completions"
	_logUtils.Info("url=" + url)

	if tmpl != "" {
		pth := filepath.Join("res", "tmpl", tmpl+".txt")
		bts, err1 := deeptest.ReadResData(pth)

		if err1 == nil {
			str := string(bts)
			content = fmt.Sprintf(str, content)
		} else {
			return
		}
	}

	req := v1.ChatCompletionReq{
		Stream: false,
		Messages: []v1.ChatCompletionMsg{{
			Role:    "user",
			Content: content,
		}},
	}
	reqBts, err := json.Marshal(req)

	reqReader := bytes.NewReader(reqBts)
	request, err := http.NewRequest("POST", url, reqReader)
	if err != nil {
		return
	}

	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	transport := &http.Transport{}
	transport.DisableCompression = true
	client.Transport = transport

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	respReader := resp.Body
	respBts, _ := io.ReadAll(respReader)

	llmNlpResp := v1.LlmNlpResp{}
	json.Unmarshal(respBts, &llmNlpResp)

	if len(llmNlpResp.Choices) > 0 {
		llmNlpCotent := llmNlpResp.Choices[0].Message.Content

		tcNlpResult := v1.TcNlpResult{}
		json.Unmarshal([]byte(llmNlpCotent), &tcNlpResult)

		instruction = tcNlpResult.Instruction
		slots = tcNlpResult.Slots
	}

	return
}
