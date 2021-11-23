package sexenum

var (
	sexEnum = []map[string]interface{}{
		{
			"key":   "male",
			"label": "Pria",
		},
		{
			"key":   "female",
			"label": "Wanita",
		},
	}
)

//Get enums
func GetEnums() []map[string]interface{} {
	return sexEnum
}

//Get enum from key
func GetEnumFromKey(key string) map[string]interface{} {
	for _, sex := range sexEnum {
		if sex["key"] == key {
			return sex
		}
	}

	return nil
}

//Get key
func GetKey(key string) string {
	for _, sex := range sexEnum {
		if sex["key"] == key {
			return sex["key"].(string)
		}
	}

	return ""
}

//Get label from key
func GetLabelFromKey(key string) string {
	for _, sex := range sexEnum {
		if sex["key"] == key {
			return sex["label"].(string)
		}
	}

	return ""
}
