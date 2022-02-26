package azuredevops

import (
	"encoding/json"
	"errors"
	"reflect"
	"sonarci/decoration/azuredevops/models"
	"sonarci/testing/mocks"
	"testing"
)

func TestPullRequestDecorator_loadMyPullRequestThreads_CheckErrorOnRequest(t *testing.T) {
	wantError := errors.New("failure")

	mockConn := &mocks.MockConnection{
		RequestMock: func(route string) (<-chan []byte, <-chan error) {
			chError := make(chan error, 1)
			defer close(chError)

			chError <- wantError
			return nil, chError
		},
	}

	decorator := NewPullRequestDecorator(mockConn, &mocks.MockEngine{}, "project-test", "repo-test")
	_, gotErr := decorator.loadMyPullRequestThreads("anything")

	if gotErr != wantError {
		t.FailNow()
	}
}

func TestPullRequestDecorator_loadMyPullRequestThreads_CheckErrorOnReadResponse(t *testing.T) {
	mockConn := &mocks.MockConnection{
		RequestMock: func(route string) (<-chan []byte, <-chan error) {
			chBuff := make(chan []byte, 1) // buffered channel
			chErr := make(chan error)      // unbuffered channel - just to remember =)
			defer close(chBuff)
			defer close(chErr)

			chBuff <- []byte("any invalid json response")
			return chBuff, chErr
		},
	}

	decorator := NewPullRequestDecorator(mockConn, &mocks.MockEngine{}, "project-test", "repo-test")
	_, gotErr := decorator.loadMyPullRequestThreads("anything")

	if gotErr == nil {
		t.FailNow()
	}
}

func TestPullRequestDecorator_loadMyPullRequestThreads_CheckResult(t *testing.T) {
	threads := models.ThreadModelWrapper{
		Value: []models.ThreadModel{
			{Id: "1", IsDeleted: false, Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "True"}}},
			{Id: "2", IsDeleted: true, Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "True"}}},
			{Id: "3", IsDeleted: false, Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "true"}}},
			{Id: "4", IsDeleted: false, Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "false"}}},
			{Id: "5", IsDeleted: false, Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "anything"}}},
		},
	}
	wantThreads := []string{"1", "3"}

	mockConn := &mocks.MockConnection{
		RequestMock: func(route string) (<-chan []byte, <-chan error) {
			chBuff := make(chan []byte, 1)
			chErr := make(chan error)
			defer close(chBuff)
			defer close(chErr)

			resp, _ := json.Marshal(threads)
			chBuff <- resp
			return chBuff, chErr
		},
	}

	decorator := NewPullRequestDecorator(mockConn, &mocks.MockEngine{}, "project-test", "repo-test")
	gotThreads, err := decorator.loadMyPullRequestThreads("anything")

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(gotThreads, wantThreads) {
		t.Fail()
	}

}
