// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Replacement messages.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemReplacemsgDeviceDetectionPortalList() *schema.Resource {
	return &schema.Resource{
		Description: "Replacement messages.",

		ReadContext: dataSourceSystemReplacemsgDeviceDetectionPortalListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"filter": {
				Type:        schema.TypeString,
				Description: "Filter to apply to query",
				Optional:    true,
				ForceNew:    true,
			},
			"msg_typelist": {
				Type:        schema.TypeList,
				Description: "List of results",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSystemReplacemsgDeviceDetectionPortalListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	setfilter := ""
	filter := []string{}
	if v, ok := d.GetOk("filter"); ok {
		if s, ok := v.(string); ok {
			setfilter = s
			filter = []string{s}
		}
	}
	urlparams.Filter = &filter

	format := []string{"msg-type"}
	urlparams.Format = &format

	o, err := c.Cmdb.ListSystemReplacemsgDeviceDetectionPortal(urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgDeviceDetectionPortal dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if len(*o) == 0 {
		return nil
	}

	results := []string{}
	for _, v := range *o {
		v2 := v.MsgType
		results = append(results, *v2)
	}

	d.Set("msg_typelist", results)

	d.SetId("DataSourceSystemReplacemsgDeviceDetectionPortalList" + setfilter)

	return nil
}
