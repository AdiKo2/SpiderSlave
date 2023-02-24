package File

import (
	"SlaveSpider/Repository"
	"SlaveSpider/Utils"
	"encoding/xml"
	"fmt"
	"os"
)

type Xml struct {
	fileName string
}

type HomeUsersXML struct {
	HomeUsers Utils.Slice2D
}

func InitXml() Repository.File {
	var xmlStruct Repository.File = &Xml{
		"test",
	}
	return xmlStruct
}

func (m *Xml) Export(homeUsers Utils.Slice2D) {
	file, err := os.Create(fmt.Sprint(m.fileName + ".xml"))
	if err != nil {
		fmt.Println("error in writing the file", err)
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")

	err = encoder.Encode(HomeUsersXML{homeUsers})
	if err != nil {
		fmt.Println("error in marshaling data and writing to file", err)
	}
}
