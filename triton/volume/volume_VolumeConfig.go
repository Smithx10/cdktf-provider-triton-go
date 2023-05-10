package volume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#id Volume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Friendly name for volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#name Volume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Desired network IDs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#networks Volume#networks}
	Networks *[]*string `field:"optional" json:"networks" yaml:"networks"`
	// The size of the volume (Mb).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#size Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Volume tags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#tags Volume#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#timeouts Volume#timeouts}
	Timeouts *VolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Type of volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/volume#type Volume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

