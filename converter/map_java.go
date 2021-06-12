package converter

func FromMysqlType(typ string) string {
	switch typ {
	case "bit":
		return "Boolean"
	case "byte":
		return " Byte"
	case "short":
		return " Short"
	case "int":
		return "Integer"
	case "smallint", "tinyint":
		return "Integer"
	case "bigint":
		return "Long"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "decimal", "numeric":
		return "BigDecimal"
	case "date":
		return "Date"
	case "datetime", "timestamp":
		return "Timestamp"
	case "time":
		return "Time"
	case "year":
		return "Short"
	case "varchar", "char", "text":
		return "String"
	}
	return typ
}
