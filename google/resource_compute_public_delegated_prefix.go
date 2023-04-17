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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceComputePublicDelegatedPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputePublicDelegatedPrefixCreate,
		Read:   resourceComputePublicDelegatedPrefixRead,
		Delete: resourceComputePublicDelegatedPrefixDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputePublicDelegatedPrefixImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The IPv4 address range, in CIDR format, represented by this public advertised prefix.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
which means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"parent_prefix": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A region where the prefix will reside.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"is_live_migration": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, the prefix will be live migrated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputePublicDelegatedPrefixCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicDelegatedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	isLiveMigrationProp, err := expandComputePublicDelegatedPrefixIsLiveMigration(d.Get("is_live_migration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_live_migration"); !isEmptyValue(reflect.ValueOf(isLiveMigrationProp)) && (ok || !reflect.DeepEqual(v, isLiveMigrationProp)) {
		obj["isLiveMigration"] = isLiveMigrationProp
	}
	nameProp, err := expandComputePublicDelegatedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	parentPrefixProp, err := expandComputePublicDelegatedPrefixParentPrefix(d.Get("parent_prefix"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_prefix"); !isEmptyValue(reflect.ValueOf(parentPrefixProp)) && (ok || !reflect.DeepEqual(v, parentPrefixProp)) {
		obj["parentPrefix"] = parentPrefixProp
	}
	ipCidrRangeProp, err := expandComputePublicDelegatedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PublicDelegatedPrefix: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating PublicDelegatedPrefix: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating PublicDelegatedPrefix", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PublicDelegatedPrefix: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PublicDelegatedPrefix %q: %#v", d.Id(), res)

	return resourceComputePublicDelegatedPrefixRead(d, meta)
}

func resourceComputePublicDelegatedPrefixRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputePublicDelegatedPrefix %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}

	if err := d.Set("description", flattenComputePublicDelegatedPrefixDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("is_live_migration", flattenComputePublicDelegatedPrefixIsLiveMigration(res["isLiveMigration"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("name", flattenComputePublicDelegatedPrefixName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("parent_prefix", flattenComputePublicDelegatedPrefixParentPrefix(res["parentPrefix"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenComputePublicDelegatedPrefixIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}

	return nil
}

func resourceComputePublicDelegatedPrefixDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting PublicDelegatedPrefix %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "PublicDelegatedPrefix")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting PublicDelegatedPrefix", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PublicDelegatedPrefix %q: %#v", d.Id(), res)
	return nil
}

func resourceComputePublicDelegatedPrefixImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/publicDelegatedPrefixes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputePublicDelegatedPrefixDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixIsLiveMigration(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixParentPrefix(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixIpCidrRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputePublicDelegatedPrefixDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIsLiveMigration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixParentPrefix(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
