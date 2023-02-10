package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-triton-go/triton/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-triton-go/triton/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/triton triton}.
type TritonProvider interface {
	cdktf.TerraformProvider
	Account() *string
	SetAccount(val *string)
	AccountInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	InsecureSkipTlsVerify() interface{}
	SetInsecureSkipTlsVerify(val interface{})
	InsecureSkipTlsVerifyInput() interface{}
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
	KeyMaterial() *string
	SetKeyMaterial(val *string)
	KeyMaterialInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccount()
	ResetAlias()
	ResetInsecureSkipTlsVerify()
	ResetKeyId()
	ResetKeyMaterial()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUrl()
	ResetUser()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TritonProvider
type jsiiProxy_TritonProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_TritonProvider) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) InsecureSkipTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) InsecureSkipTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) KeyMaterial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyMaterial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) KeyMaterialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyMaterialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TritonProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/triton triton} Resource.
func NewTritonProvider(scope constructs.Construct, id *string, config *TritonProviderConfig) TritonProvider {
	_init_.Initialize()

	if err := validateNewTritonProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TritonProvider{}

	_jsii_.Create(
		"@cdktf/provider-triton.provider.TritonProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/triton triton} Resource.
func NewTritonProvider_Override(t TritonProvider, scope constructs.Construct, id *string, config *TritonProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-triton.provider.TritonProvider",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TritonProvider)SetAccount(val *string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetInsecureSkipTlsVerify(val interface{}) {
	if err := j.validateSetInsecureSkipTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSkipTlsVerify",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetKeyMaterial(val *string) {
	_jsii_.Set(
		j,
		"keyMaterial",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_TritonProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TritonProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTritonProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.provider.TritonProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TritonProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTritonProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.provider.TritonProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TritonProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTritonProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.provider.TritonProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TritonProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-triton.provider.TritonProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TritonProvider) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TritonProvider) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TritonProvider) ResetAccount() {
	_jsii_.InvokeVoid(
		t,
		"resetAccount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		t,
		"resetAlias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetInsecureSkipTlsVerify() {
	_jsii_.InvokeVoid(
		t,
		"resetInsecureSkipTlsVerify",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetKeyMaterial() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyMaterial",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) ResetUser() {
	_jsii_.InvokeVoid(
		t,
		"resetUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TritonProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TritonProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TritonProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TritonProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

