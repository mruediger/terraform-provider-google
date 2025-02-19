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

package integrations_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccIntegrationsClient_integrationsClientBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsClient_integrationsClientBasicExample(context),
			},
			{
				ResourceName:            "google_integrations_client.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cloud_kms_config", "create_sample_workflows", "provision_gmek", "run_as_service_account", "location"},
			},
		},
	})
}

func testAccIntegrationsClient_integrationsClientBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "example" {
  location = "us-central1"
}
`, context)
}

func TestAccIntegrationsClient_integrationsClientAdvanceExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsClient_integrationsClientAdvanceExample(context),
			},
			{
				ResourceName:            "google_integrations_client.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cloud_kms_config", "create_sample_workflows", "provision_gmek", "run_as_service_account", "location"},
			},
		},
	})
}

func testAccIntegrationsClient_integrationsClientAdvanceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "test_project" {
}

resource "google_kms_key_ring" "keyring" {
  name     = "tf-test-my-keyring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "cryptokey" {
  name = "crypto-key-example"
  key_ring = google_kms_key_ring.keyring.id
  rotation_period = "7776000s"
  depends_on = [google_kms_key_ring.keyring]
}

resource "google_kms_crypto_key_version" "test_key" {
  crypto_key = google_kms_crypto_key.cryptokey.id
  depends_on = [google_kms_crypto_key.cryptokey]
}

resource "google_integrations_client" "example" {
  location = "us-central1"
  create_sample_workflows = true
  provision_gmek = true
  run_as_service_account = "radndom-service-account"
  cloud_kms_config {
    kms_location = "us-central1"
    kms_ring = google_kms_key_ring.keyring.id
    key = google_kms_crypto_key.cryptokey.id
    key_version = google_kms_crypto_key_version.test_key.id
    kms_project_id = data.google_project.test_project.id
  }
  depends_on = [google_kms_crypto_key_version.test_key]
}
`, context)
}
