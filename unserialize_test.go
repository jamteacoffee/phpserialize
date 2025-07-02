package phpserialize_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elliotchance/orderedmap/v3"
	"github.com/jamteacoffee/phpserialize"
)

func expectErrorToNotHaveOccurred(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func expectErrorToEqual(t *testing.T, err1, err2 error) {
	if err1 == nil {
		t.Error("err1 is nil")
	}

	if err2 == nil {
		t.Error("err2 is nil")
	}

	if err1.Error() != err2.Error() {
		t.Errorf("Expected '%s' to be '%s'", err1, err2)
	}
}

func TestUnmarshalInt(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        int
		expectedError error
	}{
		"0":              {[]byte("i:0;"), 0, nil},
		"5":              {[]byte("i:5;"), 5, nil},
		"-8":             {[]byte("i:-8;"), -8, nil},
		"1000000":        {[]byte("i:1000000;"), 1000000, nil},
		"not an integer": {[]byte("N;"), 0, errors.New("not an integer")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Run("int", func(t *testing.T) {
				var result int
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != test.output {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int8", func(t *testing.T) {
				var result int8
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int8(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int16", func(t *testing.T) {
				var result int16
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int16(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int32", func(t *testing.T) {
				var result int32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("int64", func(t *testing.T) {
				var result int64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != int64(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint", func(t *testing.T) {
				var result uint
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint8", func(t *testing.T) {
				var result uint8
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint8(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint16", func(t *testing.T) {
				var result uint16
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint16(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint32", func(t *testing.T) {
				var result uint32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("uint64", func(t *testing.T) {
				var result uint64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != uint64(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})
		})
	}
}

func TestUnmarshalFloat(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        float64
		expectedError error
	}{
		"3.2":         {[]byte("d:3.2;"), 3.2, nil},
		"10.0":        {[]byte("d:10;"), 10.0, nil},
		"123.456789":  {[]byte("d:123.456789;"), 123.456789, nil},
		"1.23e9":      {[]byte("d:1230000000;"), 1.23e9, nil},
		"-17.23":      {[]byte("d:3.2;"), 3.2, nil},
		"not a float": {[]byte("N;"), 0.0, errors.New("not a float")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Run("float32", func(t *testing.T) {
				var result float32
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != float32(test.output) {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})

			t.Run("float64", func(t *testing.T) {
				var result float64
				err := phpserialize.Unmarshal(test.input, &result)

				if test.expectedError == nil {
					expectErrorToNotHaveOccurred(t, err)
					if result != test.output {
						t.Errorf("Expected '%v', got '%v'", result, test.output)
					}
				} else {
					expectErrorToEqual(t, err, test.expectedError)
				}
			})
		})
	}
}

func TestUnmarshalString(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        string
		expectedError error
	}{
		"''":            {[]byte("s:0:\"\";"), "", nil},
		"'Hello world'": {[]byte("s:11:\"Hello world\";"), "Hello world", nil},
		"'Björk Guðmundsdóttir'": {
			[]byte(`s:23:"Björk Guðmundsdóttir";`),
			"Björk Guðmundsdóttir",
			nil,
		},
		"not a string": {[]byte("N;"), "", errors.New("not a string")},
		"Backslash":    {[]byte("s:1:\"\\\";"), "\\", nil},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result string
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if result != test.output {
					t.Errorf("Expected '%v', got '%v'", test.output, result)
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalBinary(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        []byte
		expectedError error
	}{
		"[]byte: \\x01\\x02\\x03": {
			[]byte("s:3:\"\x01\x02\x03\";"),
			[]byte{1, 2, 3},
			nil,
		},
		"not a string": {[]byte("N;"), []byte{}, errors.New("not a string")},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result []byte
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if string(result) != string(test.output) {
					t.Errorf("Expected '%v', got '%v'", result, test.output)
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalArray(t *testing.T) {
	tests := map[string]struct {
		input         []byte
		output        []interface{}
		expectedError error
	}{
		"[]interface{}: [7.89]": {
			[]byte("a:1:{i:0;d:7.89;}"),
			[]interface{}{7.89},
			nil,
		},
		"[]interface{}: [7, 8, 9]": {
			[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"),
			[]interface{}{int64(7), int64(8), int64(9)},
			nil,
		},
		"[]interface{}: [7.2, 'foo']": {
			[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"),
			[]interface{}{7.2, "foo"},
			nil,
		},
		"[]interface{}: [null]": {
			[]byte("a:1:{i:0;N;}"),
			[]interface{}{nil},
			nil,
		},
		"[]interface{}: [true, false]": {
			[]byte("a:2:{i:0;b:1;i:1;b:0;}"),
			[]interface{}{true, false},
			nil,
		},
		"[]interface{}: [1, 2, 'foo']": {
			[]byte(`a:3:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";}`),
			[]interface{}{int64(1), int64(2), "foo"},
			nil,
		},
		"[]interface{}: [1, 2, 'foo', '中文']": {
			[]byte(`a:4:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";}`),
			[]interface{}{int64(1), int64(2), "foo", "中文"},
			nil,
		},
		// "[]interface{}: [1, 2, 'foo', '中文', ['a' => 'a']]": {
		// 	[]byte(`a:5:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";i:4;a:1:{s:1:"a";s:1:"a";}}`),
		// 	[]interface{}{int64(1), int64(2), "foo", "中文", map[interface{}]interface{}{"a": "a"}},
		// 	nil,
		// },
		// "[]interface{}: [1, 2, 'foo', ['a' => 'a']]": {
		// 	[]byte(`a:4:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;a:1:{s:1:"a";s:1:"a";}}`),
		// 	[]interface{}{int64(1), int64(2), "foo", map[interface{}]interface{}{"a": "a"}},
		// 	nil,
		// },
		// "[]interface{}: [1, 2, 'foo', ['a' => 'a'], ['a' => 'a']]": {
		// 	[]byte(`a:5:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;a:1:{s:1:"a";s:1:"a";}i:4;a:1:{s:1:"a";s:1:"a";}}`),
		// 	[]interface{}{int64(1), int64(2), "foo", map[interface{}]interface{}{"a": "a"}, map[interface{}]interface{}{"a": "a"}},
		// 	nil,
		// },
		// "[]interface{}: [1, 2, 'foo', '中文', ['a' => 'a'], ['a' => 'a']]": {
		// 	[]byte(`a:6:{i:0;i:1;i:1;i:2;i:2;s:3:"foo";i:3;s:6:"中文";i:4;a:1:{s:1:"a";s:1:"a";}i:5;a:1:{s:1:"a";s:1:"a";}}`),
		// 	[]interface{}{int64(1), int64(2), "foo", "中文", map[interface{}]interface{}{"a": "a"}, map[interface{}]interface{}{"a": "a"}},
		// 	nil,
		// },
		// "[]interface{}: [['id'=> '1'], ['id'=> '2']]": {
		// 	[]byte(`a:2:{i:0;a:1:{s:2:"id";s:1:"1";}i:1;a:1:{s:2:"id";s:1:"2";}}`),
		// 	[]interface{}{map[interface{}]interface{}{"id": "1"}, map[interface{}]interface{}{"id": "2"}},
		// 	nil,
		// },
		// "[]interface{}: [['id'=> '1', 'name' => '1'], ['id'=> '2', 'name' => '2'], ['id'=> '3', 'name' => '3']]": {
		// 	[]byte(`a:3:{i:0;a:2:{s:2:"id";s:1:"1";s:4:"name";s:1:"1";}i:1;a:2:{s:2:"id";s:1:"2";s:4:"name";s:1:"2";}i:2;a:2:{s:2:"id";s:1:"3";s:4:"name";s:1:"3";}}`),
		// 	[]interface{}{map[interface{}]interface{}{"id": "1", "name": "1"}, map[interface{}]interface{}{"id": "2", "name": "2"}, map[interface{}]interface{}{"id": "3", "name": "3"}},
		// 	nil,
		// },
		"cannot decode map as slice": {
			[]byte("a:2:{i:0;b:1;i:5;b:0;}"),
			[]interface{}{},
			errors.New("cannot decode map as slice"),
		},
		"not an array": {
			[]byte("N;"),
			[]interface{}{},
			errors.New("not an array"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var result []interface{}
			err := phpserialize.Unmarshal(test.input, &result)

			if test.expectedError == nil {
				expectErrorToNotHaveOccurred(t, err)
				if len(result) != len(test.output) {
					t.Errorf("Expected %v, got %v", len(result), len(test.output))
				}

				for k, _ := range result {
					if !reflect.DeepEqual(result[k], test.output[k]) {
						t.Errorf("Expected %v (%s), got %v (%s) for #%d",
							result[k], reflect.TypeOf(result[k]).Name(),
							test.output[k], reflect.TypeOf(test.output[k]).Name(), k)
					}
				}
			} else {
				expectErrorToEqual(t, err, test.expectedError)
			}
		})
	}
}

func TestUnmarshalAssociativeArrayStringKeys(t *testing.T) {
	input := []byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}")
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 2 elements
	if result.Len() != 2 {
		t.Errorf("Expected map length 2, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("foo"); !ok || val != int64(10) {
		t.Errorf("Expected foo: 10, got %v", val)
	}

	if val, ok := result.Get("bar"); !ok || val != int64(20) {
		t.Errorf("Expected bar: 20, got %v", val)
	}
}

func TestUnmarshalAssociativeArrayIntegerKeys(t *testing.T) {
	input := []byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}")
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 2 elements
	if result.Len() != 2 {
		t.Errorf("Expected map length 2, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get(int64(1)); !ok || val != int64(10) {
		t.Errorf("Expected 1: 10, got %v", val)
	}

	if val, ok := result.Get(int64(2)); !ok || val != "foo" {
		t.Errorf("Expected 2: foo, got %v", val)
	}
}

func TestUnmarshalAssociativeArrayNested(t *testing.T) {
	input := []byte(`a:3:{s:3:"foo";i:10;s:3:"bar";i:20;s:6:"foobar";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}}`)
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 3 elements
	if result.Len() != 3 {
		t.Errorf("Expected map length 3, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("foo"); !ok || val != int64(10) {
		t.Errorf("Expected foo: 10, got %v", val)
	}

	if val, ok := result.Get("bar"); !ok || val != int64(20) {
		t.Errorf("Expected bar: 20, got %v", val)
	}

	// Verify nested map
	if val, ok := result.Get("foobar"); !ok {
		t.Errorf("Expected foobar key to exist")
	} else {
		nestedMap, ok := val.(*orderedmap.OrderedMap[any, any])
		if !ok {
			t.Errorf("Expected foobar value to be OrderedMap, got %T", val)
		} else {
			if nestedMap.Len() != 2 {
				t.Errorf("Expected nested map length 2, got %d", nestedMap.Len())
			}

			if nestedVal, ok := nestedMap.Get("foo"); !ok || nestedVal != int64(10) {
				t.Errorf("Expected nested foo: 10, got %v", nestedVal)
			}

			if nestedVal, ok := nestedMap.Get("bar"); !ok || nestedVal != int64(20) {
				t.Errorf("Expected nested bar: 20, got %v", nestedVal)
			}
		}
	}
}

func TestUnmarshalAssociativeArrayMultipleNested(t *testing.T) {
	input := []byte(`a:4:{s:3:"foo";i:10;s:3:"bar";i:20;s:6:"foobar";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}s:7:"foobar1";a:2:{s:3:"foo";i:10;s:3:"bar";i:20;}}`)
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 4 elements
	if result.Len() != 4 {
		t.Errorf("Expected map length 4, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("foo"); !ok || val != int64(10) {
		t.Errorf("Expected foo: 10, got %v", val)
	}

	if val, ok := result.Get("bar"); !ok || val != int64(20) {
		t.Errorf("Expected bar: 20, got %v", val)
	}

	// Verify first nested map
	if val, ok := result.Get("foobar"); !ok {
		t.Errorf("Expected foobar key to exist")
	} else {
		nestedMap, ok := val.(*orderedmap.OrderedMap[any, any])
		if !ok {
			t.Errorf("Expected foobar value to be OrderedMap, got %T", val)
		} else {
			if nestedMap.Len() != 2 {
				t.Errorf("Expected nested map length 2, got %d", nestedMap.Len())
			}

			if nestedVal, ok := nestedMap.Get("foo"); !ok || nestedVal != int64(10) {
				t.Errorf("Expected nested foo: 10, got %v", nestedVal)
			}

			if nestedVal, ok := nestedMap.Get("bar"); !ok || nestedVal != int64(20) {
				t.Errorf("Expected nested bar: 20, got %v", nestedVal)
			}
		}
	}

	// Verify second nested map
	if val, ok := result.Get("foobar1"); !ok {
		t.Errorf("Expected foobar1 key to exist")
	} else {
		nestedMap, ok := val.(*orderedmap.OrderedMap[any, any])
		if !ok {
			t.Errorf("Expected foobar1 value to be OrderedMap, got %T", val)
		} else {
			if nestedMap.Len() != 2 {
				t.Errorf("Expected nested map length 2, got %d", nestedMap.Len())
			}

			if nestedVal, ok := nestedMap.Get("foo"); !ok || nestedVal != int64(10) {
				t.Errorf("Expected nested foo: 10, got %v", nestedVal)
			}

			if nestedVal, ok := nestedMap.Get("bar"); !ok || nestedVal != int64(20) {
				t.Errorf("Expected nested bar: 20, got %v", nestedVal)
			}
		}
	}
}

func TestUnmarshalAssociativeArrayNotAnArray(t *testing.T) {
	input := []byte("N;")
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	expectedError := errors.New("not an array")
	expectErrorToEqual(t, err, expectedError)
}

var inputNull = []byte("N;")
var inputBoolFalse = []byte("b:0;")
var inputBoolTrue = []byte("b:1;")

func TestUnmarshalWithNull(t *testing.T) {
	result := interface{}(nil)
	err := phpserialize.Unmarshal(inputNull, &result)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestUnmarshalNilWithNull(t *testing.T) {
	err := phpserialize.UnmarshalNil(inputNull)
	if err != nil {
		t.Error(err)
	}
}

func TestBadUnmarshalNilWithNull(t *testing.T) {
	err := phpserialize.UnmarshalNil(inputBoolFalse)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestUnmarshalWithBooleanTrue(t *testing.T) {
	var result bool
	err := phpserialize.Unmarshal(inputBoolTrue, &result)

	if result != true {
		t.Errorf("expected true")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshalObject(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	var result struct1
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Foo != 10 {
		t.Errorf("Expected %v, got %v", 10, result.Foo)
	}

	if result.Bar.Qux != 1.23 {
		t.Errorf("Expected %v, got %v", 1.23, result.Bar.Qux)
	}

	if result.Baz != "yay" {
		t.Errorf("Expected %v, got %v", "yay", result.Baz)
	}
}

func TestUnmarshalObjectWithTags(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	var result structTag
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Bar != 10 {
		t.Errorf("Expected %v, got %v", 10, result.Bar)
	}

	if result.Foo.Qux != 1.23 {
		t.Errorf("Expected %v, got %v", 1.23, result.Foo.Qux)
	}

	if result.Balu != "yay" {
		t.Errorf("Expected %v, got %v", "yay", result.Balu)
	}

	if result.Ignored != "" {
		t.Errorf("Expected %v, got %v", "yay", result.Ignored)
	}
}

func TestUnmarshalPointers(t *testing.T) {
	data := "O:8:\"Nillable\":4:{s:3:\"foo\";s:3:\"yay\";s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:10;}s:6:\"fooPtr\";s:3:\"hey\";s:6:\"barPtr\";O:7:\"Struct2\":1:{s:3:\"qux\";d:0;}}"
	target := &Nillable{
		Foo: "yay",
		Bar: Struct2{
			Qux: 10,
		},
		FooPtr: &heyStr,
		BarPtr: &Struct2{
			Qux: 0,
		},
	}

	var result Nillable
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Foo != target.Foo {
		t.Errorf("Expected %v, got %v for Foo", target.Foo, result.Foo)
	}

	if result.Bar.Qux != target.Bar.Qux {
		t.Errorf("Expected %v, got %v for Bar", target.Bar, result.Bar)
	}

	if result.FooPtr == nil || *result.FooPtr != *target.FooPtr {
		t.Errorf("Expected %v, got %v for FooPtr", *target.FooPtr, *result.FooPtr)
	}

	if result.BarPtr == nil || result.BarPtr.Qux != result.BarPtr.Qux {
		t.Errorf("Expected %v, got %v for BarPtr", target.BarPtr, result.BarPtr)
	}

}

func TestUnmarshalPointersWithNull(t *testing.T) {
	data := "O:8:\"Nillable\":4:{s:3:\"foo\";s:0:\"\";s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:0;}s:6:\"fooPtr\";N;s:6:\"barPtr\";N;}"

	target := &Nillable{
		Foo: "",
		Bar: Struct2{
			Qux: 0,
		},
		FooPtr: nil,
		BarPtr: nil,
	}

	var result Nillable
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if result.Foo != target.Foo {
		t.Errorf("Expected %v, got %v for Foo", target.Foo, result.Foo)
	}

	if result.Bar.Qux != target.Bar.Qux {
		t.Errorf("Expected %v, got %v for Bar", target.Bar, result.Bar)
	}

	if result.FooPtr != target.FooPtr {
		t.Errorf("Expected %v, got %v for FooPtr", target.FooPtr, result.FooPtr)
	}

	if result.BarPtr != target.BarPtr {
		t.Errorf("Expected %v, got %v for BarPtr", target.BarPtr, result.BarPtr)
	}
}

func TestUnmarshalObjectIntoMap(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 3 elements
	if result.Len() != 3 {
		t.Errorf("Expected map length 3, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("baz"); !ok || val != "yay" {
		t.Errorf("Expected baz: yay, got %v", val)
	}

	if val, ok := result.Get("foo"); !ok || val != int64(10) {
		t.Errorf("Expected foo: 10, got %v", val)
	}

	// Verify nested object/map
	if val, ok := result.Get("bar"); !ok {
		t.Errorf("Expected bar key to exist")
	} else {
		nestedMap, ok := val.(*orderedmap.OrderedMap[any, any])
		if !ok {
			t.Errorf("Expected bar value to be OrderedMap, got %T", val)
		} else {
			if nestedMap.Len() != 1 {
				t.Errorf("Expected nested map length 1, got %d", nestedMap.Len())
			}

			if nestedVal, ok := nestedMap.Get("qux"); !ok || nestedVal != 1.23 {
				t.Errorf("Expected nested qux: 1.23, got %v", nestedVal)
			}
		}
	}
}

func TestUnmarshalObjectIntoMapContainingArray(t *testing.T) {
	data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}s:3:\"baz\";s:3:\"yay\";}"
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 3 elements
	if result.Len() != 3 {
		t.Errorf("Expected map length 3, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("baz"); !ok || val != "yay" {
		t.Errorf("Expected baz: yay, got %v", val)
	}

	if val, ok := result.Get("foo"); !ok || val != int64(10) {
		t.Errorf("Expected foo: 10, got %v", val)
	}

	// Verify nested array
	if val, ok := result.Get("bar"); !ok {
		t.Errorf("Expected bar key to exist")
	} else {
		arr, ok := val.([]interface{})
		if !ok {
			t.Errorf("Expected bar value to be []interface{}, got %T", val)
		} else {
			expectedArray := []interface{}{int64(7), int64(8), int64(9)}
			if len(arr) != len(expectedArray) {
				t.Errorf("Expected array length %d, got %d", len(expectedArray), len(arr))
			}

			for i, expectedVal := range expectedArray {
				if i < len(arr) && arr[i] != expectedVal {
					t.Errorf("Expected array[%d]: %v, got %v", i, expectedVal, arr[i])
				}
			}
		}
	}
}

// https://github.com/elliotchance/phpserialize/issues/7
func TestUnmarshalArrayThatContainsInteger(t *testing.T) {
	data := `a:3:{s:4:"name";s:2:"tw";s:3:"age";i:123;s:4:"wife";a:1:{s:1:"x";s:1:"y";}}`
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	// Verify the map has 3 elements
	if result.Len() != 3 {
		t.Errorf("Expected map length 3, got %d", result.Len())
	}

	// Verify specific key-value pairs
	if val, ok := result.Get("name"); !ok || val != "tw" {
		t.Errorf("Expected name: tw, got %v", val)
	}

	if val, ok := result.Get("age"); !ok || val != int64(123) {
		t.Errorf("Expected age: 123, got %v", val)
	}

	// Verify nested map
	if val, ok := result.Get("wife"); !ok {
		t.Errorf("Expected wife key to exist")
	} else {
		nestedMap, ok := val.(*orderedmap.OrderedMap[any, any])
		if !ok {
			t.Errorf("Expected wife value to be OrderedMap, got %T", val)
		} else {
			if nestedMap.Len() != 1 {
				t.Errorf("Expected nested map length 1, got %d", nestedMap.Len())
			}

			if nestedVal, ok := nestedMap.Get("x"); !ok || nestedVal != "y" {
				t.Errorf("Expected nested x: y, got %v", nestedVal)
			}
		}
	}
}

func TestUnmarshalObjectThatContainsArray(t *testing.T) {
	data := "O:7:\"Struct3\":4:{s:11:\"objectArray\";a:2:{i:0;O:7:\"Struct2\":1:{s:3:\"qux\";d:1.1;}i:1;O:7:\"Struct2\":1:{s:3:\"qux\";d:2.2;}}s:8:\"intArray\";a:2:{i:0;i:1;i:1;i:2;}s:10:\"floatArray\";a:2:{i:0;d:1;i:1;d:2;}s:11:\"stringArray\";a:2:{i:0;s:1:\"a\";i:1;s:1:\"b\";}}"
	var result Struct3
	err := phpserialize.Unmarshal([]byte(data), &result)
	expectErrorToNotHaveOccurred(t, err)

	if len(result.ObjectArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.ObjectArray))
	}
	if len(result.IntArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.IntArray))
	}
	if len(result.FloatArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.FloatArray))
	}
	if len(result.StringArray) == 0 {
		t.Errorf("Expected %v, got %v", 2, len(result.StringArray))
	}
}

var escapeTests = map[string]struct {
	Unserialized, Serialized string
}{
	"SingleQuote": {
		"foo'bar", `s:8:"foo\'bar";`,
	},
	"DoubleQuote": {
		"foo\"bar", `s:7:"foo"bar";`,
	},
	"Backslash": {
		"foo\\bar", `s:7:"foo\bar";`,
	},
	"Dollar": {
		"foo$bar", `s:7:"foo$bar";`,
	},
	"NewLine": {
		"foo\nbar", "s:7:\"foo\nbar\";",
	},
	"HorizontalTab": {
		"foo\tbar", "s:7:\"foo\tbar\";",
	},
	"CarriageReturn": {
		"foo\rbar", "s:7:\"foo\rbar\";",
	},
}

func TestUnmarshalEscape(t *testing.T) {
	for testName, test := range escapeTests {
		t.Run(testName, func(t *testing.T) {
			var result string
			err := phpserialize.Unmarshal([]byte(test.Serialized), &result)
			expectErrorToNotHaveOccurred(t, err)

			if test.Unserialized != result {
				t.Errorf("Expected:\n  %#+v\nGot:\n  %#+v", test.Unserialized, result)
			}
		})
	}
}
