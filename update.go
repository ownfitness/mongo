package mongo

// Update object by id in collection
func (u *Mongo) Update(id string, data interface{}) error {
	err := db.C(u.Collection).UpdateId(id, &data)
	return err
}
