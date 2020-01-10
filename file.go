package goal

import "os"

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func DirectoryExists(path string) bool {
	if st, err := os.Stat(path); err == nil {
		return st.IsDir()
	}

	return false
}
