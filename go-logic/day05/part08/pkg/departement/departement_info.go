package departement

import (
	"encoding/json"
	"fmt"
)

func (d *Departement) ToString() string {
	return fmt.Sprintf("Id : %d, DepartementName : %s",
		d.id, d.departementName)
}

func (d *Departement) ToJson() (string, error) {
	data := map[string]any{
		"id":        d.id,
		"departementName": d.departementName,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}