package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAttachmentResource(t *testing.T) {
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
				Config: resourceDefinitionAttachment(file),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "id"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "created"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "file_type"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "file_size"),
					resource.TestCheckResourceAttr("udoma_attachment.test_attachment", "file_name", "test_file.txt"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "file_sha256"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "url"),
				),
			},
			{
				PreConfig: func() {
					_ = appendToTestFile(file)
				},
				Config: resourceDefinitionAttachment(file),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "id"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "created"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "file_type"),
					resource.TestCheckResourceAttr("udoma_attachment.test_attachment", "file_size", "8"),
					resource.TestCheckResourceAttr("udoma_attachment.test_attachment", "file_name", "test_file.txt"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "file_sha256"),
					resource.TestCheckResourceAttrSet("udoma_attachment.test_attachment", "url"),
				),
			},
		},
	})
}

func appendToTestFile(file string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	tflog.Info(context.Background(), "Appending to file: "+file)

	if _, err = f.WriteString("test"); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func resourceDefinitionAttachment(file string) string {
	return `
resource "udoma_attachment" "test_attachment" {
	source = "` + file + `"
}
`
}
