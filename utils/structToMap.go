package utils

import "encoding/json"

func StructToMap(in interface{}) map[string]interface{} {
	var myMap map[string]interface{}
	data, _ := json.Marshal(in)
	json.Unmarshal(data, &myMap)

	for key:= range myMap {
		if myMap[key] == nil {
			delete(myMap, key)
		}
	}

	return myMap
}