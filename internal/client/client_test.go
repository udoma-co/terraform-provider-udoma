package client

import (
	"context"
	"os"
	"testing"

	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestConfig(t *testing.T) {

	cfg := Config{
		Endpoint:     os.Getenv("UDOMA_ENDPOINT"),
		ApiKey:       os.Getenv("UDOMA_API_KEY"),
		ApiSecretKey: os.Getenv("UDOMA_KEY_SECRET"),
	}

	cl, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	api := cl.GetApi()

	ctx := context.Background()

	r := api.GetAccountSummary(ctx).GetSummaryRequest(v1.GetSummaryRequest{})

	_, _, err = r.Execute()
	if err != nil {
		t.Fatal(err)
	}

}
