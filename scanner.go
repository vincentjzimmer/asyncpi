package asyncpi

import (
	"bufio"
	"bytes"
	"io"
)

// Scanner is a lexical scanner.
type Scanner struct {
	r   *bufio.Reader
	pos TokenPos
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), pos: TokenPos{Char: 0, Lines: []int{}}}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if reached the end or error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		s.pos.Lines = append(s.pos.Lines, s.pos.Char)
		s.pos.Char = 0
	} else {
		s.pos.Char++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.pos.Char == 0 {
		s.pos.Char = s.pos.Lines[len(s.pos.Lines)-1]
		s.pos.Lines = s.pos.Lines[:len(s.pos.Lines)-1]
	} else {
		s.pos.Char--
	}
}

// Scan returns the next token and parsed value.
func (s *Scanner) Scan() (token Token, value string, startPos, endPos TokenPos) {
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
	}
	if isAlphaNum(ch) {
		s.unread()
		return s.scanName()
	}

	// Track token positions.
	startPos = s.pos
	defer func() { endPos = s.pos }()

	switch ch {
	case eof:
		return 0, "", startPos, endPos
	case '<':
		return LANGLE, string(ch), startPos, endPos
	case '>':
		return RANGLE, string(ch), startPos, endPos
	case '(':
		return LPAREN, string(ch), startPos, endPos
	case ')':
		return RPAREN, string(ch), startPos, endPos
	case '.':
		return PREFIX, string(ch), startPos, endPos
	case ';':
		return SEMICOLON, string(ch), startPos, endPos
	case ':':
		return COLON, string(ch), startPos, endPos
	case '|':
		return PAR, string(ch), startPos, endPos
	case '!':
		return REPEAT, string(ch), startPos, endPos
	case ',':
		return COMMA, string(ch), startPos, endPos
	}

	return ILLEGAL, string(ch), startPos, endPos
}

func (s *Scanner) scanName() (token Token, value string, startPos, endPos TokenPos) {
	var buf bytes.Buffer
	startPos = s.pos
	defer func() { endPos = s.pos }()
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isAlphaNum(ch) && !isNameSymbols(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch buf.String() {
	case "0":
		return NIL, buf.String(), startPos, endPos
	case "new":
		return NEW, buf.String(), startPos, endPos
	}
	return NAME, buf.String(), startPos, endPos
}

func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
