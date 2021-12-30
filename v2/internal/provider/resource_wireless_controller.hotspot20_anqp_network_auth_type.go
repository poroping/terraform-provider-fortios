// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure network authentication type.

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

func resourceWirelessControllerhotspot20AnqpNetworkAuthType() *schema.Resource {
	return &schema.Resource{
		Description: "Configure network authentication type.",

		CreateContext: resourceWirelessControllerhotspot20AnqpNetworkAuthTypeCreate,
		ReadContext:   resourceWirelessControllerhotspot20AnqpNetworkAuthTypeRead,
		UpdateContext: resourceWirelessControllerhotspot20AnqpNetworkAuthTypeUpdate,
		DeleteContext: resourceWirelessControllerhotspot20AnqpNetworkAuthTypeDelete,

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
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"acceptance-of-terms", "online-enrollment", "http-redirection", "dns-redirection"}, false),

				Description: "Network authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication type name.",
				ForceNew:    true,
				Required:    true,
			},
			"url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Redirect URL.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerhotspot20AnqpNetworkAuthTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerhotspot20AnqpNetworkAuthType resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerhotspot20AnqpNetworkAuthType(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerhotspot20AnqpNetworkAuthType(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpNetworkAuthType")
	}

	return resourceWirelessControllerhotspot20AnqpNetworkAuthTypeRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpNetworkAuthTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerhotspot20AnqpNetworkAuthType(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerhotspot20AnqpNetworkAuthType(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerhotspot20AnqpNetworkAuthType resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpNetworkAuthType")
	}

	return resourceWirelessControllerhotspot20AnqpNetworkAuthTypeRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpNetworkAuthTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerhotspot20AnqpNetworkAuthType(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerhotspot20AnqpNetworkAuthType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerhotspot20AnqpNetworkAuthTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerhotspot20AnqpNetworkAuthType(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerhotspot20AnqpNetworkAuthType resource: %v", err)
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

	diags := refreshObjectWirelessControllerhotspot20AnqpNetworkAuthType(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerhotspot20AnqpNetworkAuthType(d *schema.ResourceData, o *models.WirelessControllerhotspot20AnqpNetworkAuthType, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
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

func getObjectWirelessControllerhotspot20AnqpNetworkAuthType(d *schema.ResourceData, sv string) (*models.WirelessControllerhotspot20AnqpNetworkAuthType, diag.Diagnostics) {
	obj := models.WirelessControllerhotspot20AnqpNetworkAuthType{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_type", sv)
				diags = append(diags, e)
			}
			obj.AuthType = &v2
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
