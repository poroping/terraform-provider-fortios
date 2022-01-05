// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure a remote cache device as Web cache storage.

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

func resourceWanoptRemoteStorage() *schema.Resource {
	return &schema.Resource{
		Description: "Configure a remote cache device as Web cache storage.",

		CreateContext: resourceWanoptRemoteStorageCreate,
		ReadContext:   resourceWanoptRemoteStorageRead,
		UpdateContext: resourceWanoptRemoteStorageUpdate,
		DeleteContext: resourceWanoptRemoteStorageDelete,

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
			"local_cache_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ID that this device uses to connect to the remote device.",
				Optional:    true,
				Computed:    true,
			},
			"remote_cache_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ID of the remote device to which the device connects.",
				Optional:    true,
				Computed:    true,
			},
			"remote_cache_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the remote device to which the device connects.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using remote device as Web cache storage.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWanoptRemoteStorageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptRemoteStorage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptRemoteStorage(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptRemoteStorage")
	}

	return resourceWanoptRemoteStorageRead(ctx, d, meta)
}

func resourceWanoptRemoteStorageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptRemoteStorage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptRemoteStorage(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptRemoteStorage resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptRemoteStorage")
	}

	return resourceWanoptRemoteStorageRead(ctx, d, meta)
}

func resourceWanoptRemoteStorageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWanoptRemoteStorage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWanoptRemoteStorage(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptRemoteStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptRemoteStorageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptRemoteStorage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptRemoteStorage resource: %v", err)
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

	diags := refreshObjectWanoptRemoteStorage(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWanoptRemoteStorage(d *schema.ResourceData, o *models.WanoptRemoteStorage, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.LocalCacheId != nil {
		v := *o.LocalCacheId

		if err = d.Set("local_cache_id", v); err != nil {
			return diag.Errorf("error reading local_cache_id: %v", err)
		}
	}

	if o.RemoteCacheId != nil {
		v := *o.RemoteCacheId

		if err = d.Set("remote_cache_id", v); err != nil {
			return diag.Errorf("error reading remote_cache_id: %v", err)
		}
	}

	if o.RemoteCacheIp != nil {
		v := *o.RemoteCacheIp

		if err = d.Set("remote_cache_ip", v); err != nil {
			return diag.Errorf("error reading remote_cache_ip: %v", err)
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

func getObjectWanoptRemoteStorage(d *schema.ResourceData, sv string) (*models.WanoptRemoteStorage, diag.Diagnostics) {
	obj := models.WanoptRemoteStorage{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("local_cache_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_cache_id", sv)
				diags = append(diags, e)
			}
			obj.LocalCacheId = &v2
		}
	}
	if v1, ok := d.GetOk("remote_cache_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_cache_id", sv)
				diags = append(diags, e)
			}
			obj.RemoteCacheId = &v2
		}
	}
	if v1, ok := d.GetOk("remote_cache_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_cache_ip", sv)
				diags = append(diags, e)
			}
			obj.RemoteCacheIp = &v2
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
func getEmptyObjectWanoptRemoteStorage(d *schema.ResourceData, sv string) (*models.WanoptRemoteStorage, diag.Diagnostics) {
	obj := models.WanoptRemoteStorage{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
