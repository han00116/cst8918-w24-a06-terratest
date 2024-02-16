package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type VMImage struct {
	Publisher string
	Offer     string
	SKU       string
	Version   string
}

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "f297c7ae-463c-406d-9c62-7eb1d616382a"

func GetVirtualMachineImageE(vmName string, resGroupName string, subscriptionID string) (VMImage, error) {
	// Implement this function to interact with Azure and retrieve VM image information
	return VMImage{}, nil
}

func GetVirtualMachineNicsE(vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	// Implement this function to interact with Azure and retrieve VM NIC information
	return []string{}, nil
}

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // Path to Terraform code
		Vars: map[string]interface{}{
			"labelPrefix": "han00116",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	// Get VM image
	vmImage, err := GetVirtualMachineImageE(vmName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	assert.NotNil(t, vmImage)

	// Get VM NICs
	nicList, err := GetVirtualMachineNicsE(vmName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	assert.NotNil(t, nicList)
}
