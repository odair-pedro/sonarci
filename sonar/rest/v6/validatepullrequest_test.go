package v6

import (
	"errors"
	"sonarci/testing/mocks"
	"testing"
)

func Test_ValidatePullRequest_CheckError(t *testing.T) {
	restApi := NewRestApi(&mocks.MockConnection{
		RequestMock: func(route string) (<-chan []byte, <-chan error) {
			chErr := make(chan error, 1)
			chErr <- errors.New("error-test")
			return nil, chErr
		},
	})

	err := restApi.ValidatePullRequest("project", "pull-request")
	if err == nil || err.Error() != "error-test" {
		t.Errorf("ValidatePullRequest() not returned expected error")
	}
}
