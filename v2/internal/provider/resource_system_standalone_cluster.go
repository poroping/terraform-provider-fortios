// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

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

func resourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.",

		CreateContext: resourceSystemStandaloneClusterCreate,
		ReadContext:   resourceSystemStandaloneClusterRead,
		UpdateContext: resourceSystemStandaloneClusterUpdate,
		DeleteContext: resourceSystemStandaloneClusterDelete,

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
			"encryption": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable encryption when synchronizing sessions.",
				Optional:    true,
				Computed:    true,
			},
			"group_member_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),

				Description: "Cluster member ID (0 - 15).",
				Optional:    true,
				Computed:    true,
			},
			"layer2_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"available", "unavailable"}, false),

				Description: "Indicate whether layer 2 connections are present among FGSP members.",
				Optional:    true,
				Computed:    true,
			},
			"psksecret": {
				Type: schema.TypeString,

				Description: "Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"session_sync_dev": {
				Type: schema.TypeString,

				Description: "Offload session-sync process to kernel and sync sessions using connected interface(s) directly.",
				Optional:    true,
				Computed:    true,
			},
			"standalone_group_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Cluster group ID (0 - 255). Must be the same for all members.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemStandaloneClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemStandaloneCluster(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(ctx, d, meta)
}

func resourceSystemStandaloneClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemStandaloneCluster(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemStandaloneCluster resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(ctx, d, meta)
}

func resourceSystemStandaloneClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemStandaloneCluster(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemStandaloneCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemStandaloneCluster(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStandaloneCluster resource: %v", err)
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

	diags := refreshObjectSystemStandaloneCluster(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemStandaloneCluster(d *schema.ResourceData, o *models.SystemStandaloneCluster, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Encryption != nil {
		v := *o.Encryption

		if err = d.Set("encryption", v); err != nil {
			return diag.Errorf("error reading encryption: %v", err)
		}
	}

	if o.GroupMemberId != nil {
		v := *o.GroupMemberId

		if err = d.Set("group_member_id", v); err != nil {
			return diag.Errorf("error reading group_member_id: %v", err)
		}
	}

	if o.Layer2Connection != nil {
		v := *o.Layer2Connection

		if err = d.Set("layer2_connection", v); err != nil {
			return diag.Errorf("error reading layer2_connection: %v", err)
		}
	}

	if o.Psksecret != nil {
		v := *o.Psksecret

		if v == "ENC XXXX" {
		} else if err = d.Set("psksecret", v); err != nil {
			return diag.Errorf("error reading psksecret: %v", err)
		}
	}

	if o.SessionSyncDev != nil {
		v := *o.SessionSyncDev

		if err = d.Set("session_sync_dev", v); err != nil {
			return diag.Errorf("error reading session_sync_dev: %v", err)
		}
	}

	if o.StandaloneGroupId != nil {
		v := *o.StandaloneGroupId

		if err = d.Set("standalone_group_id", v); err != nil {
			return diag.Errorf("error reading standalone_group_id: %v", err)
		}
	}

	return nil
}

func getObjectSystemStandaloneCluster(d *schema.ResourceData, sv string) (*models.SystemStandaloneCluster, diag.Diagnostics) {
	obj := models.SystemStandaloneCluster{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("encryption"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encryption", sv)
				diags = append(diags, e)
			}
			obj.Encryption = &v2
		}
	}
	if v1, ok := d.GetOk("group_member_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_member_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GroupMemberId = &tmp
		}
	}
	if v1, ok := d.GetOk("layer2_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("layer2_connection", sv)
				diags = append(diags, e)
			}
			obj.Layer2Connection = &v2
		}
	}
	if v1, ok := d.GetOk("psksecret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("psksecret", sv)
				diags = append(diags, e)
			}
			obj.Psksecret = &v2
		}
	}
	if v1, ok := d.GetOk("session_sync_dev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_sync_dev", sv)
				diags = append(diags, e)
			}
			obj.SessionSyncDev = &v2
		}
	}
	if v1, ok := d.GetOk("standalone_group_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("standalone_group_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StandaloneGroupId = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemStandaloneCluster(d *schema.ResourceData, sv string) (*models.SystemStandaloneCluster, diag.Diagnostics) {
	obj := models.SystemStandaloneCluster{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
