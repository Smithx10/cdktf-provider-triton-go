package machine


type MachineLocality struct {
	// UUIDs of other instances to attempt to provision alongside.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#close_to Machine#close_to}
	CloseTo *[]*string `field:"optional" json:"closeTo" yaml:"closeTo"`
	// UUIDs of other instances to attempt not to provision alongside.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#far_from Machine#far_from}
	FarFrom *[]*string `field:"optional" json:"farFrom" yaml:"farFrom"`
}

