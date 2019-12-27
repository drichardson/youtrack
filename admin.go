package youtrack

import (
	"context"
)

type ProjectID struct {
	ID string `json:"id"`
}

func (api *Api) lookupProjectIdFromShortName(ctx context.Context, shortName string) (string, error) {
	panic("no")
	return "", nil
}
