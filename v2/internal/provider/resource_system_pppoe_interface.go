// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure the PPPoE interfaces.

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

func resourceSystemPppoeInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure the PPPoE interfaces.",

		CreateContext: resourceSystemPppoeInterfaceCreate,
		ReadContext:   resourceSystemPppoeInterfaceRead,
		UpdateContext: resourceSystemPppoeInterfaceUpdate,
		DeleteContext: resourceSystemPppoeInterfaceDelete,

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
			"ac_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "PPPoE AC name.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "pap", "chap", "mschapv1", "mschapv2"}, false),

				Description: "PPP authentication type to use.",
				Optional:    true,
				Computed:    true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name for the physical interface.",
				Optional:    true,
				Computed:    true,
			},
			"dial_on_demand": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.",
				Optional:    true,
				Computed:    true,
			},
			"disc_retry_timeout": {
				Type: schema.TypeInt,

				Description: "PPPoE discovery init timeout value in (0-4294967295 sec).",
				Optional:    true,
				Computed:    true,
			},
			"idle_timeout": {
				Type: schema.TypeInt,

				Description: "PPPoE auto disconnect after idle timeout (0-4294967295 sec).",
				Optional:    true,
				Computed:    true,
			},
			"ipunnumbered": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "PPPoE unnumbered IP.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 Control Protocol (IPv6CP).",
				Optional:    true,
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Optional:    true,
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Maximum missed LCP echo messages before disconnect.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name of the PPPoE interface.",
				ForceNew:    true,
				Required:    true,
			},
			"padt_retry_timeout": {
				Type: schema.TypeInt,

				Description: "PPPoE terminate timeout value in (0-4294967295 sec).",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Enter the password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"pppoe_unnumbered_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPPoE unnumbered negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"service_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "PPPoE service name.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "User name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemPppoeInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemPppoeInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemPppoeInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemPppoeInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPppoeInterface")
	}

	return resourceSystemPppoeInterfaceRead(ctx, d, meta)
}

func resourceSystemPppoeInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPppoeInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemPppoeInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemPppoeInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPppoeInterface")
	}

	return resourceSystemPppoeInterfaceRead(ctx, d, meta)
}

func resourceSystemPppoeInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemPppoeInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemPppoeInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPppoeInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemPppoeInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPppoeInterface resource: %v", err)
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

	diags := refreshObjectSystemPppoeInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemPppoeInterface(d *schema.ResourceData, o *models.SystemPppoeInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcName != nil {
		v := *o.AcName

		if err = d.Set("ac_name", v); err != nil {
			return diag.Errorf("error reading ac_name: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.DialOnDemand != nil {
		v := *o.DialOnDemand

		if err = d.Set("dial_on_demand", v); err != nil {
			return diag.Errorf("error reading dial_on_demand: %v", err)
		}
	}

	if o.DiscRetryTimeout != nil {
		v := *o.DiscRetryTimeout

		if err = d.Set("disc_retry_timeout", v); err != nil {
			return diag.Errorf("error reading disc_retry_timeout: %v", err)
		}
	}

	if o.IdleTimeout != nil {
		v := *o.IdleTimeout

		if err = d.Set("idle_timeout", v); err != nil {
			return diag.Errorf("error reading idle_timeout: %v", err)
		}
	}

	if o.Ipunnumbered != nil {
		v := *o.Ipunnumbered

		if err = d.Set("ipunnumbered", v); err != nil {
			return diag.Errorf("error reading ipunnumbered: %v", err)
		}
	}

	if o.Ipv6 != nil {
		v := *o.Ipv6

		if err = d.Set("ipv6", v); err != nil {
			return diag.Errorf("error reading ipv6: %v", err)
		}
	}

	if o.LcpEchoInterval != nil {
		v := *o.LcpEchoInterval

		if err = d.Set("lcp_echo_interval", v); err != nil {
			return diag.Errorf("error reading lcp_echo_interval: %v", err)
		}
	}

	if o.LcpMaxEchoFails != nil {
		v := *o.LcpMaxEchoFails

		if err = d.Set("lcp_max_echo_fails", v); err != nil {
			return diag.Errorf("error reading lcp_max_echo_fails: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PadtRetryTimeout != nil {
		v := *o.PadtRetryTimeout

		if err = d.Set("padt_retry_timeout", v); err != nil {
			return diag.Errorf("error reading padt_retry_timeout: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PppoeUnnumberedNegotiate != nil {
		v := *o.PppoeUnnumberedNegotiate

		if err = d.Set("pppoe_unnumbered_negotiate", v); err != nil {
			return diag.Errorf("error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if o.ServiceName != nil {
		v := *o.ServiceName

		if err = d.Set("service_name", v); err != nil {
			return diag.Errorf("error reading service_name: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	return nil
}

func getObjectSystemPppoeInterface(d *schema.ResourceData, sv string) (*models.SystemPppoeInterface, diag.Diagnostics) {
	obj := models.SystemPppoeInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ac_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_name", sv)
				diags = append(diags, e)
			}
			obj.AcName = &v2
		}
	}
	if v1, ok := d.GetOk("auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_type", sv)
				diags = append(diags, e)
			}
			obj.AuthType = &v2
		}
	}
	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("dial_on_demand"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dial_on_demand", sv)
				diags = append(diags, e)
			}
			obj.DialOnDemand = &v2
		}
	}
	if v1, ok := d.GetOk("disc_retry_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("disc_retry_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DiscRetryTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("idle_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ipunnumbered"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipunnumbered", sv)
				diags = append(diags, e)
			}
			obj.Ipunnumbered = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6", sv)
				diags = append(diags, e)
			}
			obj.Ipv6 = &v2
		}
	}
	if v1, ok := d.GetOk("lcp_echo_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lcp_echo_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpEchoInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("lcp_max_echo_fails"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lcp_max_echo_fails", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpMaxEchoFails = &tmp
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
	if v1, ok := d.GetOk("padt_retry_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padt_retry_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PadtRetryTimeout = &tmp
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
	if v1, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pppoe_unnumbered_negotiate", sv)
				diags = append(diags, e)
			}
			obj.PppoeUnnumberedNegotiate = &v2
		}
	}
	if v1, ok := d.GetOk("service_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_name", sv)
				diags = append(diags, e)
			}
			obj.ServiceName = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	return &obj, diags
}
