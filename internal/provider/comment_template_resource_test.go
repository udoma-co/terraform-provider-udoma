package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCommentTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionCommentTemplate("some display name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "display_name", "some display name"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.#", "2"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.0", "ct-somethingsomething"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.1", "ct-somethingsomethinn"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "is_deny_list", "true"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "template", "Hello, world!"),
				),
			},
			{
				ResourceName:      "udoma_comment_template.test_comment_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: resourceDefinitionCommentTemplate("other display name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "display_name", "other display name"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.#", "2"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.0", "ct-somethingsomething"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "access_list.1", "ct-somethingsomethinn"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "is_deny_list", "true"),
					resource.TestCheckResourceAttr("udoma_comment_template.test_comment_template", "template", "Hello, world!"),
				),
			},
		},
	})
}

func resourceDefinitionCommentTemplate(displayName string) string {
	return `
resource "udoma_comment_template" "test_comment_template" {
	display_name = "` + displayName + `"
	access_list = ["ct-somethingsomething", "ct-somethingsomethinn"]
	is_deny_list = true
	template = "Hello, world!"
}
`
}
