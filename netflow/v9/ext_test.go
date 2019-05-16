package netflow9

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPatchFieldName(t *testing.T) {
	model, err := ParseDefinition(`D:\workspace\code\go\src\github.com\VerizonDigital\vflow\scripts`)
	if err != nil {
		t.Error("err: open file")
	}
	var dataSets [][]DecodedField
	var dataSetsStr = `
[[{"I":8,"V":"10.0.0.2"},{"I":12,"V":"10.0.0.3"},{"I":15,"V":"0.0.0.0"},{"I":10,"V":3},{"I":14,"V":5},{"I":2,"V":"0x00000001"},{"I":1,"V":"0x00000040"},{"I":7,"V":4242},{"I":11,"V":80},{"I":6,"V":"0x00"},{"I":4,"V":17},{"I":5,"V":1},{"I":17,"V":"0x0003"},{"I":16,"V":"0x0002"},{"I":9,"V":32},{"I":13,"V":31},{"I":21,"V":45964903},{"I":22,"V":45904903}]]	
`
	if err := json.Unmarshal([]byte(dataSetsStr), &dataSets); err != nil {
		t.Error("err: unmarshal")
	}

	tm, err := PatchFieldName(&model, dataSets)
	if err != nil {
		t.Error("err: PatchFieldName()")
	}
	c, _ := json.MarshalIndent(tm, "  ", "  ")
	fmt.Printf("%s", c)
}
