package refl

import (
	"strings"

	"github.com/goccy/go-reflect"
	"github.com/just-do-halee/refl/kind"
)

type Type = reflect.Type
type Value = reflect.Value
type Tag = reflect.StructTag
type Kind = reflect.Kind

func IsBothType(a, b any) bool {
	if _, ok := a.(Type); ok {
		if _, ok := b.(Type); ok {
			return true
		}
	}
	return false
}

func IsBothValue(a, b any) bool {
	if _, ok := a.(Value); ok {
		if _, ok := b.(Value); ok {
			return true
		}
	}
	return false
}

func NameEqAny(a, b any) bool {
	if IsBothType(a, b) {
		return NameEq(a.(Type), b.(Type))
	}
	return NameEq(reflect.TypeOf(a), reflect.TypeOf(b))
}

func NameEqAnyWithGeneric(a, b any) bool {
	if IsBothType(a, b) {
		return NameEqWithGeneric(a.(Type), b.(Type))
	}
	return NameEqWithGeneric(reflect.TypeOf(a), reflect.TypeOf(b))
}

func NameEq(a, b Type) bool {
	aName, _ := GetTypeName(a)
	bName, _ := GetTypeName(b)
	return aName == bName
}

func NameEqWithGeneric(a, b Type) bool {
	aName, aGeneric := GetTypeName(a)
	bName, bGeneric := GetTypeName(b)
	return aName == bName && aGeneric == bGeneric
}

func GetTypeNameAny(i any) (name string, generic string) {
	if t, ok := i.(Type); ok {
		return GetTypeName(t)
	}
	return GetTypeName(reflect.TypeOf(i))
}

func GetTypeName(t Type) (origin string, generic string) {
	origin, generic, _ = strings.Cut(t.Name(), "[")
	return
}

func TypeEq(a, b any) bool {
	switch a.(type) {
	case Type:
		switch b.(type) {
		case Type:
			if a == b {
				return true
			}
		}
	}
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func ValueEq(a, b any) bool {
	switch a.(type) {
	case Value:
		switch b.(type) {
		case Value:
			if a == b {
				return true
			}
		}
	}
	return reflect.ValueOf(a) == reflect.ValueOf(b)
}

type Option struct {
	Unwrap bool
}

func Get(value any, o *Option) (t Type, v Value) {
	t = reflect.TypeOf(value)
	v = reflect.ValueOf(value)
	if o.Unwrap {
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
			v = v.Elem()
		}
	}
	return
}

type Struct struct {
	Type  Type
	Value Value
	Size  int
}

func GetStruct(i any) Struct {
	t, v := Get(i, &Option{Unwrap: true})
	if t.Kind() != kind.Struct {
		panic("The type " + t.Name() + " must be a struct")
	}
	s := t.NumField()
	return Struct{
		Type:  t,
		Value: v,
		Size:  s,
	}
}

type StructField struct {
	reflect.StructField
	Parent Struct
	Index  int
	Type   Type
	Value  Value
	Tag    Tag
}

func UnwrapType(t Type) (u Type) {
	u = t
	for u.Kind() == kind.Ptr {
		u = u.Elem()
	}
	return
}

func UnwrapValue(v Value) (u Value) {
	u = v
	for u.Kind() == kind.Ptr {
		u = u.Elem()
	}
	return
}
