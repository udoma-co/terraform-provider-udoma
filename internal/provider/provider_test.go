package provider

import (
	"os"
	"testing"

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

func TestMain(m *testing.M) {

	os.Setenv("TF_ACC", "1")
	os.Setenv("TF_LOG", "TRACE")

	os.Setenv("UDOMA_ENDPOINT", "http://hostlocal")
	os.Setenv("UDOMA_API_KEY", "AAAAAAAA053HK3YXTO1EVFHSMSXE")
	os.Setenv("UDOMA_KEY_SECRET", "Hpz_VUJ1j-UA_eW.UxhD5EpYQdUK_1owzfN7YSOGfT.ZnuAnUDwzz62hCY.1.OMg")

	// Run the tests.
	os.Exit(m.Run())
}
