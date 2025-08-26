package logen

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
)

func Test_UUIDRegex(t *testing.T) {
	t.Run("UUIDv4", func(t *testing.T) {
		re := regexp.MustCompile(regexpUUID)
		for i := 0; i < 10; i++ {
			id := uuid.NewString()
			if !re.MatchString(id) {
				t.Errorf("UUID regexp does not match %v", id)
			}
		}
	})
}

func Test_Sanitizer(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Empty",
			input:  "",
			output: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			st, err := NewSanitizer()
			if err != nil {
				t.Fatal(err)
			}

			event := st.Sanitized(tc.input)

			if event != tc.output {
				t.Errorf("unexpected sanitized output: %s, expecting %s for input %s", event, tc.input, tc.output)
			}
		})
	}
}
