package rules

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var EXPECTED_FILES []string = []string{"_init.tf", "_variables.tf", "_outputs.tf", "_locals.tf"}

// TerraformKb4FileStructureRule checks whether modules adhere to Terraform's standard module structure
type TerraformKb4FileStructureRule struct {
	tflint.DefaultRule
}

// NewTerraformKb4ModuleStructureRule returns a new rule
func NewTerraformKb4FileStructureRule() *TerraformKb4FileStructureRule {
	return &TerraformKb4FileStructureRule{}
}

// Name returns the rule name
func (r *TerraformKb4FileStructureRule) Name() string {
	return "terraform_kb4_module_structure"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformKb4FileStructureRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformKb4FileStructureRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *TerraformKb4FileStructureRule) Link() string {
	return "https://engineering.internal.knowbe4.com/tech-stack/terraform/style-guide/#standard-files-names-and-usage"
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (r *TerraformKb4FileStructureRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	r.checkFiles(runner)
	r.checkVariables(runner)
	r.checkOutputs(runner)
	r.checkProviders(runner)
	r.checkTerraformBlock(runner)
	r.checkLocals(runner)
	r.checkTerraformRemoteState(runner)

	return nil
}

func (r *TerraformKb4FileStructureRule) checkFiles(runner tflint.Runner) error {
	files, err := runner.GetFiles()

	if err != nil {
		return err
	}

	for _, name := range EXPECTED_FILES {
		if files[name] == nil {
			runner.EmitIssue(
				r,
				fmt.Sprintf("Module should include a %s file.", name),
				hcl.Range{
					Filename: name,
					Start:    hcl.InitialPos,
				},
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkVariables(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, variable := range content.Blocks {
		if variable.DefRange.Filename != "_variables.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("variable %q should be moved from %s to %s", variable.Labels[0], variable.DefRange.Filename, "_variables.tf"),
				variable.DefRange,
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkOutputs(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "output",
				LabelNames: []string{"name"},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, output := range content.Blocks {
		if output.DefRange.Filename != "_outputs.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("output %q should be moved from %s to %s", output.Labels[0], output.DefRange.Filename, "_outputs.tf"),
				output.DefRange,
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkProviders(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "provider",
				LabelNames: []string{"name"},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, provider := range content.Blocks {
		if provider.DefRange.Filename != "_init.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("provider %q should be moved from %s to %s", provider.Labels[0], provider.DefRange.Filename, "_init.tf"),
				provider.DefRange,
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkTerraformBlock(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "terraform",
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, terraformBlock := range content.Blocks {
		if terraformBlock.DefRange.Filename != "_init.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("terraform block %q should be moved from %s to %s", terraformBlock.Labels[0], terraformBlock.DefRange.Filename, "_init.tf"),
				terraformBlock.DefRange,
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkLocals(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "locals",
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, locals := range content.Blocks {
		if locals.DefRange.Filename != "_init.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("locals block %q should be moved from %s to %s", locals.Labels[0], locals.DefRange.Filename, "_init.tf"),
				locals.DefRange,
			)
		}
	}

	return nil
}

func (r *TerraformKb4FileStructureRule) checkTerraformRemoteState(runner tflint.Runner) error {

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "data",
				LabelNames: []string{"name"},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, data := range content.Blocks {
		if data.DefRange.Filename != "_init.tf" {
			if data.Type == "terraform_remote_state" {
				runner.EmitIssue(
					r,
					fmt.Sprintf("data terraform_remote_state %q should be moved from %s to %s", data.Name, data.DefRange.Filename, "_init.tf"),
					data.DefRange,
				)
			}
		}
	}

	return nil
}
