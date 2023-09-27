package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccJsonSourceDataSource(t *testing.T) {

	expectedContent := `{
  "id": "1",
  "source": "const greeting = 'Hello, world!';\nconst name = 'John Doe';\nconst message = greeting + \" My name is \" + name;\n\n// log message to console\nconsole.log(message);",
  "elements": [
  {
    "id": "elem-1",
    "name": "Element 1"
  },
  {
    "id": "elem-2",
    "name": "Element 2"
  }
]
}`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: `
data udoma_json_source "test" {
	source = "../../testdata/json_source/definition.json"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("data.udoma_json_source.test", "content", expectedContent),
				),
			},
		},
	})
}
