// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure logical storage.

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

func resourceSystemStorage() *schema.Resource {
	return &schema.Resource{
		Description: "Configure logical storage.",

		CreateContext: resourceSystemStorageCreate,
		ReadContext:   resourceSystemStorageRead,
		UpdateContext: resourceSystemStorageUpdate,
		DeleteContext: resourceSystemStorageDelete,

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
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "Partition device.",
				Optional:    true,
				Computed:    true,
			},
			"media_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "fail"}, false),

				Description: "The physical status of current media.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Storage name.",
				ForceNew:    true,
				Required:    true,
			},
			"order": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Set storage order.",
				Optional:    true,
				Computed:    true,
			},
			"partition": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "Label of underlying partition.",
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Type: schema.TypeInt,

				Description: "Partition size.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable storage.",
				Optional:    true,
				Computed:    true,
			},
			"usage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"log", "wanopt"}, false),

				Description: "Use hard disk for logging or WAN Optimization (default = log).",
				Optional:    true,
				Computed:    true,
			},
			"wanopt_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mix", "wanopt", "webcache"}, false),

				Description: "WAN Optimization mode (default = mix).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemStorageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemStorage resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemStorage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemStorage(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStorage")
	}

	return resourceSystemStorageRead(ctx, d, meta)
}

func resourceSystemStorageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStorage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemStorage(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemStorage resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStorage")
	}

	return resourceSystemStorageRead(ctx, d, meta)
}

func resourceSystemStorageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemStorage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStorageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemStorage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStorage resource: %v", err)
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

	diags := refreshObjectSystemStorage(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemStorage(d *schema.ResourceData, o *models.SystemStorage, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.MediaStatus != nil {
		v := *o.MediaStatus

		if err = d.Set("media_status", v); err != nil {
			return diag.Errorf("error reading media_status: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Order != nil {
		v := *o.Order

		if err = d.Set("order", v); err != nil {
			return diag.Errorf("error reading order: %v", err)
		}
	}

	if o.Partition != nil {
		v := *o.Partition

		if err = d.Set("partition", v); err != nil {
			return diag.Errorf("error reading partition: %v", err)
		}
	}

	if o.Size != nil {
		v := *o.Size

		if err = d.Set("size", v); err != nil {
			return diag.Errorf("error reading size: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Usage != nil {
		v := *o.Usage

		if err = d.Set("usage", v); err != nil {
			return diag.Errorf("error reading usage: %v", err)
		}
	}

	if o.WanoptMode != nil {
		v := *o.WanoptMode

		if err = d.Set("wanopt_mode", v); err != nil {
			return diag.Errorf("error reading wanopt_mode: %v", err)
		}
	}

	return nil
}

func getObjectSystemStorage(d *schema.ResourceData, sv string) (*models.SystemStorage, diag.Diagnostics) {
	obj := models.SystemStorage{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("media_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("media_status", sv)
				diags = append(diags, e)
			}
			obj.MediaStatus = &v2
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
	if v1, ok := d.GetOk("order"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("order", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Order = &tmp
		}
	}
	if v1, ok := d.GetOk("partition"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("partition", sv)
				diags = append(diags, e)
			}
			obj.Partition = &v2
		}
	}
	if v1, ok := d.GetOk("size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Size = &tmp
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
	if v1, ok := d.GetOk("usage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("usage", sv)
				diags = append(diags, e)
			}
			obj.Usage = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt_mode", sv)
				diags = append(diags, e)
			}
			obj.WanoptMode = &v2
		}
	}
	return &obj, diags
}
