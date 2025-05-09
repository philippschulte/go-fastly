package status

import (
	"strings"
	"testing"

	"github.com/fastly/go-fastly/v10/fastly"
)

func TestClient_DomainToolsStatus(t *testing.T) {
	t.Parallel()

	var err error
	var status *Status
	domain := "fastly-sdk-gofastly-testing.com"
	fastly.Record(t, "get", func(client *fastly.Client) {
		status, err = Get(client, &GetInput{
			Domain: domain,
		})
	})

	if err != nil {
		t.Fatal(err)
	}

	if status.Domain != domain {
		t.Errorf("incorrect domain, expected %s, got %s", domain, status.Domain)
	}

	if status.Zone != "com" {
		t.Errorf("incorrect zone, expected %s, got %s", "com", status.Zone)
	}

	if len(status.Offers) > 0 {
		t.Errorf("incorrect offers, expected %d, got %d", 0, len(status.Offers))
	}

	if !strings.Contains(status.Status, "inactive") {
		t.Errorf("incorrect status, expected %s within status, got %s", "inactive", status.Status)
	}

	if !strings.Contains(status.Tags, "generic") {
		t.Errorf("incorrect tags, expected %s within tags, got %s", "generic", status.Tags)
	}

	if status.Scope != nil {
		t.Errorf("incorrect scope, unexpected presence of scope, got %s", *status.Scope)
	}
}

func TestClient_DomainToolsStatusEstimate(t *testing.T) {
	t.Parallel()

	var err error
	var status *Status
	domain := "fastly-sdk-gofastly-testing.com"
	fastly.Record(t, "get_estimate", func(client *fastly.Client) {
		status, err = Get(client, &GetInput{
			Domain: domain,
			Scope:  fastly.ToPointer(ScopeEstimate),
		})
	})

	if err != nil {
		t.Fatal(err)
	}

	if status.Domain != domain {
		t.Errorf("incorrect domain, expected %s, got %s", domain, status.Domain)
	}

	if status.Zone != "com" {
		t.Errorf("incorrect zone, expected %s, got %s", "com", status.Zone)
	}

	if status.Scope == nil || *status.Scope != ScopeEstimate {
		t.Errorf("incorrect scope, expected %s, got %v", ScopeEstimate, status.Scope)
	}
}

func TestClient_DomainToolsStatusOffers(t *testing.T) {
	t.Parallel()

	var err error
	var status *Status
	domain := "domainr-testing.org"
	fastly.Record(t, "get_offers", func(client *fastly.Client) {
		status, err = Get(client, &GetInput{
			Domain: domain,
			Scope:  fastly.ToPointer(ScopeEstimate),
		})
	})

	if err != nil {
		t.Fatal(err)
	}

	if status.Domain != domain {
		t.Errorf("incorrect domain, expected %s, got %s", domain, status.Domain)
	}

	if status.Zone != "org" {
		t.Errorf("incorrect zone, expected %s, got %s", "org", status.Zone)
	}

	if status.Scope == nil || *status.Scope != ScopeEstimate {
		t.Errorf("incorrect scope, expected %s, got %v", ScopeEstimate, status.Scope)
	}

	if !strings.Contains(status.Status, "priced") {
		t.Errorf("incorrect status, expected %s within status, got %s", "priced", status.Status)
	}

	if len(status.Offers) == 0 {
		t.Errorf("incorrect offers, expected at least one offer, got %d", len(status.Offers))
	}

	if status.Offers[0].Currency != "USD" {
		t.Errorf("incorrect currency, expected %s, got %s", "USD", status.Offers[0].Currency)
	}
}
