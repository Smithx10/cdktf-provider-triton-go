//go:build no_runtime_type_checking

package machine

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MachineVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MachineVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MachineVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MachineVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MachineVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MachineVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMachineVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

