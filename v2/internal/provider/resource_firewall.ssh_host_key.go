// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSH proxy host public keys.

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

func resourceFirewallsshHostKey() *schema.Resource {
	return &schema.Resource{
		Description: "SSH proxy host public keys.",

		CreateContext: resourceFirewallsshHostKeyCreate,
		ReadContext:   resourceFirewallsshHostKeyRead,
		UpdateContext: resourceFirewallsshHostKeyUpdate,
		DeleteContext: resourceFirewallsshHostKeyDelete,

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
			"hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Hostname of the SSH server, to match SSH certificate principals. ",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the SSH server.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSH public key name.",
				ForceNew:    true,
				Required:    true,
			},
			"nid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"256", "384", "521"}, false),

				Description: "Set the nid of the ECDSA key.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type: schema.TypeInt,

				Description: "Port of the SSH server.",
				Optional:    true,
				Computed:    true,
			},
			"public_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),

				Description: "SSH public key.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"trusted", "revoked"}, false),

				Description: "Set the trust status of the public key.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"RSA", "DSA", "ECDSA", "ED25519", "RSA-CA", "DSA-CA", "ECDSA-CA", "ED25519-CA"}, false),

				Description: "Set the type of the public key.",
				Optional:    true,
				Computed:    true,
			},
			"usage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"transparent-proxy", "access-proxy"}, false),

				Description: "Usage for this public key.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallsshHostKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallsshHostKey resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallsshHostKey(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallsshHostKey(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallsshHostKey")
	}

	return resourceFirewallsshHostKeyRead(ctx, d, meta)
}

func resourceFirewallsshHostKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallsshHostKey(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallsshHostKey(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallsshHostKey resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallsshHostKey")
	}

	return resourceFirewallsshHostKeyRead(ctx, d, meta)
}

func resourceFirewallsshHostKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallsshHostKey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallsshHostKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallsshHostKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallsshHostKey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallsshHostKey resource: %v", err)
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

	diags := refreshObjectFirewallsshHostKey(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallsshHostKey(d *schema.ResourceData, o *models.FirewallsshHostKey, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Hostname != nil {
		v := *o.Hostname

		if err = d.Set("hostname", v); err != nil {
			return diag.Errorf("error reading hostname: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nid != nil {
		v := *o.Nid

		if err = d.Set("nid", v); err != nil {
			return diag.Errorf("error reading nid: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.PublicKey != nil {
		v := *o.PublicKey

		if err = d.Set("public_key", v); err != nil {
			return diag.Errorf("error reading public_key: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Usage != nil {
		v := *o.Usage

		if err = d.Set("usage", v); err != nil {
			return diag.Errorf("error reading usage: %v", err)
		}
	}

	return nil
}

func getObjectFirewallsshHostKey(d *schema.ResourceData, sv string) (*models.FirewallsshHostKey, diag.Diagnostics) {
	obj := models.FirewallsshHostKey{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname", sv)
				diags = append(diags, e)
			}
			obj.Hostname = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
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
	if v1, ok := d.GetOk("nid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nid", sv)
				diags = append(diags, e)
			}
			obj.Nid = &v2
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
	if v1, ok := d.GetOk("public_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("public_key", sv)
				diags = append(diags, e)
			}
			obj.PublicKey = &v2
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
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("usage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("usage", sv)
				diags = append(diags, e)
			}
			obj.Usage = &v2
		}
	}
	return &obj, diags
}
