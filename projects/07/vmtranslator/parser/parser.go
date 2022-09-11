package parser

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	C_ARITHMETIC int8 = 0 // C_ARITHMETIC == 0
	C_PUSH       int8 = 1
	C_POP        int8 = 2
	C_LABEL      int8 = 3
	C_GOTO       int8 = 4
	C_IF         int8 = 5
	C_FUNCTION   int8 = 6
	C_RETURN     int8 = 7
	C_CALL       int8 = 8
)

type Parser struct {
	commands      []string
	currentCmdIdx int
}

func NewParser(filePath string) *Parser {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fileLines := strings.Split(string(fileContent), "\n")
	commands := make([]string, 0, len(fileLines))
	for _, line := range fileLines {
		unCommentedStr := strings.Split(line, "//")
		codeLine := strings.TrimSpace(unCommentedStr[0])
		if !strings.EqualFold(codeLine, "") {
			commands = append(commands, line)
		}
	}
	return &Parser{
		currentCmdIdx: 0,
		commands:      commands,
	}

}

func (p *Parser) HasMoreCommands() bool {
	return p.currentCmdIdx < len(p.commands)
}

func (p *Parser) Advance() {
	p.currentCmdIdx++
}

func (p *Parser) CommandType() int8 {
	commandComponents := strings.Split(p.commands[p.currentCmdIdx], " ")
	if len(commandComponents) == 0 {
		panic(fmt.Errorf("invalid command %s", p.commands[p.currentCmdIdx]))
	}
	command := commandComponents[0]
	switch strings.ToLower(command) {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ARITHMETIC
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	default:
		panic(fmt.Errorf("invalid command %s", p.commands[p.currentCmdIdx]))
	}
	return -1
}

func (p *Parser) Arg1() string {
	commandComponents := strings.Split(p.commands[p.currentCmdIdx], " ")
	if len(commandComponents) == 0 {
		panic(fmt.Errorf("invalid command %s", p.commands[p.currentCmdIdx]))
	}
	commandType := p.CommandType()
	if commandType == C_RETURN {
		return ""
	}
	if commandType == C_ARITHMETIC {
		return commandComponents[0]
	}
	if len(commandComponents) < 2 {
		panic(fmt.Errorf(" command with invalid arguments %s", p.commands[p.currentCmdIdx]))
	}
	return commandComponents[1]
}

func (p *Parser) Arg2() string {
	commandComponents := strings.Split(p.commands[p.currentCmdIdx], " ")
	if len(commandComponents) == 0 {
		panic(fmt.Errorf("invalid command %s", p.commands[p.currentCmdIdx]))
	}
	commandType := p.CommandType()
	if commandType == C_POP || commandType == C_PUSH || commandType == C_FUNCTION || commandType == C_CALL {
		return commandComponents[2]
	}
	return ""
}
