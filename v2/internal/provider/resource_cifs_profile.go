// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure CIFS profile.

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

func resourceCifsProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure CIFS profile.",

		CreateContext: resourceCifsProfileCreate,
		ReadContext:   resourceCifsProfileRead,
		UpdateContext: resourceCifsProfileUpdate,
		DeleteContext: resourceCifsProfileDelete,

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
			"domain_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain for which to decrypt CIFS traffic.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter": {
				Type:        schema.TypeList,
				Description: "File filter.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries": {
							Type:        schema.TypeList,
							Description: "File filter entries.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"log", "block"}, false),

										Description: "Action taken for matched file.",
										Optional:    true,
										Computed:    true,
									},
									"comment": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "Comment.",
										Optional:    true,
										Computed:    true,
									},
									"direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"incoming", "outgoing", "any"}, false),

										Description: "Match files transmitted in the session's originating or reply direction.",
										Optional:    true,
										Computed:    true,
									},
									"file_type": {
										Type:        schema.TypeList,
										Description: "Select file type.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 39),

													Description: "File type name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"filter": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Add a file filter.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable file filter logging.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable file filter.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"server_credential_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "credential-replication", "credential-keytab"}, false),

				Description: "CIFS server credential type.",
				Optional:    true,
				Computed:    true,
			},
			"server_keytab": {
				Type:        schema.TypeList,
				Description: "Server keytab.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keytab": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 8191),

							Description: "Base64 encoded keytab file containing credential of the server.",
							Optional:    true,
							Computed:    true,
						},
						"principal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Service principal. For example, host/cifsserver.example.com@example.com.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceCifsProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating CifsProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectCifsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateCifsProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CifsProfile")
	}

	return resourceCifsProfileRead(ctx, d, meta)
}

func resourceCifsProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectCifsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateCifsProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating CifsProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CifsProfile")
	}

	return resourceCifsProfileRead(ctx, d, meta)
}

func resourceCifsProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteCifsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting CifsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCifsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCifsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CifsProfile resource: %v", err)
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

	diags := refreshObjectCifsProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenCifsProfileFileFilter(d *schema.ResourceData, v *models.CifsProfileFileFilter, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.CifsProfileFileFilter{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Entries; tmp != nil {
				v["entries"] = flattenCifsProfileFileFilterEntries(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "entries"), sort)
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenCifsProfileFileFilterEntries(d *schema.ResourceData, v *[]models.CifsProfileFileFilterEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.FileType; tmp != nil {
				v["file_type"] = flattenCifsProfileFileFilterEntriesFileType(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "file_type"), sort)
			}

			if tmp := cfg.Filter; tmp != nil {
				v["filter"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "filter")
	}

	return flat
}

func flattenCifsProfileFileFilterEntriesFileType(d *schema.ResourceData, v *[]models.CifsProfileFileFilterEntriesFileType, prefix string, sort bool) interface{} {
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

func flattenCifsProfileServerKeytab(d *schema.ResourceData, v *[]models.CifsProfileServerKeytab, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Keytab; tmp != nil {
				v["keytab"] = *tmp
			}

			if tmp := cfg.Principal; tmp != nil {
				v["principal"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "principal")
	}

	return flat
}

func refreshObjectCifsProfile(d *schema.ResourceData, o *models.CifsProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DomainController != nil {
		v := *o.DomainController

		if err = d.Set("domain_controller", v); err != nil {
			return diag.Errorf("error reading domain_controller: %v", err)
		}
	}

	if _, ok := d.GetOk("file_filter"); ok {
		if o.FileFilter != nil {
			if err = d.Set("file_filter", flattenCifsProfileFileFilter(d, o.FileFilter, "file_filter", sort)); err != nil {
				return diag.Errorf("error reading file_filter: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ServerCredentialType != nil {
		v := *o.ServerCredentialType

		if err = d.Set("server_credential_type", v); err != nil {
			return diag.Errorf("error reading server_credential_type: %v", err)
		}
	}

	if o.ServerKeytab != nil {
		if err = d.Set("server_keytab", flattenCifsProfileServerKeytab(d, o.ServerKeytab, "server_keytab", sort)); err != nil {
			return diag.Errorf("error reading server_keytab: %v", err)
		}
	}

	return nil
}

func expandCifsProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.CifsProfileFileFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.CifsProfileFileFilter

	for i := range l {
		tmp := models.CifsProfileFileFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandCifsProfileFileFilterEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.CifsProfileFileFilterEntries
			// 	}
			tmp.Entries = v2

		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandCifsProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.CifsProfileFileFilterEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.CifsProfileFileFilterEntries

	for i := range l {
		tmp := models.CifsProfileFileFilterEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandCifsProfileFileFilterEntriesFileType(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.CifsProfileFileFilterEntriesFileType
			// 	}
			tmp.FileType = v2

		}

		pre_append = fmt.Sprintf("%s.%d.filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Filter = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandCifsProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.CifsProfileFileFilterEntriesFileType, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.CifsProfileFileFilterEntriesFileType

	for i := range l {
		tmp := models.CifsProfileFileFilterEntriesFileType{}
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

func expandCifsProfileServerKeytab(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.CifsProfileServerKeytab, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.CifsProfileServerKeytab

	for i := range l {
		tmp := models.CifsProfileServerKeytab{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keytab", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Keytab = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.principal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Principal = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectCifsProfile(d *schema.ResourceData, sv string) (*models.CifsProfile, diag.Diagnostics) {
	obj := models.CifsProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("domain_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain_controller", sv)
				diags = append(diags, e)
			}
			obj.DomainController = &v2
		}
	}
	if v, ok := d.GetOk("file_filter"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("file_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandCifsProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FileFilter = t
		}
	} else if d.HasChange("file_filter") {
		old, new := d.GetChange("file_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FileFilter = &models.CifsProfileFileFilter{}
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
	if v1, ok := d.GetOk("server_credential_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_credential_type", sv)
				diags = append(diags, e)
			}
			obj.ServerCredentialType = &v2
		}
	}
	if v, ok := d.GetOk("server_keytab"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server_keytab", sv)
			diags = append(diags, e)
		}
		t, err := expandCifsProfileServerKeytab(d, v, "server_keytab", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerKeytab = t
		}
	} else if d.HasChange("server_keytab") {
		old, new := d.GetChange("server_keytab")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerKeytab = &[]models.CifsProfileServerKeytab{}
		}
	}
	return &obj, diags
}
