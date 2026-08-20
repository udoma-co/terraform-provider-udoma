package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPublicChatbotInstructionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPublicChatbotInstructionConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "name", "Rent FAQs"),
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "instruction", "Always explain rent due dates first."),
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "priority", "10"),
				),
			},
			{
				ResourceName:      "udoma_public_chatbot_instruction.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
			resource "udoma_public_chatbot_instruction" "test" {
				name        = "Lease FAQs"
				instruction = "Always explain lease renewal timing first."
				priority    = 20
			}
			`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "name", "Lease FAQs"),
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "instruction", "Always explain lease renewal timing first."),
					resource.TestCheckResourceAttr("udoma_public_chatbot_instruction.test", "priority", "20"),
				),
			},
		},
	})
}

func testAccPublicChatbotInstructionConfig() string {
	return `
resource "udoma_public_chatbot_instruction" "test" {
  name        = "Rent FAQs"
  instruction = "Always explain rent due dates first."
  priority    = 10
}
`
}
