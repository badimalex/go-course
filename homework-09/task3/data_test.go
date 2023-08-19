package homework09

import (
	"bytes"
	"testing"
)

func TestSave(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		expected string
	}{
		{
			name:     "Test with strings",
			args:     []interface{}{"Hello", "World", "Go"},
			expected: "Hello\nWorld\nGo\n",
		},
		{
			name:     "Test with mixed types",
			args:     []interface{}{"Hello", 123, "World", 3.14, "Go"},
			expected: "Hello\nWorld\nGo\n",
		},
		{
			name:     "Test with no strings",
			args:     []interface{}{123, 3.14, true, []string{"a", "b", "c"}},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			Save(&buf, test.args...)

			result := buf.String()
			if result != test.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", test.expected, result)
			}
		})
	}
}
