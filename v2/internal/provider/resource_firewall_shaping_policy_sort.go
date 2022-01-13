package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
)

// Doc notes: Will sort, if sort field equal will retain order. Name/Comment are sorted as string, do 0001-x 0010-x.
func resourceFirewallShapingPolicySort() *schema.Resource {
	return &schema.Resource{
		Description: "Sort firewall shaping policies.",

		CreateContext: resourceFirewallShapingPolicySortCreate,
		ReadContext:   resourceFirewallShapingPolicySortRead,
		UpdateContext: resourceFirewallShapingPolicySortCreate,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"sortby": {
				Description:  "Attribute to sort by.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"name", "policyid", "comments"}, false),
			},
			"sortdirection": {
				Description:  "Sort direction",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ascending", "descending"}, false),
			},
			"sorted": {
				Description: "Policy set is sorted.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"auto_sort": {
				Description: "If true will attempt to auto_sort if policy list no longer sorted.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"force_recreate": {
				Description: "Change field to force a recreation.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceFirewallShapingPolicySortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	err := c.Utils.SortFirewallPolicy(urlparams, sortby, sortdirection)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sortby + sortdirection)

	return resourceFirewallShapingPolicySortRead(ctx, d, meta)
}

func resourceFirewallShapingPolicySortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	sorted, err := c.Utils.FirewallPolicyListIsSorted(urlparams, sortby, sortdirection)
	if err != nil {
		return diag.FromErr(err)
	}

	auto_sort := d.Get("auto_sort").(bool)

	if !sorted && auto_sort {
		d.Set("auto_sort", sorted)
	}
	d.Set("sorted", sorted)

	return nil
}
