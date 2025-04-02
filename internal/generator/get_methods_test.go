package generator

import (
	"testing"
)

func assertMethodSetsEqual(
	t *testing.T,
	entityType string, name string,
	expectedUnique, expectedFull, expectedEmbedded []Method,
	actual MethodSets,
) {
	if !methodsEqual(actual.UniqueSets[name], expectedUnique) {
		t.Errorf("[%s %s] UniqueSet mismatch:\nExpected: %+v\nGot:      %+v", entityType, name, expectedUnique, actual.UniqueSets[name])
	}
	if !methodsEqual(actual.FullSets[name], expectedFull) {
		t.Errorf("[%s %s] FullSet mismatch:\nExpected: %+v\nGot:      %+v", entityType, name, expectedFull, actual.FullSets[name])
	}
	if !methodsEqual(actual.EmbeddedSets[name], expectedEmbedded) {
		t.Errorf("[%s %s] EmbeddedSet mismatch:\nExpected: %+v\nGot:      %+v", entityType, name, expectedEmbedded, actual.EmbeddedSets[name])
	}
}

func methodsEqual(expected, actual []Method) bool {
	if len(expected) != len(actual) {
		return false
	}

	expectedSet := map[string]Method{}
	for _, m := range expected {
		expectedSet[m.Name] = m
	}

	for _, m := range actual {
		if _, ok := expectedSet[m.Name]; !ok {
			return false
		}
	}

	return true
}

func TestGetMethods(t *testing.T) {
	tests := []struct {
		name        string
		interfaces  []InterfaceSpec
		structs     []StructSpec
		expectIface map[string]struct {
			Unique   []Method
			Full     []Method
			Embedded []Method
		}
		expectStruct map[string]struct {
			Unique   []Method
			Full     []Method
			Embedded []Method
		}
	}{
		{
			// test 1: No embedded structs or interfaces
			name: "Test 1: Basic Interface and Struct",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
			},
			structs: []StructSpec{
				{Name: "MyReader", Implements: []string{"Reader"}},
			},

			// Expect all interface methods to be provided by interface itself
			expectIface: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"Reader": {
					Unique:   []Method{{Name: "Read"}},
					Full:     []Method{{Name: "Read"}},
					Embedded: []Method{},
				},
			},

			// Expect all interface methods to be provided by struct itself
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MyReader": {
					Unique:   []Method{{Name: "Read"}},
					Full:     []Method{{Name: "Read"}},
					Embedded: []Method{},
				},
			},
		},
		{
			// test 2: embedded interface but no embedded struct
			name: "Test 2: Interface Embedding With Basic Struct",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
				{Name: "ReadCloser", Methods: []Method{{Name: "Close"}}, Embedded: []string{"Reader"}},
			},
			structs: []StructSpec{
				{Name: "MyReadCloser", Implements: []string{"ReadCloser"}},
			},

			// expect ReadCloser to inherit Read method from embedded Reader, but provide its own close method
			expectIface: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"ReadCloser": {
					Unique:   []Method{{Name: "Close"}},
					Full:     []Method{{Name: "Close"}, {Name: "Read"}},
					Embedded: []Method{{Name: "Read"}},
				},
			},

			// expect Struct to provide all required methods since it does not embed anything
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MyReadCloser": {
					Unique:   []Method{{Name: "Close"}, {Name: "Read"}},
					Full:     []Method{{Name: "Close"}, {Name: "Read"}},
					Embedded: []Method{},
				},
			},
		},
		{
			// test 3: embedded struct which implements the interface already
			name: "Test 3: Basic Interface With Struct Embedding",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
			},
			structs: []StructSpec{
				{Name: "BaseReader", Implements: []string{"Reader"}},
				{Name: "AdvancedReader", Embedded: []string{"BaseReader"}},
			},

			// expect the parent struct to provide no methods since they're provided by embedded
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"AdvancedReader": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Read"}},
					Embedded: []Method{{Name: "Read"}},
				},
			},
		},
		{
			// test 4: multiple embedded interfaces providing all methods
			name: "Test 4: Interface Embeds Multiple Interfaces",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
				{Name: "Writer", Methods: []Method{{Name: "Write"}}},
				{Name: "ReadWriter", Embedded: []string{"Reader", "Writer"}},
			},
			structs: []StructSpec{
				{Name: "MyReadWriter", Implements: []string{"ReadWriter"}},
			},
			// expect parent interface to provide no methods since union of embedded provides both
			expectIface: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"ReadWriter": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Read"}, {Name: "Write"}},
					Embedded: []Method{{Name: "Read"}, {Name: "Write"}},
				},
			},
			// expect struct to provide both methods as there are no embedded structs
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MyReadWriter": {
					Unique:   []Method{{Name: "Read"}, {Name: "Write"}},
					Full:     []Method{{Name: "Read"}, {Name: "Write"}},
					Embedded: []Method{},
				},
			},
		},
		{
			// test 5: multiple structs embedded by parent struct
			name: "Test 5: Struct Embeds Multiple Structs",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
				{Name: "Writer", Methods: []Method{{Name: "Write"}}},
			},
			structs: []StructSpec{
				{Name: "BaseReader", Implements: []string{"Reader"}},
				{Name: "BaseWriter", Implements: []string{"Writer"}},
				{Name: "AdvancedReadWriter", Embedded: []string{"BaseReader", "BaseWriter"}},
			},
			// expect struct to have no unique methods as embedded structs provide all required methods
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"AdvancedReadWriter": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Read"}, {Name: "Write"}},
					Embedded: []Method{{Name: "Read"}, {Name: "Write"}},
				},
			},
		},

		{
			name: "Test 6: Multiple Layers of Interface Embedding",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
				{Name: "Closer", Methods: []Method{{Name: "Close"}}},
				{Name: "ReadCloser", Embedded: []string{"Reader", "Closer"}},
				{Name: "AdvancedIO", Embedded: []string{"ReadCloser"}},
			},

			structs: []StructSpec{
				{Name: "MyAdvancedIO", Implements: []string{"AdvancedIO"}},
			},
			// expect interface to have no unique methods as embedded interfaces provide all required methods
			expectIface: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"AdvancedIO": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Read"}, {Name: "Close"}},
					Embedded: []Method{{Name: "Read"}, {Name: "Close"}},
				},
			},
			// expect struct to have  unique methods as no embedded structs
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MyAdvancedIO": {
					Unique:   []Method{{Name: "Read"}, {Name: "Close"}},
					Full:     []Method{{Name: "Read"}, {Name: "Close"}},
					Embedded: []Method{},
				},
			},
		},

		{
			name: "Test 7: Overlapping Methods via Embedded Interfaces",
			interfaces: []InterfaceSpec{
				{Name: "LoggerA", Methods: []Method{{Name: "Log"}}},
				{Name: "LoggerB", Methods: []Method{{Name: "Log"}}},
				{Name: "UnifiedLogger", Embedded: []string{"LoggerA", "LoggerB"}},
			},
			structs: []StructSpec{
				{Name: "MyLogger", Implements: []string{"UnifiedLogger"}},
			},
			// expect interface to inherit only one copy of "Log"
			expectIface: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"UnifiedLogger": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Log"}},
					Embedded: []Method{{Name: "Log"}},
				},
			},
			// expect struct to provide only one copy of "Log"
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MyLogger": {
					Unique:   []Method{{Name: "Log"}},
					Full:     []Method{{Name: "Log"}},
					Embedded: []Method{},
				},
			},
		},

		{
			name: "Test 8: Multiple Layers of Struct Embedding",
			interfaces: []InterfaceSpec{
				{Name: "Reader", Methods: []Method{{Name: "Read"}}},
			},
			structs: []StructSpec{
				{Name: "BaseReader", Implements: []string{"Reader"}},
				{Name: "MidReader", Embedded: []string{"BaseReader"}},
				{Name: "TopReader", Embedded: []string{"MidReader"}},
			},
			// expect struct to inherit from Mid Reader which inherits from Reader
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"TopReader": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Read"}},
					Embedded: []Method{{Name: "Read"}},
				},
			},
		},

		{
			name: "Test 9: Overlapping Methods from Embedded Structs",
			interfaces: []InterfaceSpec{
				{Name: "Logger", Methods: []Method{{Name: "Log"}}},
			},
			structs: []StructSpec{
				{Name: "LoggerA", Implements: []string{"Logger"}},
				{Name: "LoggerB", Implements: []string{"Logger"}},
				{Name: "MultiLogger", Embedded: []string{"LoggerA", "LoggerB"}},
			},
			// expect deduplication of methods in inheritance
			expectStruct: map[string]struct {
				Unique, Full, Embedded []Method
			}{
				"MultiLogger": {
					Unique:   []Method{},
					Full:     []Method{{Name: "Log"}},
					Embedded: []Method{{Name: "Log"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifaceMethods, structMethods := GetMethods(tt.structs, tt.interfaces)

			for name, expected := range tt.expectIface {
				assertMethodSetsEqual(t, "interface", name, expected.Unique, expected.Full, expected.Embedded, ifaceMethods)
			}
			for name, expected := range tt.expectStruct {
				assertMethodSetsEqual(t, "struct", name, expected.Unique, expected.Full, expected.Embedded, structMethods)
			}
		})
	}
}
