package machine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/Smithx10/cdktf-provider-triton-go/triton/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/Smithx10/cdktf-provider-triton-go/triton/machine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/triton/r/machine triton_machine}.
type Machine interface {
	cdktf.TerraformResource
	AdministratorPw() *string
	SetAdministratorPw(val *string)
	AdministratorPwInput() *string
	Affinity() *[]*string
	SetAffinity(val *[]*string)
	AffinityInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudConfig() *string
	SetCloudConfig(val *string)
	CloudConfigInput() *string
	Cns() MachineCnsOutputReference
	CnsInput() *MachineCns
	ComputeNode() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Created() *string
	Dataset() *string
	DelegateDataset() interface{}
	SetDelegateDataset(val interface{})
	DelegateDatasetInput() interface{}
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disk() *float64
	DomainNames() *[]*string
	FirewallEnabled() interface{}
	SetFirewallEnabled(val interface{})
	FirewallEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	Ips() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locality() MachineLocalityOutputReference
	LocalityInput() *MachineLocality
	Memory() *float64
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Networks() *[]*string
	SetNetworks(val *[]*string)
	NetworksInput() *[]*string
	Nic() MachineNicList
	NicInput() interface{}
	// The tree node.
	Node() constructs.Node
	Package() *string
	SetPackage(val *string)
	PackageInput() *string
	Primaryip() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RootAuthorizedKeys() *string
	SetRootAuthorizedKeys(val *string)
	RootAuthorizedKeysInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	Updated() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	UserScript() *string
	SetUserScript(val *string)
	UserScriptInput() *string
	Volume() MachineVolumeList
	VolumeInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCns(value *MachineCns)
	PutLocality(value *MachineLocality)
	PutNic(value interface{})
	PutTimeouts(value *MachineTimeouts)
	PutVolume(value interface{})
	ResetAdministratorPw()
	ResetAffinity()
	ResetCloudConfig()
	ResetCns()
	ResetDelegateDataset()
	ResetDeletionProtectionEnabled()
	ResetFirewallEnabled()
	ResetId()
	ResetLocality()
	ResetMetadata()
	ResetName()
	ResetNetworks()
	ResetNic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRootAuthorizedKeys()
	ResetTags()
	ResetTimeouts()
	ResetUserData()
	ResetUserScript()
	ResetVolume()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Machine
type jsiiProxy_Machine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Machine) AdministratorPw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) AdministratorPwInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPwInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Affinity() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) AffinityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) CloudConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) CloudConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Cns() MachineCnsOutputReference {
	var returns MachineCnsOutputReference
	_jsii_.Get(
		j,
		"cns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) CnsInput() *MachineCns {
	var returns *MachineCns
	_jsii_.Get(
		j,
		"cnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) ComputeNode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DelegateDataset() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delegateDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DelegateDatasetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delegateDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Disk() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) FirewallEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) FirewallEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Locality() MachineLocalityOutputReference {
	var returns MachineLocalityOutputReference
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) LocalityInput() *MachineLocality {
	var returns *MachineLocality
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Networks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) NetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Nic() MachineNicList {
	var returns MachineNicList
	_jsii_.Get(
		j,
		"nic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) NicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Package() *string {
	var returns *string
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) PackageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Primaryip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) RootAuthorizedKeys() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAuthorizedKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) RootAuthorizedKeysInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAuthorizedKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Timeouts() MachineTimeoutsOutputReference {
	var returns MachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Updated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) UserScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) UserScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) Volume() MachineVolumeList {
	var returns MachineVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Machine) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/triton/r/machine triton_machine} Resource.
func NewMachine(scope constructs.Construct, id *string, config *MachineConfig) Machine {
	_init_.Initialize()

	if err := validateNewMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Machine{}

	_jsii_.Create(
		"@cdktf/provider-triton.machine.Machine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/triton/r/machine triton_machine} Resource.
func NewMachine_Override(m Machine, scope constructs.Construct, id *string, config *MachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-triton.machine.Machine",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Machine)SetAdministratorPw(val *string) {
	if err := j.validateSetAdministratorPwParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPw",
		val,
	)
}

func (j *jsiiProxy_Machine)SetAffinity(val *[]*string) {
	if err := j.validateSetAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinity",
		val,
	)
}

func (j *jsiiProxy_Machine)SetCloudConfig(val *string) {
	if err := j.validateSetCloudConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudConfig",
		val,
	)
}

func (j *jsiiProxy_Machine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Machine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Machine)SetDelegateDataset(val interface{}) {
	if err := j.validateSetDelegateDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegateDataset",
		val,
	)
}

func (j *jsiiProxy_Machine)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_Machine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Machine)SetFirewallEnabled(val interface{}) {
	if err := j.validateSetFirewallEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallEnabled",
		val,
	)
}

func (j *jsiiProxy_Machine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Machine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Machine)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_Machine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Machine)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Machine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Machine)SetNetworks(val *[]*string) {
	if err := j.validateSetNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networks",
		val,
	)
}

func (j *jsiiProxy_Machine)SetPackage(val *string) {
	if err := j.validateSetPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"package",
		val,
	)
}

func (j *jsiiProxy_Machine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Machine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Machine)SetRootAuthorizedKeys(val *string) {
	if err := j.validateSetRootAuthorizedKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootAuthorizedKeys",
		val,
	)
}

func (j *jsiiProxy_Machine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Machine)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_Machine)SetUserScript(val *string) {
	if err := j.validateSetUserScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userScript",
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
func Machine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.machine.Machine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Machine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.machine.Machine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Machine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-triton.machine.Machine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Machine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-triton.machine.Machine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Machine) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Machine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Machine) PutCns(value *MachineCns) {
	if err := m.validatePutCnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCns",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Machine) PutLocality(value *MachineLocality) {
	if err := m.validatePutLocalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLocality",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Machine) PutNic(value interface{}) {
	if err := m.validatePutNicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNic",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Machine) PutTimeouts(value *MachineTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Machine) PutVolume(value interface{}) {
	if err := m.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVolume",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Machine) ResetAdministratorPw() {
	_jsii_.InvokeVoid(
		m,
		"resetAdministratorPw",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetAffinity() {
	_jsii_.InvokeVoid(
		m,
		"resetAffinity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetCloudConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetCns() {
	_jsii_.InvokeVoid(
		m,
		"resetCns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetDelegateDataset() {
	_jsii_.InvokeVoid(
		m,
		"resetDelegateDataset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetFirewallEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetFirewallEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetLocality() {
	_jsii_.InvokeVoid(
		m,
		"resetLocality",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetNetworks() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworks",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetNic() {
	_jsii_.InvokeVoid(
		m,
		"resetNic",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetRootAuthorizedKeys() {
	_jsii_.InvokeVoid(
		m,
		"resetRootAuthorizedKeys",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetUserData() {
	_jsii_.InvokeVoid(
		m,
		"resetUserData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetUserScript() {
	_jsii_.InvokeVoid(
		m,
		"resetUserScript",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) ResetVolume() {
	_jsii_.InvokeVoid(
		m,
		"resetVolume",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Machine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Machine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

