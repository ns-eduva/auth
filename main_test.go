package main

import (
	"testing"

	"github.com/nsevenpack/env/env"
	"github.com/stretchr/testify/assert"
)

func TestExtractStringInBacktick(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"`hello`", "hello"},
		{"no backticks", ""},
		{"`multiple` `backticks`", "multiple` `backticks"},
	}

	for _, test := range tests {
		result := extractStringInBacktick(test.input)
		assert.Equalf(t, test.expected, result, "OK")
	}
}

func TestInitEnv(t *testing.T) {
	tests := []struct {
		key      string
		expected string
	}{
		{"APP_ENV", "test"}, // ici la package e,v prend le .env.test car on est en test
	}

	for _, test := range tests {
		result := env.Get(test.key)
		assert.Equalf(t, test.expected, result, "OK")
	}
}