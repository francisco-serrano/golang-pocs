package dynamicjson

import (
	"encoding/json"
	"fmt"
)

type DynamicStructure struct {
	Filters []FilterParams `json:"filters"`
}

type FilterParams struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func Run() {
	myJson := `{"filter":[{"key":"id_source","value":1},{"key":"name","value":"lucas"}]}`

	sampleStructure := DynamicStructure{Filters: []FilterParams{
		{
			Key:   "id_source",
			Value: 1,
		},
		{
			Key:   "name",
			Value: "lucas",
		},
	}}

	fmt.Println(myJson)
	fmt.Println(sampleStructure)

	byteArr, err := json.Marshal(&sampleStructure)
	if err != nil {
		panic(err)
	}

	fmt.Println(byteArr)

	var anotherStructure DynamicStructure
	if err := json.Unmarshal([]byte(byteArr), &anotherStructure); err != nil {
		panic(err)
	}

	fmt.Println(anotherStructure)
}
