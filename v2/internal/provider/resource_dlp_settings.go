// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Designate logical storage for DLP fingerprint database.

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

func resourceDlpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Designate logical storage for DLP fingerprint database.",

		CreateContext: resourceDlpSettingsCreate,
		ReadContext:   resourceDlpSettingsRead,
		UpdateContext: resourceDlpSettingsUpdate,
		DeleteContext: resourceDlpSettingsDelete,

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
			"cache_mem_percent": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),

				Description: "Maximum percentage of available memory allocated to caching (1 - 15%).",
				Optional:    true,
				Computed:    true,
			},
			"chunk_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 100000),

				Description: "Maximum fingerprint chunk size.  **Changing will flush the entire database**.",
				Optional:    true,
				Computed:    true,
			},
			"db_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"stop-adding", "remove-modified-then-oldest", "remove-oldest"}, false),

				Description: "Behaviour when the maximum size is reached.",
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Type: schema.TypeInt,

				Description: "Maximum total size of files within the storage (MB).",
				Optional:    true,
				Computed:    true,
			},
			"storage_device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Storage device name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceDlpSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectDlpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateDlpSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpSettings")
	}

	return resourceDlpSettingsRead(ctx, d, meta)
}

func resourceDlpSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectDlpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateDlpSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating DlpSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpSettings")
	}

	return resourceDlpSettingsRead(ctx, d, meta)
}

func resourceDlpSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteDlpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting DlpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDlpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpSettings resource: %v", err)
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

	diags := refreshObjectDlpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectDlpSettings(d *schema.ResourceData, o *models.DlpSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CacheMemPercent != nil {
		v := *o.CacheMemPercent

		if err = d.Set("cache_mem_percent", v); err != nil {
			return diag.Errorf("error reading cache_mem_percent: %v", err)
		}
	}

	if o.ChunkSize != nil {
		v := *o.ChunkSize

		if err = d.Set("chunk_size", v); err != nil {
			return diag.Errorf("error reading chunk_size: %v", err)
		}
	}

	if o.DbMode != nil {
		v := *o.DbMode

		if err = d.Set("db_mode", v); err != nil {
			return diag.Errorf("error reading db_mode: %v", err)
		}
	}

	if o.Size != nil {
		v := *o.Size

		if err = d.Set("size", v); err != nil {
			return diag.Errorf("error reading size: %v", err)
		}
	}

	if o.StorageDevice != nil {
		v := *o.StorageDevice

		if err = d.Set("storage_device", v); err != nil {
			return diag.Errorf("error reading storage_device: %v", err)
		}
	}

	return nil
}

func getObjectDlpSettings(d *schema.ResourceData, sv string) (*models.DlpSettings, diag.Diagnostics) {
	obj := models.DlpSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cache_mem_percent"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_mem_percent", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CacheMemPercent = &tmp
		}
	}
	if v1, ok := d.GetOk("chunk_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("chunk_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ChunkSize = &tmp
		}
	}
	if v1, ok := d.GetOk("db_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("db_mode", sv)
				diags = append(diags, e)
			}
			obj.DbMode = &v2
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
	if v1, ok := d.GetOk("storage_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("storage_device", sv)
				diags = append(diags, e)
			}
			obj.StorageDevice = &v2
		}
	}
	return &obj, diags
}
