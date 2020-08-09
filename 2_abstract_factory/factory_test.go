package factory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstrcutors(env1)
	db2, fs2 := SetupConstrcutors(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fs1.CreateFile("test.txt")
	fmt.Println(fs1.FindFile("test.txt"))

	fs2.CreateFile("test.txt")
	fmt.Println(fs2.FindFile("test.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(fs1).Name())
	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(fs2).Name())
	fmt.Println(reflect.TypeOf(&fs2).Elem())

}
