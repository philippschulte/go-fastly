package fastly

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// Purge is a response from a purge request.
type Purge struct {
	// PurgeID is the unique PurgeID of the purge request.
	PurgeID *string `mapstructure:"id"`
	// Status is the status of the purge, usually "ok".
	Status *string `mapstructure:"status"`
}

// PurgeInput is used as input to the Purge function.
type PurgeInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// Soft performs a soft purge.
	Soft bool
	// URL is the URL to purge (required).
	URL string
}

// Purge instantly purges an individual URL.
func (c *Client) Purge(i *PurgeInput) (*Purge, error) {
	if i.URL == "" {
		return nil, ErrMissingURL
	}

	requestOptions := CreateRequestOptions(i.Context)
	requestOptions.Parallel = true
	if i.Soft {
		requestOptions.Headers = map[string]string{
			"Fastly-Soft-Purge": "1",
		}
	}

	var err error
	// nosemgrep: trailofbits.go.questionable-assignment.questionable-assignment
	requestOptions.Params, err = constructRequestOptionsParam(i.URL)
	if err != nil {
		return nil, err
	}

	resp, err := c.Post("purge/"+i.URL, requestOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r *Purge
	if err := DecodeBodyMap(resp.Body, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// constructRequestOptionsParam prevents Client.RawRequest from incorrectly
// clearing the URL query string, by manually constructing a ro.Params.
func constructRequestOptionsParam(us string) (map[string]string, error) {
	m := make(map[string]string)
	u, err := url.Parse(us)
	if err != nil {
		return nil, err
	}
	v := u.Query()
	// NOTE: we can't coerce a url.Values into the underlying map[string]string
	// type, so we have to manually loop over the url.Values and copy the
	// key/value pairs into a new map instance.
	for k, v := range v {
		m[k] = v[0]
	}
	return m, nil
}

// PurgeKeyInput is used as input to the PurgeKey function.
type PurgeKeyInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// Key is the key to purge (required).
	Key string
	// ServiceID is the ID of the service (required).
	ServiceID string
	// Soft performs a soft purge.
	Soft bool
}

// PurgeKey instantly purges a particular service of items tagged with a key.
func (c *Client) PurgeKey(i *PurgeKeyInput) (*Purge, error) {
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if i.Key == "" {
		return nil, ErrMissingKey
	}

	path := ToSafeURL("service", i.ServiceID, "purge", i.Key)

	requestOptions := CreateRequestOptions(i.Context)
	requestOptions.Parallel = true

	req, err := c.RawRequest(http.MethodPost, path, requestOptions)
	if err != nil {
		return nil, err
	}

	if i.Soft {
		req.Header.Set("Fastly-Soft-Purge", "1")
	}

	resp, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r *Purge
	if err := DecodeBodyMap(resp.Body, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// PurgeKeysInput is used as input to the PurgeKeys function.
type PurgeKeysInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// Keys are the keys to purge (required).
	Keys []string
	// ServiceID is the ID of the service (required).
	ServiceID string
	// Soft performs a soft purge.
	Soft bool
}

// PurgeKeys instantly purges a particular service of items tagged with a key.
func (c *Client) PurgeKeys(i *PurgeKeysInput) (map[string]string, error) {
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if len(i.Keys) == 0 {
		return nil, ErrMissingKeys
	}

	path := ToSafeURL("service", i.ServiceID, "purge")

	requestOptions := CreateRequestOptions(i.Context)
	requestOptions.Parallel = true

	req, err := c.RawRequest(http.MethodPost, path, requestOptions)
	if err != nil {
		return nil, err
	}

	if i.Soft {
		req.Header.Set("Fastly-Soft-Purge", "1")
	}

	req.Header.Set("Surrogate-Key", strings.Join(i.Keys, " "))

	resp, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r map[string]string
	if err := DecodeBodyMap(resp.Body, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// PurgeAllInput is used as input to the Purge function.
type PurgeAllInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// ServiceID is the ID of the service (required).
	ServiceID string
}

// PurgeAll instantly purges everything from a service.
func (c *Client) PurgeAll(i *PurgeAllInput) (*Purge, error) {
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}

	path := ToSafeURL("service", i.ServiceID, "purge_all")

	req, err := c.RawRequest(http.MethodPost, path, CreateRequestOptions(i.Context))
	if err != nil {
		return nil, err
	}

	resp, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r *Purge
	if err := DecodeBodyMap(resp.Body, &r); err != nil {
		return nil, err
	}
	return r, nil
}
