package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_DigitSumOfNumbers(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"745\n", "16\n"},
		{"111\n", "3\n"},
		{"628\n", "16\n"},
		{"435\n", "12\n"},
		{"123\n", "6\n"},
		{"999\n", "27\n"},
	}

	for _, tt := range tests {
		// Подменяем stdin
		oldStdin := os.Stdin
		r, w, _ := os.Pipe()
		os.Stdin = r
		w.WriteString(tt.input)
		w.Close()

		// Перехватываем stdout
		oldStdout := os.Stdout
		ro, wo, _ := os.Pipe()
		os.Stdout = wo

		main()

		wo.Close()
		os.Stdout = oldStdout
		os.Stdin = oldStdin

		var buf bytes.Buffer
		io.Copy(&buf, ro)

		if buf.String() != tt.want {
			t.Errorf("input %q: got %q, want %q", tt.input, buf.String(), tt.want)
		}
	}
}
