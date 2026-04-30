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
		{"20\n3\n5\n", "3\n"},
		{"50\n8\n4\n", ""},
		{"1000\n13\n3\n", "13\n"},
		{"10000\n9999\n10000\n", "9999\n"},
		{"5\n10\n3\n", ""},
		{"30\n6\n9\n", "6\n"},
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
