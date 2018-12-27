package mongo

// Delete document in collection
func (u *Mongo) Delete(data interface{}) error {
	err := db.C(u.Collection).Remove(&data)
	return err
}
