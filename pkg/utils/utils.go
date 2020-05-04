package utils

import "strings"

// just check if a path end with /
func EndsWithSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

// just check if a path end with /
func DelInitialSlash(path string) string {
	if path[0] == '/' {
		return path[1:]
	}
	return path
}