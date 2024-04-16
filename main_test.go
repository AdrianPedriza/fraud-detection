package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeEmail(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "email to lower",
			input:  "thisISAtest@gMAIL.COM",
			output: "thisisatest@gmail.com",
		},
		{
			name:   "ignored chars",
			input:  "bugs.1@bunny.com,",
			output: "bugs1@bunny.com,",
		},
		{
			name:   "ignored chars after + char",
			input:  "bugs+10@bunny.com,",
			output: "bugs@bunny.com,",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeEmail(tc.input)
			require.Equal(t, result, tc.output)
		})
	}
}
