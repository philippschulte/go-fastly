package fastly

import (
	"context"
	"strconv"
	"time"
)

// Condition represents a condition response from the Fastly API.
type Condition struct {
	Comment        *string    `mapstructure:"comment"`
	CreatedAt      *time.Time `mapstructure:"created_at"`
	DeletedAt      *time.Time `mapstructure:"deleted_at"`
	Name           *string    `mapstructure:"name"`
	Priority       *int       `mapstructure:"priority"`
	ServiceID      *string    `mapstructure:"service_id"`
	ServiceVersion *int       `mapstructure:"version"`
	Statement      *string    `mapstructure:"statement"`
	Type           *string    `mapstructure:"type"`
	UpdatedAt      *time.Time `mapstructure:"updated_at"`
}

// ListConditionsInput is used as input to the ListConditions function.
type ListConditionsInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// ServiceID is the ID of the service (required).
	ServiceID string
	// ServiceVersion is the specific configuration version (required).
	ServiceVersion int
}

// ListConditions retrieves all resources.
func (c *Client) ListConditions(i *ListConditionsInput) ([]*Condition, error) {
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if i.ServiceVersion == 0 {
		return nil, ErrMissingServiceVersion
	}

	path := ToSafeURL("service", i.ServiceID, "version", strconv.Itoa(i.ServiceVersion), "condition")
	resp, err := c.Get(path, CreateRequestOptions(i.Context))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cs []*Condition
	if err := DecodeBodyMap(resp.Body, &cs); err != nil {
		return nil, err
	}
	return cs, nil
}

// CreateConditionInput is used as input to the CreateCondition function.
type CreateConditionInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context `url:"-"`
	// Name is the name of the condition.
	Name *string `url:"name,omitempty"`
	// Priority is a numeric string. Priority determines execution order. Lower numbers execute first.
	Priority *int `url:"priority,omitempty"`
	// ServiceID is the ID of the service (required).
	ServiceID string `url:"-"`
	// ServiceVersion is the specific configuration version (required).
	ServiceVersion int `url:"-"`
	// Statement is a conditional expression in VCL used to determine if the condition is met.
	Statement *string `url:"statement,omitempty"`
	// Type is the type of the condition (REQUEST, CACHE, RESPONSE, PREFETCH).
	Type *string `url:"type,omitempty"`
}

// CreateCondition creates a new resource.
func (c *Client) CreateCondition(i *CreateConditionInput) (*Condition, error) {
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if i.ServiceVersion == 0 {
		return nil, ErrMissingServiceVersion
	}

	path := ToSafeURL("service", i.ServiceID, "version", strconv.Itoa(i.ServiceVersion), "condition")
	resp, err := c.PostForm(path, i, CreateRequestOptions(i.Context))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var co *Condition
	if err := DecodeBodyMap(resp.Body, &co); err != nil {
		return nil, err
	}
	return co, nil
}

// GetConditionInput is used as input to the GetCondition function.
type GetConditionInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// Name is the name of the condition to fetch (required).
	Name string
	// ServiceID is the ID of the service (required).
	ServiceID string
	// ServiceVersion is the specific configuration version (required).
	ServiceVersion int
}

// GetCondition retrieves the specified resource.
func (c *Client) GetCondition(i *GetConditionInput) (*Condition, error) {
	if i.Name == "" {
		return nil, ErrMissingName
	}
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if i.ServiceVersion == 0 {
		return nil, ErrMissingServiceVersion
	}

	path := ToSafeURL("service", i.ServiceID, "version", strconv.Itoa(i.ServiceVersion), "condition", i.Name)
	resp, err := c.Get(path, CreateRequestOptions(i.Context))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var co *Condition
	if err := DecodeBodyMap(resp.Body, &co); err != nil {
		return nil, err
	}
	return co, nil
}

// UpdateConditionInput is used as input to the UpdateCondition function.
type UpdateConditionInput struct {
	// Comment is a freeform descriptive note.
	Comment *string `url:"comment,omitempty"`
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context `url:"-"`
	// Name is the name of the condition to update (required).
	Name string `url:"-"`
	// Priority is a numeric string. Priority determines execution order. Lower numbers execute first.
	Priority *int `url:"priority,omitempty"`
	// ServiceID is the ID of the service (required).
	ServiceID string `url:"-"`
	// ServiceVersion is the specific configuration version (required).
	ServiceVersion int `url:"-"`
	// Statement is a conditional expression in VCL used to determine if the condition is met.
	Statement *string `url:"statement,omitempty"`
	// Type is the type of the condition (REQUEST, CACHE, RESPONSE, PREFETCH).
	Type *string `url:"type,omitempty"`
}

// UpdateCondition updates the specified resource.
func (c *Client) UpdateCondition(i *UpdateConditionInput) (*Condition, error) {
	if i.Name == "" {
		return nil, ErrMissingName
	}
	if i.ServiceID == "" {
		return nil, ErrMissingServiceID
	}
	if i.ServiceVersion == 0 {
		return nil, ErrMissingServiceVersion
	}

	path := ToSafeURL("service", i.ServiceID, "version", strconv.Itoa(i.ServiceVersion), "condition", i.Name)
	resp, err := c.PutForm(path, i, CreateRequestOptions(i.Context))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var co *Condition
	if err := DecodeBodyMap(resp.Body, &co); err != nil {
		return nil, err
	}
	return co, nil
}

// DeleteConditionInput is the input parameter to DeleteCondition.
type DeleteConditionInput struct {
	// Context, if supplied, will be used as the Request's context.
	Context *context.Context
	// Name is the name of the condition to delete (required).
	Name string
	// ServiceID is the ID of the service (required).
	ServiceID string
	// ServiceVersion is the specific configuration version (required).
	ServiceVersion int
}

// DeleteCondition deletes the specified resource.
func (c *Client) DeleteCondition(i *DeleteConditionInput) error {
	if i.Name == "" {
		return ErrMissingName
	}
	if i.ServiceID == "" {
		return ErrMissingServiceID
	}
	if i.ServiceVersion == 0 {
		return ErrMissingServiceVersion
	}

	path := ToSafeURL("service", i.ServiceID, "version", strconv.Itoa(i.ServiceVersion), "condition", i.Name)
	resp, err := c.Delete(path, CreateRequestOptions(i.Context))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r *statusResp
	if err := DecodeBodyMap(resp.Body, &r); err != nil {
		return err
	}
	if !r.Ok() {
		return ErrNotOK
	}
	return nil
}
