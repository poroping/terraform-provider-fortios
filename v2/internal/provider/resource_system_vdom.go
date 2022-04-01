// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual domain.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual domain.",

		CreateContext: resourceSystemVdomCreate,
		ReadContext:   resourceSystemVdomRead,
		UpdateContext: resourceSystemVdomUpdate,
		DeleteContext: resourceSystemVdomDelete,

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
			"flag": {
				Type: schema.TypeInt,

				Description: "Flag.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "VDOM name.",
				ForceNew:    true,
				Required:    true,
			},
			"short_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),

				Description: "VDOM short name.",
				Optional:    true,
				Computed:    true,
			},
			"vcluster_id": {
				Type: schema.TypeInt,

				Description: "Virtual cluster ID (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVdomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemVdom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemVdom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVdom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(ctx, d, meta)
}

func resourceSystemVdomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVdom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVdom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(ctx, d, meta)
}

func resourceSystemVdomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemVdom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdom resource: %v", err)
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

	diags := refreshObjectSystemVdom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemVdom(d *schema.ResourceData, o *models.SystemVdom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Flag != nil {
		v := *o.Flag

		if err = d.Set("flag", v); err != nil {
			return diag.Errorf("error reading flag: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ShortName != nil {
		v := *o.ShortName

		if err = d.Set("short_name", v); err != nil {
			return diag.Errorf("error reading short_name: %v", err)
		}
	}

	if o.VclusterId != nil {
		v := *o.VclusterId

		if err = d.Set("vcluster_id", v); err != nil {
			return diag.Errorf("error reading vcluster_id: %v", err)
		}
	}

	return nil
}

func getObjectSystemVdom(d *schema.ResourceData, sv string) (*models.SystemVdom, diag.Diagnostics) {
	obj := models.SystemVdom{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("flag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("flag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Flag = &tmp
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
	if v1, ok := d.GetOk("short_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("short_name", sv)
				diags = append(diags, e)
			}
			obj.ShortName = &v2
		}
	}
	if v1, ok := d.GetOk("vcluster_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vcluster_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VclusterId = &tmp
		}
	}
	return &obj, diags
}
