package v6

import (
	"errors"
	"testing"
)

func Test_ValidateBranch_CheckError(t *testing.T) {
	restApi := NewRestApi(&mockConnection{
		doGet: func(route string) (<-chan []byte, <-chan error) {
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
