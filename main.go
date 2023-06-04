package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func RtnXMLStructs(xdata string) string {
	bld := []xml{}
	blda := xml{}
	blda.tag = "test"
	bld = append(bld, blda)
	ydata := ""

	z := RtnXMLMaxTagDepth(xdata, 0)
	for i := 1; i < z+1; i++ {
		ii := RtnXMLMaxTagDepth(xdata, i)

		fmt.Printf("  Level = %d Max Depth = %d\n", i, ii)
		for iii := 1; iii < ii+1; iii++ {
			tag, pltag := RtnXMLItemName(xdata, i, iii)
			fmt.Printf(" Item Name = [%s - %s]  Level: %d Position: %d \n ", tag, pltag, i, iii)
		}
	}

	return ydata

}

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

func RtnXMLMaxTagDepth(xdata string, xlev int) int {
	lev := 0
	levcnt := 0
	maxlev := 0
	son := false
	eon := false
	hon := 0
	for i := 0; i < len(xdata); i++ {
		if i < len(xdata)-1 {
			switch {
			case xdata[i:i+2] == "<?" && hon == 0:
				hon = 1
			case xdata[i:i+2] == "?>" && hon == 1:
				hon = 2
			case xdata[i:i+1] == ">" && eon == true && hon == 2:
				eon = false
				//             fmt.Printf("Off Out %d\n", lev)
			case xdata[i:i+2] == "</" && eon == false && hon == 2:
				lev--
				eon = true
				//				fmt.Printf("On Out %d\n", lev)
			case xdata[i:i+1] == ">" && son == true && hon == 2:
				son = false

				//				fmt.Printf("Off In %d\n", lev)
			case xdata[i:i+1] == "<" && son == false && hon == 2:
				lev++
				son = true
				if lev == xlev {
					levcnt++
				}
				//				fmt.Printf("On In %d\n", lev)
			}
			if lev > maxlev && hon == 2 {
				maxlev = lev
			}
		}
	}
	if xlev == 0 {
		return maxlev
	} else {
		return levcnt
	}

}

func RtnXMLItemName(xdata string, xlev int, xpos int) (string, string) {
	tag := ""
	pltag := ""
	lev := 0
	levcnt := 0
	son := false
	eon := false
	hon := 0
	xon := false
	pxon := false
	found := false
	for i := 0; i < len(xdata); i++ {
		if i < len(xdata)-1 {
			switch {
			case xdata[i:i+2] == "<?" && hon == 0:
				hon = 1
			case xdata[i:i+2] == "?>" && hon == 1:
				hon = 2
			case xdata[i:i+1] == ">" && eon == true && hon == 2:
				eon = false
				//				fmt.Printf("Off Out %d\n", lev)
			case xdata[i:i+2] == "</" && eon == false && hon == 2:
				lev--
				eon = true
				//				fmt.Printf("On Out %d\n", lev)
			case xdata[i:i+1] == ">" && son == true && hon == 2:
				son = false
				if xon {
					found = true
				}
				xon = false
				pxon = false
				//				fmt.Printf("Off In %d\n", lev)
			case xdata[i:i+1] == "<" && son == false && hon == 2:
				lev++
				son = true
				if lev == xlev {
					levcnt++
					if levcnt == xpos {
						xon = true
						i++

					}

				}
				if lev == xlev-1 {
					if found == false {
						pxon = true
						pltag = ""
						i++

					}
				}
				//				fmt.Printf("On In %d\n", lev)
			}
			if xon {
				tag = tag + xdata[i:i+1]
			}
			if pxon {
				pltag = pltag + xdata[i:i+1]
			}

		}
	}
	return tag, pltag

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

type xml struct {
	structure string
	tag       string
	level     int
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func BuildApp(xFile string) {
	bld := []xml{}
	blda := xml{}
	blda.tag = "test"
	bld = append(bld, blda)

	xmlFile, err := os.Open(xFile)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	fmt.Println("Successfully Opened xml")
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Printf("Size file %d\n", len(byteValue))
	//------------------------------------------------------------------------------
	//fmt.Printf("-- %s\n", xmlFile)
	//fmt.Printf(" -- %s\n", byteValue)
	// z := RtnXMLTagCount(string(byteValue), RtnXMLLevelOneTag(string(byteValue)))
	// z := RtnXMLLevelOneTag(string(byteValue))
	// z := RtnXMLLevelOneTag(string(byteValue))
	//	z := RtnXMLTagData(string(byteValue), "users")
	//	z := RtnXMLMaxTagDepth(string(byteValue), 0)
	//	zz, plt := RtnXMLItemName(string(byteValue), 4, 1)
	//	fmt.Printf("------ > max %d level item [%s] PLT %s\n", z, zz, plt)
	//-----------------------------------------------------------------------------
	z := RtnXMLMaxTagDepth(string(byteValue), 0)
	for i := 1; i < z+1; i++ {
		ii := RtnXMLMaxTagDepth(string(byteValue), i)

		fmt.Printf("  Level = %d Max Depth = %d\n", i, ii)
		for iii := 1; iii < ii+1; iii++ {
			tag, pltag := RtnXMLItemName(string(byteValue), i, iii)
			fmt.Printf(" Item Name = [%s - %s]  Level: %d Position: %d \n ", tag, pltag, i, iii)
		}
	}
	//-----------------------------------------------------------------------------

	// fmt.Printf("structs \n--------- \n %s \n--------------\n", RtnXMLStructs(string(byteValue)))

	//-----------------------------------------------------------------------------
	xdata := "package main\n\n"
	xdata = xdata + "import (\n"
	xdata = xdata + fmt.Sprintf("     %q\n", "fmt")
	xdata = xdata + fmt.Sprintf("     %q\n", "os")
	xdata = xdata + fmt.Sprintf("     %q\n", "io/ioutil")
	xdata = xdata + fmt.Sprintf("     %q\n", "encoding/xml")

	xdata = xdata + ")\n\n"
	//---------------------------------------------------------------------------------------
	xdata = xdata + "type x" + capitalize(RtnXMLLevelOneTag(string(byteValue))) + " struct {\n"
	xdata = xdata + "   XMLName xml.Name "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "users")
	xdata = xdata + "	Users   []User "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "user")
	xdata = xdata + "}\n\n"

	xdata = xdata + "type Users struct {\n"
	xdata = xdata + "   XMLName xml.Name "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "users")
	xdata = xdata + "	Users   []User "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "user")
	xdata = xdata + "}\n\n"
	//---------------------------------------------------------------------------------------
	xdata = xdata + "type User struct {\n"
	xdata = xdata + "   XMLName xml.Name "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "user")
	xdata = xdata + "	 Name   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "name")
	xdata = xdata + "	 Address   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "address")
	xdata = xdata + "	 City   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "city")
	xdata = xdata + "	 State   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "state")
	xdata = xdata + "	 Teritory   Teritory "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "teritory")
	xdata = xdata + "	 Product   Product "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "product")
	xdata = xdata + "}\n\n"

	xdata = xdata + "type Teritory struct {\n"
	xdata = xdata + "   XMLName xml.Name "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "teritory")
	xdata = xdata + "	 Location   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "location")
	xdata = xdata + "	 Contact   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "contact")
	xdata = xdata + "}\n\n"

	xdata = xdata + "type Product struct {\n"
	xdata = xdata + "   XMLName xml.Name "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "product")
	xdata = xdata + "	 Item   string "
	xdata = xdata + fmt.Sprintf("`xml:%q`\n", "item")
	xdata = xdata + "}\n\n"

	xdata = xdata + "func main() {\n"
	xdata = xdata + fmt.Sprintf(" xFile:=%q\n", "test.xml")

	xdata = xdata + fmt.Sprintf("    fmt.Println( %q )\n", "XML to Strucs Test Output")
	//	xdata = xdata + fmt.Sprintf("    fmt.Printf( %q )\n\n )\n", "Using "+xFile)

	xdata = xdata + fmt.Sprintf(" xmlFile, err := os.Open(xFile)\n")
	xdata = xdata + fmt.Sprintf("if err != nil {\n")
	xdata = xdata + fmt.Sprintf("fmt.Println(err)\n")
	xdata = xdata + fmt.Sprintf("	}\n")
	xdata = xdata + fmt.Sprintf("    fmt.Println( %q )\n", "XML File Successfuly Opened")
	xdata = xdata + fmt.Sprintf("	defer xmlFile.Close()\n")
	xdata = xdata + fmt.Sprintf("	byteValue, _ := ioutil.ReadAll(xmlFile)\n")
	xdata = xdata + fmt.Sprintf("	var users Users\n")
	xdata = xdata + fmt.Sprintf("	xml.Unmarshal(byteValue, &users)\n")

	xdata = xdata + fmt.Sprintf("for i := 0; i < len(users.Users); i++ {\n")

	//xdata = xdata + "	fmt.Println("User Name: " + users.Users[i].Name)"
	xdata = xdata + fmt.Sprintf("fmt.Println(%q + users.Users[i].Name)\n", "UserName: \n")
	xdata = xdata + "}\n"

	xdata = xdata + "}\n"

	f, err := os.Create("app/main.go")
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	l, err := f.WriteString(xdata)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

}

func main() {
	fmt.Println("Convert XML file to Go Structs")
	xFile := "app/test.xml"

	BuildApp(xFile)

}
