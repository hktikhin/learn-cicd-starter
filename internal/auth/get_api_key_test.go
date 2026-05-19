package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		wantKey string
		wantErr bool
	}{
		"valid":          {headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}}, wantKey: "my-secret-key", wantErr: false},
		"missing header": {headers: http.Header{}, wantKey: "", wantErr: true},
		"wrong prefix":   {headers: http.Header{"Authorization": []string{"Bearer my-secret-key"}}, wantKey: "", wantErr: true},
		"malformed":      {headers: http.Header{"Authorization": []string{"ApiKey"}}, wantKey: "", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if (err != nil) != tc.wantErr {
				t.Fatalf("unexpected error status: %v", err)
			}
			if got != tc.wantKey {
				t.Fatalf("want %q, got %q", tc.wantKey, got)
			}
		})
	}
}
