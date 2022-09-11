package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yadunandanc137.com/assembler/code"
	"yadunandanc137.com/assembler/parser"
	symbol_table "yadunandanc137.com/assembler/symbol-table"
)

var (
	targetFileName *string
	filePath       string
)

func main() {
	targetFileName = flag.String("o", "", "-f : enter target hack file name")
	flag.Parse()
	if len(os.Args) == 0 {
		panic("No assembly input file provided")
	}
	filePath = os.Args[1]
	asmParser := parser.NewParser(filePath)
	opFileName := getOpFileName()
	opFile, err := os.Create(fmt.Sprintf("%s.hack", opFileName))
	if err != nil {
		panic(err)
	}
	defer opFile.Close()
	symTable := getSymbolParser(asmParser)
	for ; asmParser.HasMoreCommands(); asmParser.Advance() {
		line := getLine(asmParser, symTable)
		if strings.EqualFold(line, "") {
			continue
		}
		_, err = fmt.Fprintf(opFile, "%s\n", line)
		if err != nil {
			panic(err)
		}
	}
}

func getOpFileName() string {
	if strings.EqualFold(*targetFileName, "") {
		splitPath := strings.Split(filePath, "/")
		fileName := splitPath[len(splitPath)-1]
		return strings.Split(fileName, ".")[0]
	}
	return *targetFileName
}
func getSymbolParser(asmParser *parser.Parser) *symbol_table.SymbolTable {
	cmdNo := 0
	symTable := symbol_table.NewSymbolTable()
	for ; asmParser.HasMoreCommands(); asmParser.Advance() {
		switch asmParser.CommandType() {
		case parser.ACommand, parser.CCommand:
			cmdNo++
		case parser.LCommand:
			symTable.AddEntry(asmParser.Symbol(), cmdNo)
		}
	}
	asmParser.Reset()
	return symTable
}

func getLine(asmParser *parser.Parser, symTable *symbol_table.SymbolTable) string {
	var codeLine string
	switch asmParser.CommandType() {
	case parser.ACommand:
		address, err := strconv.Atoi(asmParser.Symbol())
		if err != nil {
			if symTable.Contains(asmParser.Symbol()) {
				address = symTable.GetAddress(asmParser.Symbol())
			} else {
				address = symTable.AddVariable(asmParser.Symbol())
			}
		}
		codeLine = fmt.Sprintf("0%.15b", address)
	case parser.CCommand: // dest=comp;jmp
		comp := asmParser.Comp()
		dest := asmParser.Dest()
		jmp := asmParser.Jump()
		codeLine = fmt.Sprintf("111%s%s%s", code.Comp(comp), code.Dest(dest), code.Jump(jmp))
	case parser.LCommand:
	default:
		panic(fmt.Errorf("Command Type not recognized: %v", asmParser.CommandType()))
	}
	return codeLine
}
