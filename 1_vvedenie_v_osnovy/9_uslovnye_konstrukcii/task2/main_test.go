package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_NumCheckDigit(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"237\n", "YES\n"},
		{"117\n", "NO\n"},
		{"111\n", "NO\n"},
		{"525\n", "NO\n"},
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
