package goal

import "os"

//FileExists check file exists or not.
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

//DirectoryExists check directory exists or not.
func DirectoryExists(path string) bool {
	if st, err := os.Stat(path); err == nil {
		return st.IsDir()
	}

	return false
}
