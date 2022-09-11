package code

import (
	"fmt"
	"strings"
)

var (
	destNames       = []string{"", "M", "D", "MD", "A", "AM", "AD", "AMD"}
	jumpNames       = []string{"", "JGT", "JEQ", "JGE", "JLT", "JNE", "JLE", "JMP"}
	compNameCodeMap = map[string]int{
		"0":   42,
		"1":   63,
		"-1":  58,
		"D":   12,
		"A":   48,
		"M":   112,
		"!D":  13,
		"!A":  49,
		"!M":  113,
		"-D":  15,
		"-A":  51,
		"-M":  115,
		"D+1": 31,
		"A+1": 55,
		"M+1": 119,
		"D-1": 14,
		"A-1": 50,
		"M-1": 114,
		"D+A": 2,
		"D+M": 66,
		"D-A": 19,
		"D-M": 83,
		"A-D": 7,
		"M-D": 71,
		"D&A": 0,
		"D&M": 64,
		"D|A": 21,
		"D|M": 85,
	}
)

func Dest(destName string) string {
	for i := range destNames {
		if strings.EqualFold(destName, destNames[i]) {
			return fmt.Sprintf("%.3b", i)
		}
	}
	return ""
}

func Comp(compName string) string {
	return fmt.Sprintf("%.7b", compNameCodeMap[compName])
}

func Jump(jumpName string) string {
	for i := range jumpNames {
		if strings.EqualFold(jumpName, jumpNames[i]) {
			return fmt.Sprintf("%.3b", i)
		}
	}
	return ""
}
