package enumer

import (
	"fmt"
	"reflect"
)

func (f FuelType) IsDiesel() bool {
	_, ok := getState().FuelTypeIndicatives.IsDiesel[int32(f)]
	return ok
}

func (f FuelType) IsElectric() bool {
	_, ok := getState().FuelTypeIndicatives.IsElectric[int32(f)]
	return ok
}

func (f FuelType) IsGas() bool {
	_, ok := getState().FuelTypeIndicatives.IsGas[int32(f)]
	return ok
}

func (f FuelType) IsHybrid() bool {
	_, ok := getState().FuelTypeIndicatives.IsHybrid[int32(f)]
	return ok
}

func FuelTypeEncoder(v reflect.Value) string {
	s, ok := v.Interface().(FuelType)
	if !ok {
		panic(fmt.Errorf("cannot assert %T as FuelType", v.Interface()))
	}
	return s.String()
}

func DieselFuelType() FuelType {
	for k := range getState().FuelTypeNames {
		if FuelType(k).IsDiesel() {
			return FuelType(k)
		}
	}
	panic("expected diesel to exist")
}
