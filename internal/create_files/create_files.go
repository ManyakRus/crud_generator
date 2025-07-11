package create_files

import (
	"errors"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/mini_func"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/gobeam/stringy"
	"github.com/jinzhu/inflection"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Find_SingularName - возвращает наименование в единственном числе
func Find_SingularName(s string) string {
	var Otvet string

	if s == "" {
		return Otvet
	}

	Otvet = inflection.Singular(s)

	return Otvet
}

// FormatName - возвращает наименование в формате PascalCase
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

	//Otvet = strcase.ToCamel(Otvet) //не понимает русские буквы

	str := stringy.New(Otvet)
	Otvet = str.PascalCase("?", "").Get()
	Otvet = pascalCaseReformatNumbers(Otvet)

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

// pascalCaseReformatNumbers - после цифр буквы должны быть в верхнем регистре
func pascalCaseReformatNumbers(s string) string {
	Otvet := ""

	var last rune
	for _, s1 := range s {
		if last >= 48 && last <= 57 {
			s1 = unicode.ToUpper(s1)
		}
		Otvet = Otvet + string(s1)
		last = s1
	}

	return Otvet
}

// DeleteFuncFromFuncName - удаляет функцию из текста начиная с объявления функции
func DeleteFuncFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "\nfunc " + FuncName + "("
	//TextFind2 := "\nfunc " + FuncName + "("
	//pos1 := micro.FindPos(Otvet, TextFind, TextFind2)
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

	TextFind := Comment
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

// Find_PrimaryKeyNameTypeGo - возвращает наименование и тип golang колонки PrimaryKey
func Find_PrimaryKeyNameTypeGo(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		if Column1.IsPrimaryKey == true {
			return Column1.NameGo, Column1.TypeGo
		}
	}

	return Otvet, Type
}

// Find_PrimaryKeyNameType - возвращает наименование и тип БД колонки PrimaryKey
func Find_PrimaryKeyNameType(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		if Column1.IsPrimaryKey == true {
			return Column1.Name, Column1.Type
		}
	}

	return Otvet, Type
}

// Find_PrimaryKeyColumn - возвращает Column для колонки PrimaryKey
func Find_PrimaryKeyColumn(Table1 *types.Table) (Column1 *types.Column) {
	var Otvet *types.Column

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		if Column1.IsPrimaryKey == true {
			return Column1
		}
	}

	return Otvet
}

// Find_PrimaryKeyColumns - возвращает несколько Column для колонки PrimaryKey
func Find_PrimaryKeyColumns(Table1 *types.Table) []*types.Column {
	Otvet := make([]*types.Column, 0)

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		if Column1.IsPrimaryKey == true {
			Otvet = append(Otvet, Column1)
		}
	}

	return Otvet
}

//// Find_PrimaryKeyNameTypeGo - возвращает наименование и тип golang колонки PrimaryKey
//func Find_PrimaryKeyNameTypeGo(Table1 *types.Table) (string, string) {
//	Otvet := ""
//	Type := ""
//
//	for _, Column1 := range Table1.MapColumns {
//		if Column1.IsIdentity == true {
//			return Column1.NameGo, Column1.TypeGo
//		}
//	}
//
//	return Otvet, Type
//}

// Replace_PrimaryKeyOtvetID - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_PrimaryKeyOtvetID_Many(Otvet, Table1)

	return Otvet
}

// FindText_Equal0 - возвращает текст сравнения с нулевым значением
func FindText_Equal0(Column1 *types.Column) string {
	Otvet := " == 0"

	switch Column1.TypeGo {
	case "string":
		Otvet = ` == ""`
	case "uuid.UUID":
		Otvet = ` == uuid.Nil`
	case "time.Time":
		Otvet = `.IsZero() == true`
	case "bool":
		Otvet = ` == false`
	}

	return Otvet
}

// Replace_PrimaryKeyOtvetID1 - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID1(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnNamePK, ColumnTypeGoPK := Find_PrimaryKeyNameTypeGo(Table1)
	ColumnPK := Find_PrimaryKeyColumn(Table1)

	//заменим ID-Alias на ID
	TableName := Table1.Name
	IDName, _ := Find_PrimaryKeyNameType(Table1)
	Alias, ok := types.MapConvertID[TableName+"."+IDName]
	OtvetColumnName := "Otvet." + ColumnNamePK
	if ok == true {
		OtvetColumnName = Alias + "(" + OtvetColumnName + ")"
	}

	Otvet = strings.ReplaceAll(Otvet, "Otvet.ID = AliasFromInt(ID)", "Otvet."+ColumnNamePK+" = "+ColumnNamePK)

	//заменим int64(Otvet.ID) на Otvet.ID
	if mini_func.IsNumberType(ColumnTypeGoPK) == true {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "string" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+" == \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+" != \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "uuid.UUID" || ColumnTypeGoPK == "uuid.NullUUID" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+" == uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+" != uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "time.Time" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+".IsZero() == true")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+".IsZero() == false")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	}
	//Otvet = strings.ReplaceAll(Otvet, "Otvet.ID = ", OtvetColumnName+" = ")
	//Otvet = strings.ReplaceAll(Otvet, "Otvet.ID != ", OtvetColumnName+" != ")
	//Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " "+OtvetColumnName+")")
	Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " Otvet."+ColumnNamePK+")")

	//Alias преобразуем в int64, и наоборот
	if Alias != "" {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(Otvet.ID)", ColumnTypeGoPK+"(Otvet."+ColumnNamePK+")")
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(Otvet.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(ID)", Alias+"("+ColumnNamePK+")")
	} else {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(Otvet.ID)", "Otvet."+ColumnNamePK+"")
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(Otvet.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(ID)", "ID")
	}

	Value, GolangCode := Convert_ProtobufVariableToGolangVariable(Table1, ColumnPK, "Request.")
	if GolangCode == "" {
		TextNew := "\t" + ColumnPK.NameGo + " := " + Value + "\n"
		Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextNew)
	} else {
		TextNew := "\t" + GolangCode + "\n"
		Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextNew)
	}

	Value, GolangCode = Convert_ProtobufTypeToGolangAlias(Table1, ColumnPK, "Request.")
	if GolangCode == "" {
		TextNew := "\t" + ColumnPK.NameGo + " := " + Value + "\n"
		Otvet = strings.ReplaceAll(Otvet, "\tID := AliasFromProto(Request.ID)\n", TextNew)
	} else {
		TextNew := "\t" + GolangCode + "\n"
		Otvet = strings.ReplaceAll(Otvet, "\tID := AliasFromProto(Request.ID)\n", TextNew)
	}

	return Otvet
}

// Replace_PrimaryKeyOtvetID_Many - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID_Many(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "Otvet")
	Otvet = Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "Request")
	Otvet = Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "m")

	Otvet = Replace_PrimaryKeyOtvetID1(Otvet, Table1) //для тестов

	return Otvet
}

// Replace_PrimaryKeyOtvetID_ManyPK1 - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID_ManyPK1(Text string, Table1 *types.Table, VariableName string) string {
	Otvet := Text

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDID)
	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = int64(ID)", TextRequestIDInt64ID)
	Otvet = strings.ReplaceAll(Otvet, "\tOtvet.ID = "+VariableName+".ID\n", TextOtvetIDmID)
	Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) == 0", TextMID0)
	Otvet = strings.ReplaceAll(Otvet, "\tm2.ID = int64(m.ID)", TextM2ID)
	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", TextIfMId)
	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", TextIfMIdNot0)
	//Value := Convert_GolangVariableToProtobufVariableID(Table1, ColumnPK, "m")
	//Otvet = strings.ReplaceAll(Otvet, "ProtoFromInt(m.ID)", Value) //protobuf

	//заменим ID := Request.ID
	Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextIDRequestID)

	return Otvet
}

// Replace_PrimaryKeyM_ID - заменяет "m.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyM_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_PrimaryKeyM_ManyPK(Otvet, Table1)

	return Otvet
}

// Replace_PrimaryKeyM_ManyPK - заменяет "m.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyM_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "m")
	Otvet = Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "Request")

	//заменим int64(m.ID)
	//Otvet = ReplacePrimaryKeyM_ID1(Otvet, Table1)

	return Otvet
}

// ReplaceIntFromAlias - заменяет "m.ID" на текст m.ID или int64(m.ID)
func ReplaceIntFromAlias(Text string, Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := Text

	Value := ConvertFromAliasID(Table1, Column1, VariableName)
	Otvet = strings.ReplaceAll(Otvet, "IntFromAlias("+VariableName+".ID)", Value)

	return Otvet
}

// ConvertFromAliasID - возвращает текст m.ID или int64(m.ID)
func ConvertFromAliasID(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := VariableName + "." + Column1.NameGo

	TextConvert, ok := types.MapConvertID[Table1.Name+"."+Column1.Name]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	Otvet = Column1.TypeGo + "(" + VariableName + "." + Column1.NameGo + ")"

	return Otvet
}

// ConvertFromAlias - возвращает текст VariableName или int64(VariableName)
func ConvertFromAlias(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := VariableName

	TextConvert, ok := types.MapConvertID[Table1.Name+"."+Column1.Name]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	Otvet = Column1.TypeGo + "(" + VariableName + ")"

	return Otvet
}

// AddSkipNow - добавляет строку t.SkipNow()
func AddSkipNow(Text string, Table1 *types.Table) string {
	Otvet := Text

	Columns := Find_PrimaryKeyColumns(Table1)
	if Columns == nil {
		return Otvet
	}

	//проверка ИД=""
	is_no_zero := true
	for _, Column1 := range Columns {
		if Column1.IDMinimum == "" || Column1.IDMinimum == "0" {
			is_no_zero = false
			break
		}
	}

	//если нет пустых то возврат
	if is_no_zero == true {
		return Otvet
	}

	//добавляем t.SkipNow()
	TextFind := "(t *testing.T) {"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextFind+"\n\tt.SkipNow() //now rows in DB\n")

	return Otvet
}

// IsGood_Table - возвращает ошибку если таблица неправильная
func IsGood_Table(Table1 *types.Table) error {
	var err error

	//есть Primary Key
	err = IsGood_PrimaryKeyColumnsCount(Table1)
	if err != nil {
		return err
	}

	//название УДАЛИТЬ
	err = IsGood_TableName(Table1)
	if err != nil {
		return err
	}

	return err
}

// IsGood_TableName - возвращает ошибку если таблица неправильная
func IsGood_TableName(Table1 *types.Table) error {
	var err error

	err = IsGood_TableNamePrefix(Table1)
	if err != nil {
		return err
	}

	err = IsGood_TableCommentPrefix(Table1)
	if err != nil {
		return err
	}

	return err
}

// IsGood_PrimaryKeyColumnsCount - возвращает ошибку если количество колонок PrimaryKey неправильное
func IsGood_PrimaryKeyColumnsCount(Table1 *types.Table) error {
	var err error

	if Table1.PrimaryKeyColumnsCount <= 0 {
		TextError := fmt.Sprint("Wrong table: ", Table1.Name, ", error: can not use many Primary key columns count: ", Table1.PrimaryKeyColumnsCount)
		err = errors.New(TextError)
	}

	return err
}

// IsGood_TableNamePrefix - возвращает ошибку если префикс таблицы = "DELETED_"
func IsGood_TableNamePrefix(Table1 *types.Table) error {
	var err error

	TableName := Table1.Name
	HasPrefix := strings.HasPrefix(TableName, config.Settings.TEXT_DELETED_TABLE)
	HasPrefixRus := strings.HasPrefix(TableName, config.Settings.TEXT_DELETED_TABLE_RUS)
	if HasPrefix == true || HasPrefixRus == true {
		TextError := fmt.Sprint("Wrong table: ", Table1.Name, ", warning: name = DELETED_")
		err = errors.New(TextError)
	}

	return err
}

// IsGood_TableNamePrefix - возвращает ошибку если префикс таблицы = "DELETED_"
func IsGood_TableCommentPrefix(Table1 *types.Table) error {
	var err error

	TableComment := Table1.Comment
	HasPrefix := strings.HasPrefix(TableComment, config.Settings.TEXT_DELETED_TABLE)
	HasPrefixRus := strings.HasPrefix(TableComment, config.Settings.TEXT_DELETED_TABLE_RUS)
	if HasPrefix == true || HasPrefixRus == true {
		TextError := fmt.Sprint("Wrong table: ", Table1.Name, ", warning: comment= ", TableComment)
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
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	Otvet = DirTemplatesDB
	return Otvet
}

// FindText_DefaultGORMValue - возвращает значение по умолчанию для столбца Column1 для тегов в GORM
func FindText_DefaultGORMValue(Column1 *types.Column) string {
	var Otvet string

	ColumnName := Column1.Name

	//ищем в файле настроек nullable.json
	is_nullable_config, _ := types.MapNullableFileds[ColumnName]

	//
	sValue := ""
	Type_go := Column1.TypeGo
	//	if Column1.TableKey != "" && Column1.IsNullable == true && (is_nullable_config == true || has_is_nullable_config == false) {
	if Column1.IsNullable == true || (is_nullable_config == true) {
		sValue = "null"
	} else {

		switch Type_go {
		case "string":
			sValue = "\\\"\\\""
		case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
			sValue = "0"
		case "time.Time":
			sValue = "null"
		case "uuid.UUID", "uuid.NullUUID":
			sValue = "null"
		}
	}

	if sValue != "" {
		Otvet = ";default:" + sValue
	}

	return Otvet
}

// FindText_DefaultValue - возвращает golang значение по умолчанию для типа
func FindText_DefaultValue(Type_go string) string {
	var Otvet string

	switch Type_go {
	case "string":
		Otvet = `""`
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		Otvet = "0"
	case "time.Time":
		Otvet = "time.Time{}"
	case "bool":
		Otvet = "false"
	case "uuid.UUID", "uuid.NullUUID":
		Otvet = "uuid.Nil"
	}

	return Otvet
}

// FindText_DefaultValueSQL - возвращает значение по умолчанию для типа
func FindText_DefaultValueSQL(Type_go string) string {
	var Otvet string

	switch Type_go {
	case "string":
		Otvet = `''`
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		Otvet = "0"
	case "time.Time":
		Otvet = "null"
	case "bool":
		Otvet = "false"
	case "uuid.UUID", "uuid.NullUUID":
		Otvet = "null"
	}

	return Otvet
}

// FindText_DefaultValueSQL_NotNull - возвращает значение по умолчанию для типа, не null
func FindText_DefaultValueSQL_NotNull(Type_go string) string {
	var Otvet string

	switch Type_go {
	case "string":
		Otvet = `''`
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		Otvet = "0"
	case "time.Time":
		Otvet = "'0001-01-01 00:00:00+0000'"
	case "bool":
		Otvet = "false"
	case "uuid.UUID", "uuid.NullUUID":
		Otvet = "''"
	}

	return Otvet
}

// Replace_ModelAndTableName - заменяет имя модели и имя таблицы в шаблоне на новые
func Replace_ModelAndTableName(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	return Otvet
}

// Find_ModelComment - возвращает комментарий для модели
func Find_ModelComment(Table1 *types.Table) string {
	Otvet := ""

	TableName := Table1.Name
	ModelName := Table1.NameGo
	COMMENT_MODEL_STRUCT := config.Settings.COMMENT_MODEL_STRUCT

	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment

	return Otvet
}

// Find_ModelNameComment - возвращает комментарий для названия модели
func Find_ModelNameComment(ModelName string, Table1 *types.Table) string {
	Otvet := ""

	TableName := Table1.Name
	//ModelName := Table1.NameGo
	COMMENT_MODEL_STRUCT := config.Settings.COMMENT_MODEL_STRUCT

	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment

	return Otvet
}

// Replace_PackageName - заменяет имя пакета в шаблоне на новое
func Replace_PackageName(Text, PackageName string) string {
	Otvet := Text

	//найдём имя каталога, это будет имя пакета
	PackageName = micro.DeleteEndSlash(PackageName)
	PackageName = micro.LastWord(PackageName)

	//
	TextFind := "package "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		log.Error("not found word: package ")
		return Otvet
	}

	s2 := Otvet[pos1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		log.Error("not found word: \n")
		return Otvet
	}

	Otvet = Otvet[:pos1+len(TextFind)] + PackageName + Otvet[pos1+posEnd:]

	return Otvet
}

// Delete_EmptyImport - удаляет пустой импорт
func Delete_EmptyImport(Text string) string {
	Otvet := Text

	sFind := `import (
)
`
	Otvet = strings.ReplaceAll(Otvet, sFind, "")

	return Otvet
}

// DeleteFunc_Delete - удаляет функцию Delete()
func DeleteFunc_Delete(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Delete ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Delete")

	return Otvet
}

// DeleteFunc_Restore - удаляет функцию Restore()
func DeleteFunc_Restore(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Restore ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Restore")

	return Otvet
}

// DeleteFunc_Find_byExtID - удаляет функцию Find_ByExtID()
func DeleteFunc_Find_byExtID(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверка есть ли колонки ExtID и ConnectionID
	if Has_Column_ExtID_ConnectionID(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Find_ByExtID ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Find_ByExtID")

	return Otvet
}

// Delete_EmptyLines - удаляет пустые строки
func Delete_EmptyLines(Text string) string {
	Otvet := Text
	Otvet = strings.ReplaceAll(Otvet, "\n\n\n", "\n\n")
	Otvet = strings.ReplaceAll(Otvet, "\n//\n\n", "\n\n")
	Otvet = strings.ReplaceAll(Otvet, "\n\t//\n\n", "\n\n")
	//Otvet = strings.ReplaceAll(Otvet, "\r\r", "\r")
	//Otvet = strings.ReplaceAll(Otvet, "\r\n", "\n")
	//Otvet = strings.ReplaceAll(Otvet, "}\n\n", "}\n")
	pos1 := strings.Index(Otvet, "\n\n\n")
	if pos1 >= 0 {
		Otvet = Delete_EmptyLines(Otvet)
	}

	//удалим последние 2 пустые строки
	HasSuffix := strings.HasSuffix(Otvet, "\n\n")
	if HasSuffix == true {
		Otvet = Otvet[:len(Otvet)-1]
	}

	return Otvet
}

// Delete_LastUnderline - удаляет последний символ подчёркивания
func Delete_LastUnderline(s string) string {
	Otvet := s
	if s == "" {
		return Otvet
	}

	Otvet = strings.TrimSuffix(Otvet, "_")

	return Otvet
}

// Find_LastGoodPos - возвращает позицию последнего нахождения, с новой строки
func Find_LastGoodPos(s, TextFind string) int {
	Otvet := -1
	pos1 := strings.LastIndex(s, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	pos2 := strings.Index(s[pos1:], "\n")
	if pos2 < 0 {
		return Otvet
	}
	Otvet = pos1 + pos2 + 1

	return Otvet
}

// AddInterfaceFunction - добавляет функцию в интерфейс
func AddInterfaceFunction(s, TextAdd string) string {
	Otvet := s

	//Проверим такая функция уже есть
	pos1 := strings.Index(Otvet, TextAdd)
	if pos1 >= 0 {
		return Otvet
	}

	//найдём начало интефейса
	sFind := " interface {"
	pos1 = Find_LastGoodPos(Otvet, sFind)
	if pos1 < 0 {
		log.Error("Find_LastGoodPos() error: not found: ", sFind)
		return Otvet
	}

	s2 := Otvet[pos1:]
	pos2 := strings.Index(s2, "\n}")
	if pos2 < 0 {
		log.Error("Find_LastGoodPos() error: not found: \\n")
		return Otvet
	}
	PosStart := pos1 + pos2

	Otvet = Otvet[:PosStart] + TextAdd + Otvet[PosStart:]

	return Otvet
}

// FindText_ProtobufRequest - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
func FindText_ProtobufRequest(Table1 *types.Table) (OtvetRequestType string, OtvetRequestName string) {

	if Table1.PrimaryKeyColumnsCount == 1 {
		OtvetRequestType, OtvetRequestName = FindText_ProtobufRequest1(Table1)
	} else {
		OtvetRequestType = FindText_ProtobufRequest_ManyPK(Table1)
	}

	return OtvetRequestType, OtvetRequestName
}

// FindText_ProtobufRequest1 - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
func FindText_ProtobufRequest1(Table1 *types.Table) (OtvetRequestType string, OtvetRequestName string) {
	//OtvetRequestType := ""
	//OtvetRequestName := ""

	TextRequest := "Request_"

	PrimaryKeyColumn := Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return "", ""
	}

	PrimaryKeyTypeGo := PrimaryKeyColumn.TypeGo

	switch PrimaryKeyTypeGo {
	case "string", "uuid.UUID":
		OtvetRequestType = TextRequest + "String"
		OtvetRequestName = "String_1"
	case "time.Time":
		OtvetRequestType = TextRequest + "Date"
		OtvetRequestName = "Date"
	default:
		OtvetRequestName = PrimaryKeyColumn.TypeGo
		OtvetRequestName = micro.StringFromUpperCase(OtvetRequestName)
		OtvetRequestType = TextRequest + OtvetRequestName
	}

	return OtvetRequestType, OtvetRequestName
}

// FindText_ProtobufRequest_ManyPK - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
func FindText_ProtobufRequest_ManyPK(Table1 *types.Table) string {
	Otvet := "Request"

	MassPrimaryKeyColumns := Find_PrimaryKeyColumns(Table1)

	if len(MassPrimaryKeyColumns) == 0 {
		return Otvet
	}

	Otvet = FindText_ProtobufRequest_Column_ManyPK(Table1, MassPrimaryKeyColumns[0])

	//for _, ColumnPK1 := range MassPrimaryKeyColumns {
	//	Type1, _, _, _ := FindText_ProtobufRequest_ID_Type(Table1, ColumnPK1, "")
	//	Otvet = Otvet + Type1
	//}

	return Otvet
}

// FindText_ProtobufRequest_Column_ManyPK - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
func FindText_ProtobufRequest_Column_ManyPK(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := "Request"

	MassPrimaryKeyColumns := Find_PrimaryKeyColumns(Table1)

	if len(MassPrimaryKeyColumns) == 0 {
		return Otvet
	}

	IsPrimaryKey := false
	for _, ColumnPK1 := range MassPrimaryKeyColumns {
		if Column1 == ColumnPK1 {
			IsPrimaryKey = true
		}
		Type1, _, _, _ := FindText_ProtobufRequest_ID_Type(Table1, ColumnPK1, "")
		Otvet = Otvet + Type1
	}

	if IsPrimaryKey == false {
		Type1, _, _, _ := FindText_ProtobufRequest_ID_Type(Table1, Column1, "")
		Otvet = Otvet + Type1
	}

	return Otvet
}

// FindText_ProtobufRequest_ID_Type - возвращает имя message из .proto для двух параметров ID + Type,в зависимости от типа, а также название поля
// возвращает:
// RequestName - имя message из .proto
// TextRequestFieldName - название поля в Request
// TextRequestFieldGolang - название поля в Request с преобразованием в тип гоу
// TextGolangLine - замена всей строки в го
func FindText_ProtobufRequest_ID_Type(Table1 *types.Table, Column1 *types.Column, VariableName string) (RequestName string, RequestFieldName string, RequestFieldGolang string, GolangLine string) {
	//RequestName := "RequestId"
	//RequestFieldName := "ID"
	//RequestFieldGolang := "ID"
	//GolangLine := ""

	TypeGo := Column1.TypeGo
	TableName := Table1.Name
	ColumnName := Column1.Name

	//найдём тип колонки PrimaryKey
	PrimaryKeyColumns := Find_PrimaryKeyColumns(Table1)
	if len(PrimaryKeyColumns) == 0 {
		return "", "", "", ""
	}

	isPrimaryKey := false
	Number := 0
	for _, ColumnPK1 := range PrimaryKeyColumns {
		if Column1.TypeGo == ColumnPK1.TypeGo {
			Number = Number + 1
		} else if IsStringOrUUID(ColumnPK1.TypeGo) == true && IsStringOrUUID(Column1.TypeGo) == true {
			Number = Number + 1
		}
		if ColumnPK1 == Column1 {
			isPrimaryKey = true
			break
		}
	}
	if isPrimaryKey == false {
		Number = Number + 1
	}
	sNumber := "_" + strconv.Itoa(Number)
	//sNumber := "_"
	//if IsStringOrUUID(Column1.TypeGo) == true {
	//	sNumber = strconv.Itoa(Number)
	//	sNumber = "_" + sNumber
	//} else {
	//	//} else if Number != 1 {
	//	sNumber = strconv.Itoa(Number)
	//	sNumber = "_" + sNumber
	//}

	//RequestName, _ = FindText_ProtobufRequest(Table1)

	TextRequestProtoName := ""

	//найдём строку по типу колонки
	switch TypeGo {
	case "int", "int64":
		{
			TextRequestProtoName = "Int64"
			RequestFieldName = "Int64" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}

	case "int32":
		{
			TextRequestProtoName = "Int32"
			RequestFieldName = "Int32" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}
	case "string":
		{
			TextRequestProtoName = "String"
			RequestFieldName = "String" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}
	case "uuid.UUID":
		{
			TextRequestProtoName = "String"
			RequestFieldName = "String" + sNumber
			GolangLine = Column1.NameGo + ", err := uuid.Parse(" + VariableName + "" + RequestFieldName + ")" + `
	if Request.` + RequestFieldName + ` == "" {
		` + Column1.NameGo + ` = uuid.Nil
		err = nil
	}
	if err != nil {
		return &Otvet, err
	}
`
			RequestFieldGolang = VariableName + RequestFieldName
		}
	case "time.Time":
		{
			TextRequestProtoName = "Date"
			RequestFieldName = "Date" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName + ".AsTime()"
		}
	case "float32":
		{
			TextRequestProtoName = "Float32"
			RequestFieldName = "Float32" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}
	case "float64":
		{
			TextRequestProtoName = "Float64"
			RequestFieldName = "Float64" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}
	case "bool":
		{
			TextRequestProtoName = "Bool"
			RequestFieldName = "Bool" + sNumber
			RequestFieldGolang = VariableName + RequestFieldName
		}
	}

	TextConvert, ok := types.MapConvertID[TableName+"."+ColumnName]
	if ok == true {
		RequestFieldGolang = TextConvert + "(" + VariableName + RequestFieldName + ")"
	}

	RequestName = RequestName + "_" + TextRequestProtoName

	return RequestName, RequestFieldName, RequestFieldGolang, GolangLine
}

// IsStringOrUUID - проверяет является ли тип String или UUID
func IsStringOrUUID(TypeGo string) bool {
	if TypeGo == "string" || TypeGo == "uuid.UUID" {
		return true
	}
	return false
}

// Convert_IDToAlias - заменяет "ID" на "alias.Name(ID)"
func Convert_IDToAlias(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := VariableName

	TableName := Table1.Name
	TextConvert, ok := types.MapConvertID[TableName+"."+Column1.Name]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	Otvet = TextConvert + "(" + VariableName + ")"

	return Otvet
}

// Convert_ColumnToAlias - заменяет "Request.Int64_1" на "alias.Name(Request.Int64_1)"
// VariableName = "Request.Int64_1"
func Convert_ColumnToAlias(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := ""

	//Dot := ""
	//if VariableName != "" {
	//	Dot = "."
	//}
	Otvet = VariableName
	//Otvet = VariableName + Dot + Column1.NameGo

	TableName := Table1.Name
	IDName := Column1.Name
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	Otvet = TextConvert + "(" + VariableName + ")"
	//Otvet = TextConvert + "(" + VariableName + Dot + Column1.NameGo + ")"

	//добавим испорт
	//Text = CheckAndAdd_ImportAlias(Text)

	return Otvet
}

// DeleteCommentFromString - удаляет комментарий из строки //
func DeleteCommentFromString(TextFrom string) string {
	Otvet := TextFrom

	pos1 := strings.Index(Otvet, "//")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1]

	return Otvet
}

// Convert_GolangVariableToProtobufVariableID - возвращает имя переменной  + "." +  имя колонки, преобразованное в тип protobuf
func Convert_GolangVariableToProtobufVariableID(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := ""

	if Column1 == nil {
		return Otvet
	}

	Dot := ""
	if VariableName != "" {
		Dot = "."
	}
	Otvet = VariableName + Dot + Column1.NameGo

	//найдём alias
	_, HasAlias := types.MapConvertID[Table1.Name+"."+Column1.Name]

	//преобразуем alias в обычный тип, и дату в timestamp
	if HasAlias == true {
		switch Column1.TypeGo {
		case "time.Time":
			Otvet = "timestamppb.New(" + VariableName + Dot + Column1.NameGo + ")"
		case "string":
			Otvet = "string(" + VariableName + Dot + Column1.NameGo + ")"
		case "int64":
			Otvet = "int64(" + VariableName + Dot + Column1.NameGo + ")"
		case "int32":
			Otvet = "int32(" + VariableName + Dot + Column1.NameGo + ")"
		case "bool":
			Otvet = "bool(" + VariableName + Dot + Column1.NameGo + ")"
		case "float32":
			Otvet = "float32(" + VariableName + Dot + Column1.NameGo + ")"
		case "float64":
			Otvet = "float64(" + VariableName + Dot + Column1.NameGo + ")"
		case "uuid.UUID":
			Otvet = VariableName + Dot + Column1.NameGo + ".String()"
		}

	} else {
		switch Column1.TypeGo {
		case "time.Time":
			Otvet = "timestamppb.New(" + VariableName + Dot + Column1.NameGo + ")"
		case "string":
			Otvet = VariableName + Dot + Column1.NameGo
		case "int64":
			Otvet = VariableName + Dot + Column1.NameGo
		case "int32":
			Otvet = VariableName + Dot + Column1.NameGo
		case "bool":
			Otvet = VariableName + Dot + Column1.NameGo
		case "float32":
			Otvet = VariableName + Dot + Column1.NameGo
		case "float64":
			Otvet = VariableName + Dot + Column1.NameGo
		case "uuid.UUID":
			Otvet = VariableName + Dot + Column1.NameGo + ".String()"
		}
	}

	return Otvet
}

// Convert_GolangVariableToProtobufVariableType - возвращает имя переменной  преобразованное в тип protobuf
func Convert_GolangVariableToProtobufVariableType(Table1 *types.Table, Column1 *types.Column, VariableName, VariableType string) string {
	Otvet := VariableName

	if Column1 == nil {
		return Otvet
	}

	//Dot := ""
	//if VariableName != "" {
	//	Dot = "."
	//}
	//Otvet = VariableName + Dot + Column1.NameGo

	//найдём alias
	_, HasAlias := types.MapConvertID[Table1.Name+"."+Column1.Name]

	TextVariableName := VariableName
	if VariableType != Column1.TypeGo {
		TextVariableName = VariableType + "(" + VariableName + ")"
	}

	//преобразуем alias в обычный тип, и дату в timestamp
	if HasAlias == true {
		switch Column1.TypeGo {
		case "time.Time":
			Otvet = "timestamppb.New(" + VariableName + ")"
		case "string":
			Otvet = VariableType + "(" + VariableName + ")"
		case "int64":
			Otvet = VariableType + "(" + VariableName + ")"
		case "int32":
			Otvet = VariableType + "(" + VariableName + ")"
		case "bool":
			Otvet = VariableType + "(" + VariableName + ")"
		case "float32":
			Otvet = VariableType + "(" + VariableName + ")"
		case "float64":
			Otvet = VariableType + "(" + VariableName + ")"
		case "uuid.UUID":
			Otvet = VariableName + ".String()"
		}

	} else {
		switch Column1.TypeGo {
		case "time.Time":
			Otvet = "timestamppb.New(" + VariableName + ")"
		case "string":
			Otvet = TextVariableName
		case "int64":
			Otvet = TextVariableName
		case "int32":
			Otvet = TextVariableName
		case "bool":
			Otvet = TextVariableName
		case "float32":
			Otvet = TextVariableName
		case "float64":
			Otvet = TextVariableName
		case "uuid.UUID":
			Otvet = VariableName + ".String()"
		}
	}

	return Otvet
}

//// ConvertVariableToProtobufType - возвращает имя переменной +  имя колонки, преобразованное в тип protobuf
//func ConvertVariableToProtobufType(Column1 *types.Column, VariableName string) string {
//	Otvet := ""
//
//	if Column1 == nil {
//		return Otvet
//	}
//
//	Otvet = VariableName
//
//	switch Column1.TypeGo {
//	case "time.Time":
//		Otvet = "timestamppb.New(" + VariableName + ")"
//	case "uuid.UUID":
//		Otvet = VariableName + ".String()"
//	}
//
//	return Otvet
//}

// Convert_ProtobufVariableToGolangVariable - возвращает имя переменной +  имя колонки, преобразованное в тип golang из protobuf
func Convert_ProtobufVariableToGolangVariable(Table1 *types.Table, Column1 *types.Column, VariableName string) (VariableColumn string, GolangCode string) {
	RequestColumnName := Find_RequestFieldName(Table1, Column1)
	VariableColumn = VariableName + RequestColumnName
	//time.Time в timestamppb
	switch Column1.TypeGo {
	case "time.Time":
		{
			VariableColumn = VariableName + RequestColumnName + ".AsTime()"
			return VariableColumn, GolangCode
		}
	case "uuid.UUID":
		{
			VariableColumn = "uuid.FromBytes([]byte(" + VariableName + RequestColumnName + "))"
			GolangCode = Column1.NameGo + `, err := uuid.FromBytes([]byte(Request.` + RequestColumnName + `))
	if err != nil {
		return &Otvet, err
	}
`
			return VariableColumn, GolangCode
		}
	}

	return VariableColumn, GolangCode
}

// Convert_ProtobufTypeToGolangAlias - возвращает имя переменной +  имя колонки, преобразованное в тип golang(alias) из protobuf
func Convert_ProtobufTypeToGolangAlias(Table1 *types.Table, Column1 *types.Column, VariableName string) (VariableColumn string, GolangCode string) {
	RequestColumnName := Find_RequestFieldName(Table1, Column1)
	VariableColumn = VariableName + RequestColumnName
	//GolangCode := ""

	TableName := Table1.Name
	IDName := Column1.Name

	//alias в Int64
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == true {
		VariableColumn = TextConvert + "(" + VariableName + RequestColumnName + ")"
		return VariableColumn, GolangCode
	}

	//time.Time в timestamppb
	switch Column1.TypeGo {
	case "time.Time":
		{
			VariableColumn = VariableName + RequestColumnName + ".AsTime()"
			return VariableColumn, GolangCode
		}
	case "uuid.UUID":
		{
			VariableColumn = "uuid.FromBytes([]byte(" + VariableName + RequestColumnName + "))"
			GolangCode = Column1.NameGo + `, err := uuid.FromBytes([]byte(Request.` + RequestColumnName + `))
	if err != nil {
		return &Otvet, err
	}
`
			return VariableColumn, GolangCode
		}
	}

	return VariableColumn, GolangCode
}

// FindText_EqualEmpty - находит текст сравнение с пустым значением
func FindText_EqualEmpty(Column1 *types.Column, VariableName string) string {
	Otvet := ""

	DefaultValue := FindText_DefaultValue(Column1.TypeGo)
	Otvet = VariableName + " == " + DefaultValue

	if DefaultValue == "time.Time{}" {
		Otvet = VariableName + ".IsZero() == true"
	}

	return Otvet
}

// FindText_NotEqualEmpty - находит текст сравнение с пустым значением
func FindText_NotEqualEmpty(Column1 *types.Column, VariableName string) string {
	Otvet := ""

	DefaultValue := FindText_DefaultValue(Column1.TypeGo)
	Otvet = VariableName + " != " + DefaultValue

	if DefaultValue == "time.Time{}" {
		Otvet = VariableName + ".IsZero() == false"
	}

	return Otvet
}

//// AddSkipNowEveryFunc - добавляет функцию SkipNow() для каждой тестовой функции
//func AddSkipNowEveryFunc(Text string) string {
//	Otvet := ""
//
//	sFind := "\nfunc "
//	Mass := make([]string, 0)
//	Mass = strings.Split(Text, sFind)
//	for _, v := range Mass {
//		pos1 := strings.Index(v, sFind)
//		if pos1 < 0 {
//			continue
//		}
//
//		s2 := Text[pos1:]
//		pos2 := strings.Index(s2, "\n")
//		if pos2 < 0 {
//			continue
//		}
//		v = v[:pos1+pos2] + "\n\tt.SkipNow() //нет строк в БД \n" + v[pos1+pos2:]
//		Otvet = Otvet + v
//	}
//
//	return Otvet
//}

// FilenameWithoutLastUnderline - удаляет последний символ, если он = "_"
func FilenameWithoutLastUnderline(Filename string) string {
	Otvet := Filename

	if strings.HasSuffix(Filename, "_") == true {
		Otvet = Filename[:len(Filename)-1]
	}

	return Otvet
}

// FillVariable - заменяет переменную в тексте
func FillVariable(Text, VariableName, Value string) string {
	Otvet := Text

	sFind := VariableName + " = "
	pos1 := strings.Index(Otvet, sFind)
	if pos1 < 0 {
		return Otvet
	}

	s2 := Text[pos1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1] + VariableName + " = " + Value + Otvet[pos1+posEnd:]

	return Otvet
}

// Is_UUID_Type - проверяет является ли тип UUID
func Is_UUID_Type(TypeGo string) bool {
	Otvet := TypeGo == "uuid.UUID" || TypeGo == "uuid.NullUUID"
	return Otvet
}

// Replace_Postgres_ID_Test - заменяет текст "const Postgres_ID_Test = 0" на нужный ИД
func Replace_Postgres_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_Postgres_ID_Test_ManyPK(Otvet, Table1)

	return Otvet
}

// Replace_Postgres_ID_Test_ManyPK - заменяет текст "const Postgres_ID_Test = 0" на нужные ИД, для много колонок PrimaryKey
func Replace_Postgres_ID_Test_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassPK := Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	//заменим const Postgres_ID_Test = 0
	TextFind := "const Postgres_ID_Test = 0\n"
	TextNew := ""
	for _, PrimaryKey1 := range MassPK {
		TextNew = TextNew + Replace_Postgres_ID_Test1(TextFind, Table1, PrimaryKey1)
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим Otvet.ID = Postgres_ID_Test
	TextFind = "\tOtvet.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "Otvet."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим m.ID = Postgres_ID_Test
	TextFind = "\tm.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим m1.ID = Postgres_ID_Test
	TextFind = "\tm1.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m1."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим ReadFromCache(Postgres_ID_Test)
	TextFind = "ReadFromCache(Postgres_ID_Test)"
	TextNew = "ReadFromCache("
	Comma := ""
	for _, PrimaryKey1 := range MassPK {
		Name := FindText_ColumnNameTest(PrimaryKey1)
		TextNew = TextNew + Comma + Name
		Comma = ", "
	}
	TextNew = TextNew + ")"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//	//удалим лишний код
	//	TextDelete := `	if Otvet.ID != Postgres_ID_Test {
	//		t.Error(TableName + "_test.TestRead() error ID != ", Postgres_ID_Test)
	//	} else {
	//		t.Log(TableName+"_test.TestRead() Otvet: ", Otvet.ID)
	//	}
	//`
	//	Otvet = strings.ReplaceAll(Otvet, TextDelete, "")

	//заменим ненужные Otvet.ID на Otvet.Name
	PrimaryKey1 := MassPK[0]
	Otvet = strings.ReplaceAll(Otvet, " Otvet.ID ", " Otvet."+PrimaryKey1.NameGo+" ")
	Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " Otvet."+PrimaryKey1.NameGo+")")
	Name := FindText_ColumnNameTest(PrimaryKey1)
	Otvet = strings.ReplaceAll(Otvet, "Postgres_ID_Test", Name)

	return Otvet
}

// FindText_ColumnNameTest - находит имя переменной для тестов
func FindText_ColumnNameTest(Column1 *types.Column) string {
	Otvet := ""
	Otvet = strings.ToUpper(Column1.NameGo) + "_Test"
	return Otvet
}

// FindText_NameTest_ManyPK - находит текст "ID, ID" для тестов
func FindText_NameTest_ManyPK(Table1 *types.Table) string {
	Otvet := ""

	MassPK := Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	for _, PrimaryKey1 := range MassPK {
		Otvet = Otvet + FindText_ColumnNameTest(PrimaryKey1) + ", "
	}
	Otvet = strings.TrimSuffix(Otvet, ", ")

	return Otvet
}

// Replace_Postgres_ID_Test1 - заменяет текст "const Postgres_ID_Test = 0" на нужный ИД
func Replace_Postgres_ID_Test1(Text string, Table1 *types.Table, PrimaryKeyColumn *types.Column) string {
	Otvet := Text

	TextFind := "const Postgres_ID_Test = 0"
	IDMinimum := PrimaryKeyColumn.IDMinimum
	if IDMinimum == "" {
		IDMinimum = FindText_DefaultValue(PrimaryKeyColumn.TypeGo)
	}

	Name := FindText_ColumnNameTest(PrimaryKeyColumn)
	sIDMinimum := FindText_IDMinimum(PrimaryKeyColumn)
	switch PrimaryKeyColumn.TypeGo {
	case "uuid.UUID":
		{
			if PrimaryKeyColumn.IDMinimum == "" {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+Name+` = `+sIDMinimum)
			} else {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+Name+`, _ = `+sIDMinimum)
			}
		}
	case "time.Time":
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+Name+` = `+sIDMinimum+``)
		}
	case "string":
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+Name+` = `+sIDMinimum+``)
		}
	default:
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+Name+` = `+sIDMinimum)
		}
	}

	return Otvet
}

// FindText_VariableEqual_ColumnName_Test - возвращает текст для присваивания переменной IDMinimum
func FindText_VariableEqual_ColumnName_Test(Column1 *types.Column, VariableName string) string {
	Otvet := ""

	//TextValue := FindText_IDMinimum(Column1)
	TextValue := FindText_ColumnNameTest(Column1)

	IDMinimum := Column1.IDMinimum
	switch Column1.TypeGo {
	case "uuid.UUID":
		{
			if IDMinimum == "" {
				Otvet = VariableName + " = " + "uuid.Nil"
			} else {
				Otvet = VariableName + ", _ = " + `uuid.Parse("` + IDMinimum + `")`
			}
		}
	default:
		{
			Otvet = VariableName + " = " + TextValue
		}
	}

	return Otvet
}

// FindText_IDMinimum - возвращает текст для IDMinimum, в зависимости от типа
func FindText_IDMinimum(Column1 *types.Column) string {
	Otvet := ""

	IDMinimum := Column1.IDMinimum
	if IDMinimum == "" {
		IDMinimum = FindText_DefaultValue(Column1.TypeGo)
	}

	switch Column1.TypeGo {
	case "uuid.UUID":
		{
			if Column1.IDMinimum == "" {
				Otvet = "uuid.Nil"
			} else {
				Otvet = `uuid.Parse("` + IDMinimum + `")`
			}
		}
	case "string":
		{
			if Column1.IDMinimum == "" {
				return `""`
			}
			Otvet = `"` + IDMinimum + `"`
		}
	default:
		{
			Otvet = `` + IDMinimum + ``
		}
	}

	return Otvet
}

// Replace_Model_ID_Test1 - заменяет текст "const LawsuitStatusType_ID_Test = 0" на нужный ИД
func Replace_Model_ID_Test1(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	TEXT_TEMPLATE_MODEL := config.Settings.TEXT_TEMPLATE_MODEL
	//ModelName := Table1.NameGo
	TextFind := "const " + TEXT_TEMPLATE_MODEL + "_ID_Test = 0"

	IDMinimum := Column1.IDMinimum
	if IDMinimum == "" {
		IDMinimum = FindText_DefaultValue(Column1.TypeGo)
	}
	DefaultModelName := config.Settings.TEXT_TEMPLATE_MODEL

	ModelName := Table1.NameGo
	Name := FindText_ColumnNameTest(Column1)
	switch Column1.TypeGo {
	case "uuid.UUID":
		{
			if Column1.IDMinimum == "" {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+ModelName+"_"+Name+` = `+IDMinimum+``)
			} else {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+ModelName+"_"+Name+`, _ = uuid.Parse("`+IDMinimum+`")`)
			}
			Otvet = strings.ReplaceAll(Otvet, ``+DefaultModelName+`_ID_Test`, ``+ModelName+`_`+Name+`.String()`)
		}
	case "time.Time":
		{
			if Column1.IDMinimum == "" {
				IDMinimum = "time.Time{}"
			}
			Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+ModelName+`_`+Name+` = `+IDMinimum+``)
			//Otvet = strings.ReplaceAll(Otvet, ``+DefaultModelName+`_ID_Test`, `timestamppb.New(`+ModelName+`_`+Name+`)`)
		}
	case "string":
		{
			if Column1.IDMinimum == "" {
				IDMinimum = ""
			}
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+ModelName+`_`+Name+` = "`+IDMinimum+`"`)
		}
	default:
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+ModelName+`_`+Name+` = `+IDMinimum)
		}
	}

	return Otvet
}

// ReplaceText_RequestID_PrimaryKey - заменяет RequestId{} на RequestString{}
func ReplaceText_RequestID_PrimaryKey(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = ReplaceText_RequestID_PrimaryKey_ManyPK(Otvet, Table1)

	return Otvet
}

//// ReplaceText_RequestID_PrimaryKey1 - заменяет RequestId{} на RequestString{}
//func ReplaceText_RequestID_PrimaryKey1(Text string, Table1 *types.Table, VariableName string) string {
//	Otvet := Text
//
//	ColumnPK := Find_PrimaryKeyColumn(Table1)
//	if ColumnPK == nil {
//		return Otvet
//	}
//
//	//TypeGo := ColumnPK.TypeGo
//
//	TextRequestID, TextID := FindText_ProtobufRequest(Table1)
//
//	_, GolangCode := Convert_ProtobufVariableToGolangVariable(Table1, ColumnPK, "Request.")
//	if GolangCode != "" {
//		Otvet = strings.ReplaceAll(Otvet, "ID := "+VariableName+".ID", GolangCode)
//		Otvet = strings.ReplaceAll(Otvet, VariableName+".ID = ", VariableName+"."+TextID+" = ")
//	}
//
//	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
//	Otvet = strings.ReplaceAll(Otvet, "*grpc_proto.RequestId", "*grpc_proto."+TextRequestID)
//	//Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextID)
//	Otvet = strings.ReplaceAll(Otvet, "ID := "+VariableName+".ID", ColumnPK.NameGo+" := "+VariableName+"."+TextID)
//	Otvet = strings.ReplaceAll(Otvet, VariableName+".ID", VariableName+"."+TextID)
//
//	return Otvet
//}

// ReplaceText_RequestID_PrimaryKey_ManyPK - заменяет RequestId{} на RequestString{}
func ReplaceText_RequestID_PrimaryKey_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextRequestID := FindText_ProtobufRequest_ManyPK(Table1)
	TextProto := TextProto()

	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	Otvet = strings.ReplaceAll(Otvet, "*grpc_proto.RequestId", "*"+TextProto+"."+TextRequestID)

	return Otvet
}

// FindText_IDMany - находит все PrimaryKey строкой
func FindText_IDMany(Table1 *types.Table) (TextNames, TextNamesTypes, TextProtoNames string) {
	//TextProtoNames := ""
	//TextNamesTypes := ""
	//TextNames := ""

	TextNames, TextNamesTypes, TextProtoNames = FindText_ID_VariableName_Many(Table1, "")

	//Comma := ""
	//MassPrimaryKey := Find_PrimaryKeyColumns(Table1)
	//for _, PrimaryKey1 := range MassPrimaryKey {
	//	OtvetColumnName := Convert_GolangVariableToProtobufVariableID(Table1, PrimaryKey1, "")
	//	if OtvetColumnName == "" {
	//		continue
	//	}
	//
	//	TextProtoNames = TextProtoNames + Comma + OtvetColumnName
	//	TextNamesTypes = TextNamesTypes + Comma + PrimaryKey1.NameGo + " " + PrimaryKey1.TypeGo
	//	TextNames = TextNames + Comma + PrimaryKey1.NameGo
	//
	//	Comma = ", "
	//}

	return
}

// FindText_ID_VariableName_Many - находит все PrimaryKey строкой
func FindText_ID_VariableName_Many(Table1 *types.Table, VariableName string) (TextNames, TextNamesTypes, TextProtoNames string) {
	//TextProtoNames := ""
	//TextNamesTypes := ""
	//TextNames := ""

	Dot := ""
	if VariableName != "" {
		Dot = "."
	}
	Comma := ""
	MassPrimaryKey := Find_PrimaryKeyColumns(Table1)
	for _, PrimaryKey1 := range MassPrimaryKey {
		OtvetColumnName := Convert_GolangVariableToProtobufVariableID(Table1, PrimaryKey1, "")
		if OtvetColumnName == "" {
			continue
		}

		TextProtoNames = TextProtoNames + Comma + OtvetColumnName
		TextNamesTypes = TextNamesTypes + Comma + PrimaryKey1.NameGo + " " + PrimaryKey1.TypeGo
		TextNames = TextNames + Comma + VariableName + Dot + PrimaryKey1.NameGo

		Comma = ", "
	}

	return
}

//// Replace_OtvetIDEqual1 - заменяет Otvet.ID = -1
//func Replace_OtvetIDEqual1(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	Otvet = Replace_OtvetIDEqual1_ManyPK(Text, Table1)
//
//	return Otvet
//}

//// Replace_OtvetIDEqual1_1 - заменяет Otvet.ID = -1
//func Replace_OtvetIDEqual1_1(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	PrimaryKeyColumn := Find_PrimaryKeyColumn(Table1)
//	if PrimaryKeyColumn == nil {
//		return Otvet
//	}
//
//	Value := Find_NegativeValue(PrimaryKeyColumn.TypeGo)
//
//	Otvet = strings.ReplaceAll(Otvet, "Otvet.ID = -1", "Otvet.ID = "+Value)
//
//	return Otvet
//}

//// Replace_OtvetIDEqual1_ManyPK - заменяет Otvet.ID = -1
//func Replace_OtvetIDEqual1_ManyPK(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	PrimaryKeyColumns := Find_PrimaryKeyColumns(Table1)
//	if len(PrimaryKeyColumns) == 0 {
//		return Otvet
//	}
//
//	TextFind := "\tOtvet.ID = -1\n"
//	TextNew := ""
//	for _, ColumnPK1 := range PrimaryKeyColumns {
//		Value := Find_NegativeValue(ColumnPK1.TypeGo)
//		TextNew = TextNew + "\tOtvet." + ColumnPK1.NameGo + " = " + Value + "\n"
//	}
//
//	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)
//
//	return Otvet
//}

//// Replace_ModelIDEqual1 - заменяет Otvet.ID = -1
//func ReplaceModelIDEqual1_1(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	PrimaryKeyColumn := Find_PrimaryKeyColumn(Table1)
//	if PrimaryKeyColumn == nil {
//		return Otvet
//	}
//
//	Value := Find_NegativeValue(PrimaryKeyColumn.TypeGo)
//
//	Otvet = strings.ReplaceAll(Otvet, "m.ID = -1", "m.ID = "+Value)
//
//	return Otvet
//}

// Find_NegativeValue - возвращает -1 для числовых типов
func Find_NegativeValue(TypeGo string) string {
	Otvet := ""

	Otvet = FindText_DefaultValue(TypeGo)
	if mini_func.IsNumberType(TypeGo) == true {
		Otvet = "-1"
	}

	return Otvet
}

// Find_RequestFieldName - возвращает название колонки в Request
func Find_RequestFieldName(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""

	_, Otvet, _, _ = FindText_ProtobufRequest_ID_Type(Table1, Column1, "")

	return Otvet
}

// Replace_Connect_WithApplicationName - заменяет Connect_WithApplicationName() на Connect_WithApplicationName_SingularTableName()
func Replace_Connect_WithApplicationName(Text string) string {
	Otvet := Text

	if config.Settings.SINGULAR_TABLE_NAMES == false {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "postgres_gorm.Connect_WithApplicationName(", "postgres_gorm.Connect_WithApplicationName_SingularTableName(")
	Otvet = strings.ReplaceAll(Otvet, "postgres_gorm.Start(", "postgres_gorm.Start_SingularTableName(")

	return Otvet
}

// FindText_ConvertToString - возвращает имя переменной +  имя колонки, преобразованное в тип string
func FindText_ConvertToString(Column1 *types.Column, VariableName string) string {
	Otvet := ""

	if Column1 == nil {
		return Otvet
	}
	Otvet = VariableName
	//Otvet = VariableName + "." + Column1.NameGo
	switch Column1.TypeGo {
	case "time.Time":
		Otvet = VariableName + ".String()"
	case "int64":
		Otvet = "strconv.Itoa(int(" + VariableName + "))"
	case "int32":
		Otvet = "strconv.Itoa(int(" + VariableName + "))"
	case "bool":
		Otvet = "strconv.FormatBool(" + VariableName + ")"
	case "float32":
		Otvet = "fmt.Sprintf(%f," + VariableName + ")"
	case "float64":
		Otvet = "fmt.Sprintf(%f," + VariableName + ")"
	case "uuid.UUID":
		Otvet = VariableName + ".String()"
	}

	return Otvet
}

// IsPrimaryKeyColumn - проверяет является ли колонка PrimaryKey
func IsPrimaryKeyColumn(Table1 *types.Table, Column *types.Column) bool {
	Otvet := false

	for _, Column1 := range Table1.MapColumns {
		if Column1.IsPrimaryKey == true && Column1 == Column {
			Otvet = true
			break
		}
	}

	return Otvet
}

// Replace_ColumnNameM - заменяет "Replace_ColumnNameM(m.ID)" на "m.ID"
func Replace_ColumnNameM(Text string, Column *types.Column) string {
	Otvet := Text

	TextNew := "m." + Column.Name
	Otvet = strings.ReplaceAll(Otvet, "Replace_ColumnNameM(m.ID)", TextNew)

	return Otvet
}

// Replace_IntFromProtoRequest - заменяет "IntFromProto(Request.ID)" на "Request.ID"
func Replace_IntFromProtoRequest(Text string, Table1 *types.Table, Column *types.Column) string {
	Otvet := Text

	Variable, Golang_code := Convert_ProtobufVariableToGolangVariable(Table1, Column, "Request.")
	TextNew := "Request." + Variable
	if Golang_code != "" {

	}
	Otvet = strings.ReplaceAll(Otvet, "IntFromProto(Request.ID)", TextNew)

	return Otvet
}

// Find_RequestFieldName_FromMass - возвращает имя поля для Request protobuf
func Find_RequestFieldName_FromMass(Column *types.Column, MassColumns []*types.Column) string {
	Otvet := ""

	TypeProto := Convert_GolangTypeNameToProtobufTypeName(Column.TypeGo)

	Number := 1
	for _, Column1 := range MassColumns {
		TypeProto1 := Convert_GolangTypeNameToProtobufTypeName(Column1.TypeGo)
		if TypeProto == TypeProto1 && Column != Column1 {
			Number = Number + 1
		}

		if Column == Column1 {
			break
		}

	}
	Suffix := "_" + strconv.Itoa(Number)

	Otvet = Convert_GolangTypeNameToProtobufFieldName(Column.TypeGo)
	//if Number > 1 {
	Otvet = Otvet + Suffix
	//}

	return Otvet
}

// Convert_GolangTypeNameToProtobufFieldName - возвращает имя поля для protobuf
func Convert_GolangTypeNameToProtobufFieldName(TypeGo string) string {
	Otvet := ""

	switch TypeGo {
	case "time.Time":
		Otvet = "Date"
	case "string":
		Otvet = "String"
	case "int64", "int":
		Otvet = "Int64"
	case "int32":
		Otvet = "Int32"
	case "int16":
		Otvet = "Int32"
	case "int8":
		Otvet = "Int32"
	case "uint64":
		Otvet = "Uint64"
	case "uint32":
		Otvet = "Uint32"
	case "uint16":
		Otvet = "Uint32"
	case "uint8":
		Otvet = "Uint32"
	case "bool":
		Otvet = "Bool"
	case "float32":
		Otvet = "Float32"
	case "float64":
		Otvet = "Float64"
	case "uuid.UUID":
		Otvet = "String"
	}

	return Otvet
}

// Convert_GolangTypeNameToProtobufTypeName - возвращает имя типа для protobuf
func Convert_GolangTypeNameToProtobufTypeName(TypeGo string) string {
	Otvet := ""

	switch TypeGo {
	case "time.Time":
		Otvet = "google.protobuf.Timestamp"
	case "string":
		Otvet = "string"
	case "int64", "int":
		Otvet = "int64"
	case "int32":
		Otvet = "int32"
	case "uint64":
		Otvet = "uint64"
	case "uint32":
		Otvet = "uint32"
	case "byte":
		Otvet = "uint32"
	case "[]byte":
		Otvet = "bytes"
	case "bool":
		Otvet = "bool"
	case "float32":
		Otvet = "float"
	case "float64":
		Otvet = "double"
	case "uuid.UUID":
		Otvet = "string"
	}

	return Otvet
}

// FindMass_Columns_from_MassColumnsString - преобразует массив строк названий колонок в массив столбцов
func FindMass_Columns_from_MassColumnsString(Table1 *types.Table, MassColumnsString []string) []*types.Column {
	Otvet := make([]*types.Column, len(MassColumnsString))

	for i, ColumnName := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName]
		if ok == false {
			log.Panic(Table1.Name + " error: not found column: " + ColumnName)
			//log.Panic(Table1.Name + " .MapColumns[" + ColumnName + "] = false")
		}
		Otvet[i] = Column1
	}
	return Otvet
}

// FindMass_TableColumns - преобразует TableColumns_String в TableColumns
func FindMass_TableColumns(MapAll map[string]*types.Table, MassTableColumns_String []types.TableColumns_String) []types.TableColumns {
	Otvet := make([]types.TableColumns, 0)

	for _, TableColumns_string1 := range MassTableColumns_String {
		Table1, ok := MapAll[TableColumns_string1.TableName]
		if ok == false {
			log.Warn(" Error: not found table: ", TableColumns_string1.TableName)
			continue
		}
		Columns1 := FindMass_Columns_from_MassColumnsString(Table1, TableColumns_string1.ColumnNames)
		TableColumns1 := types.TableColumns{}
		TableColumns1.Table = Table1
		TableColumns1.Columns = Columns1
		Otvet = append(Otvet, TableColumns1)
	}

	return Otvet
}

// Convert_ProtobufVariableToGolangVariable_with_MassColumns - возвращает имя переменной +  имя колонки, преобразованное в тип golang из protobuf
func Convert_ProtobufVariableToGolangVariable_with_MassColumns(Column *types.Column, MassColumns []*types.Column, VariableName string) (VariableField string, GolangCode string) {
	RequestFieldName := Find_RequestFieldName_FromMass(Column, MassColumns)
	VariableField = VariableName + RequestFieldName

	switch Column.TypeGo {
	case "time.Time":
		{
			VariableField = VariableName + RequestFieldName + ".AsTime()"
			return VariableField, GolangCode
		}
	case "uuid.UUID":
		{
			VariableField = "uuid.FromBytes([]byte(" + VariableName + RequestFieldName + "))"
			GolangCode = Column.NameGo + `, err := uuid.FromBytes([]byte(Request.` + RequestFieldName + `))
	if err != nil {
		return &Otvet, err
	}
`
			return VariableField, GolangCode
		}
	}

	return VariableField, GolangCode
}

//// ConvertGolangVariableToProtobufVariable_with_MassColumns - преобразованное в тип protobuf из golang
//func ConvertGolangVariableToProtobufVariable_with_MassColumns(Column *types.Column, MassColumns []*types.Column) string {
//	Otvet := ""
//
//	Otvet = Find_RequestFieldName_FromMass(Column, MassColumns)
//	return Otvet
//}

// Find_RequestFieldNames_FromMass - возвращает строку с именами колонок для Protobuf
func Find_RequestFieldNames_FromMass(MassColumns []*types.Column) string {
	Otvet := ""

	TextFields := ""
	//TextRequest := ""
	Underline := ""
	for _, Column1 := range MassColumns {
		TextFields = TextFields + Underline + Column1.NameGo
		TextRequest1 := Convert_GolangTypeNameToProtobufTypeName(Column1.TypeGo)
		TextRequest1 = micro.StringFromUpperCase(TextRequest1)
		//TextRequest1 := Find_RequestFieldName_FromMass(Column1, MassColumns)
		Otvet = Otvet + Underline + TextRequest1
		Underline = "_"
	}

	return Otvet
}

// Replace_TemplateModel_to_Model - заменяет текст имя модели в шаблоне на имя модели новое
func Replace_TemplateModel_to_Model(Text, ModelName string) string {
	Otvet := Text

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)

	return Otvet
}

// Replace_TemplateTableName_to_TableName - заменяет текст имя таблицы в шаблоне на имя таблицы новое
func Replace_TemplateTableName_to_TableName(Text, TableName string) string {
	Otvet := Text

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, TableName)

	return Otvet
}

// Replace_ObjectTemplateModel_to_Model - заменяет текст имя модели в шаблоне на имя модели новое
func Replace_ObjectTemplateModel_to_Model(Text, ModelName string) string {
	Otvet := Text

	prefix_object := config.Settings.PREFIX_OBJECT_MODEL
	TextFrom := "Object" + config.Settings.TEXT_TEMPLATE_MODEL
	TextTo := prefix_object + ModelName
	Otvet = strings.ReplaceAll(Otvet, " "+TextFrom, " "+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "*"+TextFrom, "*"+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "."+TextFrom, "."+TextTo)
	//Otvet = strings.ReplaceAll(Otvet, "_"+TextFrom, "_"+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "("+TextFrom, "("+TextTo)

	return Otvet
}

// Replace_ObjectTemplateTableName_to_TableName - заменяет текст имя таблицы в шаблоне на имя таблицы новое
func Replace_ObjectTemplateTableName_to_TableName(Text, TableName string) string {
	Otvet := Text

	prefix_object := config.Settings.PREFIX_READOBJECT
	TextFrom := "object_" + config.Settings.TEXT_TEMPLATE_TABLENAME
	TextTo := prefix_object + TableName
	Otvet = strings.ReplaceAll(Otvet, " "+TextFrom, " "+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "\t"+TextFrom, "\t"+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "("+TextFrom, "("+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "*"+TextFrom, "*"+TextTo)
	Otvet = strings.ReplaceAll(Otvet, "&"+TextFrom, "&"+TextTo)

	return Otvet
}

// AddText_ModuleGenerated - добавляет текст: "Файл создан автоматически"
func AddText_ModuleGenerated(Text string) string {
	Otvet := Text

	Otvet = config.Settings.TEXT_MODULE_GENERATED + Otvet

	return Otvet
}

// CommentLineInText - закомментирует строку в коде, добавит // в начало
func CommentLineInText(Text, TextFind string) string {
	Otvet := Text

	Otvet = strings.ReplaceAll(Otvet, `//`+TextFind, TextFind)

	return Otvet
}

// Find_ColumnNamesWithComma - возвращает строку с именами колонок БД, через запятую
func Find_ColumnNamesWithComma(MassColumns []*types.Column) string {
	Otvet := ""

	Comma := ""
	for _, v := range MassColumns {
		Otvet = Otvet + Comma + v.Name

		Comma = ", "
	}

	return Otvet
}

// Find_PrimaryKeyNamesWithComma - возвращает строку с именами колонок БД Primary key, через запятую
func Find_PrimaryKeyNamesWithComma(Table1 *types.Table) string {
	Otvet := ""

	MassPK := Find_PrimaryKeyColumns(Table1)
	Comma := ""
	for _, v := range MassPK {
		Otvet = Otvet + Comma + v.NameGo
		Comma = ", "
	}

	return Otvet
}

// Find_PrimaryKeysDefaultValues - возвращает строку с значениями по умолчанию колонок БД Primary key, через запятую
func Find_PrimaryKeysDefaultValues(Table1 *types.Table) string {
	Otvet := ""

	MassPrimaryKeys := Find_PrimaryKeyColumns(Table1)

	Comma := ""
	for _, Column1 := range MassPrimaryKeys {
		DefaultValue := FindText_DefaultValue(Column1.TypeGo)
		Otvet = Otvet + Comma + DefaultValue

		Comma = ", "
	}

	return Otvet
}

// AppendColumn - добавляет колонку в слайс, если ее нет
func AppendColumn(MassPK []*types.Column, Column1 *types.Column) []*types.Column {
	Otvet := make([]*types.Column, len(MassPK))
	copy(Otvet, MassPK)

	if slices.Contains(MassPK, Column1) == false {
		Otvet = append(Otvet, Column1)
	}

	return Otvet
}

// ColumnNamesGo_WithComma - возвращает строку с именами колонок БД, через запятую
func ColumnNamesGo_WithComma(ColumnsPK []*types.Column) string {
	Otvet := ""

	Comma := ""
	for _, v := range ColumnsPK {
		Otvet = Otvet + Comma + v.NameGo
		Comma = ", "
	}

	return Otvet
}

// ColumnNamesGo_WithQuotes - возвращает строку с именами колонок БД в кавычках, через запятую
func ColumnNamesGo_WithQuotes(ColumnsPK []*types.Column) string {
	Otvet := ""

	Comma := ""
	for _, v := range ColumnsPK {
		Otvet = Otvet + Comma + `"` + v.NameGo + `"`
		Comma = ", "
	}

	return Otvet
}

// IsGood_Column - возвращает ошибку если колонка неправильная
func IsGood_Column(Column1 *types.Column) error {
	var err error

	err = IsGood_ColumnNamePrefix(Column1)
	if err != nil {
		return err
	}

	err = IsGood_ColumnCommentPrefix(Column1)
	if err != nil {
		return err
	}

	return err
}

// IsGood_ColumnNamePrefix - возвращает ошибку если префикс таблицы = "DELETED_"
func IsGood_ColumnNamePrefix(Column1 *types.Column) error {
	var err error

	ColumnName := Column1.Name
	HasPrefix := strings.HasPrefix(ColumnName, config.Settings.TEXT_DELETED_COLUMN)
	HasPrefixRus := strings.HasPrefix(ColumnName, config.Settings.TEXT_DELETED_COLUMN_RUS)
	if HasPrefix == true || HasPrefixRus == true {
		TextError := fmt.Sprint("Wrong column: ", ColumnName, ", error: name prefix= "+config.Settings.TEXT_DELETED_COLUMN)
		err = errors.New(TextError)
	}

	return err
}

// IsGood_ColumnCommentPrefix - возвращает ошибку если префикс комментария колонки = "DELETED_"
func IsGood_ColumnCommentPrefix(Column1 *types.Column) error {
	var err error

	ColumnComment := Column1.Description
	HasPrefix := strings.HasPrefix(ColumnComment, config.Settings.TEXT_DELETED_COLUMN)
	HasPrefixRus := strings.HasPrefix(ColumnComment, config.Settings.TEXT_DELETED_COLUMN_RUS)
	if HasPrefix == true || HasPrefixRus == true {
		TextError := fmt.Sprint("Wrong column: ", Column1.Name, ", error: comment prefix: ", ColumnComment)
		err = errors.New(TextError)
	}

	return err
}

// CreateDirectory - создает каталог на диске, если его нет
func CreateDirectory(DirectoryName string) {
	var err error

	//создадим каталог
	ok, err := micro.FileExists(DirectoryName)
	if ok == false {
		err = os.MkdirAll(DirectoryName, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirectoryName, " error: ", err)
		}
	}
}

// IsForeignColumn - возвращает true если у этой колонки есть ссылка на другую таблицу (foreign key)
func IsForeignColumn(MapAll map[string]*types.Table, Column1 *types.Column) bool {
	Otvet := false

	TableNameF := Column1.TableKey
	ColumnNameF := Column1.ColumnKey
	if TableNameF == "" || ColumnNameF == "" {
		return Otvet
	}

	Otvet = true

	return Otvet
}

// Find_TableF_ColumnF - для колонки с foreign keys возвращает связанные Table и Column
func Find_TableF_ColumnF(MapAll map[string]*types.Table, Column1 *types.Column) (*types.Table, *types.Column) {
	TableF := &types.Table{}
	ColumnF := &types.Column{}

	TableNameF := Column1.TableKey
	ColumnNameF := Column1.ColumnKey
	if TableNameF == "" || ColumnNameF == "" {
		return TableF, ColumnF
	}

	TableF, ok := MapAll[TableNameF]
	if ok == false {
		return TableF, ColumnF
	}

	ColumnF, ok = TableF.MapColumns[ColumnNameF]
	if ok == false {
		return TableF, ColumnF
	}

	return TableF, ColumnF
}

// Find_FieldNamesWithPercent_from_Mass - возвращает строку в формате "Имя1: %v, Имя2: %v"
func Find_FieldNamesWithPercent_from_Mass(MassColumns []*types.Column) string {
	Otvet := ""

	Comma := ""
	for _, ColumnPK1 := range MassColumns {
		Otvet = Otvet + Comma + ColumnPK1.NameGo + ": %v"
		Comma = ", "
	}

	return Otvet
}

// Find_FieldNamesWithComma_from_Mass - возвращает строку в формате "Имя1: %v, Имя2: %v"
func Find_FieldNamesWithComma_from_Mass(MassColumns []*types.Column) string {
	Otvet := ""
	Comma := ""
	for _, ColumnPK1 := range MassColumns {
		Otvet = Otvet + Comma + ColumnPK1.NameGo
		Comma = ", "
	}

	return Otvet
}

// Find_FieldNamesWithComma_from_Mass_VariableName - возвращает строку в формате "Имя1: %v, Имя2: %v"
func Find_FieldNamesWithComma_from_Mass_VariableName(MassColumns []*types.Column, VariableName string) string {
	Otvet := ""
	Comma := ""

	if VariableName != "" && strings.HasSuffix(VariableName, ".") == false {
		VariableName = VariableName + "."
	}

	for _, ColumnPK1 := range MassColumns {
		Otvet = Otvet + Comma + VariableName + ColumnPK1.NameGo
		Comma = ", "
	}

	return Otvet
}

// Find_FieldNamesWithPercent_from_Table - возвращает строку в формате "Имя1: %v, Имя2: %v"
func Find_FieldNamesWithPercent_from_Table(Table1 *types.Table) string {
	Otvet := ""

	MassPK := Find_PrimaryKeyColumns(Table1)
	Otvet = Find_FieldNamesWithPercent_from_Mass(MassPK)

	return Otvet
}

// Find_FieldNamesWithComma_from_Table - возвращает строку в формате "Имя1: %v, Имя2: %v"
func Find_FieldNamesWithComma_from_Table(Table1 *types.Table) string {
	Otvet := ""

	MassPK := Find_PrimaryKeyColumns(Table1)
	Otvet = Find_FieldNamesWithComma_from_Mass(MassPK)

	return Otvet
}

// Find_FieldNamesWithComma_from_Table_VariableName - возвращает строку в формате "VariableName.Имя1, VariableName.Имя2"
func Find_FieldNamesWithComma_from_Table_VariableName(Table1 *types.Table, VariableName string) string {
	Otvet := ""

	MassPK := Find_PrimaryKeyColumns(Table1)
	Otvet = Find_FieldNamesWithComma_from_Mass_VariableName(MassPK, VariableName)

	return Otvet
}

// Find_ObjectColumnModelName - возвращает имя модели для колонки у Object = "ModelИмяКолонкиБезИД"
func Find_ObjectColumnModelName(Table1 *types.Table, ColumnName string) string {
	Otvet := ColumnName

	//
	SuffixModel := "_model"

	//добавим _model если не кончается на ID
	Otvet = Otvet + SuffixModel

	//
	len1 := len(ColumnName)
	if len1 >= 3 && strings.HasSuffix(ColumnName, "ID") == true {
		Otvet1 := strings.TrimSuffix(ColumnName, "ID")
		len2 := len(Otvet1)
		if len2 > 0 {
			s2 := Otvet1[len2-1:]
			if s2 == strings.ToLower(s2) {
				Otvet = Otvet1
			}
		}
	}

	//
	if Otvet == Table1.NameGo {
		Otvet = Otvet + SuffixModel
	}

	return Otvet
}

//// DeleteSuffixID_small_previous - убирает ID суффикс, только если последняя буква маленькая
//func DeleteSuffixID_small_previous(ModelName string) string {
//	Otvet := ModelName
//
//	//убираем ID суффикс
//	//только если последняя буква маленькая
//	Otvet1 := strings.TrimSuffix(ModelName, "ID")
//	len1 := len(Otvet1)
//	if len1 > 0 {
//		s2 := Otvet1[len1-2:]
//		if s2 == strings.ToLower(s2) {
//			Otvet = Otvet1
//		}
//	}
//
//	return Otvet
//}

// SnakeCase_lower - возвращает строку в формате snake_case, в нижнем регистре
func SnakeCase_lower(Text string) string {
	Otvet := Text

	str := stringy.New(Otvet)
	Otvet = str.SnakeCase("?", "").Get()
	Otvet = strings.ToLower(Otvet)

	return Otvet
}

// Replace_ServiceName_CamelCase - заменяет ServiceNameTemplate на ServiceName + CamelCase
func Replace_ServiceName_CamelCase(Text string) string {
	Otvet := Text

	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME

	ServiceNameTemplate = FormatName(ServiceNameTemplate)
	ServiceName = FormatName(ServiceName)
	Otvet = strings.ReplaceAll(Otvet, ServiceNameTemplate, ServiceName)
	Otvet = strings.ReplaceAll(Otvet, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	return Otvet
}

// Replace_ServiceName - заменяет ServiceNameTemplate на ServiceName
func Replace_ServiceName(Text string) string {
	Otvet := Text

	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME

	Otvet = strings.ReplaceAll(Otvet, ServiceNameTemplate, ServiceName)
	Otvet = strings.ReplaceAll(Otvet, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	return Otvet
}

// TextProto - возвращает текст "grpc_proto"
func TextProto() string {
	Otvet := micro.LastWord(config.Settings.FOLDERNAME_GRPC_PROTO)
	return Otvet
}

// Find_TableAlias - возвращает алиас названия таблицы из 1-3 первых букв
func Find_TableAlias(Table1 *types.Table) string {
	Otvet := ""

	TableName := Table1.Name
	len1 := len(TableName)
	if len1 == 0 {
		return Otvet
	}

	MassRunes := []rune(TableName)
	Otvet = string(MassRunes[0])
	Otvet = strings.ToLower(Otvet)

	var s1 rune
	for _, Rune1 := range MassRunes {

		if s1 == '_' {
			s2 := string(Rune1)
			s2 = strings.ToLower(s2)
			Otvet = Otvet + s2
		}

		s1 = Rune1
	}

	return Otvet
}

// FindText_NullValue - возвращает текст для NullValue (sql.NullString, sql.NullBool, sql.NullInt64)
func FindText_NullValue(TypeGo string, TextVariable string) string {
	Otvet := TextVariable

	switch TypeGo {
	case "time.Time":
		Otvet = "postgres_func.NullTime_DefaultNull(" + TextVariable + ")"
	case "string":
		Otvet = "postgres_func.NullString_DefaultNull(" + TextVariable + ")"
	case "int64", "int":
		Otvet = "postgres_func.NullInt64_DefaultNull(" + TextVariable + ")"
	case "int32":
		Otvet = "postgres_func.NullInt32_DefaultNull(" + TextVariable + ")"
	case "int16":
		Otvet = "postgres_func.NullInt32_DefaultNull(" + TextVariable + ")"
	case "int8":
		Otvet = "postgres_func.NullInt32_DefaultNull(" + TextVariable + ")"
	case "bool":
		Otvet = "postgres_func.NullBool_DefaultNull(" + TextVariable + ")"
	case "float64":
		Otvet = "postgres_func.NullFloat64_DefaultNull(" + TextVariable + ")"
	case "float32":
		Otvet = "postgres_func.NullFloat64_DefaultNull(" + TextVariable + ")"
	case "uuid.UUID":
		Otvet = "postgres_func.NullString_DefaultNull(" + TextVariable + ".String())"
	default:
		log.Error("FindText_NullValue() - неизвестный тип: ", TypeGo)
	}

	return Otvet
}

// FindText_NilValue - возвращает текст, = ссылка на переменную или nil
func FindText_NilValue(TypeGo string, TextVariable string) string {
	Otvet := TextVariable

	switch TypeGo {
	case "time.Time":
		Otvet = "micro.Time_DefaultNil(" + TextVariable + ")"
	case "string":
		Otvet = "micro.String_DefaultNil(" + TextVariable + ")"
	case "int64", "int":
		Otvet = "micro.Int64_DefaultNil(" + TextVariable + ")"
	case "int32":
		Otvet = "micro.Int32_DefaultNil(" + TextVariable + ")"
	case "int16":
		Otvet = "micro.Int32_DefaultNil(" + TextVariable + ")"
	case "int8":
		Otvet = "micro.Int32_DefaultNil(" + TextVariable + ")"
	case "bool":
		Otvet = "micro.Bool_DefaultNil(" + TextVariable + ")"
	case "float64":
		Otvet = "micro.Float64_DefaultNil(" + TextVariable + ")"
	case "float32":
		Otvet = "micro.Float64_DefaultNil(" + TextVariable + ")"
	case "uuid.UUID":
		Otvet = "micro.String_DefaultNil(" + TextVariable + ".String())"
	default:
		log.Error("FindText_NilValue() - неизвестный тип: ", TypeGo)
	}

	return Otvet
}
