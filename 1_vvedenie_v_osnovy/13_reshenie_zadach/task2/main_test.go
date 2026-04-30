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
		{"745\n", "547\n"},
		{"843\n", "348\n"},
		{"628\n", "826\n"},
		{"435\n", "534\n"},
		{"123\n", "321\n"},
		{"926\n", "629\n"},
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
