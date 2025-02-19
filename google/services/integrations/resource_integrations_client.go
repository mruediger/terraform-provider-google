// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package integrations

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIntegrationsClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegrationsClientCreate,
		Read:   resourceIntegrationsClientRead,
		Delete: resourceIntegrationsClientDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIntegrationsClientImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location in which client needs to be provisioned.`,
			},
			"cloud_kms_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Cloud KMS config for AuthModule to encrypt/decrypt credentials.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `A Cloud KMS key is a named object containing one or more key versions, along
with metadata for the key. A key exists on exactly one key ring tied to a
specific location.`,
						},
						"kms_location": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Location name of the key ring, e.g. "us-west1".`,
						},
						"kms_ring": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `A key ring organizes keys in a specific Google Cloud location and allows you to
manage access control on groups of keys. A key ring's name does not need to be
unique across a Google Cloud project, but must be unique within a given location.`,
						},
						"key_version": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Each version of a key contains key material used for encryption or signing.
A key's version is represented by an integer, starting at 1. To decrypt data
or verify a signature, you must use the same key version that was used to
encrypt or sign the data.`,
						},
						"kms_project_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The Google Cloud project id of the project where the kms key stored. If empty,
the kms key is stored at the same project as customer's project and ecrypted
with CMEK, otherwise, the kms key is stored in the tenant project and
encrypted with GMEK.`,
						},
					},
				},
			},
			"create_sample_workflows": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Indicates if sample workflow should be created along with provisioning.`,
			},
			"provision_gmek": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Indicates provision with GMEK or CMEK.`,
			},
			"run_as_service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `User input run-as service account, if empty, will bring up a new default service account.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIntegrationsClientCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})

	lockName, err := tpgresource.ReplaceVars(d, config, "Client/{{location}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationsBasePath}}projects/{{project}}/locations/{{location}}/clients:provision")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Client: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Client: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Client: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Client %q: %#v", d.Id(), res)

	return resourceIntegrationsClientRead(d, meta)
}

func resourceIntegrationsClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationsBasePath}}projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Client: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IntegrationsClient %q", d.Id()))
	}

	res, err = resourceIntegrationsClientDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IntegrationsClient because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}

	return nil
}

func resourceIntegrationsClientDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Integrations Client resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceIntegrationsClientImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/clients$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)$",
		"^(?P<location>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func resourceIntegrationsClientDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Since Client resource doesnt have any properties,
	// Adding this decoder as placeholder else the linter will
	// complain that the returned `res` is never used afterwards.
	return res, nil
}
