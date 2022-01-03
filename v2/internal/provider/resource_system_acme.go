// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ACME client.

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

func resourceSystemAcme() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ACME client.",

		CreateContext: resourceSystemAcmeCreate,
		ReadContext:   resourceSystemAcmeRead,
		UpdateContext: resourceSystemAcmeUpdate,
		DeleteContext: resourceSystemAcmeDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"accounts": {
				Type:        schema.TypeList,
				Description: "ACME accounts list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Account ca_url.",
							Optional:    true,
							Computed:    true,
						},
						"email": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Account email.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Account id.",
							Optional:    true,
							Computed:    true,
						},
						"privatekey": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 8191),

							Description: "Account Private Key.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Account status.",
							Optional:    true,
							Computed:    true,
						},
						"url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Account url.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "Interface(s) on which the ACME client will listen for challenges.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemAcmeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAcme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAcme(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAcme")
	}

	return resourceSystemAcmeRead(ctx, d, meta)
}

func resourceSystemAcmeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAcme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAcme(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAcme resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAcme")
	}

	return resourceSystemAcmeRead(ctx, d, meta)
}

func resourceSystemAcmeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemAcme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemAcme(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAcme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAcmeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAcme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAcme resource: %v", err)
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

	diags := refreshObjectSystemAcme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAcmeAccounts(v *[]models.SystemAcmeAccounts, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Ca_url; tmp != nil {
				v["ca_url"] = *tmp
			}

			if tmp := cfg.Email; tmp != nil {
				v["email"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Privatekey; tmp != nil {
				v["privatekey"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Url; tmp != nil {
				v["url"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemAcmeInterface(v *[]models.SystemAcmeInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func refreshObjectSystemAcme(d *schema.ResourceData, o *models.SystemAcme, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Accounts != nil {
		if err = d.Set("accounts", flattenSystemAcmeAccounts(o.Accounts, sort)); err != nil {
			return diag.Errorf("error reading accounts: %v", err)
		}
	}

	if o.Interface != nil {
		if err = d.Set("interface", flattenSystemAcmeInterface(o.Interface, sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	return nil
}

func expandSystemAcmeAccounts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAcmeAccounts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAcmeAccounts

	for i := range l {
		tmp := models.SystemAcmeAccounts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ca_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ca_url = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.email", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Email = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.privatekey", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Privatekey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Url = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAcmeInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAcmeInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAcmeInterface

	for i := range l {
		tmp := models.SystemAcmeInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemAcme(d *schema.ResourceData, sv string) (*models.SystemAcme, diag.Diagnostics) {
	obj := models.SystemAcme{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("accounts"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("accounts", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAcmeAccounts(d, v, "accounts", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Accounts = t
		}
	} else if d.HasChange("accounts") {
		old, new := d.GetChange("accounts")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Accounts = &[]models.SystemAcmeAccounts{}
		}
	}
	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAcmeInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.SystemAcmeInterface{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemAcme(d *schema.ResourceData, sv string) (*models.SystemAcme, diag.Diagnostics) {
	obj := models.SystemAcme{}
	diags := diag.Diagnostics{}

	obj.Accounts = &[]models.SystemAcmeAccounts{}
	obj.Interface = &[]models.SystemAcmeInterface{}

	return &obj, diags
}
