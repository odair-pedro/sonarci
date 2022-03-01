package v6

import (
	"sonarci/testing/mocks"
	"testing"
)

func Test_NewRestApi(t *testing.T) {
	if got := NewRestApi(&mocks.MockConnection{}); got == nil {
		t.Errorf("NewRestApi() return nil")
	}
}
