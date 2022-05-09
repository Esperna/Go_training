// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Pack(ptr interface{}) (*url.URL, error) {
	var url url.URL
	url.Scheme = "http"
	url.Host = "localhost:12345"
	url.Path = "search"
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			url.RawQuery += "&"
		}
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get(url.Scheme)
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		if v.Field(i).Kind() == reflect.Slice {
			for j := 0; j < v.Field(i).Len(); j++ {
				if j > 0 {
					url.RawQuery += "&"
				}
				url.RawQuery += fmt.Sprintf("%s=%s", name, v.Field(i).Index(j))
			}
		} else {
			value := v.Field(i)
			url.RawQuery += fmt.Sprintf("%s=%v", name, value)
		}
	}
	return &url, nil
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	options := make(map[string]string)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		args := strings.Split(name, ",")
		length := len(args)
		if length == 1 {
			if name == "" {
				name = strings.ToLower(fieldInfo.Name)
			}
		} else if length == 2 {
			name = args[0]
			options[name] = args[1]
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if options[name] != "" {
					if !isValid(options[name], value) {
						return fmt.Errorf("invalid %s:%s", options[name], value)
					}
				}

				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func isValid(option, value string) bool {
	if option == "mail" {
		r := regexp.MustCompile(`^[0-9A-Za-z.+-][0-9A-Za-z.+-]+@([A-Za-z]+\.)[A-Za-z]+`)
		return r.MatchString(value)
	} else if option == "number" {
		r := regexp.MustCompile(`^(\d{4}-)(\d{4}-){2}\d{4}`)
		return r.MatchString(value)
	} else if option == "code" {
		r := regexp.MustCompile(`^\d\d{2}-\d{4}`)
		return r.MatchString(value)
	}
	return true
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
