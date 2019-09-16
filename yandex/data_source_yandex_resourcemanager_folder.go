package yandex

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/sdkresolvers"
)

const yandexResourceManagerFolderDefaultTimeout = 1 * time.Minute

func dataSourceYandexResourceManagerFolder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceYandexResourceManagerFolderRead,
		Schema: map[string]*schema.Schema{
			"folder_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type cloudID string

func (id cloudID) folderResolver(name string, opts ...sdkresolvers.ResolveOption) ycsdk.Resolver {
	opts = append(opts, sdkresolvers.CloudID(string(id)))
	return sdkresolvers.FolderResolver(name, opts...)
}

func dataSourceYandexResourceManagerFolderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ctx := config.ContextWithClientTraceID()

	err := checkOneOf(d, "folder_id", "name")
	if err != nil {
		return err
	}

	folderID := d.Get("folder_id").(string)
	folderName, folderNameOk := d.GetOk("name")

	if folderNameOk {
		dsCloudID, err := getCloudID(d, config)
		if err != nil {
			return fmt.Errorf("error getting cloud ID to resolve data source for folder: %s", err)
		}

		resolver := cloudID(dsCloudID).folderResolver
		folderID, err = resolveObjectID(ctx, config, folderName.(string), resolver)
		if err != nil {
			return fmt.Errorf("failed to resolve data source folder by name: %v", err)
		}
	}

	folder, err := config.sdk.ResourceManager().Folder().Get(ctx, &resourcemanager.GetFolderRequest{
		FolderId: folderID,
	})

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("folder with ID %q", folderID))
	}

	createdAt, err := getTimestamp(folder.CreatedAt)
	if err != nil {
		return err
	}

	d.Set("folder_id", folder.Id)
	d.Set("name", folder.Name)
	d.Set("description", folder.Description)
	d.Set("cloud_id", folder.CloudId)
	d.Set("status", strings.ToLower(folder.Status.String()))
	d.Set("created_at", createdAt)

	if err := d.Set("labels", folder.Labels); err != nil {
		return err
	}

	d.SetId(folder.Id)

	return nil
}
