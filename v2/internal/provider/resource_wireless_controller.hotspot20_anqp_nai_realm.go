// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure network access identifier (NAI) realm.

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

func resourceWirelessControllerHotspot20AnqpNaiRealm() *schema.Resource {
	return &schema.Resource{
		Description: "Configure network access identifier (NAI) realm.",

		CreateContext: resourceWirelessControllerHotspot20AnqpNaiRealmCreate,
		ReadContext:   resourceWirelessControllerHotspot20AnqpNaiRealmRead,
		UpdateContext: resourceWirelessControllerHotspot20AnqpNaiRealmUpdate,
		DeleteContext: resourceWirelessControllerHotspot20AnqpNaiRealmDelete,

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
			"nai_list": {
				Type:        schema.TypeList,
				Description: "NAI list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eap_method": {
							Type:        schema.TypeList,
							Description: "EAP Methods.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_param": {
										Type:        schema.TypeList,
										Description: "EAP auth param.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"non-eap-inner-auth", "inner-auth-eap", "credential", "tunneled-credential"}, false),

													Description: "ID of authentication parameter.",
													Optional:    true,
													Computed:    true,
												},
												"index": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 4),

													Description: "Param index.",
													Optional:    true,
													Computed:    true,
												},
												"val": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"eap-identity", "eap-md5", "eap-tls", "eap-ttls", "eap-peap", "eap-sim", "eap-aka", "eap-aka-prime", "non-eap-pap", "non-eap-chap", "non-eap-mschap", "non-eap-mschapv2", "cred-sim", "cred-usim", "cred-nfc", "cred-hardware-token", "cred-softoken", "cred-certificate", "cred-user-pwd", "cred-none", "cred-vendor-specific", "tun-cred-sim", "tun-cred-usim", "tun-cred-nfc", "tun-cred-hardware-token", "tun-cred-softoken", "tun-cred-certificate", "tun-cred-user-pwd", "tun-cred-anonymous", "tun-cred-vendor-specific"}, false),

													Description: "Value of authentication parameter.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"index": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 5),

										Description: "EAP method index.",
										Optional:    true,
										Computed:    true,
									},
									"method": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"eap-identity", "eap-md5", "eap-tls", "eap-ttls", "eap-peap", "eap-sim", "eap-aka", "eap-aka-prime"}, false),

										Description: "EAP method type.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"encoding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable format in accordance with IETF RFC 4282.",
							Optional:    true,
							Computed:    true,
						},
						"nai_realm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Configure NAI realms (delimited by a semi-colon character).",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "NAI realm name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "NAI realm list name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpNaiRealmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20AnqpNaiRealm resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20AnqpNaiRealm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20AnqpNaiRealm(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpNaiRealm")
	}

	return resourceWirelessControllerHotspot20AnqpNaiRealmRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20AnqpNaiRealm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20AnqpNaiRealm(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpNaiRealm")
	}

	return resourceWirelessControllerHotspot20AnqpNaiRealmRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20AnqpNaiRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNaiRealmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20AnqpNaiRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20AnqpNaiRealm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiList(v *[]models.WirelessControllerHotspot20AnqpNaiRealmNaiList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EapMethod; tmp != nil {
				v["eap_method"] = flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(tmp, sort)
			}

			if tmp := cfg.Encoding; tmp != nil {
				v["encoding"] = *tmp
			}

			if tmp := cfg.NaiRealm; tmp != nil {
				v["nai_realm"] = *tmp
			}

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

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(v *[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthParam; tmp != nil {
				v["auth_param"] = flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(tmp, sort)
			}

			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Method; tmp != nil {
				v["method"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(v *[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Val; tmp != nil {
				v["val"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func refreshObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData, o *models.WirelessControllerHotspot20AnqpNaiRealm, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.NaiList != nil {
		if err = d.Set("nai_list", flattenWirelessControllerHotspot20AnqpNaiRealmNaiList(o.NaiList, sort)); err != nil {
			return diag.Errorf("error reading nai_list: %v", err)
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

func expandWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20AnqpNaiRealmNaiList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20AnqpNaiRealmNaiList

	for i := range l {
		tmp := models.WirelessControllerHotspot20AnqpNaiRealmNaiList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.eap_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod
			// 	}
			tmp.EapMethod = v2

		}

		pre_append = fmt.Sprintf("%s.%d.encoding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Encoding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nai_realm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NaiRealm = &v2
			}
		}

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

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod

	for i := range l {
		tmp := models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_param", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam
			// 	}
			tmp.AuthParam = v2

		}

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Index = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Method = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam

	for i := range l {
		tmp := models.WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Index = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.val", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Val = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20AnqpNaiRealm, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20AnqpNaiRealm{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("nai_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nai_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiList(d, v, "nai_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NaiList = t
		}
	} else if d.HasChange("nai_list") {
		old, new := d.GetChange("nai_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NaiList = &[]models.WirelessControllerHotspot20AnqpNaiRealmNaiList{}
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
