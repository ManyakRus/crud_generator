package entities_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"

	"os"
	"sort"
	"strings"
)

// CreateFiles_UpdateEveryColumn - создаёт 1 файл в папке model, для каждой колонки функция Update()
func CreateFiles_UpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)
	//ModelName := Table1.NameGo

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile() + TableName + micro.SeparatorFile()

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyModel)

	FilenameTemplateModel := DirTemplatesModel + config.Settings.MODEL_TABLE_UPDATE_FILENAME
	FilenameReadyModel := DirReadyModel + config.Settings.PREFIX_MODEL + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyModel)

	//чтение файла шаблона
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyModel)

	//замена импортов на новые URL
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

		ConstantsURL := create_files.Find_DBConstantsURL()
		TextModel = create_files.AddImport(TextModel, ConstantsURL)
	}

	TextModel = create_files.CheckAndAdd_ImportTime_FromText(TextModel)

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	TextNew := ""
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("CreateFiles_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}
		TextNew1 := FindTextUpdateEveryColumn(Table1, Column1)
		TextNew = TextNew + TextNew1
	}

	// пустой файл не нужен
	if TextNew == "" {
		return err
	}

	//
	TextModel = TextModel + TextNew

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), config.Settings.FILE_PERMISSIONS)

	return err
}

func FindTextUpdateEveryColumn(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo

	Otvet = `
// Update_` + ColumnName + ` - изменяет объект в БД по ID, присваивает ` + ColumnName + `
func (m *` + ModelName + `) Update_` + ColumnName + `() error {
	if Crud_` + ModelName + ` == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_` + ModelName + `.Update_` + ColumnName + `(m)

	return err
}
`

	return Otvet
}

func AddInterface_UpdateEveryColumn(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	//TextNew := ""
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("CreateFiles_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}
		TextNew1 := FindTextInterfaceUpdateEveryColumn(Table1, Column1)
		Otvet = create_files.AddInterfaceFunction(Otvet, TextNew1)
		//pos1 := strings.Index(TextModel, TextNew1)
		//if pos1 >= 0 { //только если нет такой функции в интерфейсе
		//	TextNew = TextNew + TextNew1
		//}
	}

	return Otvet
}

func FindTextInterfaceUpdateEveryColumn(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo

	Otvet = `
	Update_` + ColumnName + `(*` + ModelName + `) error`

	return Otvet
}
