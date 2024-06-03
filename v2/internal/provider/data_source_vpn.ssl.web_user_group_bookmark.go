// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL-VPN user group bookmark.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnSslWebUserGroupBookmark() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL-VPN user group bookmark.",

		ReadContext: dataSourceVpnSslWebUserGroupBookmarkRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bookmarks": {
				Type:        schema.TypeList,
				Description: "Bookmark table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_params": {
							Type:        schema.TypeString,
							Description: "Additional parameters.",
							Computed:    true,
						},
						"apptype": {
							Type:        schema.TypeString,
							Description: "Application type.",
							Computed:    true,
						},
						"color_depth": {
							Type:        schema.TypeString,
							Description: "Color depth per pixel.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"domain": {
							Type:        schema.TypeString,
							Description: "Login domain.",
							Computed:    true,
						},
						"folder": {
							Type:        schema.TypeString,
							Description: "Network shared file folder parameter.",
							Computed:    true,
						},
						"form_data": {
							Type:        schema.TypeList,
							Description: "Form data.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Name.",
										Computed:    true,
									},
									"value": {
										Type:        schema.TypeString,
										Description: "Value.",
										Computed:    true,
									},
								},
							},
						},
						"height": {
							Type:        schema.TypeInt,
							Description: "Screen height (range from 0 - 65535, default = 0).",
							Computed:    true,
						},
						"host": {
							Type:        schema.TypeString,
							Description: "Host name/IP parameter.",
							Computed:    true,
						},
						"keyboard_layout": {
							Type:        schema.TypeString,
							Description: "Keyboard layout.",
							Computed:    true,
						},
						"listening_port": {
							Type:        schema.TypeInt,
							Description: "Listening port (0 - 65535).",
							Computed:    true,
						},
						"load_balancing_info": {
							Type:        schema.TypeString,
							Description: "The load balancing information or cookie which should be provided to the connection broker.",
							Computed:    true,
						},
						"logon_password": {
							Type:        schema.TypeString,
							Description: "Logon password.",
							Computed:    true,
							Sensitive:   true,
						},
						"logon_user": {
							Type:        schema.TypeString,
							Description: "Logon user.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Bookmark name.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Remote port.",
							Computed:    true,
						},
						"preconnection_blob": {
							Type:        schema.TypeString,
							Description: "An arbitrary string which identifies the RDP source.",
							Computed:    true,
						},
						"preconnection_id": {
							Type:        schema.TypeInt,
							Description: "The numeric ID of the RDP source (0-4294967295).",
							Computed:    true,
						},
						"remote_port": {
							Type:        schema.TypeInt,
							Description: "Remote port (0 - 65535).",
							Computed:    true,
						},
						"restricted_admin": {
							Type:        schema.TypeString,
							Description: "Enable/disable restricted admin mode for RDP.",
							Computed:    true,
						},
						"security": {
							Type:        schema.TypeString,
							Description: "Security mode for RDP connection.",
							Computed:    true,
						},
						"send_preconnection_id": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending of preconnection ID.",
							Computed:    true,
						},
						"server_layout": {
							Type:        schema.TypeString,
							Description: "Server side keyboard layout.",
							Computed:    true,
						},
						"show_status_window": {
							Type:        schema.TypeString,
							Description: "Enable/disable showing of status window.",
							Computed:    true,
						},
						"sso": {
							Type:        schema.TypeString,
							Description: "Single Sign-On.",
							Computed:    true,
						},
						"sso_credential": {
							Type:        schema.TypeString,
							Description: "Single sign-on credentials.",
							Computed:    true,
						},
						"sso_credential_sent_once": {
							Type:        schema.TypeString,
							Description: "Single sign-on credentials are only sent once to remote server.",
							Computed:    true,
						},
						"sso_password": {
							Type:        schema.TypeString,
							Description: "SSO password.",
							Computed:    true,
							Sensitive:   true,
						},
						"sso_username": {
							Type:        schema.TypeString,
							Description: "SSO user name.",
							Computed:    true,
						},
						"url": {
							Type:        schema.TypeString,
							Description: "URL parameter.",
							Computed:    true,
						},
						"vnc_keyboard_layout": {
							Type:        schema.TypeString,
							Description: "Keyboard layout.",
							Computed:    true,
						},
						"width": {
							Type:        schema.TypeInt,
							Description: "Screen width (range from 0 - 65535, default = 0).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Group name.",
				Required:    true,
			},
		},
	}
}

func dataSourceVpnSslWebUserGroupBookmarkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadVpnSslWebUserGroupBookmark(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebUserGroupBookmark dataSource: %v", err)
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

	diags := refreshObjectVpnSslWebUserGroupBookmark(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
