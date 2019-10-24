package vault

import "regexp"

// FilterDataPaths filters the provided paths based on a regex, this only filters the data paths
func FilterDataPaths(paths []string, r *regexp.Regexp) (filtered []string) {
	for _, path := range paths {
		lastChar := path[len(path)-1:]
		if lastChar != "/" && r.MatchString(path) {
			filtered = append(filtered, path)
		}
	}
	return
}

// FilterOnlyDataPaths
func FilterOnlyDataPaths(paths []string) (filtered []string) {
	r, _ := regexp.Compile(".*")
	return FilterDataPaths(paths, r)
}
