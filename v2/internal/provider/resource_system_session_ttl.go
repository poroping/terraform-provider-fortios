// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure global session TTL timers for this FortiGate.

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

func resourceSystemSessionTtl() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global session TTL timers for this FortiGate.",

		CreateContext: resourceSystemSessionTtlCreate,
		ReadContext:   resourceSystemSessionTtlRead,
		UpdateContext: resourceSystemSessionTtlUpdate,
		DeleteContext: resourceSystemSessionTtlDelete,

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
			"default": {
				Type: schema.TypeString,

				Description: "Default timeout.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeList,
				Description: "Session TTL port.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "End port number.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Table entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Protocol (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"start_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Start port number.",
							Optional:    true,
							Computed:    true,
						},
						"timeout": {
							Type: schema.TypeString,

							Description: "Session timeout (TTL).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemSessionTtlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSessionTtl(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSessionTtl(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSessionTtl")
	}

	return resourceSystemSessionTtlRead(ctx, d, meta)
}

func resourceSystemSessionTtlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSessionTtl(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSessionTtl(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSessionTtl resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSessionTtl")
	}

	return resourceSystemSessionTtlRead(ctx, d, meta)
}

func resourceSystemSessionTtlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSessionTtl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSessionTtl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSessionTtlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSessionTtl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSessionTtl resource: %v", err)
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

	diags := refreshObjectSystemSessionTtl(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSessionTtlPort(v *[]models.SystemSessionTtlPort, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EndPort; tmp != nil {
				v["end_port"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.StartPort; tmp != nil {
				v["start_port"] = *tmp
			}

			if tmp := cfg.Timeout; tmp != nil {
				v["timeout"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemSessionTtl(d *schema.ResourceData, o *models.SystemSessionTtl, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Default != nil {
		v := *o.Default

		if err = d.Set("default", v); err != nil {
			return diag.Errorf("error reading default: %v", err)
		}
	}

	if o.Port != nil {
		if err = d.Set("port", flattenSystemSessionTtlPort(o.Port, sort)); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	return nil
}

func expandSystemSessionTtlPort(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSessionTtlPort, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSessionTtlPort

	for i := range l {
		tmp := models.SystemSessionTtlPort{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.EndPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StartPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Timeout = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSessionTtl(d *schema.ResourceData, sv string) (*models.SystemSessionTtl, diag.Diagnostics) {
	obj := models.SystemSessionTtl{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("default"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default", sv)
				diags = append(diags, e)
			}
			obj.Default = &v2
		}
	}
	if v, ok := d.GetOk("port"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("port", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSessionTtlPort(d, v, "port", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Port = t
		}
	} else if d.HasChange("port") {
		old, new := d.GetChange("port")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Port = &[]models.SystemSessionTtlPort{}
		}
	}
	return &obj, diags
}
