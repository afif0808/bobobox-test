package sqls

import (
	"reflect"
	"strings"
)

func GenerateInsertQuery(table string, data interface{}) (query string, args []interface{}) {
	query = "INSERT INTO " + table + " SET "
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		field, exist := v.Type().Field(i).Tag.Lookup("db")
		if !exist {
			continue
		}
		action, exist := v.Type().Field(i).Tag.Lookup("action")
		if !exist || !strings.Contains(action, "create") {
			continue
		}
		query += field + " = ? ,"
		args = append(args, v.Field(i).Interface())
	}
	query = strings.TrimSuffix(query, ",")
	return
}

func GenerateUpdateQuery(table string, data interface{}) (query string, args []interface{}) {
	query = "UPDATE " + table + " SET "
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		field, exist := v.Type().Field(i).Tag.Lookup("db")
		if !exist {
			continue
		}
		action, exist := v.Type().Field(i).Tag.Lookup("action")
		if !exist || !strings.Contains(action, "update") {
			continue
		}
		query += field + " = ? ,"
		args = append(args, v.Field(i).Interface())
	}

	query = strings.TrimSuffix(query, ",")
	return
}
func GenerateUpdateByIDQuery(table string, data, id interface{}) (query string, args []interface{}) {
	query = "UPDATE " + table + " SET "
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		field, exist := v.Type().Field(i).Tag.Lookup("db")
		if !exist {
			continue
		}
		action, exist := v.Type().Field(i).Tag.Lookup("action")
		if !exist || !strings.Contains(action, "update") {
			continue
		}
		query += field + " = ? ,"
		args = append(args, v.Field(i).Interface())
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = ?"

	args = append(args, id)

	return
}

func GenerateArgs(data interface{}) []interface{} {
	args := []interface{}{}
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		args = append(args, v.Field(i).Interface())
	}
	return args
}
