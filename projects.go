package youtrack

import (
	"context"
	"net/url"
)

type Project struct {
	ID        string `json:"id"`
	ShortName string `json:"shortName"`
}

// ListProjects returns a list of Projects the user has access to.
// For more information, see
// https://www.jetbrains.com/help/youtrack/incloud/2019.3/resource-api-admin-projects.html
func (api *Api) ListProjects(ctx context.Context) ([]Project, error) {
	var Projects []Project
	u, err := url.Parse("admin/projects?fields=id,shortName")
	if err != nil {
		panic(err)
	}
	err = api.Get(ctx, u, &Projects)
	if err != nil {
		return nil, err
	}
	return Projects, nil
}

// Return the Project ID for the shortName (issue ID prefix in the YouTrack web UI)
func (api *Api) ProjectIDForShortName(ctx context.Context, shortName string) (string, error) {
	projects, err := api.ListProjects(ctx)
	if err != nil {
		return "", err
	}

	for _, p := range projects {
		if p.ShortName == shortName {
			return p.ID, nil
		}
	}

	api.trace(shortName, "project not found")
	return "", NotFoundError
}
