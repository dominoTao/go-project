package utils

type Any map[string]interface{}

func GetParamsStringOfMap(i Any, key string) string {
	if i != nil && i[key] != nil {
		if _, ok := i[key].(string);ok {
			return i[key].(string)
		}
	}
	return ""
}

func GetParamsIntOfMap(i Any, key string) int {
	if i != nil && i[key] != nil {
		if _, ok := i[key].(int);ok {
			return i[key].(int)
		}
	}
	return 0
}

func GetParamsFloat64OfMap(i Any, key string) float64 {
	if i != nil && i[key] != nil {
		if _, ok := i[key].(float64);ok {
			return i[key].(float64)
		}
	}
	val := GetParamsInterfaceOfMap(i, key)
	if val == nil {
		return 0
	}
	if _, ok := val.(float64); ok {
		return val.(float64)
	}
	return 0
}


func GetParamsInterfaceOfMap(i Any, key string) interface{} {
	if i != nil && i[key] != nil {
		return i[key]
	}
	return nil
}