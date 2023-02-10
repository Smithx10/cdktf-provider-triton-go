package machine


type MachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#create Machine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#delete Machine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#read Machine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#update Machine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

