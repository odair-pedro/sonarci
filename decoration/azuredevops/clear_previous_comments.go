package azuredevops

import (
	"encoding/json"
	"fmt"
	"sonarci/decoration/azuredevops/models"
	"strings"
)

const routeListPullRequestThreadsComments = "%s/_apis/git/repositories/%s/pullRequests/%s/threads?api-version=6.0"
const routeDeletePullRequestThreadComment = "%s/_apis/git/repositories/%s/pullRequests/%s/threads/%d/comments/%d?api-version=6.0"

func (decorator *PullRequestDecorator) ClearPreviousComments(pullRequest string, tag string) error {
	comments, err := decorator._loadMyPullRequestThreadsComments(pullRequest, tag)
	if err != nil {
		return err
	}

	if len(comments) > 0 {
		chErrDel := make(chan error, len(comments))
		defer close(chErrDel)

		for _, comment := range comments {
			go decorator._deletePullRequestThreadComment(comment, chErrDel)
			errDel := <-chErrDel
			if errDel != nil {
				return errDel
			}
		}
	}

	return nil
}

func (decorator *PullRequestDecorator) _loadMyPullRequestThreadsComments(pullRequest string, tag string) ([]_commentToDelete, error) {
	chBuff, chErr := decorator.Get(fmt.Sprintf(routeListPullRequestThreadsComments, formatPath(decorator.Project),
		formatPath(decorator.Repository), pullRequest))
	err := <-chErr
	if err != nil {
		return nil, err
	}

	buff := <-chBuff
	threadsWrapper := &models.ThreadModelWrapper{}
	err = json.Unmarshal(buff, threadsWrapper)
	if err != nil {
		return nil, err
	}

	var commentsToDelete []_commentToDelete
	for _, thread := range threadsWrapper.Value {
		if !thread.IsDeleted &&
			strings.ToLower(thread.Properties.GeneratedBySonarCI.Value) == "true" &&
			thread.Properties.Tag.Value == tag {
			for _, comment := range thread.Comments {
				if !comment.IsDeleted {
					commentsToDelete = append(commentsToDelete,
						_commentToDelete{PullRequest: pullRequest, CommentId: comment.Id, ThreadId: thread.Id})
				}
			}
		}
	}

	return commentsToDelete, nil
}

func (decorator *PullRequestDecorator) _deletePullRequestThreadComment(comment _commentToDelete, chErr chan<- error) {
	chErrDel := decorator.Connection.Delete(fmt.Sprintf(routeDeletePullRequestThreadComment, formatPath(decorator.Project),
		formatPath(decorator.Repository), comment.PullRequest, comment.ThreadId, comment.CommentId))

	errDel := <-chErrDel
	chErr <- errDel
}

type _commentToDelete struct {
	PullRequest string
	CommentId   int
	ThreadId    int
}
