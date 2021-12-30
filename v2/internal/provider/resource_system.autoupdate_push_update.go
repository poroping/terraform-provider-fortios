// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure push updates.

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

func resourceSystemautoupdatePushUpdate() *schema.Resource {
	return &schema.Resource{
		Description: "Configure push updates.",

		CreateContext: resourceSystemautoupdatePushUpdateCreate,
		ReadContext:   resourceSystemautoupdatePushUpdateRead,
		UpdateContext: resourceSystemautoupdatePushUpdateUpdate,
		DeleteContext: resourceSystemautoupdatePushUpdateDelete,

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
			"address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "IPv4 or IPv6 address used by FortiGuard servers to send push updates to this FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable push update override server.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Push update override port. (Do not overlap with other service ports)",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable push updates.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemautoupdatePushUpdateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemautoupdatePushUpdate(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemautoupdatePushUpdate(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemautoupdatePushUpdate")
	}

	return resourceSystemautoupdatePushUpdateRead(ctx, d, meta)
}

func resourceSystemautoupdatePushUpdateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemautoupdatePushUpdate(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemautoupdatePushUpdate(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemautoupdatePushUpdate resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemautoupdatePushUpdate")
	}

	return resourceSystemautoupdatePushUpdateRead(ctx, d, meta)
}

func resourceSystemautoupdatePushUpdateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemautoupdatePushUpdate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemautoupdatePushUpdate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemautoupdatePushUpdateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemautoupdatePushUpdate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemautoupdatePushUpdate resource: %v", err)
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

	diags := refreshObjectSystemautoupdatePushUpdate(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemautoupdatePushUpdate(d *schema.ResourceData, o *models.SystemautoupdatePushUpdate, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Address != nil {
		v := *o.Address

		if err = d.Set("address", v); err != nil {
			return diag.Errorf("error reading address: %v", err)
		}
	}

	if o.Override != nil {
		v := *o.Override

		if err = d.Set("override", v); err != nil {
			return diag.Errorf("error reading override: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemautoupdatePushUpdate(d *schema.ResourceData, sv string) (*models.SystemautoupdatePushUpdate, diag.Diagnostics) {
	obj := models.SystemautoupdatePushUpdate{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("address", sv)
				diags = append(diags, e)
			}
			obj.Address = &v2
		}
	}
	if v1, ok := d.GetOk("override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override", sv)
				diags = append(diags, e)
			}
			obj.Override = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}
