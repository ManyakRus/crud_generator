package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке db
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы crud
		if config.Settings.NEED_CREATE_DB == true {
			err = CreateFiles(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//тестовые файлы crud
		if config.Settings.NEED_CREATE_DB_TEST == true {
			err = CreateFiles_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//файлы UpdateEveryColumn
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы db update
			err = CreateFiles_UpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы db update
			err = CreateFiles_UpdateEveryColumn_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//файлы Cache
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//файлы cache
			if config.Settings.NEED_CREATE_CACHE_FILES == true {
				err = CreateFiles_Cache(Table1)
				if err != nil {
					log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

			//тестовые файлы cache
			if config.Settings.NEED_CREATE_CACHE_TEST_FILES == true {
				err = CreateFiles_Cache_Test(Table1)
				if err != nil {
					log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
					return err
				}
			}
		}
		//FindBy
		err = CreateFiles_FindBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindBy_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindMassBy
		err = CreateFiles_FindMassBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindMassBy_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//ReadAll
		err = CreateFiles_ReadAll(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_ReadAll_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindModelBy
		err = CreateFiles_FindModelBy(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindModelBy_Test(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// RenameFunctions - заменяет названия функций, на названия из файла
func RenameFunctions(TextDB string, Table1 *types.Table) string {
	Otvet := TextDB

	TableName := strings.ToLower(Table1.Name)
	Rename1, ok := types.MapRenameFunctions[TableName]
	if ok == false {
		return Otvet
	}

	for _, v := range Rename1 {
		Otvet = strings.ReplaceAll(Otvet, "func "+v.Old+"(", "func "+v.New+"(")
	}

	return Otvet
}

// ReplaceCacheRemove - заменяет cache.Remove(IntFromAlias(m.ID))
func ReplaceCacheRemove(Text string, Table1 *types.Table) string {
	Otvet := Text

	if Table1.PrimaryKeyColumnsCount == 1 {
		Otvet = ReplaceCacheRemove_1PK(Otvet, Table1)
	} else {
		Otvet = ReplaceCacheRemove_ManyPK(Otvet, Table1)
	}

	return Otvet
}

// ReplaceColumnNamePK - заменяет "ColumnNamePK" на текст имя колонки
func ReplaceColumnNamePK(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamePK", ColumnPK.NameGo)

	return Otvet
}

// ReplaceCacheRemove1PK - заменяет cache.Remove(IntFromAlias(m.ID))
func ReplaceCacheRemove_1PK(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	TextOld := "cache.Remove(IntFromAlias(m.ID))"
	Value := create_files.ConvertFromAlias(Table1, ColumnPK, "m")
	TextNew := "cache.Remove(" + Value + ")"
	Otvet = strings.ReplaceAll(Otvet, TextOld, TextNew)

	return Otvet
}

// ReplaceCacheRemove_ManyPK - заменяет cache.Remove(IntFromAlias(m.ID)) на cache.Remove(m.StringIdentifier())
func ReplaceCacheRemove_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	if Table1.PrimaryKeyColumnsCount > 1 {
		TextOld := "cache.Remove(IntFromAlias(m.ID))"
		TextNames, _, _ := create_files.FindText_ID_VariableName_Many(Table1, "m")
		TextNew := "cache.Remove(" + Table1.Name + ".StringIdentifier(" + TextNames + "))"
		Otvet = strings.ReplaceAll(Otvet, TextOld, TextNew)
	}

	return Otvet
}
