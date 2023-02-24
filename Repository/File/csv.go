package File

import (
	"SlaveSpider/Repository"
	"SlaveSpider/Utils"
	"encoding/csv"
	"fmt"
	"os"
)

type Csv struct {
	fileName  string
	batchSize int
}

func InitCsv() Repository.File {
	var csvStruct Repository.File = &Csv{
		"test",
		2,
	}
	//csvStruct.Export()
	return csvStruct
}

func (m *Csv) Export(homeUsers Utils.Slice2D) {
	file, err := os.Create(fmt.Sprint(m.fileName + ".csv"))
	defer file.Close()
	if err != nil {
		fmt.Println("error in writing the file", err)
	}

	//writer := bufio.NewWriter(file)
	w := csv.NewWriter(file)
	homeUsers.BatchUsers2(m.batchSize, w.WriteAll)
}
