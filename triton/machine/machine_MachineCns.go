package machine


type MachineCns struct {
	// Disable CNS for this instance (after create).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#disable Machine#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// Assign CNS service names to this instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#services Machine#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

