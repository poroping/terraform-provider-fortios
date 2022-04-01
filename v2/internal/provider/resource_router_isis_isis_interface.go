// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: IS-IS interface configuration.

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

func resourceRouterIsisIsisInterface() *schema.Resource {
	return &schema.Resource{
		Description: "IS-IS interface configuration.",

		CreateContext: resourceRouterIsisIsisInterfaceCreate,
		ReadContext:   resourceRouterIsisIsisInterfaceRead,
		UpdateContext: resourceRouterIsisIsisInterfaceUpdate,
		DeleteContext: resourceRouterIsisIsisInterfaceDelete,

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
			"auth_keychain_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication key-chain for level 1 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"auth_keychain_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication key-chain for level 2 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"auth_mode_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "password"}, false),

				Description: "Level 1 authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"auth_mode_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "password"}, false),

				Description: "Level 2 authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"auth_password_l1": {
				Type: schema.TypeString,

				Description: "Authentication password for level 1 PDUs.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"auth_password_l2": {
				Type: schema.TypeString,

				Description: "Authentication password for level 2 PDUs.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"auth_send_only_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication send-only for level 1 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"auth_send_only_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication send-only for level 2 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"circuit_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

				Description: "IS-IS interface's circuit type.",
				Optional:    true,
				Computed:    true,
			},
			"csnp_interval_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Level 1 CSNP interval.",
				Optional:    true,
				Computed:    true,
			},
			"csnp_interval_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Level 2 CSNP interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Level 1 hello interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Level 2 hello interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_multiplier_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "Level 1 multiplier for Hello holding time.",
				Optional:    true,
				Computed:    true,
			},
			"hello_multiplier_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "Level 2 multiplier for Hello holding time.",
				Optional:    true,
				Computed:    true,
			},
			"hello_padding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable padding to IS-IS hello packets.",
				Optional:    true,
				Computed:    true,
			},
			"lsp_interval": {
				Type: schema.TypeInt,

				Description: "LSP transmission interval (milliseconds).",
				Optional:    true,
				Computed:    true,
			},
			"lsp_retransmit_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "LSP retransmission interval (sec).",
				Optional:    true,
				Computed:    true,
			},
			"mesh_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IS-IS mesh group.",
				Optional:    true,
				Computed:    true,
			},
			"mesh_group_id": {
				Type: schema.TypeInt,

				Description: "Mesh group ID <0-4294967295>, 0: mesh-group blocked.",
				Optional:    true,
				Computed:    true,
			},
			"metric_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 63),

				Description: "Level 1 metric for interface.",
				Optional:    true,
				Computed:    true,
			},
			"metric_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 63),

				Description: "Level 2 metric for interface.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "IS-IS interface name.",
				ForceNew:    true,
				Required:    true,
			},
			"network_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"broadcast", "point-to-point", "loopback"}, false),

				Description: "IS-IS interface's network type.",
				Optional:    true,
				Computed:    true,
			},
			"priority_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 127),

				Description: "Level 1 priority.",
				Optional:    true,
				Computed:    true,
			},
			"priority_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 127),

				Description: "Level 2 priority.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable interface for IS-IS.",
				Optional:    true,
				Computed:    true,
			},
			"status6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 interface for IS-IS.",
				Optional:    true,
				Computed:    true,
			},
			"wide_metric_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Level 1 wide metric for interface.",
				Optional:    true,
				Computed:    true,
			},
			"wide_metric_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Level 2 wide metric for interface.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterIsisIsisInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterIsisIsisInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterIsisIsisInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterIsisIsisInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterIsisIsisInterface")
	}

	return resourceRouterIsisIsisInterfaceRead(ctx, d, meta)
}

func resourceRouterIsisIsisInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterIsisIsisInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterIsisIsisInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterIsisIsisInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterIsisIsisInterface")
	}

	return resourceRouterIsisIsisInterfaceRead(ctx, d, meta)
}

func resourceRouterIsisIsisInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterIsisIsisInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterIsisIsisInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisIsisInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterIsisIsisInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterIsisIsisInterface resource: %v", err)
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

	diags := refreshObjectRouterIsisIsisInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterIsisIsisInterface(d *schema.ResourceData, o *models.RouterIsisIsisInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthKeychainL1 != nil {
		v := *o.AuthKeychainL1

		if err = d.Set("auth_keychain_l1", v); err != nil {
			return diag.Errorf("error reading auth_keychain_l1: %v", err)
		}
	}

	if o.AuthKeychainL2 != nil {
		v := *o.AuthKeychainL2

		if err = d.Set("auth_keychain_l2", v); err != nil {
			return diag.Errorf("error reading auth_keychain_l2: %v", err)
		}
	}

	if o.AuthModeL1 != nil {
		v := *o.AuthModeL1

		if err = d.Set("auth_mode_l1", v); err != nil {
			return diag.Errorf("error reading auth_mode_l1: %v", err)
		}
	}

	if o.AuthModeL2 != nil {
		v := *o.AuthModeL2

		if err = d.Set("auth_mode_l2", v); err != nil {
			return diag.Errorf("error reading auth_mode_l2: %v", err)
		}
	}

	if o.AuthPasswordL1 != nil {
		v := *o.AuthPasswordL1

		if v == "ENC XXXX" {
		} else if err = d.Set("auth_password_l1", v); err != nil {
			return diag.Errorf("error reading auth_password_l1: %v", err)
		}
	}

	if o.AuthPasswordL2 != nil {
		v := *o.AuthPasswordL2

		if v == "ENC XXXX" {
		} else if err = d.Set("auth_password_l2", v); err != nil {
			return diag.Errorf("error reading auth_password_l2: %v", err)
		}
	}

	if o.AuthSendOnlyL1 != nil {
		v := *o.AuthSendOnlyL1

		if err = d.Set("auth_send_only_l1", v); err != nil {
			return diag.Errorf("error reading auth_send_only_l1: %v", err)
		}
	}

	if o.AuthSendOnlyL2 != nil {
		v := *o.AuthSendOnlyL2

		if err = d.Set("auth_send_only_l2", v); err != nil {
			return diag.Errorf("error reading auth_send_only_l2: %v", err)
		}
	}

	if o.CircuitType != nil {
		v := *o.CircuitType

		if err = d.Set("circuit_type", v); err != nil {
			return diag.Errorf("error reading circuit_type: %v", err)
		}
	}

	if o.CsnpIntervalL1 != nil {
		v := *o.CsnpIntervalL1

		if err = d.Set("csnp_interval_l1", v); err != nil {
			return diag.Errorf("error reading csnp_interval_l1: %v", err)
		}
	}

	if o.CsnpIntervalL2 != nil {
		v := *o.CsnpIntervalL2

		if err = d.Set("csnp_interval_l2", v); err != nil {
			return diag.Errorf("error reading csnp_interval_l2: %v", err)
		}
	}

	if o.HelloIntervalL1 != nil {
		v := *o.HelloIntervalL1

		if err = d.Set("hello_interval_l1", v); err != nil {
			return diag.Errorf("error reading hello_interval_l1: %v", err)
		}
	}

	if o.HelloIntervalL2 != nil {
		v := *o.HelloIntervalL2

		if err = d.Set("hello_interval_l2", v); err != nil {
			return diag.Errorf("error reading hello_interval_l2: %v", err)
		}
	}

	if o.HelloMultiplierL1 != nil {
		v := *o.HelloMultiplierL1

		if err = d.Set("hello_multiplier_l1", v); err != nil {
			return diag.Errorf("error reading hello_multiplier_l1: %v", err)
		}
	}

	if o.HelloMultiplierL2 != nil {
		v := *o.HelloMultiplierL2

		if err = d.Set("hello_multiplier_l2", v); err != nil {
			return diag.Errorf("error reading hello_multiplier_l2: %v", err)
		}
	}

	if o.HelloPadding != nil {
		v := *o.HelloPadding

		if err = d.Set("hello_padding", v); err != nil {
			return diag.Errorf("error reading hello_padding: %v", err)
		}
	}

	if o.LspInterval != nil {
		v := *o.LspInterval

		if err = d.Set("lsp_interval", v); err != nil {
			return diag.Errorf("error reading lsp_interval: %v", err)
		}
	}

	if o.LspRetransmitInterval != nil {
		v := *o.LspRetransmitInterval

		if err = d.Set("lsp_retransmit_interval", v); err != nil {
			return diag.Errorf("error reading lsp_retransmit_interval: %v", err)
		}
	}

	if o.MeshGroup != nil {
		v := *o.MeshGroup

		if err = d.Set("mesh_group", v); err != nil {
			return diag.Errorf("error reading mesh_group: %v", err)
		}
	}

	if o.MeshGroupId != nil {
		v := *o.MeshGroupId

		if err = d.Set("mesh_group_id", v); err != nil {
			return diag.Errorf("error reading mesh_group_id: %v", err)
		}
	}

	if o.MetricL1 != nil {
		v := *o.MetricL1

		if err = d.Set("metric_l1", v); err != nil {
			return diag.Errorf("error reading metric_l1: %v", err)
		}
	}

	if o.MetricL2 != nil {
		v := *o.MetricL2

		if err = d.Set("metric_l2", v); err != nil {
			return diag.Errorf("error reading metric_l2: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NetworkType != nil {
		v := *o.NetworkType

		if err = d.Set("network_type", v); err != nil {
			return diag.Errorf("error reading network_type: %v", err)
		}
	}

	if o.PriorityL1 != nil {
		v := *o.PriorityL1

		if err = d.Set("priority_l1", v); err != nil {
			return diag.Errorf("error reading priority_l1: %v", err)
		}
	}

	if o.PriorityL2 != nil {
		v := *o.PriorityL2

		if err = d.Set("priority_l2", v); err != nil {
			return diag.Errorf("error reading priority_l2: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Status6 != nil {
		v := *o.Status6

		if err = d.Set("status6", v); err != nil {
			return diag.Errorf("error reading status6: %v", err)
		}
	}

	if o.WideMetricL1 != nil {
		v := *o.WideMetricL1

		if err = d.Set("wide_metric_l1", v); err != nil {
			return diag.Errorf("error reading wide_metric_l1: %v", err)
		}
	}

	if o.WideMetricL2 != nil {
		v := *o.WideMetricL2

		if err = d.Set("wide_metric_l2", v); err != nil {
			return diag.Errorf("error reading wide_metric_l2: %v", err)
		}
	}

	return nil
}

func getObjectRouterIsisIsisInterface(d *schema.ResourceData, sv string) (*models.RouterIsisIsisInterface, diag.Diagnostics) {
	obj := models.RouterIsisIsisInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_keychain_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keychain_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthKeychainL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_keychain_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keychain_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthKeychainL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_mode_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_mode_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthModeL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_mode_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_mode_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthModeL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_password_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_password_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthPasswordL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_password_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_password_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthPasswordL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_send_only_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_send_only_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthSendOnlyL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_send_only_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_send_only_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthSendOnlyL2 = &v2
		}
	}
	if v1, ok := d.GetOk("circuit_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("circuit_type", sv)
				diags = append(diags, e)
			}
			obj.CircuitType = &v2
		}
	}
	if v1, ok := d.GetOk("csnp_interval_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("csnp_interval_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CsnpIntervalL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("csnp_interval_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("csnp_interval_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CsnpIntervalL2 = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_interval_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_interval_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloIntervalL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_interval_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_interval_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloIntervalL2 = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_multiplier_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_multiplier_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloMultiplierL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_multiplier_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_multiplier_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloMultiplierL2 = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_padding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_padding", sv)
				diags = append(diags, e)
			}
			obj.HelloPadding = &v2
		}
	}
	if v1, ok := d.GetOk("lsp_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lsp_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LspInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("lsp_retransmit_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lsp_retransmit_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LspRetransmitInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("mesh_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_group", sv)
				diags = append(diags, e)
			}
			obj.MeshGroup = &v2
		}
	}
	if v1, ok := d.GetOk("mesh_group_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_group_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MeshGroupId = &tmp
		}
	}
	if v1, ok := d.GetOk("metric_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("metric_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MetricL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("metric_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("metric_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MetricL2 = &tmp
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
	if v1, ok := d.GetOk("network_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_type", sv)
				diags = append(diags, e)
			}
			obj.NetworkType = &v2
		}
	}
	if v1, ok := d.GetOk("priority_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PriorityL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("priority_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PriorityL2 = &tmp
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
	if v1, ok := d.GetOk("status6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status6", sv)
				diags = append(diags, e)
			}
			obj.Status6 = &v2
		}
	}
	if v1, ok := d.GetOk("wide_metric_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wide_metric_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WideMetricL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("wide_metric_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wide_metric_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WideMetricL2 = &tmp
		}
	}
	return &obj, diags
}
