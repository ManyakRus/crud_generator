package create_files

import (
	"errors"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/mini_func"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"strconv"
	"strings"
)

// FindSingularName - возвращает наименование в единственном числе
func FindSingularName(s string) string {
	var Otvet string

	if s == "" {
		return Otvet
	}

	Otvet = inflection.Singular(s)

	return Otvet
}

// FormatName - возвращает наименование в формате CamelCase
func FormatName(Name string) string {
	Otvet := Name

	Otvet1, ok := types.MapReplaceName[Name]
	if ok == true {
		return Otvet1
	}

	switch strings.ToLower(Name) {
	case "id":
		Otvet = "ID"
		return Otvet
	}

	Otvet = strcase.ToCamel(Otvet)

	//_id в конце заменяем на ID
	lenName := len(Name)
	lenOtvet := len(Otvet)
	if lenName >= 3 {
		last3 := strings.ToLower(Name[lenName-3:])
		if last3 == "_id" {
			Otvet = Otvet[:lenOtvet-2] + "ID"
		}
	}

	return Otvet
}

// DeleteFuncFromFuncName - удаляет функцию из текста начиная с объявления функции
func DeleteFuncFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "\nfunc " + FuncName + "("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncFromComment - удаляет функцию из текста начиная с комментария
func DeleteFuncFromComment(Text, Comment string) string {
	Otvet := Text

	TextFind := Comment //"\n// Delete "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet

}

// DeleteCommentFromFuncName - удаляет комментарий с названием функции
func DeleteCommentFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "//" + FuncName + "()"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	//pos1 := strings.Index(Otvet, TextFind)
	//if pos1 < 0 {
	//	return Otvet
	//}
	//s2 := Otvet[pos1+1:]
	//
	//posEnd := strings.Index(s2, "\n}")
	//if posEnd < 0 {
	//	return Otvet
	//}
	//
	//Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteLineWithComment - удаляет текст от комментария до конца строки
func DeleteLineWithComment(Text, FuncName string) string {
	Otvet := Text

	TextFind := "//" + FuncName + "()"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1:]

	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+1:]

	return Otvet
}

// FindPrimaryKeyNameTypeGo - возвращает наименование колонки PrimaryKey
func FindPrimaryKeyNameTypeGo(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	for _, Column1 := range Table1.MapColumns {
		if Column1.IsIdentity == true {
			return Column1.NameGo, Column1.TypeGo
		}
	}

	return Otvet, Type
}

// ReplacePrimaryKeyID - заменяет "ID" на название колонки PrimaryKey
func ReplacePrimaryKeyID(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnName, ColumnTypeGo := FindPrimaryKeyNameTypeGo(Table1)
	if mini_func.IsNumberType(ColumnTypeGo) == true {
		Otvet = strings.ReplaceAll(Otvet, "Otvet.ID", "Otvet."+ColumnName)
	} else if ColumnTypeGo == "string" {
		Otvet = strings.ReplaceAll(Otvet, "Otvet.ID == 0", "Otvet."+ColumnName+" == \"\"")
		Otvet = strings.ReplaceAll(Otvet, "Otvet.ID", "Otvet."+ColumnName)
	}

	return Otvet
}

// AddSkipNow - добавляет строку t.SkipNow()
func AddSkipNow(Text string, Table1 *types.Table) string {
	Otvet := Text

	if Table1.IDMinimum == "" || Table1.IDMinimum == "0" {
		TextFind := "(t *testing.T) {"
		Otvet = strings.ReplaceAll(Otvet, TextFind, TextFind+"\n\tt.SkipNow() //now rows in DB\n")
	}

	return Otvet
}

// CheckGoodTable - возвращает ошибку если таблица неправильная
func CheckGoodTable(Table1 *types.Table) error {
	var err error

	ColumnName, _ := FindPrimaryKeyNameTypeGo(Table1)
	if ColumnName == "" {
		TextError := fmt.Sprint("Wrong table ", Table1.Name, " error: not found Primary key")
		err = errors.New(TextError)
	}

	return err
}

// PrintableString - возвращает строку без запрещённых символов
func PrintableString(s string) string {
	Otvet := s
	Otvet = strconv.Quote(Otvet) //экранирование символов
	len1 := len(Otvet)
	if len1 > 0 {
		Otvet = Otvet[1 : len1-1]
	}

	return Otvet
}

// Find_Template_DB_Foldername - возвращает путь к папке
func Find_Template_DB_Foldername() string {
	Otvet := ""

	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()

	Otvet = DirTemplatesDB
	return Otvet
}

// DeleteImportModel - удаляет лишний импорт модели
func DeleteImportModel(s string) string {
	Otvet := s

	ModelURL := FindModelURL()
	ImportName := micro.LastWord(ModelURL)

	pos1 := strings.Index(Otvet, ImportName+".")
	if pos1 < 0 {
		Otvet = strings.ReplaceAll(Otvet, `"`+ModelURL+`"`, "")
	}

	return Otvet
}

func FindModelURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL

	return Otvet
}

func FindTextDefaultValue(Type_go string) string {
	var Otvet string

	sValue := ""
	switch Type_go {
	case "string":
		sValue = "\\\"\\\""
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		sValue = "0"
	case "time.Time":
		sValue = "null"
	}

	if sValue != "" {
		Otvet = ";default:" + sValue
	}

	return Otvet
}

// FindURL_Alias - возвращает URL репозитория с пакетом "alias"
func FindURL_Alias() string {
	Otvet := ""
	if config.Settings.TEMPLATE_FOLDERNAME_ALIAS == "" {
		return Otvet
	}
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/"
	Otvet = Otvet + config.Settings.TEMPLATE_FOLDERNAME_ALIAS

	return Otvet
}

// AddImport - добавляет URL в секцию Import, если его там нет
func AddImport(Text, URL string) string {
	Otvet := Text

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `"`+URL+`"`)
	if pos1 >= 0 {
		return Otvet
	}

	//
	TextFind := "import ("
	LenFind := len(TextFind)
	pos1 = strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		log.Error("not found word: import (")
		return Otvet
	}

	Otvet = Otvet[:pos1+LenFind] + "\n\t" + `"` + URL + `"` + Otvet[pos1+LenFind:]

	return Otvet
}

// AddImportTime - добавляет покет в секцию Import, если его там нет
func AddImportTime(TextModel string) string {
	Otvet := TextModel

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `"time"`)
	if pos1 >= 0 {
		return Otvet
	}

	//
	pos1 = strings.Index(Otvet, "import (")
	if pos1 < 0 {
		log.Error("not found word: import (")
		return TextModel
	}

	Otvet = Otvet[:pos1+8] + "\n\t" + `"time"` + Otvet[pos1+8:]

	return Otvet
}

// CheckAndAddImportTime_FromTable - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAddImportTime_FromTable(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	HasTimeColumn := Has_ColumnType_Time(Table1)
	if HasTimeColumn == false {
		return Otvet
	}

	Otvet = AddImportTime(Otvet)

	return Otvet
}

// CheckAndAddImportTime_FromText - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAddImportTime_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, " time.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImportTime(Otvet)

	return Otvet
}
