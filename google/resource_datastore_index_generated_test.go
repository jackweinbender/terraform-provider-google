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

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDatastoreIndex_datastoreIndexExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDatastoreIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastoreIndex_datastoreIndexExample(context),
			},
			{
				ResourceName:      "google_datastore_index.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDatastoreIndex_datastoreIndexExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_datastore_index" "default" {
  kind = "foo"
  properties {
    name = "tf_test_property_a%{random_suffix}"
    direction = "ASCENDING"
  }
  properties {
    name = "tf_test_property_b%{random_suffix}"
    direction = "ASCENDING"
  }
}
`, context)
}

func testAccCheckDatastoreIndexDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_datastore_index" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DatastoreBasePath}}projects/{{project}}/indexes/{{index_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.DatastoreIndex409Contention},
			})
			if err == nil {
				return fmt.Errorf("DatastoreIndex still exists at %s", url)
			}
		}

		return nil
	}
}
