package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDocumentResource(t *testing.T) {
	file := fmt.Sprintf("%s/test_file.txt", os.TempDir())
	if err := appendToTestFile(file); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	t.Cleanup(func() {
		defer os.Remove(file)
	})

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionDocument("", file),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "id"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "updated_at"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "attachment"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "name", "test_document"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "public", "false"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "ref_type", "static"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "ref_id", ""),
				),
			},
			{
				Config: resourceDefinitionDocument("folder", file),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "id"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "updated_at"),
					resource.TestCheckResourceAttrSet("udoma_document.test_document", "attachment"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "path", "folder"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "name", "test_document"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "public", "false"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "ref_type", "static"),
					resource.TestCheckResourceAttr("udoma_document.test_document", "ref_id", ""),
				),
			},
		},
	})
}

func resourceDefinitionDocument(path, file string) string {
	realPath := ""
	if path != "" {
		realPath = fmt.Sprintf("path = \"%s\"", path)
	}

	return `
	` + resourceDefinitionAttachment(file) + `
		
resource "udoma_document" "test_document" {
	name			 = "test_document"
	attachment = udoma_attachment.test_attachment.id
	public		 = false
	ref_type	 = "static"
	ref_id		 = ""
	` + realPath + `
}
	`
}
