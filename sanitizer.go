package logen

import (
	"regexp"
	"strings"

	"github.com/saferwall/saferwall/pkg/gib"
)

const (
	regexpUUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
)

type Filter func(string) string

type Sanitizer struct {
	reUUID       *regexp.Regexp
	reNumbers    *regexp.Regexp
	reSeparators *regexp.Regexp
	isGibberish  func(string) (bool, error)
}

func NewSanitizer() (*Sanitizer, error) {
	// https://github.com/saferwall/saferwall/blob/93bb571f245a2b461366e6e01a520298d5a36109/pkg/gib/gib.go#L196
	isGibberish, err := gib.NewScorer(nil)
	if err != nil {
		return nil, err
	}

	return &Sanitizer{
		// TODO: rely on uuid.Validate
		reUUID:       regexp.MustCompile(regexpUUID),
		reNumbers:    regexp.MustCompile("([0-9]|-)+"),
		reSeparators: regexp.MustCompile(`\s|\[|\]|/`),
		isGibberish:  isGibberish,
	}, nil

}

func (st *Sanitizer) Sanitized(s string) string {
	return st.removeGibberish(st.reNumbers.ReplaceAllString(st.reUUID.ReplaceAllString(BracketFilter(s), ""), ""))
}

func (st *Sanitizer) removeGibberish(s string) string {
	newString := s

	start := 0

	separatorLocations := st.reSeparators.FindAllStringIndex(s, -1)
	if len(separatorLocations) == 0 {
		return s
	}

	for _, singleSeparatorIndex := range separatorLocations {
		if start >= len(s) {
			break
		}

		token := s[start:singleSeparatorIndex[0]]
		if gibberish, _ := st.isGibberish(token); gibberish {
			newString = strings.Replace(newString, token, "", 1)
		}
		start = singleSeparatorIndex[1]
	}

	return newString
}
