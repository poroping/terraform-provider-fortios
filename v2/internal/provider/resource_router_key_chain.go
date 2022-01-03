// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure key-chain.

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

func resourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Description: "Configure key-chain.",

		CreateContext: resourceRouterKeyChainCreate,
		ReadContext:   resourceRouterKeyChainRead,
		UpdateContext: resourceRouterKeyChainUpdate,
		DeleteContext: resourceRouterKeyChainDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit key settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_lifetime": {
							Type: schema.TypeString,

							Description: "Lifetime of received authentication key (format: hh:mm:ss day month year).",
							Optional:    true,
							Computed:    true,
						},
						"algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "hmac-sha1", "hmac-sha256", "hmac-sha384", "hmac-sha512"}, false),

							Description: "Cryptographic algorithm.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 10),

							Description: "Key ID (0 - 2147483647).",
							Optional:    true,
							Computed:    true,
						},
						"key_string": {
							Type: schema.TypeString,

							Description: "Password for the key (max. = 64 characters).",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"send_lifetime": {
							Type: schema.TypeString,

							Description: "Lifetime of sent authentication key (format: hh:mm:ss day month year).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Key-chain name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceRouterKeyChainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterKeyChain resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterKeyChain(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterKeyChain(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(ctx, d, meta)
}

func resourceRouterKeyChainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterKeyChain(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterKeyChain(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterKeyChain resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(ctx, d, meta)
}

func resourceRouterKeyChainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterKeyChain(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterKeyChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterKeyChainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterKeyChain(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterKeyChain resource: %v", err)
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

	diags := refreshObjectRouterKeyChain(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterKeyChainKey(v *[]models.RouterKeyChainKey, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AcceptLifetime; tmp != nil {
				v["accept_lifetime"] = *tmp
			}

			if tmp := cfg.Algorithm; tmp != nil {
				v["algorithm"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.KeyString; tmp != nil {
				v["key_string"] = *tmp
			}

			if tmp := cfg.SendLifetime; tmp != nil {
				v["send_lifetime"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterKeyChain(d *schema.ResourceData, o *models.RouterKeyChain, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Key != nil {
		if err = d.Set("key", flattenRouterKeyChainKey(o.Key, sort)); err != nil {
			return diag.Errorf("error reading key: %v", err)
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

func expandRouterKeyChainKey(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterKeyChainKey, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterKeyChainKey

	for i := range l {
		tmp := models.RouterKeyChainKey{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.accept_lifetime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AcceptLifetime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.algorithm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Algorithm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_lifetime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendLifetime = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterKeyChain(d *schema.ResourceData, sv string) (*models.RouterKeyChain, diag.Diagnostics) {
	obj := models.RouterKeyChain{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("key"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("key", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterKeyChainKey(d, v, "key", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Key = t
		}
	} else if d.HasChange("key") {
		old, new := d.GetChange("key")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Key = &[]models.RouterKeyChainKey{}
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
