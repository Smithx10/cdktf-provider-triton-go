package vlan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Unique name to identify VLAN.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#name Vlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Number between 0-4095 indicating VLAN ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#vlan_id Vlan#vlan_id}
	VlanId *float64 `field:"required" json:"vlanId" yaml:"vlanId"`
	// Description of the VLAN.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#description Vlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#id Vlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#timeouts Vlan#timeouts}
	Timeouts *VlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

