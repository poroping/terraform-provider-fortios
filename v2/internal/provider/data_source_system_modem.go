// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MODEM.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemModem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MODEM.",

		ReadContext: dataSourceSystemModemRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Dial up/stop MODEM.",
				Computed:    true,
			},
			"altmode": {
				Type:        schema.TypeString,
				Description: "Enable/disable altmode for installations using PPP in China.",
				Computed:    true,
			},
			"authtype1": {
				Type:        schema.TypeString,
				Description: "Allowed authentication types for ISP 1.",
				Computed:    true,
			},
			"authtype2": {
				Type:        schema.TypeString,
				Description: "Allowed authentication types for ISP 2.",
				Computed:    true,
			},
			"authtype3": {
				Type:        schema.TypeString,
				Description: "Allowed authentication types for ISP 3.",
				Computed:    true,
			},
			"auto_dial": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto-dial after a reboot or disconnection.",
				Computed:    true,
			},
			"connect_timeout": {
				Type:        schema.TypeInt,
				Description: "Connection completion timeout (30 - 255 sec, default = 90).",
				Computed:    true,
			},
			"dial_cmd1": {
				Type:        schema.TypeString,
				Description: "Dial command (this is often an ATD or ATDT command).",
				Computed:    true,
			},
			"dial_cmd2": {
				Type:        schema.TypeString,
				Description: "Dial command (this is often an ATD or ATDT command).",
				Computed:    true,
			},
			"dial_cmd3": {
				Type:        schema.TypeString,
				Description: "Dial command (this is often an ATD or ATDT command).",
				Computed:    true,
			},
			"dial_on_demand": {
				Type:        schema.TypeString,
				Description: "Enable/disable to dial the modem when packets are routed to the modem interface.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Distance of learned routes (1 - 255, default = 1).",
				Computed:    true,
			},
			"dont_send_cr1": {
				Type:        schema.TypeString,
				Description: "Do not send CR when connected (ISP1).",
				Computed:    true,
			},
			"dont_send_cr2": {
				Type:        schema.TypeString,
				Description: "Do not send CR when connected (ISP2).",
				Computed:    true,
			},
			"dont_send_cr3": {
				Type:        schema.TypeString,
				Description: "Do not send CR when connected (ISP3).",
				Computed:    true,
			},
			"extra_init1": {
				Type:        schema.TypeString,
				Description: "Extra initialization string to ISP 1.",
				Computed:    true,
			},
			"extra_init2": {
				Type:        schema.TypeString,
				Description: "Extra initialization string to ISP 2.",
				Computed:    true,
			},
			"extra_init3": {
				Type:        schema.TypeString,
				Description: "Extra initialization string to ISP 3.",
				Computed:    true,
			},
			"holddown_timer": {
				Type:        schema.TypeInt,
				Description: "Hold down timer in seconds (1 - 60 sec).",
				Computed:    true,
			},
			"idle_timer": {
				Type:        schema.TypeInt,
				Description: "MODEM connection idle time (1 - 9999 min, default = 5).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Name of redundant interface.",
				Computed:    true,
			},
			"lockdown_lac": {
				Type:        schema.TypeString,
				Description: "Allow connection only to the specified Location Area Code (LAC).",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Set MODEM operation mode to redundant or standalone.",
				Computed:    true,
			},
			"network_init": {
				Type:        schema.TypeString,
				Description: "AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).",
				Computed:    true,
			},
			"passwd1": {
				Type:        schema.TypeString,
				Description: "Password to access the specified dialup account.",
				Computed:    true,
				Sensitive:   true,
			},
			"passwd2": {
				Type:        schema.TypeString,
				Description: "Password to access the specified dialup account.",
				Computed:    true,
				Sensitive:   true,
			},
			"passwd3": {
				Type:        schema.TypeString,
				Description: "Password to access the specified dialup account.",
				Computed:    true,
				Sensitive:   true,
			},
			"peer_modem1": {
				Type:        schema.TypeString,
				Description: "Specify peer MODEM type for phone1.",
				Computed:    true,
			},
			"peer_modem2": {
				Type:        schema.TypeString,
				Description: "Specify peer MODEM type for phone2.",
				Computed:    true,
			},
			"peer_modem3": {
				Type:        schema.TypeString,
				Description: "Specify peer MODEM type for phone3.",
				Computed:    true,
			},
			"phone1": {
				Type:        schema.TypeString,
				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Computed:    true,
			},
			"phone2": {
				Type:        schema.TypeString,
				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Computed:    true,
			},
			"phone3": {
				Type:        schema.TypeString,
				Description: "Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).",
				Computed:    true,
			},
			"pin_init": {
				Type:        schema.TypeString,
				Description: "AT command to set the PIN (AT+PIN=<pin>).",
				Computed:    true,
			},
			"ppp_echo_request1": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPP echo-request to ISP 1.",
				Computed:    true,
			},
			"ppp_echo_request2": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPP echo-request to ISP 2.",
				Computed:    true,
			},
			"ppp_echo_request3": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPP echo-request to ISP 3.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority of learned routes (0 - 4294967295, default = 0).",
				Computed:    true,
			},
			"redial": {
				Type:        schema.TypeString,
				Description: "Redial limit (1 - 10 attempts, none = redial forever).",
				Computed:    true,
			},
			"reset": {
				Type:        schema.TypeInt,
				Description: "Number of dial attempts before resetting modem (0 = never reset).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable Modem support (equivalent to bringing an interface up or down).",
				Computed:    true,
			},
			"traffic_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable traffic-check.",
				Computed:    true,
			},
			"username1": {
				Type:        schema.TypeString,
				Description: "User name to access the specified dialup account.",
				Computed:    true,
			},
			"username2": {
				Type:        schema.TypeString,
				Description: "User name to access the specified dialup account.",
				Computed:    true,
			},
			"username3": {
				Type:        schema.TypeString,
				Description: "User name to access the specified dialup account.",
				Computed:    true,
			},
			"wireless_port": {
				Type:        schema.TypeInt,
				Description: "Enter wireless port number: 0 for default, 1 for first port, and so on (0 - 4294967295).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemModemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemModem"

	o, err := c.Cmdb.ReadSystemModem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemModem dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
