package fileutil

import (
	"testing"
)

func TestGetSafeFileName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic alphanumeric",
			input:    "file123",
			expected: "file123",
		},
		{
			name:     "With spaces",
			input:    "my file name",
			expected: "my-file-name",
		},
		{
			name:     "With special characters",
			input:    "file@#$%^&*()name",
			expected: "file-name",
		},
		{
			name:     "With leading and trailing special chars",
			input:    "---file-name---",
			expected: "file-name",
		},
		{
			name:     "Mixed case",
			input:    "FileNAME",
			expected: "filename",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only special characters",
			input:    "@#$%^&*()",
			expected: "",
		},
		{
			name:     "With underscores",
			input:    "file_name_with_underscores",
			expected: "file_name_with_underscores",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetSafeFileName(tt.input)
			if result != tt.expected {
				t.Errorf("GetSafeFileName(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
