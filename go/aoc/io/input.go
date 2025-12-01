package io

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadInput(day int) (string, error) {
	filename := fmt.Sprintf("input%d.txt", day)
	filepath := filepath.Join("input", filename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", filepath, err)
	}

	return string(data), nil
}

func MustReadInput(day int) string {
	input, err := ReadInput(day)
	if err != nil {
		panic(err)
	}
	return input
}
