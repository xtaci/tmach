package tmach

import (
	"unicode"
	"unicode/utf8"
)

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT
)

type Token int

func (t *Token) IsOperator() bool {
	return false
}

func (t *Token) IsLiteral() bool {
	return false
}

func (t *Token) String() string {
	return ""
}

const NoPos Pos = 0

type Pos int

func (p Pos) IsValid() bool {
	return p != NoPos
}

type Scanner struct {
	text []byte
	pos  int
}

func (s *Scanner) Init(text []byte) error {
	s.text = text
	return nil
}

func (s *Scanner) Scan() (pos Pos, tok Token, lit string) {
	for {
		r, size := utf8.DecodeRune(s.text[pos:])
		if r == utf8.RuneError {
			if size == 0 {
				return Pos(0), Token(EOF), ""
			} else {
				return Pos(0), Token(ILLEGAL), ""
			}
		} else if unicode.IsSpace(r) {
			s.pos += size
			continue
		}
	}
}
