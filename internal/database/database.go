package database

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/database/postgres"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/alexsergivan/transliterator"
)

// FillMapTable - возвращает массив MassTable данными из БД
func FillMapTable(SettingsFill types.SettingsFillFromDatabase) (map[string]*types.Table, error) {
	MapTable := make(map[string]*types.Table, 0)
	var err error

	//заполним MapTable
	MapTable, err = postgres.StartAll(SettingsFill)
	if err != nil {
		log.Error("FillMapTable1() error: ", err)
		return MapTable, err
	}

	//заполним NameGo + TypeGo + NameGo_translit
	MapDBTypes := dbmeta.GetMappings()
	err = FillNameGoAll(MapTable, MapDBTypes)

	//зазполним PrimaryKeys ручные
	err = FillPrimaryKeysAll_from_MapPrimaryKeys(MapTable, types.MapPrimaryKeys)

	return MapTable, err
}

// FillNameGoAll - заполняет NameGo во всех таблицах, а также NameGo_translit
func FillNameGoAll(MapTable map[string]*types.Table, MapDBTypes map[string]*dbmeta.SQLMapping) error {
	var err error

	trans := transliterator.NewTransliterator(nil)

	for TableName, Table1 := range MapTable {
		//заполним имя колонки
		ModelName := create_files.Find_SingularName(TableName)
		ModelName = create_files.FormatName(ModelName)
		NameGo_translit := trans.Transliterate(ModelName, "en")
		Table1.NameGo = ModelName
		Table1.NameGo_translit = NameGo_translit

		for _, Column1 := range Table1.MapColumns {
			FillNameGo(Column1)
			//postgres.FillTypeGo(MapDBTypes, Column1)
		}

	}

	return err
}

// FillNameGo - заполняет NameGo в 1 колонку
func FillNameGo(Column1 *types.Column) {
	var err error

	ColumnName := Column1.Name
	ColumnNameGo := create_files.FormatName(ColumnName)
	Column1.NameGo = ColumnNameGo
	if ColumnNameGo == "" {
		err = errors.New("FillNameGo() error: Column: " + ColumnName + " Type: " + Column1.Type + " NameGo= \"\"")
		log.Panic(err)
	}

	trans := transliterator.NewTransliterator(nil)
	NameGo_translit := trans.Transliterate(ColumnNameGo, "en")
	Column1.NameGo_translit = NameGo_translit

	return
}

// FillPrimaryKeysAll_from_MapPrimaryKeys - заполняет IsPrimaryKey во всех таблицах
func FillPrimaryKeysAll_from_MapPrimaryKeys(MapTable map[string]*types.Table, MapPrimaryKeys map[string][]string) error {
	var err error

	for TableName, Table1 := range MapTable {
		for _, Column1 := range Table1.MapColumns {
			IsPrimaryKey_manual := Find_IsPrimaryKey_manual(MapPrimaryKeys, TableName, Column1.Name)
			if IsPrimaryKey_manual == true {
				Column1.IsPrimaryKey = true
			}
		}

	}

	return err
}

// Find_IsPrimaryKey_manual - возвращает true если эта колонка является PrimaryKey (заполненных вручную)
func Find_IsPrimaryKey_manual(MapPrimaryKeys map[string][]string, TableName string, ColumnName string) bool {
	Otvet := false

	MassColumnStrings, ok := MapPrimaryKeys[TableName]
	if ok == false {
		return Otvet
	}

	for _, ColumnString := range MassColumnStrings {
		if ColumnString == ColumnName {
			Otvet = true
		}
	}

	return Otvet
}
