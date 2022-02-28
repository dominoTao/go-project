package utils

func ConvObj2String(input interface{}) string {
	if input != nil {
		if s, ok := input.(string); ok{
			return s
		}
	}
	return ""
}
