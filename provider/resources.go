// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package honeycomb

import (
	"fmt"
	"path/filepath"

	"github.com/MaterializeInc/pulumi-honeycomb/provider/pkg/version"
	"github.com/honeycombio/terraform-provider-honeycombio/honeycombio"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "honeycomb"
	// modules:
	mainMod = "index" // the xyz module
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(honeycombio.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "honeycombio",
		Description:       "A Pulumi package for creating and managing honeycomb.io resources.",
		Keywords:          []string{"pulumi", "honeycomb", "tracing", "o11y", "monitoring", "otel", "category/utility"},
		License:           "Apache-2.0",
		Publisher:         "Materialize Inc",
		LogoURL:           "https://raw.githubusercontent.com/MaterializeInc/pulumi-honeycomb/main/assets/honeycomb.svg",
		GitHubOrg:         "honeycombio",
		Homepage:          "https://github.com/honeycombio/terraform-provider-honeycombio",
		Repository:        "https://github.com/honeycombio/terraform-provider-honeycombio",
		PluginDownloadURL: "https://github.com/MaterializeInc/pulumi-honeycomb/releases/download/v${VERSION}",
		DisplayName:       "Honeycomb.io",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"honeycombio_board":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Board")},
			"honeycombio_burn_alert":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BurnAlert")},
			"honeycombio_column":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Column")},
			"honeycombio_dataset":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dataset")},
			"honeycombio_derived_column":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DerivedColumn")},
			"honeycombio_email_recipient":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EmailRecipient")},
			"honeycombio_marker":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Marker")},
			"honeycombio_pagerduty_recipient":{Tok: tfbridge.MakeResource(mainPkg, mainMod, "PagerdutyRecipient")},
			"honeycombio_query":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Query")},
			"honeycombio_query_annotation": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "QueryAnnotation")},
			"honeycombio_slo":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SLO")},
			"honeycombio_trigger":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Trigger")},
			"honeycombio_webhook_recipient":{Tok: tfbridge.MakeResource(mainPkg, mainMod, "WebhookRecipient")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"honeycombio_datasets":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetDatasets")},
			"honeycombio_query_result":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetQueryResult")},
			"honeycombio_query_specification": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetQuerySpecification")},
			"honeycombio_recipient":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetRecipient")},
			"honeycombio_trigger_recipient":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetTriggerRecipient")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@materializeinc/pulumi-honeycomb",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
