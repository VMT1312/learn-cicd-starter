package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		want      string
		expectErr error
	}{
		"no authorization header": {
			expectErr: errors.New("no authorization header included"),
		},
		"header without api": {
			key:       "Authorization",
			expectErr: errors.New("no authorization header included"),
		},
		"malformed authorization header": {
			key:       "Authorization",
			value:     "",
			expectErr: errors.New("malformed authorization header"),
		},
		"header with only bearer": {
			key:       "Authorization",
			value:     "Bearer xxxxx",
			expectErr: errors.New("malformed authorization header"),
		},
		"valid api key": {
			key:       "Authorization",
			value:     "ApiKey xxxxx",
			want:      "xxxxx",
			expectErr: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			if test.key != "" {
				header.Set(test.key, test.value)
			}

			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr.Error()) {
					return
				}
				t.Errorf("unexpected error: got %v, want %v", err, test.expectErr)
				return
			}

			if got != test.want {
				t.Errorf("expected API key %q, got %q", test.want, got)
			}
		})
	}
}
