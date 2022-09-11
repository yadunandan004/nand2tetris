package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	code_writer "vmtranslator/code-writer"
	"vmtranslator/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("vm input file path not specified")
	}
	ipFilePath := os.Args[1]

	opFile, err := os.Create(replaceFileExtension(ipFilePath))
	if err != nil {
		panic(err)
	}
	writer := code_writer.NewCodeWriter(opFile)
	defer writer.Close()
	vmParser := parser.NewParser(ipFilePath)
	for ; vmParser.HasMoreCommands(); vmParser.Advance() {
		switch vmParser.CommandType() {
		case parser.C_ARITHMETIC:
			writer.WriteArithmetic(vmParser.Arg1())
		case parser.C_PUSH, parser.C_POP:
			writer.WritePushPop(vmParser.CommandType(), vmParser.Arg1(), vmParser.Arg2())
		default:
			panic(fmt.Errorf("command not supported yet. type [%s]", vmParser.CommandType()))
		}
	}

}

func replaceFileExtension(filePath string) string {
	return strings.Replace(filePath, ".vm", ".asm", -1)
}
