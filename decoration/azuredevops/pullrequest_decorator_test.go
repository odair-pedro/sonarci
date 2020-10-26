package azuredevops

import "testing"

func Test_NewPullRequestDecorator(t *testing.T) {
	if got := NewPullRequestDecorator(nil, nil, "", ""); got == nil {
		t.Errorf("NewPullRequestDecorator() return nil")
	}
}
