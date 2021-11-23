package sellerstatusenum

var (
	sellerStatusEnum = []map[string]interface{}{
		{
			"key":   "waiting_approval",
			"label": "Menunggu Persetujuan",
		},
		{
			"key":   "reject",
			"label": "Ditolak",
		},
		{
			"key":   "active",
			"label": "Aktif",
		},
		{
			"key":   "inactive",
			"label": "Tidak Aktif",
		},
	}
)

//Get enums
func GetEnums() []map[string]interface{} {
	return sellerStatusEnum
}

//Get enum from key
func GetEnumFromKey(key string) map[string]interface{} {
	for _, sellerStatus := range sellerStatusEnum {
		if sellerStatus["key"] == key {
			return sellerStatus
		}
	}

	return nil
}

//Get key
func GetKey(key string) string {
	for _, sellerStatus := range sellerStatusEnum {
		if sellerStatus["key"] == key {
			return sellerStatus["key"].(string)
		}
	}

	return ""
}

//Get label from key
func GetLabelFromKey(key string) string {
	for _, sellerStatus := range sellerStatusEnum {
		if sellerStatus["key"] == key {
			return sellerStatus["label"].(string)
		}
	}

	return ""
}
