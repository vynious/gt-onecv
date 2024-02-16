package tests

import (
	"github.com/vynious/gt-onecv/domains/notifications"
	"testing"
)

func TestExtractEmails(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single email",
			input:    "test@example.com",
			expected: []string{"test@example.com"},
		},
		{
			name:     "multiple emails",
			input:    "first@example.com, second@test.com",
			expected: []string{"first@example.com", "second@test.com"},
		},
		{
			name:     "no emails",
			input:    "no emails here",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := notifications.ExtractEmails(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("Expected %d emails, got %d", len(tt.expected), len(got))
			}
			for i, email := range got {
				if email != tt.expected[i] {
					t.Errorf("Expected email %s, got %s", tt.expected[i], email)
				}
			}
		})
	}
}
