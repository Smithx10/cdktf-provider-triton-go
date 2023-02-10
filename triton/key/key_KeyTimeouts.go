package key


type KeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/key#create Key#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/key#delete Key#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/key#read Key#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/key#update Key#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

