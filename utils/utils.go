package utils

import (
	"os"
	"path"
	"regexp"
)

// DoesExist checks if the given path exists
func DoesExist(pth string) bool {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsAbsolute checks if the given path is absolute or not
func IsAbsolute(pth string) bool {
	return path.IsAbs(pth)
}

// IsDir checks if the path is a directory
func IsDir(pth string) bool {
	// Assuming it exists
	stat, _ := os.Stat(pth)
	if stat.IsDir() {
		return true
	}
	return false
}

// IsFile checks if the path is a regular file
func IsFile(pth string) bool {
	// Assuming it exists
	stat, _ := os.Stat(pth)
	mode := stat.Mode()
	if !mode.IsDir() && mode.IsRegular() {
		return true
	}
	return false
}

// IsValidVarName checks if the given string can be accepted as a variable name
func IsValidVarName(s string) bool {
	// A valid variable name can start with either an alpha character or an underscore
	// Other caracters can be alphanumeric or underscore
	regex := regexp.MustCompile(`^[A-Za-z_][0-9A-Za-z_]*$`)
	return regex.MatchString(s)
}

// IsValidName checks if the given string includes only alphanumeric and underscore
func IsValidName(s string) bool {
	regex := regexp.MustCompile(`^[0-9A-Za-z_]*$`)
	return regex.MatchString(s)
}