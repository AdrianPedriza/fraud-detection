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

func TestNormalizeAddress(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "address to lower",
			input:  "123 Sesame St",
			output: "123 sesame st",
		},
		{
			name:   "address abbreviations: st",
			input:  "123 Sesame Street",
			output: "123 sesame st",
		},
		{
			name:   "address abbreviations: rd",
			input:  "123 Sesame Road",
			output: "123 sesame rd",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeAddress(tc.input)
			require.Equal(t, result, tc.output)
		})
	}
}
