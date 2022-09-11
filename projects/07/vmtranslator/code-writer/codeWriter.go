package code_writer

import (
	"fmt"
	"io"
	"strings"
	"vmtranslator/parser"
)

type CodeWriter struct {
	fp     io.WriteCloser
	jgtIdx int
	jltIdx int
	negIdx int
}

func NewCodeWriter(file io.WriteCloser) *CodeWriter {
	asmCodeBlock := strings.Builder{}

	_, err := fmt.Fprintf(file, "%s", asmCodeBlock.String())
	if err != nil {
		panic(err)
	}
	return &CodeWriter{
		fp: file,
	}
}

func (cw *CodeWriter) SetFileName(fileName string) {

}

func (cw *CodeWriter) WriteArithmetic(command string) {
	asmCodeBlock := strings.Builder{}
	switch strings.ToLower(command) {
	case "add":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Add Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=M+D\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("// ---- Add End ----\n")
	case "sub":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Sub Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=M-D\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("// ---- Sub End ----\n")
	case "neg":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Neg Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=-M\n")
		asmCodeBlock.WriteString("// ---- Neg End ----\n")
	case "eq":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Eq Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=D-M\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("A=D\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString(fmt.Sprintf("@negs%d\n", cw.negIdx))
		asmCodeBlock.WriteString("D;JNE\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=-1\n")
		asmCodeBlock.WriteString("D=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(negs%d)\n", cw.negIdx))
		asmCodeBlock.WriteString(fmt.Sprintf("@nege%d\n", cw.negIdx))
		asmCodeBlock.WriteString("D;JEQ\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(nege%d)\n", cw.negIdx))
		asmCodeBlock.WriteString("// ---- Eq End ----\n")
		cw.negIdx++

	case "gt":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Gt Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n") // 12,7, sp
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n") // D=7
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=M-D\n") // M=12-7
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("A=D\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString(fmt.Sprintf("@jgts%d\n", cw.jgtIdx))
		asmCodeBlock.WriteString("D;JLE\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=-1\n")
		asmCodeBlock.WriteString("D=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(jgts%d)\n", cw.jgtIdx))
		asmCodeBlock.WriteString(fmt.Sprintf("@jgte%d\n", cw.jgtIdx))
		asmCodeBlock.WriteString("D;JEQ\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(jgte%d)\n", cw.jgtIdx))
		asmCodeBlock.WriteString("// ---- Gt End ----\n")
		cw.jgtIdx++

	case "lt":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Lt Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=M-D\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("A=D\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString(fmt.Sprintf("@jlts%d\n", cw.jltIdx))
		asmCodeBlock.WriteString("D;JGE\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=-1\n")
		asmCodeBlock.WriteString("D=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(jlts%d)\n", cw.jltIdx))
		asmCodeBlock.WriteString(fmt.Sprintf("@jlte%d\n", cw.jltIdx))
		asmCodeBlock.WriteString("D;JEQ\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=0\n")
		asmCodeBlock.WriteString(fmt.Sprintf("(jlte%d)\n", cw.jltIdx))
		asmCodeBlock.WriteString("// ---- Lt End ----\n")
		cw.jltIdx++
	case "and":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- And Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=D&M\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("// ---- And End ----\n")
	case "or":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Or Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("A=A-1\n")
		asmCodeBlock.WriteString("M=D|M\n")
		asmCodeBlock.WriteString("D=A\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=D+1\n")
		asmCodeBlock.WriteString("// ---- Or End ----\n")
	case "not":
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Not Start ---- ## %s\n", command))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("M=!M\n")
		asmCodeBlock.WriteString("// ---- Not End ----\n")
	}
	_, err := fmt.Fprintf(cw.fp, "%s", asmCodeBlock.String())
	if err != nil {
		panic(err)
	}
}
func (cw *CodeWriter) WritePushPop(commandType int8, segmentName string, indexName string) {
	asmCodeBlock := strings.Builder{}

	switch commandType {
	case parser.C_PUSH:
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Push Start ---- ## push %s %s\n", segmentName, indexName))
		asmCodeBlock.WriteString(getSegmentAddress(segmentName, indexName, true))
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M\n")
		asmCodeBlock.WriteString("M=D\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=M+1\n")
		asmCodeBlock.WriteString("// ---- Push End ----\n")
	case parser.C_POP:
		asmCodeBlock.WriteString(fmt.Sprintf("// ---- Pop Start ---- ## pop %s %s\n", segmentName, indexName))

		asmCodeBlock.WriteString(getSegmentAddress(segmentName, indexName, false))

		asmCodeBlock.WriteString("@13\n")
		asmCodeBlock.WriteString("M=D\n")
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("A=M-1\n")
		asmCodeBlock.WriteString("D=M\n")
		asmCodeBlock.WriteString("@13\n")
		asmCodeBlock.WriteString("A=M\n")
		asmCodeBlock.WriteString("M=D\n")
		if strings.EqualFold(segmentName, "pointer") {
			if strings.EqualFold(indexName, "0") {
				asmCodeBlock.WriteString("@3\n")
				asmCodeBlock.WriteString("M=D\n")
			} else {
				asmCodeBlock.WriteString("@4\n")
				asmCodeBlock.WriteString("M=D\n")
			}
		}
		asmCodeBlock.WriteString("@0\n")
		asmCodeBlock.WriteString("M=M-1\n")
		asmCodeBlock.WriteString("// ---- Pop End ----\n")
	}

	_, err := fmt.Fprintf(cw.fp, "%s", asmCodeBlock.String())
	if err != nil {
		panic(err)
	}
}

func getSegmentAddress(segmentName string, indexName string, isPush bool) string {
	asmCodeBlock := strings.Builder{}
	asmCodeBlock.WriteString(fmt.Sprintf("@%s\n", indexName))
	asmCodeBlock.WriteString("D=A\n")
	switch segmentName {
	case "local":
		asmCodeBlock.WriteString("@1\n")
		asmCodeBlock.WriteString("A=M\n")
		asmCodeBlock.WriteString("A=A+D\n")
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}
	case "static":
		asmCodeBlock.WriteString("@30\n")
		asmCodeBlock.WriteString("A=M\n")
		asmCodeBlock.WriteString("A=A+D\n")
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}
	case "argument":
		asmCodeBlock.WriteString("@2\n")
		asmCodeBlock.WriteString("A=M\n")
		asmCodeBlock.WriteString("A=A+D\n")
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}
	case "this":
		asmCodeBlock.WriteString("@3\n")
		asmCodeBlock.WriteString("A=M+D\n")
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}
	case "that":
		asmCodeBlock.WriteString("@4\n")
		asmCodeBlock.WriteString("A=M+D\n")
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}
	case "pointer":
		if strings.EqualFold(indexName, "0") {
			asmCodeBlock.WriteString("@50\n")
		} else {
			asmCodeBlock.WriteString("@51\n")
		}
		if isPush {
			asmCodeBlock.WriteString("D=M\n")
		} else {
			asmCodeBlock.WriteString("D=A\n")
		}

	case "temp":
		asmCodeBlock.WriteString("@5\n")
	case "constant":
	default:
		panic("invalid push segment")
	}
	return asmCodeBlock.String()
}
func (cw *CodeWriter) Close() {
	// close the file pointer
	cw.fp.Close()
}
