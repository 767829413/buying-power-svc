package enumer

import (
	"reflect"

	"github.com/pkg/errors"
)

func WheelPositionEncoder(v reflect.Value) string {
	s, ok := v.Interface().(WheelPosition)
	if !ok {
		panic(errors.Errorf("cannot assert %T as WheelType", v.Interface()))
	}
	return s.String()
}

func (v WheelPosition) IsRight() bool {
	_, ok := getState().WheelPositionIndicatives.IsRight[int32(v)]
	return ok
}

func (v WheelPosition) IsLeft() bool {
	_, ok := getState().WheelPositionIndicatives.IsLeft[int32(v)]
	return ok
}
