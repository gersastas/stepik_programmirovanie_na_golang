package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CountMaxNumsSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1\n3\n3\n1\n0\n", "2\n"},
		{"7\n3\n8, 4\n7\n0\n", "1\n"},
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
