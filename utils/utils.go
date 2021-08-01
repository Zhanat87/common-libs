package utils

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const (
	MegaByte = 1048576
	GigaByte = 1000 * MegaByte
)

func ReadErrorsChannel(errorsChan chan error, goroutinesCount int) error {
	var errorsCount int
	var errors []error
	for err := range errorsChan {
		if err != nil {
			errors = append(errors, err)
		}
		errorsCount++
		if errorsCount == goroutinesCount {
			close(errorsChan)
		}
	}

	return GetErrorFromErrors(errors)
}

func GetErrorFromErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	errorText := ""
	for i, err := range errs {
		errorText += fmt.Sprintf("#%d %#v\r\n", i, err)
	}

	return errors.New(errorText)
}

func GetMegabytes(mb int) int {
	return mb * MegaByte
}

func GetGigabytes(gb int) int {
	return gb * GigaByte
}

func GetCurrentDate(timeZone string) (string, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc)

	return now.Format("2006-01-02"), nil
}

func GetCurrentDateTime(timeZone string) (string, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc)

	return now.Format("2006-01-02 15:04:05"), nil
}

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

func GetTraceIDAndSpanNameFromContext(ctx context.Context, prefix string) (string, string) {
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
				if reflectValueInterfaceSlice[0] == "span" && reflectValueInterfaceSlice[2] == prefix {
					return reflectValueInterfaceSlice[1], strings.Join(reflectValueInterfaceSlice[2:], " ")
				}
			}
		}
	}

	return "", ""
}
