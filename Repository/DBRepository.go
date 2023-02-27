package Repository

import (
	"SlaveSpider/Utils"
	"fmt"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"github.com/smallnest/gen/dbmeta"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
)

type DBRepository struct {
	cfg *Utils.Config
	db  *gorm.DB
}

type tableData struct {
	ColumnName string
	DataType   string
}

type tableDataResp []tableData

func InitDBRepository(cfg *Utils.Config) *DBRepository {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DataBase.Server.Host,
		cfg.DataBase.Credentials.UserName,
		cfg.DataBase.Credentials.Password,
		cfg.DataBase.Name,
		cfg.DataBase.Server.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(err)
	}
	return &DBRepository{cfg: cfg, db: db}
}

func (api DBRepository) ReturnSQLNames2() {
	var resp []map[string]any
	//var resp map[string]interface{}
	api.db.Table("names").Find(&resp)
	api.db.Get("name")
	//confgg := api.db.Table("names").Find(&resp).Config
	//gorm
	fmt.Println(dbmeta.Exists("names"))
	fmt.Println(len(resp))
	value, index := resp[0]["name"]
	fmt.Println(value, index)

}

func (api DBRepository) TableColumns(tableName string) (*tableDataResp, error) {
	var resp tableDataResp

	tx := api.db.Raw("SELECT column_name, data_type FROM information_schema.COLUMNS WHERE TABLE_NAME = ?", tableName).Scan(&resp)

	return &resp, tx.Error
}

func (columns tableDataResp) CreateTableStruct() interface{} {
	//dictionaryTypes := map[string]interface{}{
	//	"character varying": "string",
	//	"integer":           1,
	//}

	dictionaryTypes := map[string]reflect.Type{
		"character varying": reflect.Type("l"),
		"integer":           1,
	}

	//instance := dynamicstruct.NewStruct().
	//	AddField("ColumnName", "", `json:"int"`).
	//	AddField("DataType", "", `json:"someText"`).
	//	Build().
	//	New()

	//newStruct := dynamicstruct.NewStruct()
	var newStruct *[]reflect.StructField
	for _, column := range columns {
		fmt.Println(dictionaryTypes[column.DataType])
		//newStruct = newStruct.AddField(column.ColumnName, dictionaryTypes[column.DataType], "")
		//newStruct.AddField(fmt.Sprint(index), "", "")
		*newStruct = append(*newStruct, reflect.StructField{
			Name: column.ColumnName,
			Type: dictionaryTypes[column.DataType],
			Tag: ,
		})
	}
	newStruct.Build()
	var inter1 interface{}
	return inter1
	//return any
}
