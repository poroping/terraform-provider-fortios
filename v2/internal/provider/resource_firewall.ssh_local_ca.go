// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSH proxy local CA.

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

func resourceFirewallSshLocalCa() *schema.Resource {
	return &schema.Resource{
		Description: "SSH proxy local CA.",

		CreateContext: resourceFirewallSshLocalCaCreate,
		ReadContext:   resourceFirewallSshLocalCaRead,
		UpdateContext: resourceFirewallSshLocalCaUpdate,
		DeleteContext: resourceFirewallSshLocalCaDelete,

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
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSH proxy local CA name.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password for SSH private key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"private_key": {
				Type: schema.TypeString,

				Description: "SSH proxy private key, encrypted with a password.",
				Optional:    true,
				Computed:    true,
			},
			"public_key": {
				Type: schema.TypeString,

				Description: "SSH proxy public key.",
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"built-in", "user"}, false),

				Description: "SSH proxy local CA source type.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSshLocalCaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallSshLocalCa resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallSshLocalCa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSshLocalCa(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSshLocalCa")
	}

	return resourceFirewallSshLocalCaRead(ctx, d, meta)
}

func resourceFirewallSshLocalCaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSshLocalCa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSshLocalCa(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSshLocalCa resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSshLocalCa")
	}

	return resourceFirewallSshLocalCaRead(ctx, d, meta)
}

func resourceFirewallSshLocalCaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallSshLocalCa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSshLocalCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshLocalCaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSshLocalCa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSshLocalCa resource: %v", err)
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

	diags := refreshObjectFirewallSshLocalCa(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallSshLocalCa(d *schema.ResourceData, o *models.FirewallSshLocalCa, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PrivateKey != nil {
		v := *o.PrivateKey

		if err = d.Set("private_key", v); err != nil {
			return diag.Errorf("error reading private_key: %v", err)
		}
	}

	if o.PublicKey != nil {
		v := *o.PublicKey

		if err = d.Set("public_key", v); err != nil {
			return diag.Errorf("error reading public_key: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
		}
	}

	return nil
}

func getObjectFirewallSshLocalCa(d *schema.ResourceData, sv string) (*models.FirewallSshLocalCa, diag.Diagnostics) {
	obj := models.FirewallSshLocalCa{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("private_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("private_key", sv)
				diags = append(diags, e)
			}
			obj.PrivateKey = &v2
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
	if v1, ok := d.GetOk("source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source", sv)
				diags = append(diags, e)
			}
			obj.Source = &v2
		}
	}
	return &obj, diags
}
