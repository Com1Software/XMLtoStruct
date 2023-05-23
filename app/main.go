package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func RtnXMLTagCount(xdata string, xvar string) int {
	cnt := 0
	svar := "<" + xvar + ">"
	evar := "</" + xvar + ">"
	xon := false
	for i := 0; i < len(xdata); i++ {
		if i < (len(xdata) - len(svar)) {
			if xdata[i:i+len(svar)] == svar {
				xon = true
				i = i + len(svar)
			}
		}
		if i < (len(xdata) - len(evar)) {
			if xon {
				if xdata[i:i+len(evar)] == evar {
					xon = false
					i = i + len(svar)
					cnt++
				}
			}
		}
	}
	return cnt

}

func RtnXMLLevelOneTag(xdata string) string {
	ydata := ""
	xon := false
	t1 := 0
	for i := 0; i < len(xdata); i++ {
		if i < len(xdata)-1 {
			if t1 > 0 {
				if xon {
					if xdata[i:i+1] == ">" {
						xon = false
						i = len(xdata)
					}
				}
				if xon {
					ydata = ydata + xdata[i:i+1]
				}
				if i < len(xdata)-1 {
					if xdata[i:i+1] == "<" {
						xon = true
					}
				}
			}
			if i < len(xdata)-4 {
				if xdata[i:i+5] == "<?xml" {
					t1++
				}
			}

		}
	}
	return ydata
}

func RtnXMLMaxTagDepth(xdata string) int {
	lev := 0
	maxlev := 0
	//	ton := false

	for i := 0; i < len(xdata); i++ {
		if i < len(xdata)-1 {
			if xdata[i:i+1] == "<" {
				lev++
			}
			if xdata[i:i+2] == "</" {
				lev--
			}
			if lev > maxlev {
				maxlev = lev
			}
		}
	}
	return maxlev
}
func RtnXMLTagData(xdata string, xvar string) string {
	ydata := ""
	svar := "<" + xvar + ">"
	evar := "</" + xvar + ">"
	xon := false
	for i := 0; i < len(xdata); i++ {
		if i < (len(xdata) - len(svar)) {
			if xdata[i:i+len(svar)] == svar {
				xon = true
				i = i + len(svar)
			}
		}
		if i < (len(xdata) - len(evar)) {
			if xdata[i:i+len(evar)] == evar {
				xon = false
				i = i + len(evar) - 1
			}
		}
		if xon {
			ydata = ydata + xdata[i:i+1]
		}
	}
	return ydata

}

func main() {
	fmt.Println("Convert XML file to Go Structs")
	xFile := "test.xml"
	xmlFile, err := os.Open(xFile)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	fmt.Println("Successfully Opened xml")
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Printf("Size file %d\n", len(byteValue))
	//fmt.Printf("-- %s\n", xmlFile)
	//fmt.Printf(" -- %s\n", byteValue)
	z := RtnXMLTagCount(string(byteValue), RtnXMLLevelOneTag(string(byteValue)))
	// z := RtnXMLLevelOneTag(string(byteValue))
	// z := RtnXMLLevelOneTag(string(byteValue))
	// z := RtnXMLMaxTagDepth(string(byteValue))
	//	z := RtnXMLTagData(string(byteValue), "users")
	fmt.Println(z)

}
