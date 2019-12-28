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

// CreateIssue returns the issue ID on success.
func (api *Api) CreateIssue(ctx context.Context, project, summary, description string) (string, error) {
	issue := &Issue{
		Summary:     summary,
		Description: description,
		Project: ProjectID{
			ID: project,
		},
	}

	var resp IDResult

	err := api.Post(ctx, &url.URL{Path: "issues"}, issue, &resp)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
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

	u, err := url.Parse(fmt.Sprintf("issues/%s/attachments", issueID))
	if err != nil {
		log.Print("Failed to parse attachments URL", err)
		return "", err
	}

	result := &IDResult{}
	err = api.Post(ctx, u, issueAttachment, result)
	if err != nil {
		log.Print("Failed to post attachment.", err)
		return "", err
	}

	return result.ID, nil
}
