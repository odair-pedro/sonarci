package azuredevops

import (
	"errors"
	"sonarci/sonar"
	"sonarci/testing/mocks"
	"testing"
)

func TestPullRequestDecorator_CommentQualityGate_CheckErrorOnParseTemplate(t *testing.T) {
	wantError := errors.New("failure")

	mockEngine := &mocks.MockEngine{
		ProcessTemplateMock: func(template string, dataSource interface{}) (string, error) {
			return "", wantError
		},
	}

	decorator := NewPullRequestDecorator(&mocks.MockConnection{}, mockEngine, "project-test", "repo-test")
	gotError := decorator.CommentQualityGate(sonar.QualityGate{})

	if gotError != wantError {
		t.Fail()
	}
}

func TestPullRequestDecorator_CommentQualityGate_CheckErrorOnSendRequest(t *testing.T) {
	wantError := errors.New("failure")

	mockConn := &mocks.MockConnection{
		PostMock: func(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error) {
			chOut := make(chan []byte, 1)
			chErr := make(chan error, 1)
			defer close(chOut)
			defer close(chErr)

			chErr <- wantError
			return chOut, chErr
		},
	}

	mockEngine := &mocks.MockEngine{
		ProcessTemplateMock: func(template string, dataSource interface{}) (string, error) {
			return "Is is a fake report", nil
		},
	}

	decorator := NewPullRequestDecorator(mockConn, mockEngine, "project-test", "repo-test")
	gotError := decorator.CommentQualityGate(sonar.QualityGate{})

	if gotError != wantError {
		t.Fail()
	}
}

func TestPullRequestDecorator_CommentQualityGate_CheckNoError(t *testing.T) {
	mockConn := &mocks.MockConnection{
		PostMock: func(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error) {
			chOut := make(chan []byte, 1)
			chErr := make(chan error, 1)
			defer close(chOut)
			defer close(chErr)

			return chOut, chErr
		},
	}

	mockEngine := &mocks.MockEngine{
		ProcessTemplateMock: func(template string, dataSource interface{}) (string, error) {
			return "Is is a fake report", nil
		},
	}

	decorator := NewPullRequestDecorator(mockConn, mockEngine, "project-test", "repo-test")
	gotError := decorator.CommentQualityGate(sonar.QualityGate{})

	if gotError != nil {
		t.Fail()
	}
}
