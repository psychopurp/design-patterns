package factory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	env1 := "production"
	env2 := "development"

	db1 := DatabaseFactory(env1)
	db2 := DatabaseFactory(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fmt.Println(reflect.TypeOf(db1).Elem())
	fmt.Println(reflect.TypeOf(&db1).Elem())

	fmt.Println(reflect.TypeOf(db2).Elem())
	fmt.Println(reflect.TypeOf(&db2).Elem())

}
