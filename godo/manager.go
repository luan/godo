package godo

type manager struct {
}


func (m *manager) Add(model interface{}) error {
	err := Dbmap.Insert(model)
	return err
}

func (m *manager) FindAll(array interface {}, tableName string) (err error) {
	err = Dbmap.Select(array, "select * from "+ tableName +" order by id")
	return
}

func (m *manager) Update(dest interface {}) (err error) {
	_ , err = Dbmap.Update(dest)
	return
}

func (m *manager) Find(id int, dest interface {}) (err error) {
	err = Dbmap.Get(dest, id)
	return
}


