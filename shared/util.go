package shared

// MapToProto 转换函数：将 map[string][]string 转换为 map[string]*StringList
func MapToProto(input map[string][]string) map[string]*StringList {
	result := make(map[string]*StringList)

	for key, value := range input {
		result[key] = &StringList{Values: value}
	}

	return result
}

// ProtoToMap 转换函数：将 map[string]*StringList  转换为 map[string][]string
func ProtoToMap(input map[string]*StringList) map[string][]string {
	result := make(map[string][]string)

	for key, value := range input {
		result[key] = value.Values
	}

	return result
}
