// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL-VPN user group bookmark.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceVpnSslWebUserGroupBookmark() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL-VPN user group bookmark.",

		CreateContext: resourceVpnSslWebUserGroupBookmarkCreate,
		ReadContext:   resourceVpnSslWebUserGroupBookmarkRead,
		UpdateContext: resourceVpnSslWebUserGroupBookmarkUpdate,
		DeleteContext: resourceVpnSslWebUserGroupBookmarkDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"bookmarks": {
				Type:        schema.TypeList,
				Description: "Bookmark table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_params": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "Additional parameters.",
							Optional:    true,
							Computed:    true,
						},
						"apptype": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ftp", "rdp", "sftp", "smb", "ssh", "telnet", "vnc", "web"}, false),

							Description: "Application type.",
							Optional:    true,
							Computed:    true,
						},
						"color_depth": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"32", "16", "8"}, false),

							Description: "Color depth per pixel.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "Login domain.",
							Optional:    true,
							Computed:    true,
						},
						"folder": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "Network shared file folder parameter.",
							Optional:    true,
							Computed:    true,
						},
						"form_data": {
							Type:        schema.TypeList,
							Description: "Form data.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name.",
										Optional:    true,
										Computed:    true,
									},
									"value": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Value.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"height": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(480, 65535),

							Description: "Screen height (range from 480 - 65535, default = 768).",
							Optional:    true,
							Computed:    true,
						},
						"host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "Host name/IP parameter.",
							Optional:    true,
							Computed:    true,
						},
						"keyboard_layout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ar-101", "ar-102", "ar-102-azerty", "can-mul", "cz", "cz-qwerty", "cz-pr", "da", "nl", "de", "de-ch", "de-ibm", "en-uk", "en-uk-ext", "en-us", "en-us-dvorak", "es", "es-var", "fi", "fi-sami", "fr", "fr-ca", "fr-ch", "fr-be", "hr", "hu", "hu-101", "it", "it-142", "ja", "ko", "lt", "lt-ibm", "lt-std", "lav-std", "lav-leg", "mk", "mk-std", "no", "no-sami", "pol-214", "pol-pr", "pt", "pt-br", "pt-br-abnt2", "ru", "ru-mne", "ru-t", "sl", "sv", "sv-sami", "tuk", "tur-f", "tur-q", "zh-sym-sg-us", "zh-sym-us", "zh-tr-hk", "zh-tr-mo", "zh-tr-us"}, false),

							Description: "Keyboard layout.",
							Optional:    true,
							Computed:    true,
						},
						"listening_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Listening port (0 - 65535).",
							Optional:    true,
							Computed:    true,
						},
						"load_balancing_info": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "The load balancing information or cookie which should be provided to the connection broker.",
							Optional:    true,
							Computed:    true,
						},
						"logon_password": {
							Type: schema.TypeString,

							Description: "Logon password.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"logon_user": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Logon user.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Bookmark name.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Remote port.",
							Optional:    true,
							Computed:    true,
						},
						"preconnection_blob": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "An arbitrary string which identifies the RDP source.",
							Optional:    true,
							Computed:    true,
						},
						"preconnection_id": {
							Type: schema.TypeInt,

							Description: "The numeric ID of the RDP source (0-4294967295).",
							Optional:    true,
							Computed:    true,
						},
						"remote_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Remote port (0 - 65535).",
							Optional:    true,
							Computed:    true,
						},
						"restricted_admin": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable restricted admin mode for RDP.",
							Optional:    true,
							Computed:    true,
						},
						"security": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"rdp", "nla", "tls", "any"}, false),

							Description: "Security mode for RDP connection.",
							Optional:    true,
							Computed:    true,
						},
						"send_preconnection_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sending of preconnection ID.",
							Optional:    true,
							Computed:    true,
						},
						"server_layout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"de-de-qwertz", "en-gb-qwerty", "en-us-qwerty", "es-es-qwerty", "fr-ca-qwerty", "fr-fr-azerty", "fr-ch-qwertz", "it-it-qwerty", "ja-jp-qwerty", "pt-br-qwerty", "sv-se-qwerty", "tr-tr-qwerty", "failsafe"}, false),

							Description: "Server side keyboard layout.",
							Optional:    true,
							Computed:    true,
						},
						"show_status_window": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable showing of status window.",
							Optional:    true,
							Computed:    true,
						},
						"sso": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "static", "auto"}, false),

							Description: "Single Sign-On.",
							Optional:    true,
							Computed:    true,
						},
						"sso_credential": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"sslvpn-login", "alternative"}, false),

							Description: "Single sign-on credentials.",
							Optional:    true,
							Computed:    true,
						},
						"sso_credential_sent_once": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Single sign-on credentials are only sent once to remote server.",
							Optional:    true,
							Computed:    true,
						},
						"sso_password": {
							Type: schema.TypeString,

							Description: "SSO password.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sso_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SSO user name.",
							Optional:    true,
							Computed:    true,
						},
						"url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),

							Description: "URL parameter.",
							Optional:    true,
							Computed:    true,
						},
						"width": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(640, 65535),

							Description: "Screen width (range from 640 - 65535, default = 1024).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Group name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceVpnSslWebUserGroupBookmarkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnSslWebUserGroupBookmark resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnSslWebUserGroupBookmark(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnSslWebUserGroupBookmark(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebUserGroupBookmark")
	}

	return resourceVpnSslWebUserGroupBookmarkRead(ctx, d, meta)
}

func resourceVpnSslWebUserGroupBookmarkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslWebUserGroupBookmark(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnSslWebUserGroupBookmark(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnSslWebUserGroupBookmark resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebUserGroupBookmark")
	}

	return resourceVpnSslWebUserGroupBookmarkRead(ctx, d, meta)
}

func resourceVpnSslWebUserGroupBookmarkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnSslWebUserGroupBookmark(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnSslWebUserGroupBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserGroupBookmarkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebUserGroupBookmark(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebUserGroupBookmark resource: %v", err)
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

	diags := refreshObjectVpnSslWebUserGroupBookmark(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarks(v *[]models.VpnSslWebUserGroupBookmarkBookmarks, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AdditionalParams; tmp != nil {
				v["additional_params"] = *tmp
			}

			if tmp := cfg.Apptype; tmp != nil {
				v["apptype"] = *tmp
			}

			if tmp := cfg.ColorDepth; tmp != nil {
				v["color_depth"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Domain; tmp != nil {
				v["domain"] = *tmp
			}

			if tmp := cfg.Folder; tmp != nil {
				v["folder"] = *tmp
			}

			if tmp := cfg.FormData; tmp != nil {
				v["form_data"] = flattenVpnSslWebUserGroupBookmarkBookmarksFormData(tmp, sort)
			}

			if tmp := cfg.Height; tmp != nil {
				v["height"] = *tmp
			}

			if tmp := cfg.Host; tmp != nil {
				v["host"] = *tmp
			}

			if tmp := cfg.KeyboardLayout; tmp != nil {
				v["keyboard_layout"] = *tmp
			}

			if tmp := cfg.ListeningPort; tmp != nil {
				v["listening_port"] = *tmp
			}

			if tmp := cfg.LoadBalancingInfo; tmp != nil {
				v["load_balancing_info"] = *tmp
			}

			if tmp := cfg.LogonPassword; tmp != nil {
				v["logon_password"] = *tmp
			}

			if tmp := cfg.LogonUser; tmp != nil {
				v["logon_user"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.PreconnectionBlob; tmp != nil {
				v["preconnection_blob"] = *tmp
			}

			if tmp := cfg.PreconnectionId; tmp != nil {
				v["preconnection_id"] = *tmp
			}

			if tmp := cfg.RemotePort; tmp != nil {
				v["remote_port"] = *tmp
			}

			if tmp := cfg.RestrictedAdmin; tmp != nil {
				v["restricted_admin"] = *tmp
			}

			if tmp := cfg.Security; tmp != nil {
				v["security"] = *tmp
			}

			if tmp := cfg.SendPreconnectionId; tmp != nil {
				v["send_preconnection_id"] = *tmp
			}

			if tmp := cfg.ServerLayout; tmp != nil {
				v["server_layout"] = *tmp
			}

			if tmp := cfg.ShowStatusWindow; tmp != nil {
				v["show_status_window"] = *tmp
			}

			if tmp := cfg.Sso; tmp != nil {
				v["sso"] = *tmp
			}

			if tmp := cfg.SsoCredential; tmp != nil {
				v["sso_credential"] = *tmp
			}

			if tmp := cfg.SsoCredentialSentOnce; tmp != nil {
				v["sso_credential_sent_once"] = *tmp
			}

			if tmp := cfg.SsoPassword; tmp != nil {
				v["sso_password"] = *tmp
			}

			if tmp := cfg.SsoUsername; tmp != nil {
				v["sso_username"] = *tmp
			}

			if tmp := cfg.Url; tmp != nil {
				v["url"] = *tmp
			}

			if tmp := cfg.Width; tmp != nil {
				v["width"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormData(v *[]models.VpnSslWebUserGroupBookmarkBookmarksFormData, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData, o *models.VpnSslWebUserGroupBookmark, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Bookmarks != nil {
		if err = d.Set("bookmarks", flattenVpnSslWebUserGroupBookmarkBookmarks(o.Bookmarks, sort)); err != nil {
			return diag.Errorf("error reading bookmarks: %v", err)
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

func expandVpnSslWebUserGroupBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebUserGroupBookmarkBookmarks, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebUserGroupBookmarkBookmarks

	for i := range l {
		tmp := models.VpnSslWebUserGroupBookmarkBookmarks{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.additional_params", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalParams = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.apptype", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Apptype = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.color_depth", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ColorDepth = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Domain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.folder", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Folder = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.form_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslWebUserGroupBookmarkBookmarksFormData(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslWebUserGroupBookmarkBookmarksFormData
			// 	}
			tmp.FormData = v2

		}

		pre_append = fmt.Sprintf("%s.%d.height", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Height = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Host = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keyboard_layout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyboardLayout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.listening_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ListeningPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.load_balancing_info", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LoadBalancingInfo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.logon_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogonPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.logon_user", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogonUser = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preconnection_blob", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreconnectionBlob = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preconnection_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PreconnectionId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RemotePort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.restricted_admin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RestrictedAdmin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.security", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Security = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_preconnection_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendPreconnectionId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_layout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerLayout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.show_status_window", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShowStatusWindow = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sso = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_credential", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoCredential = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_credential_sent_once", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoCredentialSentOnce = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Url = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.width", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Width = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebUserGroupBookmarkBookmarksFormData, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebUserGroupBookmarkBookmarksFormData

	for i := range l {
		tmp := models.VpnSslWebUserGroupBookmarkBookmarksFormData{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData, sv string) (*models.VpnSslWebUserGroupBookmark, diag.Diagnostics) {
	obj := models.VpnSslWebUserGroupBookmark{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("bookmarks"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("bookmarks", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebUserGroupBookmarkBookmarks(d, v, "bookmarks", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Bookmarks = t
		}
	} else if d.HasChange("bookmarks") {
		old, new := d.GetChange("bookmarks")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Bookmarks = &[]models.VpnSslWebUserGroupBookmarkBookmarks{}
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
