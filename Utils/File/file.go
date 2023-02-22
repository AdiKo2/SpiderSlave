package File

import (
	"SlaveSpider/Repository"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type File interface {
	Export(homeUsers []Repository.HomeUser)
	//Read(file []byte) any
}

type Xml struct {
}

func (m Xml) Export(homeUsers []Repository.HomeUser) {
	file, err := os.Create("test.csv")
	defer file.Close()
	if err != nil {
		fmt.Println("error in writing the file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	for _, user := range homeUsers {
		row := []string{user.DisplayName, user.FirstName, user.LastName, user.Email}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
