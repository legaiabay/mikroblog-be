package helpers

func ArrayKeys(arr map[string]map[string]interface{}) []string {
	keys := make([]string, 0, len(arr))
	for key := range arr {
		keys = append(keys, key)
	}

	return keys
}
