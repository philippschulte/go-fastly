package accesskeys

import (
	"encoding/json"
	"fmt"

	"github.com/fastly/go-fastly/v10/fastly"
)

// ListAccessKeys retrieves all access keys within object storage.
func ListAccessKeys(c *fastly.Client) (*AccessKeys, error) {
	resp, err := c.Get("/resources/object-storage/access-keys", fastly.CreateRequestOptions(nil))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accessKeys *AccessKeys

	if err := json.NewDecoder(resp.Body).Decode(&accessKeys); err != nil {
		return nil, fmt.Errorf("failed to decode json response: %w", err)
	}

	return accessKeys, nil
}
