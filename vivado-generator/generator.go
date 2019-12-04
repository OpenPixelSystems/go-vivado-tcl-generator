package vivado_generator

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type VivadoProj struct {
	Project_name     string
	Project_name_raw string

	Source_directory        string
	Ip_repository_directory string

	Board_name string
	Board_id   string
	Partnum    string

	Top_tcl_name string

	Orig_bd_path  string
	Orig_hdl_path string
	Orig_xdr_path string

	Board_id_used   bool
	Local_hdl_files bool
}

func GenerateCreateTcl(project VivadoProj) error {
	log.Println("Project settings")
	log.Print(project)

	files, err := ioutil.ReadDir("./vivado-generator/templates")
	if err != nil {
		log.Fatal("Failed to read template dir")
		return err
	}
	os.Mkdir("./output", 0755)

	for _, f := range files {
		var t_out bytes.Buffer
		outName := strings.Replace(f.Name(), "-template", "", -1)
		outName = "./output/" + outName

		b, err := ioutil.ReadFile("./vivado-generator/templates/" + f.Name())
		if err != nil {
			log.Fatal("Failed to read create template")
			return err
		}
		template_str := string(b)

		t := template.Must(template.New(outName).Parse(template_str))
		err = t.Execute(&t_out, project)
		if err != nil {
			log.Println("executing template:", err)
			return err
		}
		err = ioutil.WriteFile(outName, t_out.Bytes(), 0644)

	}

	return nil
}
