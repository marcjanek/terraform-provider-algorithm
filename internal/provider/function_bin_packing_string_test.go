// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0
package provider

//
//import (
//	"testing"
//
//	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
//	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
//	"github.com/hashicorp/terraform-plugin-testing/plancheck"
//	"github.com/hashicorp/terraform-plugin-testing/tfversion"
//)
//
//func TestRFC3339Parse_offset(t *testing.T) {
//	resource.UnitTest(t, resource.TestCase{
//		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
//			tfversion.SkipBelow(tfversion.Version1_8_0),
//		},
//		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
//		Steps: []resource.TestStep{
//			{
//				Config: `
//               output "test" {
//                   value = provider::algorithm::bin_packing_string(["dom","ania", "kasia"], 5)
//               }
//               `,
//				ConfigPlanChecks: resource.ConfigPlanChecks{
//					PreApply: []plancheck.PlanCheck{
//						plancheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
//							map[string]knownvalue.Check{
//								"day":          knownvalue.Int64Exact(19),
//								"hour":         knownvalue.Int64Exact(16),
//								"iso_week":     knownvalue.Int64Exact(51),
//								"iso_year":     knownvalue.Int64Exact(1996),
//								"minute":       knownvalue.Int64Exact(39),
//								"month":        knownvalue.Int64Exact(12),
//								"month_name":   knownvalue.StringExact("December"),
//								"second":       knownvalue.Int64Exact(57),
//								"unix":         knownvalue.Int64Exact(851042397),
//								"weekday":      knownvalue.Int64Exact(4),
//								"weekday_name": knownvalue.StringExact("Thursday"),
//								"year":         knownvalue.Int64Exact(1996),
//								"year_day":     knownvalue.Int64Exact(354),
//							},
//						)),
//					},
//				},
//			},
//		},
//	})
//}
