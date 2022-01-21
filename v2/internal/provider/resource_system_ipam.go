// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP address management services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemIpam() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP address management services.",

		CreateContext: resourceSystemIpamCreate,
		ReadContext:   resourceSystemIpamRead,
		UpdateContext: resourceSystemIpamUpdate,
		DeleteContext: resourceSystemIpamDelete,

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
			"pool_subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Configure IPAM pool subnet, Class A - Class B subnet.",
				Optional:    true,
				Computed:    true,
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cloud", "fabric-root"}, false),

				Description: "Configure the type of IPAM server to use.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP address management services.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemIpamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpam(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpam")
	}

	return resourceSystemIpamRead(ctx, d, meta)
}

func resourceSystemIpamUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpam(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpam resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpam")
	}

	return resourceSystemIpamRead(ctx, d, meta)
}

func resourceSystemIpamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemIpam(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpam(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpam resource: %v", err)
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

	diags := refreshObjectSystemIpam(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemIpam(d *schema.ResourceData, o *models.SystemIpam, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.PoolSubnet != nil {
		v := *o.PoolSubnet
		if current, ok := d.GetOk("pool_subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("pool_subnet", v); err != nil {
			return diag.Errorf("error reading pool_subnet: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemIpam(d *schema.ResourceData, sv string) (*models.SystemIpam, diag.Diagnostics) {
	obj := models.SystemIpam{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("pool_subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pool_subnet", sv)
				diags = append(diags, e)
			}
			obj.PoolSubnet = &v2
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
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemIpam(d *schema.ResourceData, sv string) (*models.SystemIpam, diag.Diagnostics) {
	obj := models.SystemIpam{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
