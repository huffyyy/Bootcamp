package departement

type Departement struct {
	id              int64
	departementName string
}

func NewDepartement(id int64, departementName string) *Departement {
	return &Departement{
		id:              id,
		departementName: departementName,
	}

}
func (d *Departement) GetId() int64 {
	if d != nil {
		return d.id
	}
	return 0
}

func (d *Departement) SetId(id int64) {
	if d != nil {
		d.id = id
	}
}

func (d *Departement) GetDepartementName() string {
	if d != nil {
		return d.departementName
	}
	return ""
}

func (d *Departement) SetDepartementName(departementName string) {
	if d != nil {
		d.departementName = departementName
	}
}
