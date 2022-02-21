package fastly

import (
	"testing"
)

func TestClient_GetStats(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/service_stats", func(c *Client) {
		_, err = c.GetStats(&GetStatsInput{
			Service: testServiceID,
			From:    "10 days ago",
			To:      "now",
			By:      "minute",
			Region:  "europe",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetStats_ByField(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/service_stats_by_field", func(c *Client) {
		_, err = c.GetStatsField(&GetStatsInput{
			Field:  "bandwidth",
			From:   "1 hour ago",
			To:     "now",
			By:     "minute",
			Region: "europe",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetStats_ByFieldAndService(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/service_stats_by_field_and_service", func(c *Client) {
		_, err = c.GetStats(&GetStatsInput{
			Service: testServiceID,
			Field:   "bandwidth",
			From:    "10 days ago",
			To:      "now",
			By:      "day",
			Region:  "usa",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetStatsJSON(t *testing.T) {
	t.Parallel()

	var ret struct {
		RenameStatus string `json:"status"`
	}

	var err error
	record(t, "stats/service_stats", func(c *Client) {
		err = c.GetStatsJSON(&GetStatsInput{
			Service: testServiceID,
			From:    "10 days ago",
			To:      "now",
			By:      "minute",
			Region:  "europe",
		}, &ret)
	})
	if err != nil {
		t.Fatal(err)
	}

	if ret.RenameStatus != "success" {
		t.Fatalf("got RenameStatus=%q, want %q", ret.RenameStatus, "success")
	}
}

func TestClient_GetRegions(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/regions", func(c *Client) {
		_, err = c.GetRegions()
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetRegionsUsage(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/regions_usage", func(c *Client) {
		_, err = c.GetUsage(&GetUsageInput{
			From:   "10 days ago",
			To:     "now",
			By:     "minute",
			Region: "usa",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetServicesByRegionsUsage(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "stats/services_usage", func(c *Client) {
		_, err = c.GetUsageByService(&GetUsageInput{
			From:   "10 days ago",
			To:     "now",
			By:     "minute",
			Region: "usa",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}
