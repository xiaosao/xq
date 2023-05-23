package scanner

type Scanner struct {
	Source string
}

type Token struct {
}

func (sc *Scanner) ScanTokens() []Token {
	return []Token{}
}
