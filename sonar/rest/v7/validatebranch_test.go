package v7

import (
	"errors"
	"sonarci/testing/mocks"
	"testing"
)

func Test_ValidateBranch_CheckError(t *testing.T) {
	restApi := NewRestApi(&mocks.MockConnection{
		GetMock: func(route string) (<-chan []byte, <-chan error) {
			chErr := make(chan error, 1)
			chErr <- errors.New("error-test")
			return nil, chErr
		},
	})

	err := restApi.ValidateBranch("project", "branch")
	if err == nil || err.Error() != "error-test" {
		t.Errorf("ValidateBranch() not returned expected error")
	}
}
