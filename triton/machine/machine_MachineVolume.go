package machine


type MachineVolume struct {
	// Where to attach the volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#mountpoint Machine#mountpoint}
	Mountpoint *string `field:"required" json:"mountpoint" yaml:"mountpoint"`
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#name Machine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The volume attachment mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#mode Machine#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The type of volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#type Machine#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

