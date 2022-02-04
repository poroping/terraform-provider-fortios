// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VXLAN devices.

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

func resourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VXLAN devices.",

		CreateContext: resourceSystemVxlanCreate,
		ReadContext:   resourceSystemVxlanRead,
		UpdateContext: resourceSystemVxlanUpdate,
		DeleteContext: resourceSystemVxlanDelete,

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
			"dstport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "VXLAN destination port (1 - 65535, default = 4789).",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Outgoing interface for VXLAN encapsulated traffic.",
				Optional:    true,
				Computed:    true,
			},
			"ip_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4-unicast", "ipv6-unicast", "ipv4-multicast", "ipv6-multicast"}, false),

				Description: "IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "VXLAN multicast TTL (1-255, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "VXLAN device or interface name. Must be a unique interface name.",
				ForceNew:    true,
				Required:    true,
			},
			"remote_ip": {
				Type:        schema.TypeList,
				Description: "IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "IPv4 address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"remote_ip6": {
				Type:        schema.TypeList,
				Description: "IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),

							Description: "IPv6 address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vni": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777215),

				Description: "VXLAN network ID.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVxlanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemVxlan resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemVxlan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVxlan(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVxlan")
	}

	return resourceSystemVxlanRead(ctx, d, meta)
}

func resourceSystemVxlanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVxlan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVxlan(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVxlan resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVxlan")
	}

	return resourceSystemVxlanRead(ctx, d, meta)
}

func resourceSystemVxlanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemVxlan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVxlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVxlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVxlan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVxlan resource: %v", err)
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

	diags := refreshObjectSystemVxlan(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVxlanRemoteIp(v *[]models.SystemVxlanRemoteIp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func flattenSystemVxlanRemoteIp6(v *[]models.SystemVxlanRemoteIp6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Ip6; tmp != nil {
				v["ip6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip6")
	}

	return flat
}

func refreshObjectSystemVxlan(d *schema.ResourceData, o *models.SystemVxlan, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Dstport != nil {
		v := *o.Dstport

		if err = d.Set("dstport", v); err != nil {
			return diag.Errorf("error reading dstport: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpVersion != nil {
		v := *o.IpVersion

		if err = d.Set("ip_version", v); err != nil {
			return diag.Errorf("error reading ip_version: %v", err)
		}
	}

	if o.MulticastTtl != nil {
		v := *o.MulticastTtl

		if err = d.Set("multicast_ttl", v); err != nil {
			return diag.Errorf("error reading multicast_ttl: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RemoteIp != nil {
		if err = d.Set("remote_ip", flattenSystemVxlanRemoteIp(o.RemoteIp, sort)); err != nil {
			return diag.Errorf("error reading remote_ip: %v", err)
		}
	}

	if o.RemoteIp6 != nil {
		if err = d.Set("remote_ip6", flattenSystemVxlanRemoteIp6(o.RemoteIp6, sort)); err != nil {
			return diag.Errorf("error reading remote_ip6: %v", err)
		}
	}

	if o.Vni != nil {
		v := *o.Vni

		if err = d.Set("vni", v); err != nil {
			return diag.Errorf("error reading vni: %v", err)
		}
	}

	return nil
}

func expandSystemVxlanRemoteIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVxlanRemoteIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVxlanRemoteIp

	for i := range l {
		tmp := models.SystemVxlanRemoteIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVxlanRemoteIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVxlanRemoteIp6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVxlanRemoteIp6

	for i := range l {
		tmp := models.SystemVxlanRemoteIp6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ip6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemVxlan(d *schema.ResourceData, sv string) (*models.SystemVxlan, diag.Diagnostics) {
	obj := models.SystemVxlan{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("dstport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Dstport = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ip_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_version", sv)
				diags = append(diags, e)
			}
			obj.IpVersion = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MulticastTtl = &tmp
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
	if v, ok := d.GetOk("remote_ip"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("remote_ip", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVxlanRemoteIp(d, v, "remote_ip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RemoteIp = t
		}
	} else if d.HasChange("remote_ip") {
		old, new := d.GetChange("remote_ip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RemoteIp = &[]models.SystemVxlanRemoteIp{}
		}
	}
	if v, ok := d.GetOk("remote_ip6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("remote_ip6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVxlanRemoteIp6(d, v, "remote_ip6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RemoteIp6 = t
		}
	} else if d.HasChange("remote_ip6") {
		old, new := d.GetChange("remote_ip6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RemoteIp6 = &[]models.SystemVxlanRemoteIp6{}
		}
	}
	if v1, ok := d.GetOk("vni"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vni", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vni = &tmp
		}
	}
	return &obj, diags
}
