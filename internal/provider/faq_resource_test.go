package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFAQResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionFAQ(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("udoma_faq.test", "keywords.0", "testing faq"),

					// Verify map attributes
					resource.TestCheckResourceAttr("udoma_faq.test", "question.de", "Testfragen"),
					resource.TestCheckResourceAttr("udoma_faq.test", "question.en", "Test Question"),

					resource.TestCheckResourceAttr("udoma_faq.test", "answer.de", "Testantwort"),
					resource.TestCheckResourceAttr("udoma_faq.test", "answer.en", "Test Answer"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_faq.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_faq.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: `
				resource udoma_faq "test" {
			
					question = {
						de = "Testfragen"
						en = "Test Question Updated"
					}
				
					answer = {
						de = "Testantwort"
						en = "Test Answer Updated"
					}
				
					keywords = ["testing faq updated"]
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_faq.test", "keywords.0", "testing faq updated"),
					resource.TestCheckResourceAttr("udoma_faq.test", "question.en", "Test Question Updated"),
					resource.TestCheckResourceAttr("udoma_faq.test", "answer.en", "Test Answer Updated"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionFAQ() string {
	return `
	resource udoma_faq "test" {

		question = {
			de = "Testfragen"
			en = "Test Question"
		}
	
		answer = {
			de = "Testantwort"
			en = "Test Answer"
		}
	
		keywords = ["testing faq"]
	}
	`
}
