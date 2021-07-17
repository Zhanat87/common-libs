package utils

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func PrintContextInternals(label string, ctx interface{}, inner bool) {
	println()
	println(label, "------------------------------------------------")
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()
	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}
	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				PrintContextInternals(label, reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
	println(label, "------------------------------------------------")
	println()
}

func GetContextInternals(label string, ctx interface{}, inner bool) {
	println()
	println(label, "------------------------------------------------")
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()
	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}
	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				PrintContextInternals(label, reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
	println(label, "------------------------------------------------")
	println()
}

func GetTraceIDFromContext(ctx context.Context, prefix string) string {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()
	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()
			reflectField := contextKeys.Field(i)

			if reflectField.Name == "val" {
				reflectValueInterfaceString := fmt.Sprintf("%s", reflectValue.Interface())
				reflectValueInterfaceSlice := strings.Split(reflectValueInterfaceString, " ")
				if reflectValueInterfaceSlice[0] == "span" && strings.HasPrefix(reflectValueInterfaceSlice[2], "\""+prefix) {
					return reflectValueInterfaceSlice[1]
				}
			}
		}
	}

	return ""
}
