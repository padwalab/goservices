package testsvc

// SomeOperation performs specific action on request data
func SomeOperation(in string) (TransformData, error) {
	return TransformData{
		Name:    in,
		Address: in + "'s Address",
	}, nil
}
