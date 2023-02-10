//go:build no_runtime_type_checking

package snapshot

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Snapshot) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snapshot) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSnapshot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSnapshot_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSnapshot_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetMachineIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Snapshot) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewSnapshotParameters(scope constructs.Construct, id *string, config *SnapshotConfig) error {
	return nil
}

