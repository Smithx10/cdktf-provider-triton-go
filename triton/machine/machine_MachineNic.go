package machine


type MachineNic struct {
	// ID of the network to which the NIC is attached.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/machine#network Machine#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

