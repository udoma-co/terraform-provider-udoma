package client

import (
	"context"
	"testing"

	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestConfig(t *testing.T) {

	cfg := Config{
		Endpoint:     "http://hostlocal",
		ApiKey:       "CBAAAAAATG919N5RAPBZ75UF8WF7",
		ApiSecretKey: "NR3qQ.kppGek19q-ZIkXzmSFyJDxhR1gTyhdwh8nFj4_TK8d.jpaMOQyuVxGjJHJ",
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
