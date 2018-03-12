package jsonlogic

func sumOp(value interface{}, other interface{}) float64 {
	floatVal := interFaceToFloat(value)
	floatOther := interFaceToFloat(other)
	return floatVal + floatOther
}

func multOp(value interface{}, other interface{}) float64 {
	floatVal := interFaceToFloat(value)
	floatOther := interFaceToFloat(other)
	return floatVal * floatOther
}

func subOp(value interface{}, other interface{}) float64 {
	floatVal := interFaceToFloat(value)
	floatOther := interFaceToFloat(other)
	return floatVal - floatOther
}

func divOp(value interface{}, other interface{}) float64 {
	floatVal := interFaceToFloat(value)
	floatOther := interFaceToFloat(other)
	return floatVal / floatOther
}

func modOp(value interface{}, other interface{}) float64 {
	val := int(interFaceToFloat(value))
	otherVal := int(interFaceToFloat(other))
	return float64(val % otherVal)
}
