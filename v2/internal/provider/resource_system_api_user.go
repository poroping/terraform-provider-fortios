// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure API users.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemApiUser() *schema.Resource {
	return &schema.Resource{
		Description: "Configure API users.",

		CreateContext: resourceSystemApiUserCreate,
		ReadContext:   resourceSystemApiUserRead,
		UpdateContext: resourceSystemApiUserUpdate,
		DeleteContext: resourceSystemApiUserDelete,

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
			"accprofile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Admin user access profile.",
				Optional:    true,
				Computed:    true,
			},
			"api_key": {
				Type: schema.TypeString,

				Description: "Admin user password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"cors_allow_origin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 269),

				Description: "Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User name.",
				ForceNew:    true,
				Required:    true,
			},
			"peer_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable peer authentication.",
				Optional:    true,
				Computed:    true,
			},
			"peer_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Peer group name.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
				Optional:    true,
				Computed:    true,
			},
			"trusthost": {
				Type:        schema.TypeList,
				Description: "Trusthost.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_trusthost": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "IPv4 trusted host address.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_trusthost": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "IPv6 trusted host address.",
							Optional:         true,
							Computed:         true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4-trusthost", "ipv6-trusthost"}, false),

							Description: "Trusthost type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domains.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Virtual domain name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemApiUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemApiUser resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemApiUser(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemApiUser(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemApiUser")
	}

	return resourceSystemApiUserRead(ctx, d, meta)
}

func resourceSystemApiUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemApiUser(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemApiUser(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemApiUser resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemApiUser")
	}

	return resourceSystemApiUserRead(ctx, d, meta)
}

func resourceSystemApiUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemApiUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemApiUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemApiUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemApiUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemApiUser resource: %v", err)
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

	diags := refreshObjectSystemApiUser(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemApiUserTrusthost(d *schema.ResourceData, v *[]models.SystemApiUserTrusthost, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ipv4Trusthost; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.ipv4_trusthost", prefix, i), *tmp); tmp != nil {
					v["ipv4_trusthost"] = *tmp
				}
			}

			if tmp := cfg.Ipv6Trusthost; tmp != nil {
				v["ipv6_trusthost"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemApiUserVdom(d *schema.ResourceData, v *[]models.SystemApiUserVdom, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemApiUser(d *schema.ResourceData, o *models.SystemApiUser, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Accprofile != nil {
		v := *o.Accprofile

		if err = d.Set("accprofile", v); err != nil {
			return diag.Errorf("error reading accprofile: %v", err)
		}
	}

	if o.ApiKey != nil {
		v := *o.ApiKey

		if v == "ENC XXXX" {
		} else if err = d.Set("api_key", v); err != nil {
			return diag.Errorf("error reading api_key: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.CorsAllowOrigin != nil {
		v := *o.CorsAllowOrigin

		if err = d.Set("cors_allow_origin", v); err != nil {
			return diag.Errorf("error reading cors_allow_origin: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PeerAuth != nil {
		v := *o.PeerAuth

		if err = d.Set("peer_auth", v); err != nil {
			return diag.Errorf("error reading peer_auth: %v", err)
		}
	}

	if o.PeerGroup != nil {
		v := *o.PeerGroup

		if err = d.Set("peer_group", v); err != nil {
			return diag.Errorf("error reading peer_group: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.Trusthost != nil {
		if err = d.Set("trusthost", flattenSystemApiUserTrusthost(d, o.Trusthost, "trusthost", sort)); err != nil {
			return diag.Errorf("error reading trusthost: %v", err)
		}
	}

	if o.Vdom != nil {
		if err = d.Set("vdom", flattenSystemApiUserVdom(d, o.Vdom, "vdom", sort)); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func expandSystemApiUserTrusthost(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemApiUserTrusthost, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemApiUserTrusthost

	for i := range l {
		tmp := models.SystemApiUserTrusthost{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_trusthost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Trusthost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_trusthost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Trusthost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemApiUserVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemApiUserVdom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemApiUserVdom

	for i := range l {
		tmp := models.SystemApiUserVdom{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemApiUser(d *schema.ResourceData, sv string) (*models.SystemApiUser, diag.Diagnostics) {
	obj := models.SystemApiUser{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("accprofile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("accprofile", sv)
				diags = append(diags, e)
			}
			obj.Accprofile = &v2
		}
	}
	if v1, ok := d.GetOk("api_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("api_key", sv)
				diags = append(diags, e)
			}
			obj.ApiKey = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("cors_allow_origin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cors_allow_origin", sv)
				diags = append(diags, e)
			}
			obj.CorsAllowOrigin = &v2
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
	if v1, ok := d.GetOk("peer_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_auth", sv)
				diags = append(diags, e)
			}
			obj.PeerAuth = &v2
		}
	}
	if v1, ok := d.GetOk("peer_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_group", sv)
				diags = append(diags, e)
			}
			obj.PeerGroup = &v2
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v, ok := d.GetOk("trusthost"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("trusthost", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemApiUserTrusthost(d, v, "trusthost", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Trusthost = t
		}
	} else if d.HasChange("trusthost") {
		old, new := d.GetChange("trusthost")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Trusthost = &[]models.SystemApiUserTrusthost{}
		}
	}
	if v, ok := d.GetOk("vdom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vdom", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemApiUserVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdom = t
		}
	} else if d.HasChange("vdom") {
		old, new := d.GetChange("vdom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdom = &[]models.SystemApiUserVdom{}
		}
	}
	return &obj, diags
}
