// Private functions that helps Ezcli.
package ezcli

func stringSliceContains(slice []string, str string) bool {
	for _, i := range slice {
		if i == str {
			return true
		}
	}

	return false
}
