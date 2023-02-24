package main

import (
	"SlaveSpider/Repository/File"
	"SlaveSpider/Utils"
)

func main() {

	type MyUsserr struct {
		name string
	}
	myhomeusers2 := Utils.Slice2D{
		{"adi", "komraz", 19, "adikomraz@gmail.com"},
		{"ravit", "komraz", 19, "rk98@gmail.com"},
		{"or", "dratler", 24, "or@dratler.com"},
		{"rima", "polsky", 54, "fairy3@gmail.com"},
	}
	myhomeusers1 := Utils.Slice2D{
		{"adi", "komraz", 19, "adikomraz@gmail.com"},
	}
	csvFile := File.InitCsv()
	xmlFile := File.InitXml()

	csvFile.Export(myhomeusers2)
	xmlFile.Export(myhomeusers2)
}
