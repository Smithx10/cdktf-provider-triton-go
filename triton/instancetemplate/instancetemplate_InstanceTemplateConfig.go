package instancetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstanceTemplateConfig struct {
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
	// UUID of the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#image InstanceTemplate#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Package name used for provisioning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#package InstanceTemplate#package}
	Package *string `field:"required" json:"package" yaml:"package"`
	// Friendly name for the instance template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#template_name InstanceTemplate#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// Whether to enable the firewall for group instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#firewall_enabled InstanceTemplate#firewall_enabled}
	FirewallEnabled interface{} `field:"optional" json:"firewallEnabled" yaml:"firewallEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#id InstanceTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata for group instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#metadata InstanceTemplate#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Network IDs for group instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#networks InstanceTemplate#networks}
	Networks *[]*string `field:"optional" json:"networks" yaml:"networks"`
	// Tags for group instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#tags InstanceTemplate#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#timeouts InstanceTemplate#timeouts}
	Timeouts *InstanceTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Data copied to instance on boot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/instance_template#userdata InstanceTemplate#userdata}
	Userdata *string `field:"optional" json:"userdata" yaml:"userdata"`
}

