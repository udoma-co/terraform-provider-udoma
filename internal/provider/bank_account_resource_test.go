package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBankAccountResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionBankAccount(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("udoma_bank_account.test", "account_holder", "test_account"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "iban", "DE29100100100987654321"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "bic", "TESTDETTXXX"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "bank_name", "test_bank"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "description", "test_description"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_bank_account.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_bank_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: `
				resource udoma_bank_account "test" {
					account_holder = "updated_test_account"
					iban 					 = "DE29100100100987654321"
					bic 					 = "TESTDETTXXX"
					bank_name 		 = "updated_test_bank"
					description 	 = "updated_test_description"
					cadence        = "monthly_cadence"
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_bank_account.test", "account_holder", "updated_test_account"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "bank_name", "updated_test_bank"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "description", "updated_test_description"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionBankAccount() string {
	return `
	resource udoma_bank_account "test" {
		account_holder = "test_account"
		iban 					 = "DE29100100100987654321"
		bic 					 = "TESTDETTXXX"
		bank_name 		 = "test_bank"
		description 	 = "test_description"
		cadence        = "monthly_cadence"
	}
	`
}
