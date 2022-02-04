// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WCCP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemWccp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WCCP.",

		CreateContext: resourceSystemWccpCreate,
		ReadContext:   resourceSystemWccpRead,
		UpdateContext: resourceSystemWccpUpdate,
		DeleteContext: resourceSystemWccpDelete,

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
			"assignment_bucket_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wccp-v2", "cisco-implementation"}, false),

				Description: "Assignment bucket format for the WCCP cache engine.",
				Optional:    true,
				Computed:    true,
			},
			"assignment_dstaddr_mask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "Assignment destination address mask.",
				Optional:    true,
				Computed:    true,
			},
			"assignment_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"HASH", "MASK", "any"}, false),

				Description: "Hash key assignment preference.",
				Optional:    true,
				Computed:    true,
			},
			"assignment_srcaddr_mask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "Assignment source address mask.",
				Optional:    true,
				Computed:    true,
			},
			"assignment_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Assignment of hash weight/ratio for the WCCP cache engine.",
				Optional:    true,
				Computed:    true,
			},
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MD5 authentication.",
				Optional:    true,
				Computed:    true,
			},
			"cache_engine_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"GRE", "L2"}, false),

				Description: "Method used to forward traffic to the routers or to return to the cache engine.",
				Optional:    true,
				Computed:    true,
			},
			"cache_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.",
				Optional:    true,
				Computed:    true,
			},
			"forward_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"GRE", "L2", "any"}, false),

				Description: "Method used to forward traffic to the cache servers.",
				Optional:    true,
				Computed:    true,
			},
			"group_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password for MD5 authentication.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ports": {
				Type: schema.TypeString,

				Description: "Service ports.",
				Optional:    true,
				Computed:    true,
			},
			"ports_defined": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"source", "destination"}, false),

				Description: "Match method.",
				Optional:    true,
				Computed:    true,
			},
			"primary_hash": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Hash method.",
				Optional:         true,
				Computed:         true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Service priority.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Service protocol.",
				Optional:    true,
				Computed:    true,
			},
			"return_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"GRE", "L2", "any"}, false),

				Description: "Method used to decline a redirected packet and return it to the FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"router_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.",
				Optional:    true,
				Computed:    true,
			},
			"router_list": {
				Type: schema.TypeString,

				Description: "IP addresses of one or more WCCP routers.",
				Optional:    true,
				Computed:    true,
			},
			"server_list": {
				Type: schema.TypeString,

				Description: "IP addresses and netmasks for up to four cache servers.",
				Optional:    true,
				Computed:    true,
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"forward", "proxy"}, false),

				Description: "Cache server type.",
				Optional:    true,
				Computed:    true,
			},
			"service_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 3),

				Description: "Service ID.",
				ForceNew:    true,
				Required:    true,
			},
			"service_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "standard", "dynamic"}, false),

				Description: "WCCP service type used by the cache server for logical interception and redirection of traffic.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemWccpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "service_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemWccp resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemWccp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemWccp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemWccp")
	}

	return resourceSystemWccpRead(ctx, d, meta)
}

func resourceSystemWccpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemWccp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemWccp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemWccp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemWccp")
	}

	return resourceSystemWccpRead(ctx, d, meta)
}

func resourceSystemWccpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemWccp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemWccp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWccpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemWccp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemWccp resource: %v", err)
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

	diags := refreshObjectSystemWccp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemWccp(d *schema.ResourceData, o *models.SystemWccp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AssignmentBucketFormat != nil {
		v := *o.AssignmentBucketFormat

		if err = d.Set("assignment_bucket_format", v); err != nil {
			return diag.Errorf("error reading assignment_bucket_format: %v", err)
		}
	}

	if o.AssignmentDstaddrMask != nil {
		v := *o.AssignmentDstaddrMask

		if err = d.Set("assignment_dstaddr_mask", v); err != nil {
			return diag.Errorf("error reading assignment_dstaddr_mask: %v", err)
		}
	}

	if o.AssignmentMethod != nil {
		v := *o.AssignmentMethod

		if err = d.Set("assignment_method", v); err != nil {
			return diag.Errorf("error reading assignment_method: %v", err)
		}
	}

	if o.AssignmentSrcaddrMask != nil {
		v := *o.AssignmentSrcaddrMask

		if err = d.Set("assignment_srcaddr_mask", v); err != nil {
			return diag.Errorf("error reading assignment_srcaddr_mask: %v", err)
		}
	}

	if o.AssignmentWeight != nil {
		v := *o.AssignmentWeight

		if err = d.Set("assignment_weight", v); err != nil {
			return diag.Errorf("error reading assignment_weight: %v", err)
		}
	}

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.CacheEngineMethod != nil {
		v := *o.CacheEngineMethod

		if err = d.Set("cache_engine_method", v); err != nil {
			return diag.Errorf("error reading cache_engine_method: %v", err)
		}
	}

	if o.CacheId != nil {
		v := *o.CacheId

		if err = d.Set("cache_id", v); err != nil {
			return diag.Errorf("error reading cache_id: %v", err)
		}
	}

	if o.ForwardMethod != nil {
		v := *o.ForwardMethod

		if err = d.Set("forward_method", v); err != nil {
			return diag.Errorf("error reading forward_method: %v", err)
		}
	}

	if o.GroupAddress != nil {
		v := *o.GroupAddress

		if err = d.Set("group_address", v); err != nil {
			return diag.Errorf("error reading group_address: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Ports != nil {
		v := *o.Ports

		if err = d.Set("ports", v); err != nil {
			return diag.Errorf("error reading ports: %v", err)
		}
	}

	if o.PortsDefined != nil {
		v := *o.PortsDefined

		if err = d.Set("ports_defined", v); err != nil {
			return diag.Errorf("error reading ports_defined: %v", err)
		}
	}

	if o.PrimaryHash != nil {
		v := *o.PrimaryHash

		if err = d.Set("primary_hash", v); err != nil {
			return diag.Errorf("error reading primary_hash: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.ReturnMethod != nil {
		v := *o.ReturnMethod

		if err = d.Set("return_method", v); err != nil {
			return diag.Errorf("error reading return_method: %v", err)
		}
	}

	if o.RouterId != nil {
		v := *o.RouterId

		if err = d.Set("router_id", v); err != nil {
			return diag.Errorf("error reading router_id: %v", err)
		}
	}

	if o.RouterList != nil {
		v := *o.RouterList

		if err = d.Set("router_list", v); err != nil {
			return diag.Errorf("error reading router_list: %v", err)
		}
	}

	if o.ServerList != nil {
		v := *o.ServerList

		if err = d.Set("server_list", v); err != nil {
			return diag.Errorf("error reading server_list: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.ServiceId != nil {
		v := *o.ServiceId

		if err = d.Set("service_id", v); err != nil {
			return diag.Errorf("error reading service_id: %v", err)
		}
	}

	if o.ServiceType != nil {
		v := *o.ServiceType

		if err = d.Set("service_type", v); err != nil {
			return diag.Errorf("error reading service_type: %v", err)
		}
	}

	return nil
}

func getObjectSystemWccp(d *schema.ResourceData, sv string) (*models.SystemWccp, diag.Diagnostics) {
	obj := models.SystemWccp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("assignment_bucket_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assignment_bucket_format", sv)
				diags = append(diags, e)
			}
			obj.AssignmentBucketFormat = &v2
		}
	}
	if v1, ok := d.GetOk("assignment_dstaddr_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assignment_dstaddr_mask", sv)
				diags = append(diags, e)
			}
			obj.AssignmentDstaddrMask = &v2
		}
	}
	if v1, ok := d.GetOk("assignment_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assignment_method", sv)
				diags = append(diags, e)
			}
			obj.AssignmentMethod = &v2
		}
	}
	if v1, ok := d.GetOk("assignment_srcaddr_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assignment_srcaddr_mask", sv)
				diags = append(diags, e)
			}
			obj.AssignmentSrcaddrMask = &v2
		}
	}
	if v1, ok := d.GetOk("assignment_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assignment_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AssignmentWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v1, ok := d.GetOk("cache_engine_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_engine_method", sv)
				diags = append(diags, e)
			}
			obj.CacheEngineMethod = &v2
		}
	}
	if v1, ok := d.GetOk("cache_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_id", sv)
				diags = append(diags, e)
			}
			obj.CacheId = &v2
		}
	}
	if v1, ok := d.GetOk("forward_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_method", sv)
				diags = append(diags, e)
			}
			obj.ForwardMethod = &v2
		}
	}
	if v1, ok := d.GetOk("group_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_address", sv)
				diags = append(diags, e)
			}
			obj.GroupAddress = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("ports"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ports", sv)
				diags = append(diags, e)
			}
			obj.Ports = &v2
		}
	}
	if v1, ok := d.GetOk("ports_defined"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ports_defined", sv)
				diags = append(diags, e)
			}
			obj.PortsDefined = &v2
		}
	}
	if v1, ok := d.GetOk("primary_hash"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("primary_hash", sv)
				diags = append(diags, e)
			}
			obj.PrimaryHash = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Protocol = &tmp
		}
	}
	if v1, ok := d.GetOk("return_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("return_method", sv)
				diags = append(diags, e)
			}
			obj.ReturnMethod = &v2
		}
	}
	if v1, ok := d.GetOk("router_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("router_id", sv)
				diags = append(diags, e)
			}
			obj.RouterId = &v2
		}
	}
	if v1, ok := d.GetOk("router_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("router_list", sv)
				diags = append(diags, e)
			}
			obj.RouterList = &v2
		}
	}
	if v1, ok := d.GetOk("server_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_list", sv)
				diags = append(diags, e)
			}
			obj.ServerList = &v2
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
		}
	}
	if v1, ok := d.GetOk("service_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_id", sv)
				diags = append(diags, e)
			}
			obj.ServiceId = &v2
		}
	}
	if v1, ok := d.GetOk("service_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_type", sv)
				diags = append(diags, e)
			}
			obj.ServiceType = &v2
		}
	}
	return &obj, diags
}
