// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MODEM.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemModem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MODEM.",

		CreateContext: resourceSystemModemCreate,
		ReadContext:   resourceSystemModemRead,
		UpdateContext: resourceSystemModemUpdate,
		DeleteContext: resourceSystemModemDelete,

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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dial", "stop", "none"}, false),

				Description: "Dial up/stop MODEM.",
				Optional:    true,
				Computed:    true,
			},
			"altmode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable altmode for installations using PPP in China.",
				Optional:    true,
				Computed:    true,
			},
			"authtype1": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed authentication types for ISP 1.",
				Optional:         true,
				Computed:         true,
			},
			"authtype2": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed authentication types for ISP 2.",
				Optional:         true,
				Computed:         true,
			},
			"authtype3": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed authentication types for ISP 3.",
				Optional:         true,
				Computed:         true,
			},
			"auto_dial": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable auto-dial after a reboot or disconnection.",
				Optional:    true,
				Computed:    true,
			},
			"connect_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 255),

				Description: "Connection completion timeout (30 - 255 sec, default = 90).",
				Optional:    true,
				Computed:    true,
			},
			"dial_cmd1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Dial command (this is often an ATD or ATDT command).",
				Optional:    true,
				Computed:    true,
			},
			"dial_cmd2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Dial command (this is often an ATD or ATDT command).",
				Optional:    true,
				Computed:    true,
			},
			"dial_cmd3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Dial command (this is often an ATD or ATDT command).",
				Optional:    true,
				Computed:    true,
			},
			"dial_on_demand": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to dial the modem when packets are routed to the modem interface.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance of learned routes (1 - 255, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"dont_send_cr1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Do not send CR when connected (ISP1).",
				Optional:    true,
				Computed:    true,
			},
			"dont_send_cr2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Do not send CR when connected (ISP2).",
				Optional:    true,
				Computed:    true,
			},
			"dont_send_cr3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Do not send CR when connected (ISP3).",
				Optional:    true,
				Computed:    true,
			},
			"extra_init1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Extra initialization string to ISP 1.",
				Optional:    true,
				Computed:    true,
			},
			"extra_init2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Extra initialization string to ISP 2.",
				Optional:    true,
				Computed:    true,
			},
			"extra_init3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Extra initialization string to ISP 3.",
				Optional:    true,
				Computed:    true,
			},
			"holddown_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Hold down timer in seconds (1 - 60 sec).",
				Optional:    true,
				Computed:    true,
			},
			"idle_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 9999),

				Description: "MODEM connection idle time (1 - 9999 min, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of redundant interface.",
				Optional:    true,
				Computed:    true,
			},
			"lockdown_lac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Allow connection only to the specified Location Area Code (LAC).",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "redundant"}, false),

				Description: "Set MODEM operation mode to redundant or standalone.",
				Optional:    true,
				Computed:    true,
			},
			"network_init": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).",
				Optional:    true,
				Computed:    true,
			},
			"passwd1": {
				Type: schema.TypeString,

				Description: "Password to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"passwd2": {
				Type: schema.TypeString,

				Description: "Password to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"passwd3": {
				Type: schema.TypeString,

				Description: "Password to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"peer_modem1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"generic", "actiontec", "ascend_TNT"}, false),

				Description: "Specify peer MODEM type for phone1.",
				Optional:    true,
				Computed:    true,
			},
			"peer_modem2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"generic", "actiontec", "ascend_TNT"}, false),

				Description: "Specify peer MODEM type for phone2.",
				Optional:    true,
				Computed:    true,
			},
			"peer_modem3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"generic", "actiontec", "ascend_TNT"}, false),

				Description: "Specify peer MODEM type for phone3.",
				Optional:    true,
				Computed:    true,
			},
			"phone1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Optional:    true,
				Computed:    true,
			},
			"phone2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Optional:    true,
				Computed:    true,
			},
			"phone3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Optional:    true,
				Computed:    true,
			},
			"pin_init": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "AT command to set the PIN (AT+PIN=<pin>).",
				Optional:    true,
				Computed:    true,
			},
			"ppp_echo_request1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPP echo-request to ISP 1.",
				Optional:    true,
				Computed:    true,
			},
			"ppp_echo_request2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPP echo-request to ISP 2.",
				Optional:    true,
				Computed:    true,
			},
			"ppp_echo_request3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPP echo-request to ISP 3.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type: schema.TypeInt,

				Description: "Priority of learned routes (0 - 4294967295, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"redial": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, false),

				Description: "Redial limit (1 - 10 attempts, none = redial forever).",
				Optional:    true,
				Computed:    true,
			},
			"reset": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),

				Description: "Number of dial attempts before resetting modem (0 = never reset).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Modem support (equivalent to bringing an interface up or down).",
				Optional:    true,
				Computed:    true,
			},
			"traffic_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable traffic-check.",
				Optional:    true,
				Computed:    true,
			},
			"username1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "User name to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
			},
			"username2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "User name to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
			},
			"username3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "User name to access the specified dialup account.",
				Optional:    true,
				Computed:    true,
			},
			"wireless_port": {
				Type: schema.TypeInt,

				Description: "Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemModemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemModem(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemModem")
	}

	return resourceSystemModemRead(ctx, d, meta)
}

func resourceSystemModemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemModem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemModem resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemModem")
	}

	return resourceSystemModemRead(ctx, d, meta)
}

func resourceSystemModemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemModem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemModem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemModemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemModem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemModem resource: %v", err)
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

	diags := refreshObjectSystemModem(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemModem(d *schema.ResourceData, o *models.SystemModem, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Altmode != nil {
		v := *o.Altmode

		if err = d.Set("altmode", v); err != nil {
			return diag.Errorf("error reading altmode: %v", err)
		}
	}

	if o.Authtype1 != nil {
		v := *o.Authtype1

		if err = d.Set("authtype1", v); err != nil {
			return diag.Errorf("error reading authtype1: %v", err)
		}
	}

	if o.Authtype2 != nil {
		v := *o.Authtype2

		if err = d.Set("authtype2", v); err != nil {
			return diag.Errorf("error reading authtype2: %v", err)
		}
	}

	if o.Authtype3 != nil {
		v := *o.Authtype3

		if err = d.Set("authtype3", v); err != nil {
			return diag.Errorf("error reading authtype3: %v", err)
		}
	}

	if o.AutoDial != nil {
		v := *o.AutoDial

		if err = d.Set("auto_dial", v); err != nil {
			return diag.Errorf("error reading auto_dial: %v", err)
		}
	}

	if o.ConnectTimeout != nil {
		v := *o.ConnectTimeout

		if err = d.Set("connect_timeout", v); err != nil {
			return diag.Errorf("error reading connect_timeout: %v", err)
		}
	}

	if o.DialCmd1 != nil {
		v := *o.DialCmd1

		if err = d.Set("dial_cmd1", v); err != nil {
			return diag.Errorf("error reading dial_cmd1: %v", err)
		}
	}

	if o.DialCmd2 != nil {
		v := *o.DialCmd2

		if err = d.Set("dial_cmd2", v); err != nil {
			return diag.Errorf("error reading dial_cmd2: %v", err)
		}
	}

	if o.DialCmd3 != nil {
		v := *o.DialCmd3

		if err = d.Set("dial_cmd3", v); err != nil {
			return diag.Errorf("error reading dial_cmd3: %v", err)
		}
	}

	if o.DialOnDemand != nil {
		v := *o.DialOnDemand

		if err = d.Set("dial_on_demand", v); err != nil {
			return diag.Errorf("error reading dial_on_demand: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DontSendCR1 != nil {
		v := *o.DontSendCR1

		if err = d.Set("dont_send_cr1", v); err != nil {
			return diag.Errorf("error reading dont_send_cr1: %v", err)
		}
	}

	if o.DontSendCR2 != nil {
		v := *o.DontSendCR2

		if err = d.Set("dont_send_cr2", v); err != nil {
			return diag.Errorf("error reading dont_send_cr2: %v", err)
		}
	}

	if o.DontSendCR3 != nil {
		v := *o.DontSendCR3

		if err = d.Set("dont_send_cr3", v); err != nil {
			return diag.Errorf("error reading dont_send_cr3: %v", err)
		}
	}

	if o.ExtraInit1 != nil {
		v := *o.ExtraInit1

		if err = d.Set("extra_init1", v); err != nil {
			return diag.Errorf("error reading extra_init1: %v", err)
		}
	}

	if o.ExtraInit2 != nil {
		v := *o.ExtraInit2

		if err = d.Set("extra_init2", v); err != nil {
			return diag.Errorf("error reading extra_init2: %v", err)
		}
	}

	if o.ExtraInit3 != nil {
		v := *o.ExtraInit3

		if err = d.Set("extra_init3", v); err != nil {
			return diag.Errorf("error reading extra_init3: %v", err)
		}
	}

	if o.HolddownTimer != nil {
		v := *o.HolddownTimer

		if err = d.Set("holddown_timer", v); err != nil {
			return diag.Errorf("error reading holddown_timer: %v", err)
		}
	}

	if o.IdleTimer != nil {
		v := *o.IdleTimer

		if err = d.Set("idle_timer", v); err != nil {
			return diag.Errorf("error reading idle_timer: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.LockdownLac != nil {
		v := *o.LockdownLac

		if err = d.Set("lockdown_lac", v); err != nil {
			return diag.Errorf("error reading lockdown_lac: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.NetworkInit != nil {
		v := *o.NetworkInit

		if err = d.Set("network_init", v); err != nil {
			return diag.Errorf("error reading network_init: %v", err)
		}
	}

	if o.Passwd1 != nil {
		v := *o.Passwd1

		if v == "ENC XXXX" {
		} else if err = d.Set("passwd1", v); err != nil {
			return diag.Errorf("error reading passwd1: %v", err)
		}
	}

	if o.Passwd2 != nil {
		v := *o.Passwd2

		if v == "ENC XXXX" {
		} else if err = d.Set("passwd2", v); err != nil {
			return diag.Errorf("error reading passwd2: %v", err)
		}
	}

	if o.Passwd3 != nil {
		v := *o.Passwd3

		if v == "ENC XXXX" {
		} else if err = d.Set("passwd3", v); err != nil {
			return diag.Errorf("error reading passwd3: %v", err)
		}
	}

	if o.PeerModem1 != nil {
		v := *o.PeerModem1

		if err = d.Set("peer_modem1", v); err != nil {
			return diag.Errorf("error reading peer_modem1: %v", err)
		}
	}

	if o.PeerModem2 != nil {
		v := *o.PeerModem2

		if err = d.Set("peer_modem2", v); err != nil {
			return diag.Errorf("error reading peer_modem2: %v", err)
		}
	}

	if o.PeerModem3 != nil {
		v := *o.PeerModem3

		if err = d.Set("peer_modem3", v); err != nil {
			return diag.Errorf("error reading peer_modem3: %v", err)
		}
	}

	if o.Phone1 != nil {
		v := *o.Phone1

		if err = d.Set("phone1", v); err != nil {
			return diag.Errorf("error reading phone1: %v", err)
		}
	}

	if o.Phone2 != nil {
		v := *o.Phone2

		if err = d.Set("phone2", v); err != nil {
			return diag.Errorf("error reading phone2: %v", err)
		}
	}

	if o.Phone3 != nil {
		v := *o.Phone3

		if err = d.Set("phone3", v); err != nil {
			return diag.Errorf("error reading phone3: %v", err)
		}
	}

	if o.PinInit != nil {
		v := *o.PinInit

		if err = d.Set("pin_init", v); err != nil {
			return diag.Errorf("error reading pin_init: %v", err)
		}
	}

	if o.PppEchoRequest1 != nil {
		v := *o.PppEchoRequest1

		if err = d.Set("ppp_echo_request1", v); err != nil {
			return diag.Errorf("error reading ppp_echo_request1: %v", err)
		}
	}

	if o.PppEchoRequest2 != nil {
		v := *o.PppEchoRequest2

		if err = d.Set("ppp_echo_request2", v); err != nil {
			return diag.Errorf("error reading ppp_echo_request2: %v", err)
		}
	}

	if o.PppEchoRequest3 != nil {
		v := *o.PppEchoRequest3

		if err = d.Set("ppp_echo_request3", v); err != nil {
			return diag.Errorf("error reading ppp_echo_request3: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Redial != nil {
		v := *o.Redial

		if err = d.Set("redial", v); err != nil {
			return diag.Errorf("error reading redial: %v", err)
		}
	}

	if o.Reset != nil {
		v := *o.Reset

		if err = d.Set("reset", v); err != nil {
			return diag.Errorf("error reading reset: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TrafficCheck != nil {
		v := *o.TrafficCheck

		if err = d.Set("traffic_check", v); err != nil {
			return diag.Errorf("error reading traffic_check: %v", err)
		}
	}

	if o.Username1 != nil {
		v := *o.Username1

		if err = d.Set("username1", v); err != nil {
			return diag.Errorf("error reading username1: %v", err)
		}
	}

	if o.Username2 != nil {
		v := *o.Username2

		if err = d.Set("username2", v); err != nil {
			return diag.Errorf("error reading username2: %v", err)
		}
	}

	if o.Username3 != nil {
		v := *o.Username3

		if err = d.Set("username3", v); err != nil {
			return diag.Errorf("error reading username3: %v", err)
		}
	}

	if o.WirelessPort != nil {
		v := *o.WirelessPort

		if err = d.Set("wireless_port", v); err != nil {
			return diag.Errorf("error reading wireless_port: %v", err)
		}
	}

	return nil
}

func getObjectSystemModem(d *schema.ResourceData, sv string) (*models.SystemModem, diag.Diagnostics) {
	obj := models.SystemModem{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("altmode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("altmode", sv)
				diags = append(diags, e)
			}
			obj.Altmode = &v2
		}
	}
	if v1, ok := d.GetOk("authtype1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtype1", sv)
				diags = append(diags, e)
			}
			obj.Authtype1 = &v2
		}
	}
	if v1, ok := d.GetOk("authtype2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtype2", sv)
				diags = append(diags, e)
			}
			obj.Authtype2 = &v2
		}
	}
	if v1, ok := d.GetOk("authtype3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtype3", sv)
				diags = append(diags, e)
			}
			obj.Authtype3 = &v2
		}
	}
	if v1, ok := d.GetOk("auto_dial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_dial", sv)
				diags = append(diags, e)
			}
			obj.AutoDial = &v2
		}
	}
	if v1, ok := d.GetOk("connect_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("connect_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConnectTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("dial_cmd1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dial_cmd1", sv)
				diags = append(diags, e)
			}
			obj.DialCmd1 = &v2
		}
	}
	if v1, ok := d.GetOk("dial_cmd2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dial_cmd2", sv)
				diags = append(diags, e)
			}
			obj.DialCmd2 = &v2
		}
	}
	if v1, ok := d.GetOk("dial_cmd3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dial_cmd3", sv)
				diags = append(diags, e)
			}
			obj.DialCmd3 = &v2
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
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
		}
	}
	if v1, ok := d.GetOk("dont_send_cr1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dont_send_cr1", sv)
				diags = append(diags, e)
			}
			obj.DontSendCR1 = &v2
		}
	}
	if v1, ok := d.GetOk("dont_send_cr2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dont_send_cr2", sv)
				diags = append(diags, e)
			}
			obj.DontSendCR2 = &v2
		}
	}
	if v1, ok := d.GetOk("dont_send_cr3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dont_send_cr3", sv)
				diags = append(diags, e)
			}
			obj.DontSendCR3 = &v2
		}
	}
	if v1, ok := d.GetOk("extra_init1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extra_init1", sv)
				diags = append(diags, e)
			}
			obj.ExtraInit1 = &v2
		}
	}
	if v1, ok := d.GetOk("extra_init2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extra_init2", sv)
				diags = append(diags, e)
			}
			obj.ExtraInit2 = &v2
		}
	}
	if v1, ok := d.GetOk("extra_init3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extra_init3", sv)
				diags = append(diags, e)
			}
			obj.ExtraInit3 = &v2
		}
	}
	if v1, ok := d.GetOk("holddown_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("holddown_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HolddownTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("idle_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdleTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("lockdown_lac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lockdown_lac", sv)
				diags = append(diags, e)
			}
			obj.LockdownLac = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("network_init"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_init", sv)
				diags = append(diags, e)
			}
			obj.NetworkInit = &v2
		}
	}
	if v1, ok := d.GetOk("passwd1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd1", sv)
				diags = append(diags, e)
			}
			obj.Passwd1 = &v2
		}
	}
	if v1, ok := d.GetOk("passwd2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd2", sv)
				diags = append(diags, e)
			}
			obj.Passwd2 = &v2
		}
	}
	if v1, ok := d.GetOk("passwd3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd3", sv)
				diags = append(diags, e)
			}
			obj.Passwd3 = &v2
		}
	}
	if v1, ok := d.GetOk("peer_modem1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_modem1", sv)
				diags = append(diags, e)
			}
			obj.PeerModem1 = &v2
		}
	}
	if v1, ok := d.GetOk("peer_modem2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_modem2", sv)
				diags = append(diags, e)
			}
			obj.PeerModem2 = &v2
		}
	}
	if v1, ok := d.GetOk("peer_modem3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_modem3", sv)
				diags = append(diags, e)
			}
			obj.PeerModem3 = &v2
		}
	}
	if v1, ok := d.GetOk("phone1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phone1", sv)
				diags = append(diags, e)
			}
			obj.Phone1 = &v2
		}
	}
	if v1, ok := d.GetOk("phone2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phone2", sv)
				diags = append(diags, e)
			}
			obj.Phone2 = &v2
		}
	}
	if v1, ok := d.GetOk("phone3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phone3", sv)
				diags = append(diags, e)
			}
			obj.Phone3 = &v2
		}
	}
	if v1, ok := d.GetOk("pin_init"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pin_init", sv)
				diags = append(diags, e)
			}
			obj.PinInit = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_echo_request1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppp_echo_request1", sv)
				diags = append(diags, e)
			}
			obj.PppEchoRequest1 = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_echo_request2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppp_echo_request2", sv)
				diags = append(diags, e)
			}
			obj.PppEchoRequest2 = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_echo_request3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppp_echo_request3", sv)
				diags = append(diags, e)
			}
			obj.PppEchoRequest3 = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("redial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redial", sv)
				diags = append(diags, e)
			}
			obj.Redial = &v2
		}
	}
	if v1, ok := d.GetOk("reset"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reset", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Reset = &tmp
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
	if v1, ok := d.GetOk("traffic_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_check", sv)
				diags = append(diags, e)
			}
			obj.TrafficCheck = &v2
		}
	}
	if v1, ok := d.GetOk("username1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username1", sv)
				diags = append(diags, e)
			}
			obj.Username1 = &v2
		}
	}
	if v1, ok := d.GetOk("username2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username2", sv)
				diags = append(diags, e)
			}
			obj.Username2 = &v2
		}
	}
	if v1, ok := d.GetOk("username3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username3", sv)
				diags = append(diags, e)
			}
			obj.Username3 = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wireless_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WirelessPort = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemModem(d *schema.ResourceData, sv string) (*models.SystemModem, diag.Diagnostics) {
	obj := models.SystemModem{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
