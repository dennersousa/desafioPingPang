package main

import (
	"strings"
	"testing"
)

func TestPingPang(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{3, "1\n2\nPing\n"},
		{5, "1\n2\nPing\n4\nPang\n"},
		{15, "1\n2\nPing\n4\nPang\nPing\n7\n8\nPing\nPang\n11\nPing\n13\n14\nPingPang\n"},
		{7, "1\n2\nPing\n4\nPang\nPing\n7\n"},
	}

	for _, test := range tests {
		result := PingPang(test.input)
		if result != test.expected {
			t.Errorf("Para %d, esperado %s, mas obtido %s", test.input, test.expected, result)
		}
	}
}

func TestPingPangStringContains(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{3, "Ping"},
		{5, "Pang"},
		{15, "PingPang"},
		{7, "7"},
	}

	for _, test := range tests {
		result := PingPang(test.input)
		if !strings.Contains(result, test.expected) {
			t.Errorf("Para %d, esperado conter %s, mas obtido %s", test.input, test.expected, result)
		}
	}
}
