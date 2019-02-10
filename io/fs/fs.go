package fs

import "os"

// FileExists checks if a file exists at the given path, the function will return an error if the file exists, but is not accessible
// otherwise it will return a true or false without any errors
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		// the file may or may not exists, but we do not have access to get file stats
		if os.IsNotExist(err) { // check the error to see if it's because the file doesn't exist
			// the file doesn't exist so we can return a false flag with no error
			return false, nil
		}
		// the file exists but we don't have permission to access it or there
		// is some other error associated with the fil
		return false, err
	}

	// if we get to here, then it's because the file exists and we have permission to access it
	return true, nil
}
