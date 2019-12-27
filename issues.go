package youtrack

import (
	"context"
)

type Issue struct {
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Project     ProjectID `json:"project"`
}

type createIssueResponse struct {
	IDReadable string `json:"idReadable"`
	dollarType string `json:"$type"`
}

func (api *Api) CreateIssueRequest(ctx context.Context, project, summary, description string) error {
	issue := &Issue{
		Summary:     summary,
		Description: description,
		Project: ProjectID{
			ID: project,
		},
	}

	var resp createIssueResponse

	err := api.Post(ctx, "issues", issue, &resp)
	if err != nil {
		return err
	}

	return nil
}
