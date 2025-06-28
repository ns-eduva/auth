package main

import (
	"testing"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
)

func TestExtractStringInBacktick(t *testing.T) {
	testup.LogNameTestInfo(t, "Test extract string in backtick")

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
	testup.LogNameTestInfo(t, "Test if to access to the .env.test file")

	tests := []struct {
		key      string
		expected string
	}{
		{"APP_ENV", "test"},
	}

	for _, test := range tests {
		result := env.Get(test.key)
		assert.Equalf(t, test.expected, result, "OK")
	}
}
