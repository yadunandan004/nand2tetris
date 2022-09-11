package symbol_table

type SymbolTable struct {
	varAddress       int
	symbolAddressMap map[string]int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		varAddress: 16,
		symbolAddressMap: map[string]int{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"R0":     0,
			"R1":     1,
			"R2":     2,
			"R3":     3,
			"R4":     4,
			"R5":     5,
			"R6":     6,
			"R7":     7,
			"R8":     8,
			"R9":     9,
			"R10":    10,
			"R11":    11,
			"R12":    12,
			"R13":    13,
			"R14":    14,
			"R15":    15,
			"SCREEN": 16384,
			"KBD":    24567,
		},
	}
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.symbolAddressMap[symbol] = address
}

func (s *SymbolTable) Contains(symbol string) bool {
	_, isPresent := s.symbolAddressMap[symbol]
	return isPresent
}

func (s *SymbolTable) GetAddress(symbol string) int {
	return s.symbolAddressMap[symbol]
}

func (s *SymbolTable) AddVariable(symbol string) int {
	s.symbolAddressMap[symbol] = s.varAddress
	s.varAddress++
	return s.varAddress - 1
}
