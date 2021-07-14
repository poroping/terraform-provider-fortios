// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAccessProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyCreate,
		Read:   resourceFirewallAccessProxyRead,
		Update: resourceFirewallAccessProxyUpdate,
		Delete: resourceFirewallAccessProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_table": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"api_gateway": {
				Type:        schema.TypeList,
				Description: "Set API Gateway.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 525600),
							Description:  "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Description:  "Domain that HTTP cookie persistence should apply to.",
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_domain_from_host": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_generation": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Description:  "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Description:  "Limit HTTP cookie persistence to the specified path.",
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_share": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"disable", "same-ip"}),
							Description:  "Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Optional:     true,
							Computed:     true,
						},
						"https_cookie_secure": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable verification that inserted HTTPS cookies are secure.",
							Optional:     true,
							Computed:     true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Description:  "API Gateway ID.",
							Optional:     true,
							Computed:     true,
						},
						"ldb_method": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"static", "round-robin", "weighted", "least-session", "least-rtt", "first-alive", "http-host"}),
							Description:  "Method used to distribute sessions to real servers.",
							Optional:     true,
							Computed:     true,
						},
						"persistence": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "http-cookie"}),
							Description:  "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Optional:     true,
							Computed:     true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Description:  "Address or address group of the real server.",
										Optional:     true,
										Computed:     true,
									},
									"health_check": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnableDisable(),
										Description:  "Enable to check the responsiveness of the real server before forwarding traffic.",
										Optional:     true,
										Computed:     true,
									},
									"health_check_proto": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"ping", "http", "tcp-connect"}),
										Description:  "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Optional:     true,
										Computed:     true,
									},
									"http_host": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Description:  "HTTP server domain name in HTTP header.",
										Optional:     true,
										Computed:     true,
									},
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Description:  "Real server ID.",
										Optional:     true,
										Computed:     true,
									},
									"ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,
										Description:  "IP address of the real server.",
										Optional:     true,
										Computed:     true,
									},
									"mappedport": {
										Type: schema.TypeString,

										Description: "Port for communicating with the real server.",
										Optional:    true,
										Computed:    true,
									},
									"port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Description:  "Port for communicating with the real server.",
										Optional:     true,
										Computed:     true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"active", "standby", "disable"}),
										Description:  "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Optional:     true,
										Computed:     true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Description:  "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"saml_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Description:  "SAML service provider configuration for VIP authentication.",
							Optional:     true,
							Computed:     true,
						},
						"service": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"http", "https", "tcp-forwarding", "samlsp"}),
							Description:  "Service.",
							Optional:     true,
							Computed:     true,
						},
						"ssl_algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"high", "medium", "low", "custom"}),
							Description:  "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Optional:     true,
							Computed:     true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"TLS-AES-128-GCM-SHA256", "TLS-AES-256-GCM-SHA384", "TLS-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-128-GCM-SHA256", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-AES-128-CBC-SHA", "TLS-RSA-WITH-AES-256-CBC-SHA", "TLS-RSA-WITH-AES-128-CBC-SHA256", "TLS-RSA-WITH-AES-128-GCM-SHA256", "TLS-RSA-WITH-AES-256-CBC-SHA256", "TLS-RSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-SEED-CBC-SHA", "TLS-DHE-DSS-WITH-SEED-CBC-SHA", "TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384", "TLS-RSA-WITH-SEED-CBC-SHA", "TLS-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-RC4-128-SHA", "TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-RC4-128-MD5", "TLS-RSA-WITH-RC4-128-SHA", "TLS-DHE-RSA-WITH-DES-CBC-SHA", "TLS-DHE-DSS-WITH-DES-CBC-SHA", "TLS-RSA-WITH-DES-CBC-SHA"}),
										Description:  "Cipher suite name.",
										Optional:     true,
										Computed:     true,
									},
									"priority": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Description:  "SSL/TLS cipher suites priority.",
										Optional:     true,
										Computed:     true,
									},
									"versions": {
										Type: schema.TypeString,

										Description: "SSL/TLS versions that the cipher suite can be used with.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"768", "1024", "1536", "2048", "3072", "4096"}),
							Description:  "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Optional:     true,
							Computed:     true,
						},
						"ssl_max_version": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}),
							Description:  "Highest SSL/TLS version acceptable from a server.",
							Optional:     true,
							Computed:     true,
						},
						"ssl_min_version": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}),
							Description:  "Lowest SSL/TLS version acceptable from a server.",
							Optional:     true,
							Computed:     true,
						},
						"url_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Description:  "URL pattern to match.",
							Optional:     true,
							Computed:     true,
						},
						"url_map_type": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"sub-string", "wildcard", "regex"}),
							Description:  "Type of url-map.",
							Optional:     true,
							Computed:     true,
						},
						"virtual_host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "Virtual host.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable to request client certificate.",
				Optional:     true,
				Computed:     true,
			},
			"empty_cert_action": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"accept", "block"}),
				Description:  "Action of an empty client certificate.",
				Optional:     true,
				Computed:     true,
			},
			"ldb_method": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"static", "round-robin", "weighted", "least-session", "least-rtt", "first-alive"}),
				Description:  "Method used to distribute sessions to SSL real servers.",
				Optional:     true,
				Computed:     true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Description:  "Access Proxy name.",
				ForceNew:     true,
				Required:     true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Select the SSL real servers that this Access Proxy will distribute traffic to.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Description:  "Real server ID.",
							Optional:     true,
							Computed:     true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,
							Description:  "IP address of the real server.",
							Optional:     true,
							Computed:     true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Description:  "Port for communicating with the real server.",
							Optional:     true,
							Computed:     true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"active", "standby", "disable"}),
							Description:  "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
							Optional:     true,
							Computed:     true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Description:  "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"server_pubkey_auth": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable SSH real server public key authentication.",
				Optional:     true,
				Computed:     true,
			},
			"server_pubkey_auth_settings": {
				Type:        schema.TypeList,
				Description: "Server SSH public key authentication settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ca": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "Name of the SSH server public key authentication CA.",
							Optional:     true,
							Computed:     true,
						},
						"cert_extension": {
							Type:        schema.TypeList,
							Description: "Configure certificate extension for user certificate.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"critical": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"no", "yes"}),
										Description:  "Critical option.",
										Optional:     true,
										Computed:     true,
									},
									"data": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Description:  "Name of certificate extension.",
										Optional:     true,
										Computed:     true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Description:  "Name of certificate extension.",
										Optional:     true,
										Computed:     true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"fixed", "user"}),
										Description:  "Type of certificate extension.",
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"permit_agent_forwarding": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending permit-agent-forwarding certificate extension.",
							Optional:     true,
							Computed:     true,
						},
						"permit_port_forwarding": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending permit-port-forwarding certificate extension.",
							Optional:     true,
							Computed:     true,
						},
						"permit_pty": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending permit-pty certificate extension.",
							Optional:     true,
							Computed:     true,
						},
						"permit_user_rc": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending permit-user-rc certificate extension.",
							Optional:     true,
							Computed:     true,
						},
						"permit_x11_forwarding": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending permit-x11-forwarding certificate extension.",
							Optional:     true,
							Computed:     true,
						},
						"source_address": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),
							Description:  "Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"vip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Description:  "Virtual IP name.",
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallAccessProxyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectFirewallAccessProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating FirewallAccessProxy resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating FirewallAccessProxy resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateFirewallAccessProxy(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateFirewallAccessProxy(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating FirewallAccessProxy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy")
	}

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectFirewallAccessProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating FirewallAccessProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAccessProxy(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating FirewallAccessProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy")
	}

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteFirewallAccessProxy(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting FirewallAccessProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadFirewallAccessProxy(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallAccessProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading FirewallAccessProxy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyApiGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {

			tmp["http_cookie_age"] = flattenFirewallAccessProxyApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {

			tmp["http_cookie_domain"] = flattenFirewallAccessProxyApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {

			tmp["http_cookie_domain_from_host"] = flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {

			tmp["http_cookie_generation"] = flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {

			tmp["http_cookie_path"] = flattenFirewallAccessProxyApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {

			tmp["http_cookie_share"] = flattenFirewallAccessProxyApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {

			tmp["https_cookie_secure"] = flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallAccessProxyApiGatewayId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {

			tmp["ldb_method"] = flattenFirewallAccessProxyApiGatewayLdbMethod(i["ldb-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {

			tmp["persistence"] = flattenFirewallAccessProxyApiGatewayPersistence(i["persistence"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {

			tmp["realservers"] = flattenFirewallAccessProxyApiGatewayRealservers(i["realservers"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {

			tmp["saml_server"] = flattenFirewallAccessProxyApiGatewaySamlServer(i["saml-server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {

			tmp["service"] = flattenFirewallAccessProxyApiGatewayService(i["service"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {

			tmp["ssl_algorithm"] = flattenFirewallAccessProxyApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {

			tmp["ssl_cipher_suites"] = flattenFirewallAccessProxyApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {

			tmp["ssl_dh_bits"] = flattenFirewallAccessProxyApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {

			tmp["ssl_max_version"] = flattenFirewallAccessProxyApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {

			tmp["ssl_min_version"] = flattenFirewallAccessProxyApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {

			tmp["url_map"] = flattenFirewallAccessProxyApiGatewayUrlMap(i["url-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {

			tmp["url_map_type"] = flattenFirewallAccessProxyApiGatewayUrlMapType(i["url-map-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {

			tmp["virtual_host"] = flattenFirewallAccessProxyApiGatewayVirtualHost(i["virtual-host"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}
func flattenFirewallAccessProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = flattenFirewallAccessProxyApiGatewayRealserversAddress(i["address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {

			tmp["health_check_proto"] = flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = flattenFirewallAccessProxyApiGatewayRealserversHttpHost(i["http-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallAccessProxyApiGatewayRealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallAccessProxyApiGatewayRealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {

			tmp["mappedport"] = flattenFirewallAccessProxyApiGatewayRealserversMappedport(i["mappedport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallAccessProxyApiGatewayRealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallAccessProxyApiGatewayRealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallAccessProxyApiGatewayRealserversWeight(i["weight"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}
func flattenFirewallAccessProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}
func flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallAccessProxyRealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallAccessProxyRealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallAccessProxyRealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallAccessProxyRealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallAccessProxyRealserversWeight(i["weight"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}
func flattenFirewallAccessProxyRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettings(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_ca"
		if _, ok := i["auth-ca"]; ok {

			tmp["auth_ca"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(i["auth-ca"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert_extension"
		if _, ok := i["cert-extension"]; ok {

			tmp["cert_extension"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(i["cert-extension"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_agent_forwarding"
		if _, ok := i["permit-agent-forwarding"]; ok {

			tmp["permit_agent_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(i["permit-agent-forwarding"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_port_forwarding"
		if _, ok := i["permit-port-forwarding"]; ok {

			tmp["permit_port_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(i["permit-port-forwarding"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_pty"
		if _, ok := i["permit-pty"]; ok {

			tmp["permit_pty"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(i["permit-pty"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_user_rc"
		if _, ok := i["permit-user-rc"]; ok {

			tmp["permit_user_rc"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(i["permit-user-rc"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_x11_forwarding"
		if _, ok := i["permit-x11-forwarding"]; ok {

			tmp["permit_x11_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(i["permit-x11-forwarding"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := i["source-address"]; ok {

			tmp["source_address"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(i["source-address"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}
func flattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {

			tmp["critical"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(i["critical"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {

			tmp["data"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(i["data"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(i["type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}
func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway"]) {
				return fmt.Errorf("error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway"]) {
					return fmt.Errorf("error reading api_gateway: %v", err)
				}
			}
		}
	}

	if err = d.Set("client_cert", flattenFirewallAccessProxyClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}
	if err = d.Set("empty_cert_action", flattenFirewallAccessProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action", sv)); err != nil {
		if !fortiAPIPatch(o["empty-cert-action"]) {
			return fmt.Errorf("error reading empty_cert_action: %v", err)
		}
	}
	if err = d.Set("ldb_method", flattenFirewallAccessProxyLdbMethod(o["ldb-method"], d, "ldb_method", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("error reading ldb_method: %v", err)
		}
	}
	if err = d.Set("name", flattenFirewallAccessProxyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}
	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallAccessProxyRealservers(o["realservers"], d, "realservers", sv)); err != nil {
			if !fortiAPIPatch(o["realservers"]) {
				return fmt.Errorf("error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallAccessProxyRealservers(o["realservers"], d, "realservers", sv)); err != nil {
				if !fortiAPIPatch(o["realservers"]) {
					return fmt.Errorf("error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_pubkey_auth", flattenFirewallAccessProxyServerPubkeyAuth(o["server-pubkey-auth"], d, "server_pubkey_auth", sv)); err != nil {
		if !fortiAPIPatch(o["server-pubkey-auth"]) {
			return fmt.Errorf("error reading server_pubkey_auth: %v", err)
		}
	}
	if isImportTable() {
		if err = d.Set("server_pubkey_auth_settings", flattenFirewallAccessProxyServerPubkeyAuthSettings(o["server-pubkey-auth-settings"], d, "server_pubkey_auth_settings", sv)); err != nil {
			if !fortiAPIPatch(o["server-pubkey-auth-settings"]) {
				return fmt.Errorf("error reading server_pubkey_auth_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_pubkey_auth_settings"); ok {
			if err = d.Set("server_pubkey_auth_settings", flattenFirewallAccessProxyServerPubkeyAuthSettings(o["server-pubkey-auth-settings"], d, "server_pubkey_auth_settings", sv)); err != nil {
				if !fortiAPIPatch(o["server-pubkey-auth-settings"]) {
					return fmt.Errorf("error reading server_pubkey_auth_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("vip", flattenFirewallAccessProxyVip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("error reading vip: %v", err)
		}
	}

	return nil
}

func expandFirewallAccessProxyApiGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-age"], _ = expandFirewallAccessProxyApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-generation"], _ = expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-path"], _ = expandFirewallAccessProxyApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-share"], _ = expandFirewallAccessProxyApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["https-cookie-secure"], _ = expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallAccessProxyApiGatewayId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ldb-method"], _ = expandFirewallAccessProxyApiGatewayLdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["persistence"], _ = expandFirewallAccessProxyApiGatewayPersistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["realservers"], _ = expandFirewallAccessProxyApiGatewayRealservers(d, i["realservers"], pre_append, sv)
		} else {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["saml-server"], _ = expandFirewallAccessProxyApiGatewaySamlServer(d, i["saml_server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service"], _ = expandFirewallAccessProxyApiGatewayService(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-algorithm"], _ = expandFirewallAccessProxyApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-cipher-suites"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxyApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-max-version"], _ = expandFirewallAccessProxyApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-min-version"], _ = expandFirewallAccessProxyApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map"], _ = expandFirewallAccessProxyApiGatewayUrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map-type"], _ = expandFirewallAccessProxyApiGatewayUrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["virtual-host"], _ = expandFirewallAccessProxyApiGatewayVirtualHost(d, i["virtual_host"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["address"], _ = expandFirewallAccessProxyApiGatewayRealserversAddress(d, i["address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-host"], _ = expandFirewallAccessProxyApiGatewayRealserversHttpHost(d, i["http_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallAccessProxyApiGatewayRealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallAccessProxyApiGatewayRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mappedport"], _ = expandFirewallAccessProxyApiGatewayRealserversMappedport(d, i["mappedport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallAccessProxyApiGatewayRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallAccessProxyApiGatewayRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallAccessProxyApiGatewayRealserversWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallAccessProxyRealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallAccessProxyRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallAccessProxyRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallAccessProxyRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallAccessProxyRealserversWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_ca"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-ca"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(d, i["auth_ca"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert_extension"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cert-extension"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(d, i["cert_extension"], pre_append, sv)
		} else {
			tmp["cert-extension"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_agent_forwarding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["permit-agent-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(d, i["permit_agent_forwarding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_port_forwarding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["permit-port-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(d, i["permit_port_forwarding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_pty"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["permit-pty"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(d, i["permit_pty"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_user_rc"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["permit-user-rc"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(d, i["permit_user_rc"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_x11_forwarding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["permit-x11-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(d, i["permit_x11_forwarding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-address"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(d, i["source_address"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["critical"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(d, i["critical"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["data"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(d, i["data"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(d, i["type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("api_gateway"); ok {

		t, err := expandFirewallAccessProxyApiGateway(d, v, "api_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}
	if v, ok := d.GetOk("client_cert"); ok {

		t, err := expandFirewallAccessProxyClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}
	if v, ok := d.GetOk("empty_cert_action"); ok {

		t, err := expandFirewallAccessProxyEmptyCertAction(d, v, "empty_cert_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}
	if v, ok := d.GetOk("ldb_method"); ok {

		t, err := expandFirewallAccessProxyLdbMethod(d, v, "ldb_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallAccessProxyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("realservers"); ok {

		t, err := expandFirewallAccessProxyRealservers(d, v, "realservers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}
	if v, ok := d.GetOk("server_pubkey_auth"); ok {

		t, err := expandFirewallAccessProxyServerPubkeyAuth(d, v, "server_pubkey_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-pubkey-auth"] = t
		}
	}
	if v, ok := d.GetOk("server_pubkey_auth_settings"); ok {

		t, err := expandFirewallAccessProxyServerPubkeyAuthSettings(d, v, "server_pubkey_auth_settings", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-pubkey-auth-settings"] = t
		}
	}
	if v, ok := d.GetOk("vip"); ok {

		t, err := expandFirewallAccessProxyVip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	return &obj, nil
}
