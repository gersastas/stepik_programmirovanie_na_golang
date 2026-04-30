package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CheckTringle(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"6 8 10\n", "Прямоугольный\n"},
		{"3 4 5\n", "Прямоугольный\n"},
		{"5 12 13\n", "Прямоугольный\n"},
		{"8 15 17\n", "Прямоугольный\n"},
		{"1 2 3\n", "Непрямоугольный\n"},
		{"3 4 6\n", "Непрямоугольный\n"},
		{"5 6 7\n", "Непрямоугольный\n"},
		{"1 1 2\n", "Непрямоугольный\n"},
		{"7 24 25\n", "Прямоугольный\n"},
		{"9 40 41\n", "Прямоугольный\n"},
		{"6 8 10\n", "Прямоугольный\n"},
		{"3 4 5\n", "Прямоугольный\n"},
		{"5 12 13\n", "Прямоугольный\n"},
		{"8 15 17\n", "Прямоугольный\n"},
		{"1 2 3\n", "Непрямоугольный\n"},
		{"3 4 6\n", "Непрямоугольный\n"},
		{"5 6 7\n", "Непрямоугольный\n"},
		{"1 1 2\n", "Непрямоугольный\n"},
		{"7 24 25\n", "Прямоугольный\n"},
		{"9 40 41\n", "Прямоугольный\n"},
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
