package main

import (
	"SlaveSpider/Repository"
	"SlaveSpider/Repository/File"
	"SlaveSpider/Utils"
	"fmt"
)

func main() {

	myhomeusers2 := Utils.Slice2D{
		{"adi", "komraz", 19, "adikomraz@gmail.com"},
		{"ravit", "komraz", 19, "rk98@gmail.com"},
		{"or", "dratler", 24, "or@dratler.com"},
		{"rima", "polsky", 54, "fairy3@gmail.com"},
	}

	csvFile := File.InitCsv()
	xmlFile := File.InitXml()

	csvFile.Export(myhomeusers2)
	xmlFile.Export(myhomeusers2)

	cfg1 := Utils.InitConfig()
	dbRepository := Repository.InitDBRepository(cfg1)
	result, err := dbRepository.TableColumns("names")
	fmt.Println(result, err)

	result.CreateTableStruct()

}
