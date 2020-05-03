package util

// CombineErrors - combine 2 or more errors
func CombineErrors(errs ...error) (err string) {
	var allError string
	for _, err := range errs {
		if err != nil {
			allError = err.Error() + "\n"
		}
	}
	return allError
}
