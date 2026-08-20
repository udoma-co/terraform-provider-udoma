package provider

import (
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"udoma": providerserver.NewProtocol6WithError(New("test")()),
	}
)

func testAccAccountRef() string {
	accountRef := os.Getenv("UDOMA_ACCOUNT_REF")
	if accountRef == "" {
		panic("UDOMA_ACCOUNT_REF environment variable is required for catalog acceptance tests")
	}
	return accountRef
}
