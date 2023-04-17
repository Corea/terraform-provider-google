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

func ResourceFirestoreDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreDatabaseCreate,
		Read:   resourceFirestoreDatabaseRead,
		Update: resourceFirestoreDatabaseUpdate,
		Delete: resourceFirestoreDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreDatabaseImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The location of the database. Available databases are listed at
https://cloud.google.com/firestore/docs/locations.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the database, which will become the final
component of the database's resource name. This value should be 4-63
characters. Valid characters are /[a-z][0-9]-/ with first character
a letter and the last a letter or a number. Must not be
UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
"(default)" database id is also valid.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEnum([]string{"FIRESTORE_NATIVE", "DATASTORE_MODE"}),
				Description: `The type of the database.
See https://cloud.google.com/datastore/docs/firestore-or-datastore
for information about how to choose. Possible values: ["FIRESTORE_NATIVE", "DATASTORE_MODE"]`,
			},
			"app_engine_integration_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"ENABLED", "DISABLED", ""}),
				Description:  `The App Engine integration mode to use for this database. Possible values: ["ENABLED", "DISABLED"]`,
			},
			"concurrency_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS", ""}),
				Description:  `The concurrency control mode to use for this database. Possible values: ["OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS"]`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp at which this database was created.`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `This checksum is computed by the server based on the value of other fields,
and may be sent on update and delete requests to ensure the client has an
up-to-date value before proceeding.`,
			},
			"key_prefix": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The keyPrefix for this database.
This keyPrefix is used, in combination with the project id ("~") to construct the application id
that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).`,
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

func resourceFirestoreDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandFirestoreDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationIdProp, err := expandFirestoreDatabaseLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !isEmptyValue(reflect.ValueOf(concurrencyModeProp)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !isEmptyValue(reflect.ValueOf(appEngineIntegrationModeProp)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases?databaseId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Database: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Database: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Database", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Database: %s", err)
	}

	if err := d.Set("name", flattenFirestoreDatabaseName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Database %q: %#v", d.Id(), res)

	return resourceFirestoreDatabaseRead(d, meta)
}

func resourceFirestoreDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirestoreDatabase %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	if err := d.Set("name", flattenFirestoreDatabaseName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("location_id", flattenFirestoreDatabaseLocationId(res["locationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("type", flattenFirestoreDatabaseType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("concurrency_mode", flattenFirestoreDatabaseConcurrencyMode(res["concurrencyMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("app_engine_integration_mode", flattenFirestoreDatabaseAppEngineIntegrationMode(res["appEngineIntegrationMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("key_prefix", flattenFirestoreDatabaseKeyPrefix(res["key_prefix"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("etag", flattenFirestoreDatabaseEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("create_time", flattenFirestoreDatabaseCreateTime(res["create_time"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	return nil
}

func resourceFirestoreDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Database %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("concurrency_mode") {
		updateMask = append(updateMask, "concurrencyMode")
	}

	if d.HasChange("app_engine_integration_mode") {
		updateMask = append(updateMask, "appEngineIntegrationMode")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
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
		return fmt.Errorf("Error updating Database %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Database %q: %#v", d.Id(), res)
	}

	err = FirestoreOperationWaitTime(
		config, res, project, "Updating Database", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceFirestoreDatabaseRead(d, meta)
}

func resourceFirestoreDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Firestore Database resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirestoreDatabaseImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/databases/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreDatabaseName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenFirestoreDatabaseLocationId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseConcurrencyMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseKeyPrefix(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseEtag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreDatabaseCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirestoreDatabaseName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
}

func expandFirestoreDatabaseLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseConcurrencyMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseEtag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
