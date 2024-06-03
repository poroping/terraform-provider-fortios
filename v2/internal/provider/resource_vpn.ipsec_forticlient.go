// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiClient policy realm.

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

func resourceVpnIpsecForticlient() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiClient policy realm.",

		CreateContext: resourceVpnIpsecForticlientCreate,
		ReadContext:   resourceVpnIpsecForticlientRead,
		UpdateContext: resourceVpnIpsecForticlientUpdate,
		DeleteContext: resourceVpnIpsecForticlientDelete,

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
			"phase2name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Phase 2 tunnel name that you defined in the FortiClient dialup configuration.",
				Optional:    true,
				Computed:    true,
			},
			"realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiClient realm name.",
				ForceNew:    true,
				Required:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this FortiClient configuration.",
				Optional:    true,
				Computed:    true,
			},
			"usergroupname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User group name for FortiClient users.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnIpsecForticlientCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "realm"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating VpnIpsecForticlient resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnIpsecForticlient(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnIpsecForticlient(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecForticlient")
	}

	return resourceVpnIpsecForticlientRead(ctx, d, meta)
}

func resourceVpnIpsecForticlientUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnIpsecForticlient(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnIpsecForticlient(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnIpsecForticlient resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecForticlient")
	}

	return resourceVpnIpsecForticlientRead(ctx, d, meta)
}

func resourceVpnIpsecForticlientDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnIpsecForticlient(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnIpsecForticlient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecForticlientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecForticlient(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecForticlient resource: %v", err)
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

	diags := refreshObjectVpnIpsecForticlient(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnIpsecForticlient(d *schema.ResourceData, o *models.VpnIpsecForticlient, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Phase2name != nil {
		v := *o.Phase2name

		if err = d.Set("phase2name", v); err != nil {
			return diag.Errorf("error reading phase2name: %v", err)
		}
	}

	if o.Realm != nil {
		v := *o.Realm

		if err = d.Set("realm", v); err != nil {
			return diag.Errorf("error reading realm: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Usergroupname != nil {
		v := *o.Usergroupname

		if err = d.Set("usergroupname", v); err != nil {
			return diag.Errorf("error reading usergroupname: %v", err)
		}
	}

	return nil
}

func getObjectVpnIpsecForticlient(d *schema.ResourceData, sv string) (*models.VpnIpsecForticlient, diag.Diagnostics) {
	obj := models.VpnIpsecForticlient{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("phase2name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phase2name", sv)
				diags = append(diags, e)
			}
			obj.Phase2name = &v2
		}
	}
	if v1, ok := d.GetOk("realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("realm", sv)
				diags = append(diags, e)
			}
			obj.Realm = &v2
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
	if v1, ok := d.GetOk("usergroupname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("usergroupname", sv)
				diags = append(diags, e)
			}
			obj.Usergroupname = &v2
		}
	}
	return &obj, diags
}
