package main

import (
	"SlaveSpider/Repository"
	"SlaveSpider/Utils/File"
)

func main() {
	myhomeusers := []Repository.HomeUser{
		{"adi22", "adi", "komraz", "adikomraz@gmail.com"},
		{"or33", "or", "dratler", "or@dratler.com"},
	}

	var file File.File
	file = File.Xml{}
	file.Export(myhomeusers)
}
