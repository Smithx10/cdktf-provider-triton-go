package servicegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceGroupConfig struct {
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
	// Friendly name for the service group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/service_group#group_name ServiceGroup#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Identifier of an instance template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/service_group#template ServiceGroup#template}
	Template *string `field:"required" json:"template" yaml:"template"`
	// Number of instances to launch and monitor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/service_group#capacity ServiceGroup#capacity}
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/service_group#id ServiceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/service_group#timeouts ServiceGroup#timeouts}
	Timeouts *ServiceGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

