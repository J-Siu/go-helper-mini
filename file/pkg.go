package file

import (
	"os"
	"path"
	"strings"
)

// Return full path of path provided.
//   - If <workPath> is empty/nil, use current path.
//   - <workPathP> not modified.
//   - Return string pointer.
func FullPath(workPath string) (fullPath string) {
	if workPath == "" {
		// Empty return current path
		fullPath = *CurrentPath()
	} else if workPath[0] == '/' {
		// Path start with / already full path, return it
		fullPath = workPath
	} else {
		// Add to current path
		fullPath = path.Join(*CurrentPath(), workPath)
	}
	return fullPath
}

func IsDir(path string) (result bool) {
	if info, e := os.Stat(path); e == nil {
		return info.IsDir()
	}
	return
}

// Check workPath is regular file
//
// Return false on stat() error
func IsRegularFile(workPath string) (result bool) {
	fileInfo, err := os.Stat(workPath)
	if err == nil {
		result = fileInfo.Mode().IsRegular()
	}
	return result
}

func ReadStr(filePath string) (*string, error) {
	var (
		byteRead []byte
		e        error
		str      string
	)
	byteRead, e = os.ReadFile(filePath)
	if e == nil {
		str = string(byteRead)
	}
	return &str, e
}

// Read file into []string.
//
// Lines are split by "\n".
func ReadStrArray(filePath string) (*[]string, error) {
	var (
		str      *string
		strArray []string
		e        error
	)
	str, e = ReadStr(filePath)
	if e == nil {
		strArray = strings.Split(*str, "\n")
	}
	return &strArray, e
}

// Write *[]byte into file
func WriteByte(filePath string, bP *[]byte, perm os.FileMode) error {
	return os.WriteFile(filePath, *bP, perm)
}

// Write *string into file
func WriteStr(filePath string, str *string, perm os.FileMode) error {
	tmp := []byte(*str)
	return WriteByte(filePath, &tmp, perm)
}

// Write *[]string into file
func WriteStrArray(filePath string, strArray *[]string, perm os.FileMode) error {
	tmp := []byte(strings.Join(*strArray, "\n"))
	return WriteByte(filePath, &tmp, perm)
}

// Get full path of current directory.
//   - Return string pointer.
func CurrentPath() *string {
	p, _ := os.Getwd()
	return &p
}

// Expand Linux `~` and environment variable in string
func TildeEnvExpand(strIn string) (strOut string) {
	if strIn == "~" {
		strOut = "$HOME"
	} else if strings.HasPrefix(strIn, "~/") {
		strOut = "$HOME" + strIn[1:]
	} else {
		strOut = strIn
	}
	return os.ExpandEnv(strOut)
}
