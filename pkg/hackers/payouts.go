package hackers

import (
	"context"
	"fmt"

	"github.com/sam-bee/security-hackerone-api-client/pkg/api"
)

type getPayoutsResponse struct {
	Data  []api.Payout `json:"data"`
	Links api.Links    `json:"links"`
}

// GetPayouts returns a list of payouts received by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetPayouts(ctx context.Context, pageOptions *api.PageOptions) (payouts []api.Payout, nextPage int, err error) {
	var response getPayoutsResponse
	path := fmt.Sprintf(
		"/hackers/payments/payouts?page[number]=%d&page[size]=%d",
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}
