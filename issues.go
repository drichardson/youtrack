package youtrack

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

type ProjectID struct {
	ID string `json:"id"`
}

type Issue struct {
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Project     ProjectID `json:"project"`
}

type IDResult struct {
	ID string `json:"id"`
}

type IssueResult struct {
	IDResult
	NumberInProject int `json:"numberInProject"`
}

// IssueURL returns a user facing (rather than REST API) URL to the issue.
// Note that because this issue uses the short project name in the URL, the link
// could be broken if the project short name changes.
func IssueURL(baseURL *url.URL, shortProjectName string, issueNumberInProject int) *url.URL {
	path := fmt.Sprintf("../issue/%s-%d", shortProjectName, issueNumberInProject)
	return baseURL.ResolveReference(&url.URL{Path: path})
}

// CreateIssue returns the issue ID on success.
func (api *Api) CreateIssue(ctx context.Context, project, summary, description string) (*IssueResult, error) {
	issue := &Issue{
		Summary:     summary,
		Description: description,
		Project: ProjectID{
			ID: project,
		},
	}

	result := new(IssueResult)
	u := &url.URL{
		Path:     "issues",
		RawQuery: "fields=id,numberInProject",
	}

	err := api.Post(ctx, u, issue, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type IssueAttachment struct {
	Name          string `json:"name"`
	Base64Content string `json:"base64Content"`
}

// CreateIssueAttachment attached a file to the given issue ID. On success, the attachment ID is returned.
func (api *Api) CreateIssueAttachment(ctx context.Context, issueID string, attachment io.Reader, name, mediaType string) (string, error) {

	data, err := ioutil.ReadAll(attachment)
	if err != nil {
		log.Print("Failed to read attachment", err)
		return "", err
	}

	issueAttachment := &IssueAttachment{
		Name:          name,
		Base64Content: "data:" + mediaType + ";base64," + base64.StdEncoding.EncodeToString(data),
	}

	u := &url.URL{
		Path: fmt.Sprintf("issues/%s/attachments", issueID),
	}
	result := &IDResult{}

	err = api.Post(ctx, u, issueAttachment, result)
	if err != nil {
		log.Print("Failed to post attachment.", err)
		return "", err
	}

	return result.ID, nil
}
