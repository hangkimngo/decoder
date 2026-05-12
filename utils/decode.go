package utils

import "strings"

func Decoder(s string) string {
	return decoder(s)
}

func decoder(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for i := 0; i < len(s); i++ {
		if s[i] != '[' {
			b.WriteByte(s[i])
			continue
		}

		i++
		count := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			count = count*10 + int(s[i]-'0')
			i++
		}

		if i < len(s) && s[i] == ' ' {
			i++
		}

		start := i
		for i < len(s) && s[i] != ']' {
			i++
		}

		part := s[start:i]
		b.Grow(len(part) * count)
		for ; count > 0; count-- {
			b.WriteString(part)
		}
	}

	return b.String()
}
