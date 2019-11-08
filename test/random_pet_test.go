package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformBasicExample(t *testing.T) {

	expPetNameLength := 3

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"pet_name_length": expPetNameLength,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	_, err := terraform.InitAndApplyE(t, terraformOptions)

	if err != nil {
		t.Fatalf("terraform init/apply failed: %v", err)
	}

	petName := terraform.Output(t, terraformOptions, "pet_name")
	petNameLength := len(strings.Split(petName, "-"))

	if expPetNameLength != petNameLength {
		t.Logf("output pet_name has wrong length, want: %v, got: %v", expPetNameLength, petNameLength)
		t.Fail()
	}
}
