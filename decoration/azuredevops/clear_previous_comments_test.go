package azuredevops

import (
	"encoding/json"
	"errors"
	"reflect"
	"sonarci/decoration/azuredevops/models"
	"sonarci/testing/mocks"
	"testing"
)

func TestPullRequestDecorator_loadMyPullRequestThreadsComments_CheckErrorOnRequest(t *testing.T) {
	wantError := errors.New("failure")

	mockConn := &mocks.MockConnection{
		GetMock: func(route string) (<-chan []byte, <-chan error) {
			chError := make(chan error, 1)
			defer close(chError)

			chError <- wantError
			return nil, chError
		},
	}

	decorator := NewPullRequestDecorator(mockConn, &mocks.MockEngine{}, "project-test", "repo-test")
	_, gotErr := decorator.loadMyPullRequestThreadsComments("anything")

	if gotErr != wantError {
		t.FailNow()
	}
}

func TestPullRequestDecorator_loadMyPullRequestThreadsComments_CheckErrorOnReadResponse(t *testing.T) {
	mockConn := &mocks.MockConnection{
		GetMock: func(route string) (<-chan []byte, <-chan error) {
			chBuff := make(chan []byte, 1) // buffered channel
			chErr := make(chan error)      // unbuffered channel - just to remember =)
			defer close(chBuff)
			defer close(chErr)

			chBuff <- []byte("any invalid json response")
			return chBuff, chErr
		},
	}

	decorator := NewPullRequestDecorator(mockConn, &mocks.MockEngine{}, "project-test", "repo-test")
	_, gotErr := decorator.loadMyPullRequestThreadsComments("anything")

	if gotErr == nil {
		t.FailNow()
	}
}

func TestPullRequestDecorator_loadMyPullRequestThreadsComments_CheckResult(t *testing.T) {
	threads := models.ThreadModelWrapper{
		Value: []models.ThreadModel{
			{Id: "1", IsDeleted: false,
				Comments: []models.ThreadCommentModel{
					{Id: "10", IsDeleted: false},
					{Id: "11", IsDeleted: true},
					{Id: "12", IsDeleted: false},
				},
				Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "True"}},
			},
			{Id: "2", IsDeleted: true,
				Comments: []models.ThreadCommentModel{
					{Id: "20", IsDeleted: false},
					{Id: "21", IsDeleted: false},
					{Id: "22", IsDeleted: false},
				},
				Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "True"}},
			},
			{Id: "3", IsDeleted: false,
				Comments: []models.ThreadCommentModel{
					{Id: "31", IsDeleted: false},
					{Id: "32", IsDeleted: false},
				},
				Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "true"}},
			},
			{Id: "4", IsDeleted: false,
				Comments: []models.ThreadCommentModel{
					{Id: "41", IsDeleted: false},
					{Id: "42", IsDeleted: false},
				},
				Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "false"}},
			},
			{Id: "5", IsDeleted: false,
				Comments: []models.ThreadCommentModel{
					{Id: "51", IsDeleted: false},
					{Id: "52", IsDeleted: false},
				},
				Properties: models.ThreadPropertyModel{GeneratedBySonarCI: models.ThreadPropertySonarCIModel{Value: "anything"}},
			},
		},
	}

	pullRequest := "pull-request-test"
	wantThreads := []commentToDelete{
		{PullRequest: pullRequest, ThreadId: "1", CommentId: "10"},
		{PullRequest: pullRequest, ThreadId: "1", CommentId: "12"},
		{PullRequest: pullRequest, ThreadId: "3", CommentId: "31"},
		{PullRequest: pullRequest, ThreadId: "3", CommentId: "32"},
	}

	mockConn := &mocks.MockConnection{
		GetMock: func(route string) (<-chan []byte, <-chan error) {
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
	gotThreads, err := decorator.loadMyPullRequestThreadsComments(pullRequest)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(gotThreads, wantThreads) {
		t.Fail()
	}

}
