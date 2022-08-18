package parser

import (
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	A_COMMAND = 0
	C_COMMAND = 1
	L_COMMAND = 2
)

var ()

type Parser struct {
	currentCommandIdx int
	commands          []string
	aCommandMatcher   *regexp.Regexp
	cCommandMatcher   *regexp.Regexp
	lCommandMatcher   *regexp.Regexp
}

func NewParser(filePath string) *Parser {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error while reading .asm file. error: %s", err)
	}
	lines := strings.Split(string(fileContent), "\n")
	commands := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.EqualFold(trimmedLine, "") {
			continue
		}
		uncommentedLine := strings.Split(trimmedLine, "//")[0]
		if strings.EqualFold(uncommentedLine, "") {
			continue
		}
		commands = append(commands, uncommentedLine)
	}
	aCommandMatcher, err := regexp.Compile("@(\\w+)")
	if err != nil {
		log.Fatalf("error while generating command parser. error: %s", err)
	}
	cCommandMatcher, err := regexp.Compile("((([A-Z])=)?(([A-Z]\\+)?[A-Z]);)?(\\w+)?")
	if err != nil {
		log.Fatalf("error while generating command parser. error: %s", err)
	}
	lCommandMatcher, err := regexp.Compile("\\((\\w+)\\)")
	if err != nil {
		log.Fatalf("error while generating command parser. error: %s", err)
	}
	return &Parser{
		commands:          commands,
		currentCommandIdx: 0,
		aCommandMatcher:   aCommandMatcher,
		cCommandMatcher:   cCommandMatcher,
		lCommandMatcher:   lCommandMatcher,
	}
}

// HasMoreCommands - are there any more commands left in the asm file?
func (p *Parser) HasMoreCommands() bool {
	return p.currentCommandIdx < len(p.commands)
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}

func (p *Parser) CommandType() int {
	if p.currentCommandIdx >= len(p.commands) {
		log.Fatalf("commands are not parsed properly")
	}
	command := p.commands[p.currentCommandIdx]

	isMatch := p.aCommandMatcher.Match([]byte(command))
	if isMatch {
		return A_COMMAND
	}

	isMatch = p.cCommandMatcher.Match([]byte(command))
	if isMatch {
		return C_COMMAND
	}

	isMatch = p.lCommandMatcher.Match([]byte(command))
	if isMatch {
		return L_COMMAND
	}
	log.Fatalf("command not recognized. command: %s", command)
	return 0
}

func (p *Parser) Symbol() string {
	command := p.commands[p.currentCommandIdx]
	if p.CommandType() == A_COMMAND {
		matchGroups := p.aCommandMatcher.FindAllStringSubmatch(command, -1)
		if len(matchGroups) == 0 || len(matchGroups[0]) == 0 {
			log.Fatalf("no matches found")
		}
		return matchGroups[0][0]
	}
	if p.CommandType() == L_COMMAND {
		matchGroups := p.lCommandMatcher.FindAllStringSubmatch(command, -1)
		if len(matchGroups) == 0 || len(matchGroups[0]) == 0 {
			log.Fatalf("no matches found")
		}
		return matchGroups[0][0]
	}
	return ""
}

func (p *Parser) Dest() string {
	command := p.commands[p.currentCommandIdx]
	if p.CommandType() == C_COMMAND {
		matchGroups := p.cCommandMatcher.FindAllStringSubmatch(command, -1)
		if len(matchGroups) == 0 || len(matchGroups[0]) < 4 {
			log.Fatalf("no matches found")
		}
		return matchGroups[0][3]
	}
	return ""
}

func (p *Parser) Comp() string {
	command := p.commands[p.currentCommandIdx]
	if p.CommandType() == C_COMMAND {
		matchGroups := p.cCommandMatcher.FindAllStringSubmatch(command, -1)
		if len(matchGroups) == 0 || len(matchGroups[0]) < 5 {
			log.Fatalf("no matches found")
		}
		return matchGroups[0][4]
	}
	return ""
}

func (p *Parser) Jump() string {
	command := p.commands[p.currentCommandIdx]
	if p.CommandType() == C_COMMAND {
		matchGroups := p.cCommandMatcher.FindAllStringSubmatch(command, -1)
		if len(matchGroups) == 0 || len(matchGroups[0]) < 6 {
			log.Fatalf("no matches found")
		}
		return matchGroups[0][6]
	}
	return ""
}
