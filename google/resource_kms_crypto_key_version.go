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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceKMSCryptoKeyVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSCryptoKeyVersionCreate,
		Read:   resourceKMSCryptoKeyVersionRead,
		Update: resourceKMSCryptoKeyVersionUpdate,
		Delete: resourceKMSCryptoKeyVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceKMSCryptoKeyVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"crypto_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of the cryptoKey associated with the CryptoKeyVersions.
Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''`,
			},
			"state": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED", ""}),
				Description:  `The current state of the CryptoKeyVersion. Possible values: ["PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED"]`,
			},
			"algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.`,
			},
			"attestation": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
Only provided for key versions with protectionLevel HSM.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_chains": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The certificate chains needed to validate the attestation`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cavium_certs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Cavium certificate chain corresponding to the attestation.`,
									},
									"google_card_certs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Google card certificate chain corresponding to the attestation.`,
									},
									"google_partition_certs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Google partition certificate chain corresponding to the attestation.`,
									},
								},
							},
						},
						"external_protection_level_options": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ekm_connection_key_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of externalKeyUri when using an EkmConnection.`,
									},
									"external_key_uri": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The URI for an external resource that this CryptoKeyVersion represents.`,
									},
								},
							},
						},
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The attestation data provided by the HSM when the key operation was performed.`,
						},
						"format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The format of the attestation data.`,
						},
					},
				},
			},
			"generate_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time this CryptoKeyVersion key material was generated`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for this CryptoKeyVersion.`,
			},
			"protection_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceKMSCryptoKeyVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	stateProp, err := expandKMSCryptoKeyVersionState(d.Get("state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("state"); !isEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}

	url, err := ReplaceVars(d, config, "{{KMSBasePath}}{{crypto_key}}/cryptoKeyVersions")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new CryptoKeyVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating CryptoKeyVersion: %s", err)
	}
	if err := d.Set("name", flattenKMSCryptoKeyVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating CryptoKeyVersion %q: %#v", d.Id(), res)

	return resourceKMSCryptoKeyVersionRead(d, meta)
}

func resourceKMSCryptoKeyVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{KMSBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("KMSCryptoKeyVersion %q", d.Id()))
	}

	if err := d.Set("name", flattenKMSCryptoKeyVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}
	if err := d.Set("state", flattenKMSCryptoKeyVersionState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}
	if err := d.Set("protection_level", flattenKMSCryptoKeyVersionProtectionLevel(res["protectionLevel"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}
	if err := d.Set("generate_time", flattenKMSCryptoKeyVersionGenerateTime(res["generateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}
	if err := d.Set("algorithm", flattenKMSCryptoKeyVersionAlgorithm(res["algorithm"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}
	if err := d.Set("attestation", flattenKMSCryptoKeyVersionAttestation(res["attestation"], d, config)); err != nil {
		return fmt.Errorf("Error reading CryptoKeyVersion: %s", err)
	}

	return nil
}

func resourceKMSCryptoKeyVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	stateProp, err := expandKMSCryptoKeyVersionState(d.Get("state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("state"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}

	url, err := ReplaceVars(d, config, "{{KMSBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating CryptoKeyVersion %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("state") {
		updateMask = append(updateMask, "state")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating CryptoKeyVersion %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating CryptoKeyVersion %q: %#v", d.Id(), res)
	}

	return resourceKMSCryptoKeyVersionRead(d, meta)
}

func resourceKMSCryptoKeyVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	cryptoKeyVersionId, err := parseKmsCryptoKeyVersionId(d.Id(), config)
	if err != nil {
		return err
	}
	if cryptoKeyVersionId == nil {
		return nil
	}
	if err := deleteCryptoKeyVersions(cryptoKeyVersionId, d, userAgent, config); err != nil {
		return nil
	}
	d.SetId("")
	return nil
}

func resourceKMSCryptoKeyVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	cryptoKeyVersionId, err := parseKmsCryptoKeyVersionId(d.Id(), config)
	if err != nil {
		return nil, err
	}
	if err := d.Set("crypto_key", cryptoKeyVersionId.CryptoKeyId.CryptoKeyId()); err != nil {
		return nil, fmt.Errorf("Error setting key_ring: %s", err)
	}
	if err := d.Set("name", cryptoKeyVersionId.Name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenKMSCryptoKeyVersionName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionProtectionLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionGenerateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAlgorithm(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["format"] =
		flattenKMSCryptoKeyVersionAttestationFormat(original["format"], d, config)
	transformed["content"] =
		flattenKMSCryptoKeyVersionAttestationContent(original["content"], d, config)
	transformed["cert_chains"] =
		flattenKMSCryptoKeyVersionAttestationCertChains(original["certChains"], d, config)
	transformed["external_protection_level_options"] =
		flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptions(original["externalProtectionLevelOptions"], d, config)
	return []interface{}{transformed}
}
func flattenKMSCryptoKeyVersionAttestationFormat(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationContent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationCertChains(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cavium_certs"] =
		flattenKMSCryptoKeyVersionAttestationCertChainsCaviumCerts(original["caviumCerts"], d, config)
	transformed["google_card_certs"] =
		flattenKMSCryptoKeyVersionAttestationCertChainsGoogleCardCerts(original["googleCardCerts"], d, config)
	transformed["google_partition_certs"] =
		flattenKMSCryptoKeyVersionAttestationCertChainsGooglePartitionCerts(original["googlePartitionCerts"], d, config)
	return []interface{}{transformed}
}
func flattenKMSCryptoKeyVersionAttestationCertChainsCaviumCerts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationCertChainsGoogleCardCerts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationCertChainsGooglePartitionCerts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["external_key_uri"] =
		flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptionsExternalKeyUri(original["externalKeyUri"], d, config)
	transformed["ekm_connection_key_path"] =
		flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptionsEkmConnectionKeyPath(original["ekmConnectionKeyPath"], d, config)
	return []interface{}{transformed}
}
func flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptionsExternalKeyUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionAttestationExternalProtectionLevelOptionsEkmConnectionKeyPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandKMSCryptoKeyVersionState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
