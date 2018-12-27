package mongo

func (u *Mongo) Insert(data interface{}) error {
	err := db.C(u.Collection).Insert(&data)
	return err
}
