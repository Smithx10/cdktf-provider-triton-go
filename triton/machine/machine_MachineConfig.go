package machine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// UUID of the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#image Machine#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// The package for use for provisioning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#package Machine#package}
	Package *string `field:"required" json:"package" yaml:"package"`
	// Administrator's initial password (Windows only).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#administrator_pw Machine#administrator_pw}
	AdministratorPw *string `field:"optional" json:"administratorPw" yaml:"administratorPw"`
	// Label based affinity rules for assisting instance placement.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#affinity Machine#affinity}
	Affinity *[]*string `field:"optional" json:"affinity" yaml:"affinity"`
	// copied to machine on boot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#cloud_config Machine#cloud_config}
	CloudConfig *string `field:"optional" json:"cloudConfig" yaml:"cloudConfig"`
	// cns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#cns Machine#cns}
	Cns *MachineCns `field:"optional" json:"cns" yaml:"cns"`
	// Whether to create a delegate dataset for this machine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#delegate_dataset Machine#delegate_dataset}
	DelegateDataset interface{} `field:"optional" json:"delegateDataset" yaml:"delegateDataset"`
	// Whether to enable deletion protection for this machine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#deletion_protection_enabled Machine#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Whether to enable the firewall for this machine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#firewall_enabled Machine#firewall_enabled}
	FirewallEnabled interface{} `field:"optional" json:"firewallEnabled" yaml:"firewallEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#id Machine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// locality block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#locality Machine#locality}
	Locality *MachineLocality `field:"optional" json:"locality" yaml:"locality"`
	// Machine metadata.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#metadata Machine#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Friendly name for machine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#name Machine#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Desired network IDs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#networks Machine#networks}
	Networks *[]*string `field:"optional" json:"networks" yaml:"networks"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#nic Machine#nic}
	Nic interface{} `field:"optional" json:"nic" yaml:"nic"`
	// Authorized keys for the root user on this machine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#root_authorized_keys Machine#root_authorized_keys}
	RootAuthorizedKeys *string `field:"optional" json:"rootAuthorizedKeys" yaml:"rootAuthorizedKeys"`
	// Machine tags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#tags Machine#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#timeouts Machine#timeouts}
	Timeouts *MachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Data copied to machine on boot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#user_data Machine#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// User script to run on boot (every boot on SmartMachines).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#user_script Machine#user_script}
	UserScript *string `field:"optional" json:"userScript" yaml:"userScript"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#volume Machine#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

