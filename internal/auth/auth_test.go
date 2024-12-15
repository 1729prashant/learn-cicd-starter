package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
		errorContains string // For matching error messages containing specific text
	}{
		{
			name:          "No authorization header",
			headers:       http.Header{},
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			expectedKey: "test-api-key",
		},
		{
			name: "Malformed header - no space",
			headers: http.Header{
				"Authorization": []string{"ApiKeytest-api-key"},
			},
			errorContains: "malformed authorization header",
		},
		{
			name: "Malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key"},
			},
			errorContains: "malformed authorization header",
		},
		{
			name: "Empty API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			errorContains: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			// Check for expected errors
			if tt.expectedError != nil {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
			} else if tt.errorContains != "" {
				if err == nil {
					t.Errorf("expected error containing '%s', got nil", tt.errorContains)
				} else if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error containing '%s', got '%s'", tt.errorContains, err.Error())
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Check for expected API key
			if key != tt.expectedKey {
				t.Errorf("expected key '%s', got '%s'", tt.expectedKey, key)
			}
		})
	}
}
