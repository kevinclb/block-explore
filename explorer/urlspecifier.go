package explorer

//urlSpecifier takes a base url and appends a path to it
//depending on the calling function's request
func urlSpecifier(urlPath string) string {
	defaultURL := coinbaseURL + urlPath
	return defaultURL
}
