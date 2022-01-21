// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemEmailServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.",

		ReadContext: dataSourceSystemEmailServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authenticate": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "SMTP server user password for authentication.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "SMTP server port.",
				Computed:    true,
			},
			"reply_to": {
				Type:        schema.TypeString,
				Description: "Reply-To email address.",
				Computed:    true,
			},
			"security": {
				Type:        schema.TypeString,
				Description: "Connection security used by the email server.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "SMTP server IP address or hostname.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "SMTP server IPv4 source IP.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "SMTP server IPv6 source IP.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Use FortiGuard Message service or custom email server.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "SMTP server user name for authentication.",
				Computed:    true,
			},
			"validate_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable validation of server certificate.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemEmailServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemEmailServer"

	o, err := c.Cmdb.ReadSystemEmailServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemEmailServer dataSource: %v", err)
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

	diags := refreshObjectSystemEmailServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
