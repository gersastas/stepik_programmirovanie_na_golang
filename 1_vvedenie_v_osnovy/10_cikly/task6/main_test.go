package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CheckNumbers(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"30\n11\n7\n101\n", "30\n11\n"},
		{"101\n50\n30\n", ""},
		{"50\n", "50\n"},
		{"9\n101\n", ""},
		{"15\n8\n23\n102\n", "15\n23\n"},
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
