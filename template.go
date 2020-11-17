package main

import (
    "io/ioutil"
	"os"
	"text/template"
    "gopkg.in/yaml.v2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Get values from the values file and store it as Values map[string]interface{} field in ValuesFiles Struct
func getValues(filename string)(ValuesFile, error) {
	data,err := ioutil.ReadFile(filename)
	check(err)
	vf := ValuesFile {}
	err = yaml.Unmarshal(data, &vf.Values)
	return vf, err
}

//Get all the template file names from a specific directory and store it in a []string to pass them
// to the rendering function 
func getTemplates(dirname string)(tmplfiles[]string, err error){
	files, err := ioutil.ReadDir(dirname)
	for _, file := range files {
		tmplfiles = append(tmplfiles, file.Name())
    }
	return tmplfiles, err
}

//Renders a template with the values from the Values field of ValuesFile struct and creates a file with
// the result inside a folder called manifests.
func executeTmpl(filename string, dirname string, data ValuesFile){
	t, err := template.ParseFiles("./"+dirname+"/"+filename)
	check(err)
	f, err := os.Create("./manifests/"+filename)
	check(err)
	err = t.Execute(f, data)
	f.Close()
	
}
// This struct is optional and just gives us a field called Values to access it from the templates instead
// of calling each value directly. ex {{ .Values.Name }} instead of {{ .Name }}.
type ValuesFile struct {
	Values map[string]interface{}
}

