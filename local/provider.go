package local

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"local_file":      resourceLocalFile(),
			"local_directory": resourceLocalDirectory(),
			"local_symlink":   resourceLocalSymlink(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"local_file":      dataSourceLocalFile(),
			"local_directory": dataSourceLocalDirectory(),
			"local_symlink":   dataSourceLocalSymlink(),
		},
	}
}
