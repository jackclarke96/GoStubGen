package cmd

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestIface_GetUniqueMethods(t *testing.T) {
// 	itm := interfaceToMethods{
// 		"Vehicle": {
// 			{Name: "GetTopSpeed"},
// 			{Name: "Turn"},
// 			{Name: "Accelerate"},
// 		},
// 		"SelfDriving": {
// 			{Name: "DriveSelf"},
// 		},
// 	}
// 	iFace := iFace{Name: "SelfDriving", Embedded: []string{"Vehicle"}}
// 	unique := iFace.getUniqueMethods(itm)
// 	fmt.Println("unique:", unique)
// 	assert.Equal(t, 1, len(unique))
// 	assert.Equal(t, "DriveSelf", unique[0].Name)
// }

// func TestStruct_GetUniqueMethods(t *testing.T) {
// 	stm := structToMethods{
// 		"Car": {
// 			{Name: "GetTopSpeed"},
// 			{Name: "Turn"},
// 		},
// 		"FutureCar": {
// 			{Name: "DriveSelf"},
// 			{Name: "Teleport"},
// 			Embedded: []string{"Car"},
// 		},
// 	}
// 	s := ss{Name: "FutureCar", Embedded: []string{"Car"}}
// 	unique := s.getUniqueMethods(stm)
// 	fmt.Println("unique:", unique)
// 	assert.Equal(t, 2, len(unique))
// 	assert.Equal(t, "DriveSelf", unique[0].Name)
// 	assert.Equal(t, "Teleport", unique[1].Name)
// }
