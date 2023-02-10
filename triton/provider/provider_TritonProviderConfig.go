package provider


type TritonProviderConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#account TritonProvider#account}.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#alias TritonProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#insecure_skip_tls_verify TritonProvider#insecure_skip_tls_verify}.
	InsecureSkipTlsVerify interface{} `field:"optional" json:"insecureSkipTlsVerify" yaml:"insecureSkipTlsVerify"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#key_id TritonProvider#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#key_material TritonProvider#key_material}.
	KeyMaterial *string `field:"optional" json:"keyMaterial" yaml:"keyMaterial"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#url TritonProvider#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton#user TritonProvider#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

