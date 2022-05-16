package main

import (
	_ "embed"

	"github.com/knowbe4/tflint-ruleset-kb4/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

/** @todo
 * These rules still need to be written:
 * For Modules
 *  - Resources should be named `this` where possible.
 *  - No providers in modules (this can be ignored on a module by module basis if needed)
 * For all terraform
 *  - The following checks need tests:
 *    - checkProviders
 *    - checkTerraformBlock
 *    - checkLocals
 *    - checkTerraformRemoteState
 */

//go:embed VERSION
var VERSION string

func main() {

	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "template",
			Version: VERSION,
			Rules: []tflint.Rule{
				rules.NewTerraformValidatedVariablesRule(),
				rules.NewTerraformKb4FileStructureRule(),
			},
		},
	})
}
