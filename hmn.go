package hmn

import (
	"bytes"
	"reflect"
	"fmt"
	"time"
	"strconv"
	"strings"
	"regexp"
)

var camelRegex = regexp.MustCompile("[0-9A-Za-z]+")

func Load(t interface{}, line string) error {
	// Find the actual value of the object
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	// Split using a whitespace
	keyVals := strings.Split(line, " ")
	for i := range keyVals {
		// Take key, val pairs so 2 at a time
		i = 2 * i
		if i >= len(keyVals) {
			break
		}

		// Convert key to CamelCase
		key := CamelCase(keyVals[i])
		val := keyVals[i + 1]

		field := v.FieldByName(key)
		if err := LoadField(t, field, key, val); err != nil {
			return err
		}
	}
	return nil
}

// LoadField ...
func LoadField(t interface{}, field reflect.Value, key, val string) error {
	// Figure out the kind, then load it up
	switch field.Kind() {
	case reflect.Int32, reflect.Int64, reflect.Int:
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(value)
	case reflect.Float64, reflect.Float32:
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		field.SetFloat(value)
	case reflect.String:
		field.SetString(val)
	case reflect.Struct:
		LoadStruct(t, field, val)
	case reflect.Invalid:
		panic(
			fmt.Sprintf(
				"Invalid type, check key name\nKey `%s\nGo to github.com/johnmcconnell/hmn/hmm.go to add handling.",
				key))
	default:
		panic(
			fmt.Sprintf(
				"Unsure how to handle Kind `%s.\nKind `%s\nGo to github.com/johnmcconnell/hmn/hmm.go to add handling.",
				field,
				field.Kind()))
	}
	return nil
}

// LoadStruct ...
func LoadStruct(t interface{}, field reflect.Value, val string) {
	// Figure out the type, then load it up
	name := field.Type().Name()
	switch name {
	case "Time":
		time, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		field.Set(reflect.ValueOf(time))
	default:
		panic(
			fmt.Sprintf(
				"Unsure how to handle `%s.\nType `%s\nGo to github.com/johnmcconnell/hmn/hmm.go to add handling.",
				name,
				t))
	}
}

// CamelCase ...
func CamelCase(src string) string {
  byteSrc := []byte(src)
  chunks := camelRegex.FindAll(byteSrc, -1)
  for idx, val := range chunks {
    chunks[idx] = bytes.Title(val)
  }
  return string(bytes.Join(chunks, nil))
}

