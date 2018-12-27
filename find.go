package mongo

import "github.com/globalsign/mgo/bson"

// Find all in collection
func (u *Mongo) FindAll() ([]interface{}, error) {
	var data []interface{}
	err := db.C(u.Collection).Find(bson.M{}).All(&data)
	return data, err
}

// Find by all by filter in collection
func (u *Mongo) FindByFilter(filter interface{}) (interface{}, error) {
	var data interface{}
	err := db.C(u.Collection).Find(filter).One(&data)
	return data, err
}

// Find document by id in collection
func (u *Mongo) FindById(id string) (interface{}, error) {
	var data interface{}
	err := db.C(u.Collection).FindId(bson.ObjectIdHex(id)).One(&data)
	return data, err
}
