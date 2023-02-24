package Repository

import "SlaveSpider/Utils"

type File interface {
	//Export(homeUsers []HomeUser)
	Export(homeUsers Utils.Slice2D)

	//Read(file []byte) any
}
