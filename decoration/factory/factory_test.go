package factory

import (
	"sonarci/decoration"
	"sonarci/decoration/azuredevops"
	"testing"
)

func Test_CreatePullRequestDecorator_AzRepos_CheckResult(t *testing.T) {
	decorator, _ := CreatePullRequestDecorator("azrepos", "project", "repository", nil, func(server string) decoration.Connection {
		return nil
	})
	if decorator == nil {
		t.Errorf("CreatePullRequestDecorator() == nil")
	}
}

func Test_CreatePullRequestDecorator_AzRepos_CheckResultType(t *testing.T) {
	decorator, _ := CreatePullRequestDecorator("azrepos", "project", "repository", nil, func(server string) decoration.Connection {
		return nil
	})

	switch tp := decorator.(type) {
	case *azuredevops.PullRequestDecorator:
		return
	default:
		t.Errorf("Invalid returned type: %T", tp)
	}
}

func Test_CreatePullRequestDecorator_AzRepos_CheckError(t *testing.T) {
	_, err := CreatePullRequestDecorator("azrepos", "project", "repository", nil, func(server string) decoration.Connection {
		return nil
	})
	if err != nil {
		t.Errorf("CreatePullRequestDecorator() error != nil")
	}
}

func Test_CreatePullRequestDecorator_GitHub_CheckResult(t *testing.T) {
	decorator, _ := CreatePullRequestDecorator("github", "project", "repository", nil, nil)
	if decorator != nil {
		t.Errorf("CreatePullRequestDecorator() != nil")
	}
}

func Test_CreatePullRequestDecorator_GitHub_CheckError(t *testing.T) {
	_, err := CreatePullRequestDecorator("github", "project", "repository", nil, nil)
	if err == nil {
		t.Errorf("CreatePullRequestDecorator() error == nil")
	}
}

func Test_CreatePullRequestDecorator_Default_CheckResult(t *testing.T) {
	decorator, _ := CreatePullRequestDecorator("default", "project", "repository", nil, nil)
	if decorator != nil {
		t.Errorf("CreatePullRequestDecorator() != nil")
	}
}

func Test_CreatePullRequestDecorator_Default_CheckError(t *testing.T) {
	_, err := CreatePullRequestDecorator("default", "project", "repository", nil, nil)
	if err == nil {
		t.Errorf("CreatePullRequestDecorator() error == nil")
	}
}
