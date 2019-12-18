package main

import (
	ana "./vivado-bd-analyzer"
	vg "./vivado-generator"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func findFileInPath(path, filename string) (string, error) {
	fileFound := ""
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Name() == filename {
			fileFound = path
			return nil
		}
		return nil
	})
	if err != nil || fileFound == "" {
		log.Fatal("Error!")
		return "", err
	}
	return fileFound, nil
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 6 {
		log.Fatal("Usage: go-vivado-tcl-generate <absolute project dir> <xpr_file_name> <bd_file_name> <bd relative path> xdc relative path> <hdl relative path>")
		return
	}

	for i, arg := range argsWithoutProg {
		argsWithoutProg[i] = strings.Replace(arg, "\"", "", -1)
	}
	fmt.Println(argsWithoutProg)

	source_dir := argsWithoutProg[0]
	xpr_file := argsWithoutProg[1]
	bd_file := argsWithoutProg[2]

	xpr_file_path, err := findFileInPath(source_dir, xpr_file)
	bd_file_path, err := findFileInPath(source_dir, bd_file)

	bd_path := source_dir + "/" + argsWithoutProg[3]
	if _, err := os.Stat(bd_path); os.IsNotExist(err) {
		log.Fatal("BD path not found!")
		return
	}

	xdr_path := source_dir + "/" + argsWithoutProg[4]
	if _, err := os.Stat(xdr_path); os.IsNotExist(err) {
		log.Fatal("Constrains path not found!")
		return
	}

	local_hdl := false
	hdl_path := source_dir + "/" + argsWithoutProg[5]
	if _, err := os.Stat(hdl_path); !os.IsNotExist(err) {
		log.Fatal("HDL path found!")
		local_hdl = true
	}

	bd, xpr, err := ana.AnalyzeBDFile(bd_file_path, xpr_file_path)
	if err != nil {
		log.Fatal("Failed to analyze project")
		return
	}

	proj_name := bd.Design.DesignInfo.Name
	proj_name_esc := "\"" + proj_name + "\""

	part := ana.FindOptionInXPR(&xpr, "Part")
	if part != bd.Design.DesignInfo.Device {
		fmt.Println("Device diffent between BD and XPR!")
		fmt.Println(part, bd.Design.DesignInfo.Device)
		return
	}

	var project vg.VivadoProj

	project.Project_name = proj_name_esc
	project.Project_name_raw = proj_name
	project.Source_directory = source_dir

	ip_dir_tmp := strings.Replace(ana.FindOptionInXPR(&xpr, "IPRepoPath"), "$PPRDIR/../", "", -1)
	project.Ip_repository_directory = source_dir + "/" + ip_dir_tmp

	project.Board_name = ana.FindOptionInXPR(&xpr, "BoardPart")
	project.Board_id = ana.FindOptionInXPR(&xpr, "DSABoardId")
	project.Partnum = part
	if project.Board_id != "" {
		project.Board_id_used = true
	}

	project.Top_tcl_name = "\"zynq_sys.tcl\""
	/* TODO: Append with / if needed! */
	project.Orig_bd_path = bd_path
	project.Orig_hdl_path = hdl_path
	project.Orig_xdr_path = xdr_path
	project.Local_hdl_files = local_hdl

	fmt.Println(project)

	fmt.Println("Running Tcl generator")
	vg.GenerateCreateTcl(project)
}
