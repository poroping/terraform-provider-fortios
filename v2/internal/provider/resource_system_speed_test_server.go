// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure speed test server list.

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

func resourceSystemSpeedTestServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure speed test server list.",

		CreateContext: resourceSystemSpeedTestServerCreate,
		ReadContext:   resourceSystemSpeedTestServerRead,
		UpdateContext: resourceSystemSpeedTestServerUpdate,
		DeleteContext: resourceSystemSpeedTestServerDelete,

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
			"host": {
				Type:        schema.TypeList,
				Description: "Hosts of the server.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Server host ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Server host IPv4 address.",
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Type: schema.TypeString,

							Description: "Speed test host password.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Server host port number to communicate with client.",
							Optional:    true,
							Computed:    true,
						},
						"user": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Speed test host user name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Speed test server name.",
				ForceNew:    true,
				Required:    true,
			},
			"timestamp": {
				Type: schema.TypeInt,

				Description: "Speed test server timestamp.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSpeedTestServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSpeedTestServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSpeedTestServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSpeedTestServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSpeedTestServer")
	}

	return resourceSystemSpeedTestServerRead(ctx, d, meta)
}

func resourceSystemSpeedTestServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSpeedTestServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSpeedTestServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSpeedTestServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSpeedTestServer")
	}

	return resourceSystemSpeedTestServerRead(ctx, d, meta)
}

func resourceSystemSpeedTestServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSpeedTestServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSpeedTestServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSpeedTestServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSpeedTestServer resource: %v", err)
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

	diags := refreshObjectSystemSpeedTestServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSpeedTestServerHost(d *schema.ResourceData, v *[]models.SystemSpeedTestServerHost, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Password; tmp != nil {
				v["password"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.User; tmp != nil {
				v["user"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemSpeedTestServer(d *schema.ResourceData, o *models.SystemSpeedTestServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Host != nil {
		if err = d.Set("host", flattenSystemSpeedTestServerHost(d, o.Host, "host", sort)); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Timestamp != nil {
		v := *o.Timestamp

		if err = d.Set("timestamp", v); err != nil {
			return diag.Errorf("error reading timestamp: %v", err)
		}
	}

	return nil
}

func expandSystemSpeedTestServerHost(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSpeedTestServerHost, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSpeedTestServerHost

	for i := range l {
		tmp := models.SystemSpeedTestServerHost{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Password = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.user", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.User = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSpeedTestServer(d *schema.ResourceData, sv string) (*models.SystemSpeedTestServer, diag.Diagnostics) {
	obj := models.SystemSpeedTestServer{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("host"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("host", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSpeedTestServerHost(d, v, "host", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Host = t
		}
	} else if d.HasChange("host") {
		old, new := d.GetChange("host")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Host = &[]models.SystemSpeedTestServerHost{}
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
	if v1, ok := d.GetOk("timestamp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timestamp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timestamp = &tmp
		}
	}
	return &obj, diags
}
