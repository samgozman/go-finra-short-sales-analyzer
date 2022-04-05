package tester

import (
	"fmt"
	"reflect"
	"testing"
)

// Helper to compare boolean filters etc.
func Compare[C comparable](t *testing.T, want C, got C, attributes ...any) {
	if want != got {
		str := "Error"

		if attributes != nil {
			str = fmt.Sprintf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		} else {
			str = fmt.Sprintf("Expected '%v', but got '%v'", want, got)
		}

		t.Errorf(str)
	}
}

// Helper to compare structs
func StructCompare(t *testing.T, want any, got any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

type comparable interface {
	int | int64 | uint | float32 | float64 | uint64 | bool | string
}
