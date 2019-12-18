# Vivado TCL Generator

The goal of this go program is to generate TCL scripts which can be used by jenkins for example to build a Vivado project in the background.
Currently TCL scripts are build for:
- Creating a project
- Running the synthesis
- Running the implementation
- Generating devicetree's from the Vivado suite

## Running
```bash
$ go build -o create_proj
$ ./create_proj <Firmware dir> <XPR Filename> <BD Filename> <BD Relative path> <XDC Relative path> <HDL Relative path>
```
Path's are relative to the Firmware directory


## Adding auto generated files

To add a template file, one can just add a file <name>-template to the templates directory.
The template files use the Go template file library (https://golang.org/pkg/text/template/)
