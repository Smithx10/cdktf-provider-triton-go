//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TritonProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TritonProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTritonProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTritonProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTritonProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TritonProvider) validateSetInsecureSkipTlsVerifyParameters(val interface{}) error {
	return nil
}

func validateNewTritonProviderParameters(scope constructs.Construct, id *string, config *TritonProviderConfig) error {
	return nil
}

