// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure firewall IP-translation.

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

func resourceFirewallIpTranslation() *schema.Resource {
	return &schema.Resource{
		Description: "Configure firewall IP-translation.",

		CreateContext: resourceFirewallIpTranslationCreate,
		ReadContext:   resourceFirewallIpTranslationRead,
		UpdateContext: resourceFirewallIpTranslationUpdate,
		DeleteContext: resourceFirewallIpTranslationDelete,

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
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"endip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"map_startip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"startip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"transid": {
				Type: schema.TypeInt,

				Description: "IP translation ID.",
				ForceNew:    true,
				Required:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"SCTP"}, false),

				Description: "IP translation type (option: SCTP).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallIpTranslationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "transid"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallIpTranslation resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallIpTranslation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallIpTranslation(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallIpTranslation")
	}

	return resourceFirewallIpTranslationRead(ctx, d, meta)
}

func resourceFirewallIpTranslationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallIpTranslation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallIpTranslation(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallIpTranslation resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallIpTranslation")
	}

	return resourceFirewallIpTranslationRead(ctx, d, meta)
}

func resourceFirewallIpTranslationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallIpTranslation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallIpTranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpTranslationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallIpTranslation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIpTranslation resource: %v", err)
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

	diags := refreshObjectFirewallIpTranslation(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallIpTranslation(d *schema.ResourceData, o *models.FirewallIpTranslation, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Endip != nil {
		v := *o.Endip

		if err = d.Set("endip", v); err != nil {
			return diag.Errorf("error reading endip: %v", err)
		}
	}

	if o.MapStartip != nil {
		v := *o.MapStartip

		if err = d.Set("map_startip", v); err != nil {
			return diag.Errorf("error reading map_startip: %v", err)
		}
	}

	if o.Startip != nil {
		v := *o.Startip

		if err = d.Set("startip", v); err != nil {
			return diag.Errorf("error reading startip: %v", err)
		}
	}

	if o.Transid != nil {
		v := *o.Transid

		if err = d.Set("transid", v); err != nil {
			return diag.Errorf("error reading transid: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	return nil
}

func getObjectFirewallIpTranslation(d *schema.ResourceData, sv string) (*models.FirewallIpTranslation, diag.Diagnostics) {
	obj := models.FirewallIpTranslation{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("endip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("endip", sv)
				diags = append(diags, e)
			}
			obj.Endip = &v2
		}
	}
	if v1, ok := d.GetOk("map_startip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("map_startip", sv)
				diags = append(diags, e)
			}
			obj.MapStartip = &v2
		}
	}
	if v1, ok := d.GetOk("startip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("startip", sv)
				diags = append(diags, e)
			}
			obj.Startip = &v2
		}
	}
	if v1, ok := d.GetOk("transid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("transid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Transid = &tmp
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	return &obj, diags
}
