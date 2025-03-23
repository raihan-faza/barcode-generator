package encoder

type Encoder struct {
}

func GenerateCodeATable() map[rune]int {
	codeATable := make(map[rune]int)
	for i := 0; i <= 95; i++ {
		codeATable[rune(i)] = i
	}
	return codeATable
}

func GenerateCodeBTable() map[rune]int {
	codeBTable := make(map[rune]int)
	for i := 32; i <= 127; i++ {
		codeBTable[rune(i)] = i
	}
	return codeBTable
}

func GenerateCodeCTable() map[rune]int {
	codeCTable := make(map[rune]int)
	return codeCTable
}

func (e *Encoder) Code128Encode(data string) (string, error) {
	return "", nil
}

func (e *Encoder) Code39Encode(data string) (string, error) {
	return "", nil
}

func (e *Encoder) UPCEncode(data string) (string, error) {
	return "", nil
}
