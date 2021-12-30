// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP address type availability.

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

func resourceWirelessControllerhotspot20AnqpIpAddressType() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP address type availability.",

		CreateContext: resourceWirelessControllerhotspot20AnqpIpAddressTypeCreate,
		ReadContext:   resourceWirelessControllerhotspot20AnqpIpAddressTypeRead,
		UpdateContext: resourceWirelessControllerhotspot20AnqpIpAddressTypeUpdate,
		DeleteContext: resourceWirelessControllerhotspot20AnqpIpAddressTypeDelete,

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
			"ipv4_address_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"not-available", "public", "port-restricted", "single-NATed-private", "double-NATed-private", "port-restricted-and-single-NATed", "port-restricted-and-double-NATed", "not-known"}, false),

				Description: "IPv4 address type.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_address_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"not-available", "available", "not-known"}, false),

				Description: "IPv6 address type.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IP type name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerhotspot20AnqpIpAddressTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerhotspot20AnqpIpAddressType resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerhotspot20AnqpIpAddressType(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerhotspot20AnqpIpAddressType(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpIpAddressType")
	}

	return resourceWirelessControllerhotspot20AnqpIpAddressTypeRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpIpAddressTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerhotspot20AnqpIpAddressType(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerhotspot20AnqpIpAddressType(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerhotspot20AnqpIpAddressType resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerhotspot20AnqpIpAddressType")
	}

	return resourceWirelessControllerhotspot20AnqpIpAddressTypeRead(ctx, d, meta)
}

func resourceWirelessControllerhotspot20AnqpIpAddressTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerhotspot20AnqpIpAddressType(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerhotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerhotspot20AnqpIpAddressTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerhotspot20AnqpIpAddressType(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerhotspot20AnqpIpAddressType resource: %v", err)
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

	diags := refreshObjectWirelessControllerhotspot20AnqpIpAddressType(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerhotspot20AnqpIpAddressType(d *schema.ResourceData, o *models.WirelessControllerhotspot20AnqpIpAddressType, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Ipv4AddressType != nil {
		v := *o.Ipv4AddressType

		if err = d.Set("ipv4_address_type", v); err != nil {
			return diag.Errorf("error reading ipv4_address_type: %v", err)
		}
	}

	if o.Ipv6AddressType != nil {
		v := *o.Ipv6AddressType

		if err = d.Set("ipv6_address_type", v); err != nil {
			return diag.Errorf("error reading ipv6_address_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerhotspot20AnqpIpAddressType(d *schema.ResourceData, sv string) (*models.WirelessControllerhotspot20AnqpIpAddressType, diag.Diagnostics) {
	obj := models.WirelessControllerhotspot20AnqpIpAddressType{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ipv4_address_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_address_type", sv)
				diags = append(diags, e)
			}
			obj.Ipv4AddressType = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_address_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_address_type", sv)
				diags = append(diags, e)
			}
			obj.Ipv6AddressType = &v2
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
	return &obj, diags
}
