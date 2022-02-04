// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

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

func resourceDlpFpDocSource() *schema.Resource {
	return &schema.Resource{
		Description: "Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.",

		CreateContext: resourceDlpFpDocSourceCreate,
		ReadContext:   resourceDlpFpDocSourceRead,
		UpdateContext: resourceDlpFpDocSourceUpdate,
		DeleteContext: resourceDlpFpDocSourceDelete,

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
			"date": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),

				Description: "Day of the month on which to scan the server (1 - 31).",
				Optional:    true,
				Computed:    true,
			},
			"file_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 119),

				Description: "Path on the server to the fingerprint files (max 119 characters).",
				Optional:    true,
				Computed:    true,
			},
			"file_pattern": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.",
				Optional:    true,
				Computed:    true,
			},
			"keep_modified": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the DLP fingerprint database.",
				ForceNew:    true,
				Required:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password required to log into the file server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"period": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "daily", "weekly", "monthly"}, false),

				Description: "Frequency for which the FortiGate checks the server for new or changed files.",
				Optional:    true,
				Computed:    true,
			},
			"remove_deleted": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to keep the fingerprint database up to date when a file is deleted from the server.",
				Optional:    true,
				Computed:    true,
			},
			"scan_on_creation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to keep the fingerprint database up to date when a file is added or changed on the server.",
				Optional:    true,
				Computed:    true,
			},
			"scan_subdirectories": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable scanning subdirectories to find files to create fingerprints from.",
				Optional:    true,
				Computed:    true,
			},
			"sensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using sensitivity.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IPv4 or IPv6 address of the server.",
				Optional:    true,
				Computed:    true,
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"samba"}, false),

				Description: "Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported.",
				Optional:    true,
				Computed:    true,
			},
			"tod_hour": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),

				Description: "Hour of the day on which to scan the server (0 - 23, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"tod_min": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 59),

				Description: "Minute of the hour on which to scan the server (0 - 59).",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User name required to log into the file server.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mgmt", "current"}, false),

				Description: "Select the VDOM that can communicate with the file server.",
				Optional:    true,
				Computed:    true,
			},
			"weekday": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}, false),

				Description: "Day of the week on which to scan the server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceDlpFpDocSourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating DlpFpDocSource resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectDlpFpDocSource(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateDlpFpDocSource(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpFpDocSource")
	}

	return resourceDlpFpDocSourceRead(ctx, d, meta)
}

func resourceDlpFpDocSourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectDlpFpDocSource(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateDlpFpDocSource(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating DlpFpDocSource resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DlpFpDocSource")
	}

	return resourceDlpFpDocSourceRead(ctx, d, meta)
}

func resourceDlpFpDocSourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteDlpFpDocSource(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting DlpFpDocSource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpFpDocSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDlpFpDocSource(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DlpFpDocSource resource: %v", err)
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

	diags := refreshObjectDlpFpDocSource(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectDlpFpDocSource(d *schema.ResourceData, o *models.DlpFpDocSource, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Date != nil {
		v := *o.Date

		if err = d.Set("date", v); err != nil {
			return diag.Errorf("error reading date: %v", err)
		}
	}

	if o.FilePath != nil {
		v := *o.FilePath

		if err = d.Set("file_path", v); err != nil {
			return diag.Errorf("error reading file_path: %v", err)
		}
	}

	if o.FilePattern != nil {
		v := *o.FilePattern

		if err = d.Set("file_pattern", v); err != nil {
			return diag.Errorf("error reading file_pattern: %v", err)
		}
	}

	if o.KeepModified != nil {
		v := *o.KeepModified

		if err = d.Set("keep_modified", v); err != nil {
			return diag.Errorf("error reading keep_modified: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Period != nil {
		v := *o.Period

		if err = d.Set("period", v); err != nil {
			return diag.Errorf("error reading period: %v", err)
		}
	}

	if o.RemoveDeleted != nil {
		v := *o.RemoveDeleted

		if err = d.Set("remove_deleted", v); err != nil {
			return diag.Errorf("error reading remove_deleted: %v", err)
		}
	}

	if o.ScanOnCreation != nil {
		v := *o.ScanOnCreation

		if err = d.Set("scan_on_creation", v); err != nil {
			return diag.Errorf("error reading scan_on_creation: %v", err)
		}
	}

	if o.ScanSubdirectories != nil {
		v := *o.ScanSubdirectories

		if err = d.Set("scan_subdirectories", v); err != nil {
			return diag.Errorf("error reading scan_subdirectories: %v", err)
		}
	}

	if o.Sensitivity != nil {
		v := *o.Sensitivity

		if err = d.Set("sensitivity", v); err != nil {
			return diag.Errorf("error reading sensitivity: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.TodHour != nil {
		v := *o.TodHour

		if err = d.Set("tod_hour", v); err != nil {
			return diag.Errorf("error reading tod_hour: %v", err)
		}
	}

	if o.TodMin != nil {
		v := *o.TodMin

		if err = d.Set("tod_min", v); err != nil {
			return diag.Errorf("error reading tod_min: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	if o.Weekday != nil {
		v := *o.Weekday

		if err = d.Set("weekday", v); err != nil {
			return diag.Errorf("error reading weekday: %v", err)
		}
	}

	return nil
}

func getObjectDlpFpDocSource(d *schema.ResourceData, sv string) (*models.DlpFpDocSource, diag.Diagnostics) {
	obj := models.DlpFpDocSource{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("date"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("date", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Date = &tmp
		}
	}
	if v1, ok := d.GetOk("file_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("file_path", sv)
				diags = append(diags, e)
			}
			obj.FilePath = &v2
		}
	}
	if v1, ok := d.GetOk("file_pattern"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("file_pattern", sv)
				diags = append(diags, e)
			}
			obj.FilePattern = &v2
		}
	}
	if v1, ok := d.GetOk("keep_modified"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keep_modified", sv)
				diags = append(diags, e)
			}
			obj.KeepModified = &v2
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
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("period"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("period", sv)
				diags = append(diags, e)
			}
			obj.Period = &v2
		}
	}
	if v1, ok := d.GetOk("remove_deleted"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remove_deleted", sv)
				diags = append(diags, e)
			}
			obj.RemoveDeleted = &v2
		}
	}
	if v1, ok := d.GetOk("scan_on_creation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scan_on_creation", sv)
				diags = append(diags, e)
			}
			obj.ScanOnCreation = &v2
		}
	}
	if v1, ok := d.GetOk("scan_subdirectories"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scan_subdirectories", sv)
				diags = append(diags, e)
			}
			obj.ScanSubdirectories = &v2
		}
	}
	if v1, ok := d.GetOk("sensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sensitivity", sv)
				diags = append(diags, e)
			}
			obj.Sensitivity = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
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
	if v1, ok := d.GetOk("tod_hour"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tod_hour", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TodHour = &tmp
		}
	}
	if v1, ok := d.GetOk("tod_min"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tod_min", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TodMin = &tmp
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			obj.Vdom = &v2
		}
	}
	if v1, ok := d.GetOk("weekday"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weekday", sv)
				diags = append(diags, e)
			}
			obj.Weekday = &v2
		}
	}
	return &obj, diags
}
