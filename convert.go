package git

import (
	"bytes"
	"strings"
	"unicode"
)

func crlf_to_git(s string) string {
	return strings.Replace(s, "\r\n", "\n", -1)
}

type textStat struct {
	nul, cr, lf, crlf int64 // NUL, CR, LF and CRLF counts
	p, np             int64 // These are just approximations!
}

func IsBinary(buf *bytes.Buffer) bool {
	var stats textStat
	var c rune
	var err error

	for {
		c, _, err = buf.ReadRune()
		if err != nil {
			return false
		}

		switch {
		case c == '\r':
			stats.cr++
			c, _, err = buf.ReadRune()
			if err != nil {
				return false
			}
			if c == '\n' {
				stats.crlf++
			}
			buf.UnreadRune()
		case c == '\n':
			stats.lf++
		case c == '\x7f': // DEL
			stats.np++
		case c == '\x1a': // EOL
		case unicode.IsControl(c):
			switch c {
			// BS, HT, ESC and FF
			case '\b', '\t', '\x1b', '\x0c':
				stats.p++
			case 0:
				stats.nul++
				fallthrough
			default:
				stats.np++
			}
		}
	}

	// If file ends with EOF then don't count this EOF as non-printable.
	buf.UnreadRune()
	c, _, err = buf.ReadRune()
	if err != nil && c == '\x1a' {
		stats.np--
	}

	// Is binary heuristics, check `git/convert.c:is_binary`
	if stats.nul == 0 {
		return true
	}
	if stats.p>>7 < stats.np {
		return true
	}
	return false
}
