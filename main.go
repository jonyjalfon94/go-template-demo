package main

import (
	"os"
)
func main() {

	// Directory where the templates are stored, it has to be in the same directory as this script
	var dirname = "sample-templates"

	// Name of the values file, also has to be in the same directory as the script
	vf,err := getValues("values.yaml")
	if err != nil {
		return
	}
	// Get the names of the files stored in the templates folder
	tmplfiles, err := getTemplates(dirname)

	// Create the manifests folder if not exists
	if _, err := os.Stat("./manifests"); os.IsNotExist(err) {
		os.Mkdir("./manifests", 0700)
	}

	// Iterate over the list of filenames and render each one of the templates and store them in the manifests folder
	for _, tmplfile := range tmplfiles {
		executeTmpl(tmplfile,dirname, vf)
	}
}