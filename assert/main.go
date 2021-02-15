package assert

import (
	"fmt"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf(fmt.Sprintf("Expected %v but got %v", expected, actual))
	}
}

func True(t *testing.T, boolean bool) {
	if !boolean {
		t.Errorf(fmt.Sprintf("Expected true but got false"))
	}
}

func False(t *testing.T, boolean bool) {
	if !boolean {
		t.Errorf(fmt.Sprintf("Expected false but got true"))
	}
}
