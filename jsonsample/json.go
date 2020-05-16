package jsonsample

import (
	"encoding/json"
	"fmt"
	"log"
)

func Run() {
	jsonString := `{"type":"web","filePath":"/data/jobs/activator/prepackage_-_startapp_pp_2020011008_v_soda_node0005.tsv_part_9083","translateSegmentsPlatform":45,"segmentsIncluded":true,"fileFormat":"withHeaders","optimized":true,"fileHeaders":{"device_type":0,"country_code":3,"segments":2,"device_id":1}}`

	var aux map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &aux); err != nil {
		log.Fatal(err)
	}

	fmt.Println(aux)
}
