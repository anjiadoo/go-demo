package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Class string `bson:"class"`
}

// get mongodb db   连接获取mongo  db
func getDB() *mgo.Database {
	session, err := mgo.Dial("172.16.27.134:10001")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("test")
	return db
}

func main() {
	//多条件查询
	var query []bson.M
	//以TOM为name条件进行模糊查询
	q1 := bson.M{"name": bson.M{"$regex": "TOM", "$options": "$im"}}
	query = append(query, q1)

	//以年龄大于10来进行查询
	q2 := bson.M{"age": bson.M{"gt": 10}}
	query = append(query, q2)

	//以班级为终极一班来查
	q3 := bson.M{"class": "终极一班"}
	query = append(query, q3)

	//查询语句
	var userArray []User
	err := getDB().C("user").Find(bson.M{"$and": query}).All(&userArray)
	if err != nil {
		fmt.Println("mongodb error!!", err)
	}
}