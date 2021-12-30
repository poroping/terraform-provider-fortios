// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Config global/VDOM Wildcard FQDN address.

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

func resourceFirewallwildcardFqdnCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Config global/VDOM Wildcard FQDN address.",

		CreateContext: resourceFirewallwildcardFqdnCustomCreate,
		ReadContext:   resourceFirewallwildcardFqdnCustomRead,
		UpdateContext: resourceFirewallwildcardFqdnCustomUpdate,
		DeleteContext: resourceFirewallwildcardFqdnCustomDelete,

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
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "GUI icon color.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Address name.",
				ForceNew:    true,
				Required:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable address visibility.",
				Optional:    true,
				Computed:    true,
			},
			"wildcard_fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Wildcard FQDN.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallwildcardFqdnCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallwildcardFqdnCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallwildcardFqdnCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallwildcardFqdnCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallwildcardFqdnCustom")
	}

	return resourceFirewallwildcardFqdnCustomRead(ctx, d, meta)
}

func resourceFirewallwildcardFqdnCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallwildcardFqdnCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallwildcardFqdnCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallwildcardFqdnCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallwildcardFqdnCustom")
	}

	return resourceFirewallwildcardFqdnCustomRead(ctx, d, meta)
}

func resourceFirewallwildcardFqdnCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallwildcardFqdnCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallwildcardFqdnCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallwildcardFqdnCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallwildcardFqdnCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallwildcardFqdnCustom resource: %v", err)
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

	diags := refreshObjectFirewallwildcardFqdnCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallwildcardFqdnCustom(d *schema.ResourceData, o *models.FirewallwildcardFqdnCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.Visibility != nil {
		v := *o.Visibility

		if err = d.Set("visibility", v); err != nil {
			return diag.Errorf("error reading visibility: %v", err)
		}
	}

	if o.WildcardFqdn != nil {
		v := *o.WildcardFqdn

		if err = d.Set("wildcard_fqdn", v); err != nil {
			return diag.Errorf("error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func getObjectFirewallwildcardFqdnCustom(d *schema.ResourceData, sv string) (*models.FirewallwildcardFqdnCustom, diag.Diagnostics) {
	obj := models.FirewallwildcardFqdnCustom{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
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
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	if v1, ok := d.GetOk("wildcard_fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wildcard_fqdn", sv)
				diags = append(diags, e)
			}
			obj.WildcardFqdn = &v2
		}
	}
	return &obj, diags
}
