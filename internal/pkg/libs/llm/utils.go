package llmUtils

import (
	"fmt"
	_int "github.com/deeptest-com/deeptest-next/pkg/libs/int"
	"strings"
)

func ReplaceFields(item InstructionItem, mp map[string][]string) (ret InstructionItem) {
	ret = item

	for key, valArr := range mp {
		key := fmt.Sprintf("${%s}", key)

		index := _int.GenUniqueRandNum(0, len(valArr), 1)[0]
		val1 := valArr[index]
		ret.Instruction = strings.ReplaceAll(ret.Instruction, key, val1)

		ret.Input = strings.ReplaceAll(ret.Input, key, val1)

		ret.Output = strings.ReplaceAll(ret.Output, key, val1)
	}

	return
}

type InstructionItem struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}
