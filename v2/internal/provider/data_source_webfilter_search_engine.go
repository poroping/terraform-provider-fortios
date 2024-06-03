// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web filter search engines.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebfilterSearchEngine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web filter search engines.",

		ReadContext: dataSourceWebfilterSearchEngineRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"charset": {
				Type:        schema.TypeString,
				Description: "Search engine charset.",
				Computed:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "Hostname (regular expression).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Search engine name.",
				Required:    true,
			},
			"query": {
				Type:        schema.TypeString,
				Description: "Code used to prefix a query (must end with an equals character).",
				Computed:    true,
			},
			"safesearch": {
				Type:        schema.TypeString,
				Description: "Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.",
				Computed:    true,
			},
			"safesearch_str": {
				Type:        schema.TypeString,
				Description: "Safe search parameter used in the URL in URL mode. In translate mode, it provides either the regex to translate the URL or the special case to translate the URL.",
				Computed:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL (regular expression).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterSearchEngineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWebfilterSearchEngine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterSearchEngine dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectWebfilterSearchEngine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
