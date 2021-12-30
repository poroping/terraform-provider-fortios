// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure terms and conditions.

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

func resourceWirelessControllerhotspot20H2qpTermsAndConditions() *schema.Resource {
	return &schema.Resource{
		Description: "Configure terms and conditions.",

		CreateContext: resourceWirelessControllerhotspot20H2qpTermsAndConditionsCreate,
		ReadContext:   resourceWirelessControllerhotspot20H2qpTermsAndConditionsRead,
		UpdateContext: resourceWirelessControllerhotspot20H2qpTermsAndConditionsUpdate,
		DeleteContext: resourceWirelessControllerhotspot20H2qpTermsAndConditionsDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"filename": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 254),

				Description: "Filename.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Terms and Conditions ID.",
				ForceNew:    true,
				Required:    true,
			},
			"timestamp": {
				Type: schema.TypeInt,

				Description: "Timestamp.",
				Optional:    true,
				Computed:    true,
			},
			"url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 253),

				Description: "URL.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerhotspot20H2qpTermsAndConditionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WirelessControllerhotspot20H2qpTermsAndConditions resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerhotspot20H2qpTermsAndConditions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerhotspot20H2qpTermsAndConditions(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20H2qpTermsAndConditions")
	}

	return resourceWirelessControllerhotspot20H2qpTermsAndConditionsRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20H2qpTermsAndConditionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerhotspot20H2qpTermsAndConditions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerhotspot20H2qpTermsAndConditions(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerhotspot20H2qpTermsAndConditions resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20H2qpTermsAndConditions")
	}

	return resourceWirelessControllerhotspot20H2qpTermsAndConditionsRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20H2qpTermsAndConditionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerhotspot20H2qpTermsAndConditions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerhotspot20H2qpTermsAndConditions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerhotspot20H2qpTermsAndConditionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerhotspot20H2qpTermsAndConditions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerhotspot20H2qpTermsAndConditions resource: %v", err)
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

	diags := refreshObjectWirelessControllerhotspot20H2qpTermsAndConditions(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerhotspot20H2qpTermsAndConditions(d *schema.ResourceData, o *models.WirelessControllerhotspot20H2qpTermsAndConditions, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Filename != nil {
		v := *o.Filename

		if err = d.Set("filename", v); err != nil {
			return diag.Errorf("error reading filename: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Timestamp != nil {
		v := *o.Timestamp

		if err = d.Set("timestamp", v); err != nil {
			return diag.Errorf("error reading timestamp: %v", err)
		}
	}

	if o.Url != nil {
		v := *o.Url

		if err = d.Set("url", v); err != nil {
			return diag.Errorf("error reading url: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerhotspot20H2qpTermsAndConditions(d *schema.ResourceData, sv string) (*models.WirelessControllerhotspot20H2qpTermsAndConditions, diag.Diagnostics) {
	obj := models.WirelessControllerhotspot20H2qpTermsAndConditions{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("filename"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filename", sv)
				diags = append(diags, e)
			}
			obj.Filename = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("timestamp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timestamp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timestamp = &tmp
		}
	}
	if v1, ok := d.GetOk("url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url", sv)
				diags = append(diags, e)
			}
			obj.Url = &v2
		}
	}
	return &obj, diags
}
