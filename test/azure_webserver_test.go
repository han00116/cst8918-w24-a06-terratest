package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "f297c7ae-463c-406d-9c62-7eb1d616382a"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "han00116",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))
}

func GetVirtualMachineImage(t testing.T, vmName string, resGroupName string, subscriptionID string) VMImage {
	vmImage, err := GetVirtualMachineImageE(vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return vmImage
}

func GetVirtualMachineNics(t testing.T, vmName string, resGroupName string, subscriptionID string) []string {
	nicList, err := GetVirtualMachineNicsE(vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return nicList
}
