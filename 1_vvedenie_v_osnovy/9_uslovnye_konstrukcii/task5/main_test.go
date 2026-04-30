package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_DigitSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"2000\n", "Yes\n"},
		{"2004\n", "Yes\n"},
		{"2008\n", "Yes\n"},
		{"2024\n", "Yes\n"},
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
