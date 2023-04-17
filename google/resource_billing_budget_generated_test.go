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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBillingBudget_billingBudgetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetBasicExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget%{random_suffix}"
  amount {
    specified_amount {
      currency_code = "USD"
      units = "100000"
    }
  }
  threshold_rules {
      threshold_percent =  0.5
  }
}
`, context)
}

func TestAccBillingBudget_billingBudgetLastperiodExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetLastperiodExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetLastperiodExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

data "google_project" "project" {
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget%{random_suffix}"
  
  budget_filter {
    projects = ["projects/${data.google_project.project.number}"]
  }

  amount {
    last_period_amount = true
  }

  threshold_rules {
      threshold_percent =  10.0
      # Typically threshold_percent would be set closer to 1.0 (100%).
      # It has been purposely set high (10.0 / 1000%) in this example
      # so it does not trigger alerts during automated testing.
  }
}
`, context)
}

func TestAccBillingBudget_billingBudgetFilterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetFilterExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetFilterExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

data "google_project" "project" {
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget%{random_suffix}"

  budget_filter {
    projects               = ["projects/${data.google_project.project.number}"]
    credit_types_treatment = "INCLUDE_SPECIFIED_CREDITS"
    services               = ["services/24E6-581D-38E5"] # Bigquery
    credit_types           = ["PROMOTION", "FREE_TIER"]
  }

  amount {
    specified_amount {
      currency_code = "USD"
      units = "100000"
    }
  }

  threshold_rules {
    threshold_percent = 0.5
  }
  threshold_rules {
    threshold_percent = 0.9
    spend_basis = "FORECASTED_SPEND"
  }
}
`, context)
}

func TestAccBillingBudget_billingBudgetNotifyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetNotifyExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetNotifyExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

data "google_project" "project" {
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name    = "Example Billing Budget%{random_suffix}"

  budget_filter {
    projects = ["projects/${data.google_project.project.number}"]
  }

  amount {
    specified_amount {
      currency_code = "USD"
      units         = "100000"
    }
  }

  threshold_rules {
    threshold_percent = 1.0
  }
  threshold_rules {
    threshold_percent = 1.0
    spend_basis       = "FORECASTED_SPEND"
  }
  
  all_updates_rule {
    monitoring_notification_channels = [
      google_monitoring_notification_channel.notification_channel.id,
    ]
    disable_default_iam_recipients = true
  }
}

resource "google_monitoring_notification_channel" "notification_channel" {
  display_name = "Example Notification Channel%{random_suffix}"
  type         = "email"
  
  labels = {
    email_address = "address@example.com"
  }
}
`, context)
}

func TestAccBillingBudget_billingBudgetCustomperiodExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetCustomperiodExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetCustomperiodExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

data "google_project" "project" {
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget%{random_suffix}"

  budget_filter {
    projects = ["projects/${data.google_project.project.number}"]
    credit_types_treatment = "EXCLUDE_ALL_CREDITS"
    services = ["services/24E6-581D-38E5"] # Bigquery
    
    custom_period { 
        start_date {
          year = 2022
          month = 1
          day = 1
        }
        end_date {
          year = 2023
          month = 12
          day = 31
        }
      }
  }

  amount {
    specified_amount {
      currency_code = "USD"
      units = "100000"
    }
  }

  threshold_rules {
    threshold_percent = 0.5
  }
  threshold_rules {
    threshold_percent = 0.9
  }
}
`, context)
}

func TestAccBillingBudget_billingBudgetOptionalExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":  GetTestMasterBillingAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBillingBudgetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBillingBudget_billingBudgetOptionalExample(context),
			},
			{
				ResourceName:            "google_billing_budget.budget",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func testAccBillingBudget_billingBudgetOptionalExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_billing_account" "account" {
  billing_account = "%{billing_acct}"
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget%{random_suffix}"

  amount {
    specified_amount {
      currency_code = "USD"
      units = "100000"
    }
  }

  all_updates_rule {
    disable_default_iam_recipients = true
    pubsub_topic = google_pubsub_topic.budget.id
  }
}

resource "google_pubsub_topic" "budget" {
  name = "tf-test-example-topic%{random_suffix}"
}
`, context)
}

func testAccCheckBillingBudgetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_billing_budget" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BillingBasePath}}billingAccounts/{{billing_account}}/budgets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("BillingBudget still exists at %s", url)
			}
		}

		return nil
	}
}
