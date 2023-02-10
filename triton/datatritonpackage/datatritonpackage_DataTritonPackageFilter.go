package datatritonpackage


type DataTritonPackageFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#disk DataTritonPackage#disk}.
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#group DataTritonPackage#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#lwps DataTritonPackage#lwps}.
	Lwps *float64 `field:"optional" json:"lwps" yaml:"lwps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#memory DataTritonPackage#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#name DataTritonPackage#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#swap DataTritonPackage#swap}.
	Swap *float64 `field:"optional" json:"swap" yaml:"swap"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#vcpus DataTritonPackage#vcpus}.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/d/package#version DataTritonPackage#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

