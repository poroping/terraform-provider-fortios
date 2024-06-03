// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure network visibility settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemNetworkVisibility() *schema.Resource {
	return &schema.Resource{
		Description: "Configure network visibility settings.",

		CreateContext: resourceSystemNetworkVisibilityCreate,
		ReadContext:   resourceSystemNetworkVisibilityRead,
		UpdateContext: resourceSystemNetworkVisibilityUpdate,
		DeleteContext: resourceSystemNetworkVisibilityDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"destination_hostname_visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of destination hostname visibility.",
				Optional:    true,
				Computed:    true,
			},
			"destination_location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of destination geographical location visibility.",
				Optional:    true,
				Computed:    true,
			},
			"destination_visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of destination visibility.",
				Optional:    true,
				Computed:    true,
			},
			"hostname_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50000),

				Description: "Limit of the number of hostname table entries (0 - 50000).",
				Optional:    true,
				Computed:    true,
			},
			"hostname_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "TTL of hostname table entries (60 - 86400).",
				Optional:    true,
				Computed:    true,
			},
			"source_location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of source geographical location visibility.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemNetworkVisibilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	obj, diags := getObjectSystemNetworkVisibility(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemNetworkVisibility(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNetworkVisibility")
	}

	return resourceSystemNetworkVisibilityRead(ctx, d, meta)
}

func resourceSystemNetworkVisibilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectSystemNetworkVisibility(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemNetworkVisibility(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemNetworkVisibility resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNetworkVisibility")
	}

	return resourceSystemNetworkVisibilityRead(ctx, d, meta)
}

func resourceSystemNetworkVisibilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	obj, diags := getEmptyObjectSystemNetworkVisibility(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemNetworkVisibility(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemNetworkVisibility resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetworkVisibilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadSystemNetworkVisibility(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNetworkVisibility resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectSystemNetworkVisibility(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemNetworkVisibility(d *schema.ResourceData, o *models.SystemNetworkVisibility, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DestinationHostnameVisibility != nil {
		v := *o.DestinationHostnameVisibility

		if err = d.Set("destination_hostname_visibility", v); err != nil {
			return diag.Errorf("error reading destination_hostname_visibility: %v", err)
		}
	}

	if o.DestinationLocation != nil {
		v := *o.DestinationLocation

		if err = d.Set("destination_location", v); err != nil {
			return diag.Errorf("error reading destination_location: %v", err)
		}
	}

	if o.DestinationVisibility != nil {
		v := *o.DestinationVisibility

		if err = d.Set("destination_visibility", v); err != nil {
			return diag.Errorf("error reading destination_visibility: %v", err)
		}
	}

	if o.HostnameLimit != nil {
		v := *o.HostnameLimit

		if err = d.Set("hostname_limit", v); err != nil {
			return diag.Errorf("error reading hostname_limit: %v", err)
		}
	}

	if o.HostnameTtl != nil {
		v := *o.HostnameTtl

		if err = d.Set("hostname_ttl", v); err != nil {
			return diag.Errorf("error reading hostname_ttl: %v", err)
		}
	}

	if o.SourceLocation != nil {
		v := *o.SourceLocation

		if err = d.Set("source_location", v); err != nil {
			return diag.Errorf("error reading source_location: %v", err)
		}
	}

	return nil
}

func getObjectSystemNetworkVisibility(d *schema.ResourceData, sv string) (*models.SystemNetworkVisibility, diag.Diagnostics) {
	obj := models.SystemNetworkVisibility{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("destination_hostname_visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("destination_hostname_visibility", sv)
				diags = append(diags, e)
			}
			obj.DestinationHostnameVisibility = &v2
		}
	}
	if v1, ok := d.GetOk("destination_location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("destination_location", sv)
				diags = append(diags, e)
			}
			obj.DestinationLocation = &v2
		}
	}
	if v1, ok := d.GetOk("destination_visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("destination_visibility", sv)
				diags = append(diags, e)
			}
			obj.DestinationVisibility = &v2
		}
	}
	if v1, ok := d.GetOk("hostname_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HostnameLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("hostname_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HostnameTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("source_location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_location", sv)
				diags = append(diags, e)
			}
			obj.SourceLocation = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemNetworkVisibility(d *schema.ResourceData, sv string) (*models.SystemNetworkVisibility, diag.Diagnostics) {
	obj := models.SystemNetworkVisibility{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
