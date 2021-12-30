// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: FortiExtender dataplan configuration.

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

func resourceExtenderControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Description: "FortiExtender dataplan configuration.",

		CreateContext: resourceExtenderControllerDataplanCreate,
		ReadContext:   resourceExtenderControllerDataplanRead,
		UpdateContext: resourceExtenderControllerDataplanUpdate,
		DeleteContext: resourceExtenderControllerDataplanDelete,

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
			"apn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "APN configuration.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "pap", "chap"}, false),

				Description: "Authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"billing_date": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),

				Description: "Billing day of the month (1 - 31).",
				Optional:    true,
				Computed:    true,
			},
			"capacity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 102400000),

				Description: "Capacity in MB (0 - 102400000).",
				Optional:    true,
				Computed:    true,
			},
			"carrier": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Carrier configuration.",
				Optional:    true,
				Computed:    true,
			},
			"iccid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "ICCID configuration.",
				Optional:    true,
				Computed:    true,
			},
			"modem_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"modem1", "modem2", "all"}, false),

				Description: "Dataplan's modem specifics, if any.",
				Optional:    true,
				Computed:    true,
			},
			"monthly_fee": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000000),

				Description: "Monthly fee of dataplan (0 - 100000, in local currency).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "FortiExtender dataplan name",
				ForceNew:    true,
				Required:    true,
			},
			"overage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dataplan overage detection.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"pdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4-only", "ipv6-only", "ipv4-ipv6"}, false),

				Description: "PDN type.",
				Optional:    true,
				Computed:    true,
			},
			"preferred_subnet": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Preferred subnet mask (0 - 32).",
				Optional:    true,
				Computed:    true,
			},
			"private_network": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dataplan private network support.",
				Optional:    true,
				Computed:    true,
			},
			"signal_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 18000),

				Description: "Signal period (600 to 18000 seconds).",
				Optional:    true,
				Computed:    true,
			},
			"signal_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 100),

				Description: "Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.",
				Optional:    true,
				Computed:    true,
			},
			"slot": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sim1", "sim2"}, false),

				Description: "SIM slot configuration.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"carrier", "slot", "iccid", "generic"}, false),

				Description: "Type preferences configuration.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Username.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceExtenderControllerDataplanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ExtenderControllerDataplan resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectExtenderControllerDataplan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateExtenderControllerDataplan(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerDataplan")
	}

	return resourceExtenderControllerDataplanRead(ctx, d, meta)
}

func resourceExtenderControllerDataplanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectExtenderControllerDataplan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateExtenderControllerDataplan(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ExtenderControllerDataplan resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerDataplan")
	}

	return resourceExtenderControllerDataplanRead(ctx, d, meta)
}

func resourceExtenderControllerDataplanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteExtenderControllerDataplan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ExtenderControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerDataplanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadExtenderControllerDataplan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ExtenderControllerDataplan resource: %v", err)
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

	diags := refreshObjectExtenderControllerDataplan(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectExtenderControllerDataplan(d *schema.ResourceData, o *models.ExtenderControllerDataplan, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Apn != nil {
		v := *o.Apn

		if err = d.Set("apn", v); err != nil {
			return diag.Errorf("error reading apn: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.BillingDate != nil {
		v := *o.BillingDate

		if err = d.Set("billing_date", v); err != nil {
			return diag.Errorf("error reading billing_date: %v", err)
		}
	}

	if o.Capacity != nil {
		v := *o.Capacity

		if err = d.Set("capacity", v); err != nil {
			return diag.Errorf("error reading capacity: %v", err)
		}
	}

	if o.Carrier != nil {
		v := *o.Carrier

		if err = d.Set("carrier", v); err != nil {
			return diag.Errorf("error reading carrier: %v", err)
		}
	}

	if o.Iccid != nil {
		v := *o.Iccid

		if err = d.Set("iccid", v); err != nil {
			return diag.Errorf("error reading iccid: %v", err)
		}
	}

	if o.ModemId != nil {
		v := *o.ModemId

		if err = d.Set("modem_id", v); err != nil {
			return diag.Errorf("error reading modem_id: %v", err)
		}
	}

	if o.MonthlyFee != nil {
		v := *o.MonthlyFee

		if err = d.Set("monthly_fee", v); err != nil {
			return diag.Errorf("error reading monthly_fee: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Overage != nil {
		v := *o.Overage

		if err = d.Set("overage", v); err != nil {
			return diag.Errorf("error reading overage: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Pdn != nil {
		v := *o.Pdn

		if err = d.Set("pdn", v); err != nil {
			return diag.Errorf("error reading pdn: %v", err)
		}
	}

	if o.PreferredSubnet != nil {
		v := *o.PreferredSubnet

		if err = d.Set("preferred_subnet", v); err != nil {
			return diag.Errorf("error reading preferred_subnet: %v", err)
		}
	}

	if o.PrivateNetwork != nil {
		v := *o.PrivateNetwork

		if err = d.Set("private_network", v); err != nil {
			return diag.Errorf("error reading private_network: %v", err)
		}
	}

	if o.SignalPeriod != nil {
		v := *o.SignalPeriod

		if err = d.Set("signal_period", v); err != nil {
			return diag.Errorf("error reading signal_period: %v", err)
		}
	}

	if o.SignalThreshold != nil {
		v := *o.SignalThreshold

		if err = d.Set("signal_threshold", v); err != nil {
			return diag.Errorf("error reading signal_threshold: %v", err)
		}
	}

	if o.Slot != nil {
		v := *o.Slot

		if err = d.Set("slot", v); err != nil {
			return diag.Errorf("error reading slot: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
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

func getObjectExtenderControllerDataplan(d *schema.ResourceData, sv string) (*models.ExtenderControllerDataplan, diag.Diagnostics) {
	obj := models.ExtenderControllerDataplan{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("apn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("apn", sv)
				diags = append(diags, e)
			}
			obj.Apn = &v2
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
	if v1, ok := d.GetOk("billing_date"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("billing_date", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BillingDate = &tmp
		}
	}
	if v1, ok := d.GetOk("capacity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capacity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Capacity = &tmp
		}
	}
	if v1, ok := d.GetOk("carrier"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("carrier", sv)
				diags = append(diags, e)
			}
			obj.Carrier = &v2
		}
	}
	if v1, ok := d.GetOk("iccid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("iccid", sv)
				diags = append(diags, e)
			}
			obj.Iccid = &v2
		}
	}
	if v1, ok := d.GetOk("modem_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("modem_id", sv)
				diags = append(diags, e)
			}
			obj.ModemId = &v2
		}
	}
	if v1, ok := d.GetOk("monthly_fee"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monthly_fee", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MonthlyFee = &tmp
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
	if v1, ok := d.GetOk("overage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("overage", sv)
				diags = append(diags, e)
			}
			obj.Overage = &v2
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
	if v1, ok := d.GetOk("pdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pdn", sv)
				diags = append(diags, e)
			}
			obj.Pdn = &v2
		}
	}
	if v1, ok := d.GetOk("preferred_subnet"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("preferred_subnet", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PreferredSubnet = &tmp
		}
	}
	if v1, ok := d.GetOk("private_network"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("private_network", sv)
				diags = append(diags, e)
			}
			obj.PrivateNetwork = &v2
		}
	}
	if v1, ok := d.GetOk("signal_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signal_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SignalPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("signal_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signal_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SignalThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("slot"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("slot", sv)
				diags = append(diags, e)
			}
			obj.Slot = &v2
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
