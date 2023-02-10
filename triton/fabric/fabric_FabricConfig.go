package fabric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FabricConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Network name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#name Fabric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Last assignable IP on the network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#provision_end_ip Fabric#provision_end_ip}
	ProvisionEndIp *string `field:"required" json:"provisionEndIp" yaml:"provisionEndIp"`
	// First IP on the network that can be assigned.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#provision_start_ip Fabric#provision_start_ip}
	ProvisionStartIp *string `field:"required" json:"provisionStartIp" yaml:"provisionStartIp"`
	// CIDR formatted string describing network address space.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#subnet Fabric#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	// VLAN on which the network exists.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#vlan_id Fabric#vlan_id}
	VlanId *float64 `field:"required" json:"vlanId" yaml:"vlanId"`
	// Description of network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#description Fabric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Gateway IP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#gateway Fabric#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#id Fabric#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether or not a NAT zone is provisioned at the Gateway IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#internet_nat Fabric#internet_nat}
	InternetNat interface{} `field:"optional" json:"internetNat" yaml:"internetNat"`
	// List of IP addresses for DNS resolvers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#resolvers Fabric#resolvers}
	Resolvers *[]*string `field:"optional" json:"resolvers" yaml:"resolvers"`
	// Map of CIDR block to Gateway IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/fabric#routes Fabric#routes}
	Routes *map[string]*string `field:"optional" json:"routes" yaml:"routes"`
}

