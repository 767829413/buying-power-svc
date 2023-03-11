package enumer

import (
	"fmt"
	"reflect"
)

func (t TransmissionType) IsManual() bool {
	_, ok := getState().TransmissionTypeIndicatives.IsManual[int32(t)]
	return ok
}

func TransmissionTypeEncoder(v reflect.Value) string {
	s, ok := v.Interface().(TransmissionType)
	if !ok {
		panic(fmt.Errorf("cannot assert %T as TransmissionType", v.Interface()))
	}
	return s.String()
}
