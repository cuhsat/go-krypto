package internal

import (
	"encoding/hex"
	"strings"
)

func HexToBytes(s string) []byte {
	var sb strings.Builder
	sb.Grow(len(s))
	s = strings.TrimPrefix(s, "0x")
	for _, c := range s {
		switch {
		case '0' <= c && c <= '9':
			sb.WriteRune(c)
		case 'a' <= c && c <= 'f':
			sb.WriteRune(c)
		case 'A' <= c && c <= 'F':
			sb.WriteRune(c)
		}
	}
	b, err := hex.DecodeString(sb.String())
	if err != nil {
		panic(sb.String())
	}
	return b
}
