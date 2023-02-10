//go:build no_runtime_type_checking

package machine

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MachineNicList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MachineNicList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MachineNicList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MachineNicList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MachineNicList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MachineNicList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMachineNicListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

