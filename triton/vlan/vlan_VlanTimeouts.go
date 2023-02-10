package vlan


type VlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#create Vlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#delete Vlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#read Vlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/vlan#update Vlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

