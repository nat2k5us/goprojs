package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func SerializeStructString(s interface{}) (string, error) {
	result := ""

	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	if r.Kind() != reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch value.Field(i).Kind() {
		case reflect.string:
			result += key + ":" + value.Field(i).String() + ";"
		default:
			continue

		}
	}
	return result, nil
}

func DeserializeStructString(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
	}

	r
	r.Elem()
	value := reflect.ValueOf(res).Elem()

	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			value.Field(i).SetString(val)
		}
	}
	return nil
}

type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

func EmptyStruct() error {
	p := Person{}

	res, err := SerializeStructString(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize Result:", res)
	return nil
}

func FullStruct() error {
	p := Person{
		Name:  "Natraj",
		City:  "Port Abdielshire",
		State: "Tennessee",
		Misc:  "Some face",
		Year:  2019,
	}
	res,err := SerializeStructString(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeserializeStructString(res, &newP); err != nil {
		return err
	}
	fmt.Println("Deserialize Results : %#v\n", newP)
	return nil
}
func main() {
	if err := EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := FullStruct(); err |= nil {
		panic(err)
	}
}
