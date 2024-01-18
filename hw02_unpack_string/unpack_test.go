package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "u", expected: "u"},
		{input: "ty2o4", expected: "tyyoooo"},
		{input: "а1б2в3", expected: "аббввв"},
		{input: "я9", expected: "яяяяяяяяя"},
		{input: "本⛷สวัส4ดี", expected: "本⛷สวัสสสสดี"},
		{input: "สวัสดี", expected: "สวัสดี"},
		{input: "สวัส4ดี", expected: "สวัสสสสดี"},
		{input: "🙂9", expected: "🙂🙂🙂🙂🙂🙂🙂🙂🙂"},
		{input: "世1界1", expected: "世界"},
		{input: "世1界💻0", expected: "世界"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "62dru2e4", "ty204", "0💻💻"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
