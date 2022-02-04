// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization peers.

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

func resourceWanoptPeer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization peers.",

		CreateContext: resourceWanoptPeerCreate,
		ReadContext:   resourceWanoptPeerRead,
		UpdateContext: resourceWanoptPeerUpdate,
		DeleteContext: resourceWanoptPeerDelete,

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
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Peer IP address.",
				Optional:    true,
				Computed:    true,
			},
			"peer_host_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Peer host ID.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWanoptPeerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "peer_host_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WanoptPeer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWanoptPeer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptPeer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(ctx, d, meta)
}

func resourceWanoptPeerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptPeer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptPeer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptPeer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(ctx, d, meta)
}

func resourceWanoptPeerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWanoptPeer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptPeerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptPeer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptPeer resource: %v", err)
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

	diags := refreshObjectWanoptPeer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWanoptPeer(d *schema.ResourceData, o *models.WanoptPeer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.PeerHostId != nil {
		v := *o.PeerHostId

		if err = d.Set("peer_host_id", v); err != nil {
			return diag.Errorf("error reading peer_host_id: %v", err)
		}
	}

	return nil
}

func getObjectWanoptPeer(d *schema.ResourceData, sv string) (*models.WanoptPeer, diag.Diagnostics) {
	obj := models.WanoptPeer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("peer_host_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer_host_id", sv)
				diags = append(diags, e)
			}
			obj.PeerHostId = &v2
		}
	}
	return &obj, diags
}
