package youtrack

import (
	"testing"
)

// The YOUTRACK_URL and YOUTRACK_TOKEN environment variables need to be setup for this to work.
func TestDefaultApi(t *testing.T) {
	api, err := NewDefaultApi()
	if err != nil {
		t.Fatal("Failed to create api.", err)
	}
	if api == nil {
		t.Fatal("API is nil")
	}
}
