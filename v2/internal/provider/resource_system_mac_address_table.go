// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MAC address tables.

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

func resourceSystemMacAddressTable() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MAC address tables.",

		CreateContext: resourceSystemMacAddressTableCreate,
		ReadContext:   resourceSystemMacAddressTableRead,
		UpdateContext: resourceSystemMacAddressTableUpdate,
		DeleteContext: resourceSystemMacAddressTableDelete,

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
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface name.",
				Optional:    true,
				Computed:    true,
			},
			"mac": {
				Type: schema.TypeString,

				Description: "MAC address.",
				ForceNew:    true,
				Required:    true,
			},
			"reply_substitute": {
				Type: schema.TypeString,

				Description: "New MAC for reply traffic.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemMacAddressTableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "mac"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemMacAddressTable resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemMacAddressTable(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemMacAddressTable(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(ctx, d, meta)
}

func resourceSystemMacAddressTableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemMacAddressTable(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemMacAddressTable(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemMacAddressTable resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(ctx, d, meta)
}

func resourceSystemMacAddressTableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemMacAddressTable(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemMacAddressTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMacAddressTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemMacAddressTable(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemMacAddressTable resource: %v", err)
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

	diags := refreshObjectSystemMacAddressTable(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemMacAddressTable(d *schema.ResourceData, o *models.SystemMacAddressTable, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Mac != nil {
		v := *o.Mac

		if err = d.Set("mac", v); err != nil {
			return diag.Errorf("error reading mac: %v", err)
		}
	}

	if o.ReplySubstitute != nil {
		v := *o.ReplySubstitute

		if err = d.Set("reply_substitute", v); err != nil {
			return diag.Errorf("error reading reply_substitute: %v", err)
		}
	}

	return nil
}

func getObjectSystemMacAddressTable(d *schema.ResourceData, sv string) (*models.SystemMacAddressTable, diag.Diagnostics) {
	obj := models.SystemMacAddressTable{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac", sv)
				diags = append(diags, e)
			}
			obj.Mac = &v2
		}
	}
	if v1, ok := d.GetOk("reply_substitute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reply_substitute", sv)
				diags = append(diags, e)
			}
			obj.ReplySubstitute = &v2
		}
	}
	return &obj, diags
}
