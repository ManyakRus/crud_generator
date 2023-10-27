package mini_func

func IsNumberType(TypeGo string) bool {
	Otvet := false

	switch TypeGo {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64", "uint", "uint8", "uint16", "uint32", "uint64", "byte":
		{
			Otvet = true
		}
	}

	return Otvet
}
