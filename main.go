package main

import (
	vg "./vivado-generator"
	"fmt"
)

func main() {
	var project = vg.VivadoProj{
		"\"bbr_bd\"",
		"bbr_bd",
		"\"../etf0002-bbr/\"",
		"\"../etf0002-bbr/IPlib/\"",
		"\"em.avnet.com:zed:part0:1.4\"",
		"\"zed\"",
		"\"xc7z020clg484-1\"",
		"\"zynq_sys.tcl\"",
		"\"../etf0002-bbr/sources/bd/\"",
		"\"../etf0002-bbr/sources/hdl/\"",
		"\"../etf0002-bbr/sources/constraints/\"",
		true,
		false,
	}

	fmt.Println("Running Tcl generator")
	vg.GenerateCreateTcl(project)
}
