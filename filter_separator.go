package logen

import (
	"strings"
)

func separatorFilter(startMark, endMark string) func(string) string {
	return func(s string) string {
		if len(s) == 0 {
			return s
		}

		var b strings.Builder

		pos := 0
		for pos < len(s) {
			searchStr := s[pos:]

			startIndex := strings.Index(searchStr, startMark)
			if startIndex < 0 {
				b.WriteString(searchStr)
				break
			}

			// Found a start mark.
			endIndex := strings.Index(searchStr[startIndex+len(startMark):], endMark)
			if endIndex < 0 {
				b.WriteString(searchStr)
				break
			}

			// Found both start and end mark.
			b.WriteString(searchStr[:startIndex])

			// move after
			pos += startIndex + len(startMark) + endIndex + len(endMark)
		}

		return b.String()
	}
}

var (
	// BracketFilter filters words separate by brackets.
	// Common logging practice is to report variables in log message using brackets, eg:
	//
	// - "parsed [58] roles from file [/usr/share/elasticsearch/config/roles.yml]"
	//
	// - "loaded file [foo.log] with checksum [jksjksjsq]"
	BracketFilter = separatorFilter("[", "]")
)
