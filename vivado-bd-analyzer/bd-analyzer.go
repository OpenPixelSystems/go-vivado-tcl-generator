package bd_generator

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
)

type BDInfo struct {
	Design struct {
		DesignInfo struct {
			BoundaryCrc   string `json:"boundary_crc"`
			Device        string `json:"device"`
			Name          string `json:"name"`
			SynthFlowMode string `json:"synth_flow_mode"`
			ToolVersion   string `json:"tool_version"`
			Validated     string `json:"validated"`
		} `json:"design_info"`
	} `json:"design"`
}

type XprInfo struct {
	XMLName       xml.Name `xml:"Project"`
	Text          string   `xml:",chardata"`
	Version       string   `xml:"Version,attr"`
	Minor         string   `xml:"Minor,attr"`
	Path          string   `xml:"Path,attr"`
	DefaultLaunch struct {
		Text string `xml:",chardata"`
		Dir  string `xml:"Dir,attr"`
	} `xml:"DefaultLaunch"`
	Configuration struct {
		Text   string `xml:",chardata"`
		Option []struct {
			Text string `xml:",chardata"`
			Name string `xml:"Name,attr"`
			Val  string `xml:"Val,attr"`
		} `xml:"Option"`
	} `xml:"Configuration"`
}

func FindOptionInXPR(xpr_proj *XprInfo, option string) string {
	for _, opt := range xpr_proj.Configuration.Option {
		if opt.Name == option {
			return opt.Val
		}
	}
	return ""
}

func AnalyzeBDFile(bd_path, xpr_path string) (BDInfo, XprInfo, error) {

	b, err := ioutil.ReadFile(bd_path)
	if err != nil {
		log.Fatal("Failed to read bd file!")
		return BDInfo{}, XprInfo{}, err
	}

	var bd_info BDInfo
	err = json.Unmarshal(b, &bd_info)
	if err != nil {
		log.Fatal(err)
		return BDInfo{}, XprInfo{}, err
	}

	b, err = ioutil.ReadFile(xpr_path)
	if err != nil {
		log.Fatal("Failed to read xpr file!")
		return BDInfo{}, XprInfo{}, err
	}

	var xpr_info XprInfo
	err = xml.Unmarshal(b, &xpr_info)
	if err != nil {
		log.Fatal(err)
		return BDInfo{}, XprInfo{}, err
	}

	return bd_info, xpr_info, nil
}
