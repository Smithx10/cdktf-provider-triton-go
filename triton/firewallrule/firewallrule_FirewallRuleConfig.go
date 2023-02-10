package firewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallRuleConfig struct {
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
	// firewall rule text.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/firewall_rule#rule FirewallRule#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Human-readable description of the rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/firewall_rule#description FirewallRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/firewall_rule#enabled FirewallRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/triton/r/firewall_rule#id FirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

