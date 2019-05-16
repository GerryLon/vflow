package netflow9

import (
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v2"
)

func init() {

}

type DefinitionModel map[uint16][]interface{}

func ParseDefinition(cfgPath string) (model DefinitionModel, err error) {
	var (
		file     = path.Join(cfgPath, "netflow_v9.yaml")
		contents []byte
	)
	contents, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	model = DefinitionModel{}
	err = yaml.Unmarshal(contents, &model)
	return
}

// add field name for netflow v9
func PatchFieldName(model *DefinitionModel, dataSets [][]DecodedField) (tm map[string]interface{}, err error) {
	tm = make(map[string]interface{}, 0)
	for _, v1 := range dataSets {
		for _, v2 := range v1 {
			for k, v3 := range *model {
				entry := v3
				length := len(entry)
				if length < 2 {
					if length == 1 && entry[0].(string) == ":skip" {
						continue
					}
					log.Printf("error:, k=%v, len(entry)=%d\n", k, len(entry))
					continue
				}

				// TODO: typ not used
				// typ := entry[0]
				name := entry[1].(string)
				name = name[1:]

				if v2.ID == k {
					tm[name] = v2.Value
				}
			}
		}
	}
	return
}
