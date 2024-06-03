// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization authentication groups.

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

func resourceWanoptAuthGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization authentication groups.",

		CreateContext: resourceWanoptAuthGroupCreate,
		ReadContext:   resourceWanoptAuthGroupRead,
		UpdateContext: resourceWanoptAuthGroupUpdate,
		DeleteContext: resourceWanoptAuthGroupDelete,

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
			"auth_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cert", "psk"}, false),

				Description: "Select certificate or pre-shared key authentication for this authentication group.",
				Optional:    true,
				Computed:    true,
			},
			"cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of certificate to identify this peer.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Auth-group name.",
				ForceNew:    true,
				Required:    true,
			},
			"peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.",
				Optional:    true,
				Computed:    true,
			},
			"peer_accept": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "defined", "one"}, false),

				Description: "Determine if this auth group accepts, any peer, a list of defined peers, or just one peer.",
				Optional:    true,
				Computed:    true,
			},
			"psk": {
				Type: schema.TypeString,

				Description: "Pre-shared key used by the peers in this authentication group.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func resourceWanoptAuthGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WanoptAuthGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWanoptAuthGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptAuthGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptAuthGroup")
	}

	return resourceWanoptAuthGroupRead(ctx, d, meta)
}

func resourceWanoptAuthGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptAuthGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptAuthGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptAuthGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptAuthGroup")
	}

	return resourceWanoptAuthGroupRead(ctx, d, meta)
}

func resourceWanoptAuthGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWanoptAuthGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptAuthGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptAuthGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptAuthGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptAuthGroup resource: %v", err)
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

	diags := refreshObjectWanoptAuthGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWanoptAuthGroup(d *schema.ResourceData, o *models.WanoptAuthGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthMethod != nil {
		v := *o.AuthMethod

		if err = d.Set("auth_method", v); err != nil {
			return diag.Errorf("error reading auth_method: %v", err)
		}
	}

	if o.Cert != nil {
		v := *o.Cert

		if err = d.Set("cert", v); err != nil {
			return diag.Errorf("error reading cert: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Peer != nil {
		v := *o.Peer

		if err = d.Set("peer", v); err != nil {
			return diag.Errorf("error reading peer: %v", err)
		}
	}

	if o.PeerAccept != nil {
		v := *o.PeerAccept

		if err = d.Set("peer_accept", v); err != nil {
			return diag.Errorf("error reading peer_accept: %v", err)
		}
	}

	if o.Psk != nil {
		v := *o.Psk

		if v == "ENC XXXX" {
		} else if err = d.Set("psk", v); err != nil {
			return diag.Errorf("error reading psk: %v", err)
		}
	}

	return nil
}

func getObjectWanoptAuthGroup(d *schema.ResourceData, sv string) (*models.WanoptAuthGroup, diag.Diagnostics) {
	obj := models.WanoptAuthGroup{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_method", sv)
				diags = append(diags, e)
			}
			obj.AuthMethod = &v2
		}
	}
	if v1, ok := d.GetOk("cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert", sv)
				diags = append(diags, e)
			}
			obj.Cert = &v2
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
	if v1, ok := d.GetOk("peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer", sv)
				diags = append(diags, e)
			}
			obj.Peer = &v2
		}
	}
	if v1, ok := d.GetOk("peer_accept"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_accept", sv)
				diags = append(diags, e)
			}
			obj.PeerAccept = &v2
		}
	}
	if v1, ok := d.GetOk("psk"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("psk", sv)
				diags = append(diags, e)
			}
			obj.Psk = &v2
		}
	}
	return &obj, diags
}
