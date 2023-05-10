package datatritonvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTritonVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#filesystem_path DataTritonVolume#filesystem_path}.
	FilesystemPath *string `field:"optional" json:"filesystemPath" yaml:"filesystemPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#id DataTritonVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#name DataTritonVolume#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#size DataTritonVolume#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#state DataTritonVolume#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#tags DataTritonVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Type of volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/volume#type DataTritonVolume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

