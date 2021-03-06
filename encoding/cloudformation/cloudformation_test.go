/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cloudformation

import (
	"testing"

	"go.awsctrl.io/manager/testutils"
)

type testCase struct {
	name    string
	keyType string
	object  interface{}
	length  int
	keyName string
	value   string
}

type dynamodb struct {
	Name           string   `cloudformation:"Name,Parameter"`
	CalculatedName string   `cloudformation:"CalcName,Output"`
	ArrayValue     []string `cloudformation:"Array,Parameter"`
	Boolean        bool     `cloudformation:"Boolean,Parameter"`
	Integer        int      `cloudformation:"Integer,Parameter"`
	NestedType     nestedType
}

type nestedType struct {
	Body string `cloudformation:"Body,Parameter"`
}

func (d dynamodb) TemplateBody() string {
	body, _ := testutils.Asset("dynamodb.yaml")
	return string(body)
}

func TestMarshalParameters(t *testing.T) {
	d := &dynamodb{Name: "foo"}

	params := MarshalParameters(d)

	if len(params) != 4 {
		t.Errorf("marshalled params '%d' not equal to '4'", len(params))
	}

	expected := map[string]string{
		"Name":    "foo",
		"Body":    "",
		"Array":   "",
		"Boolean": "false",
		"Integer": "0",
	}

	for k, v := range expected {
		for _, param := range params {
			if *param.ParameterKey == k {
				if *param.ParameterValue != v {
					t.Errorf("parameter value '%s' doesn't match expected 'foo'", *param.ParameterValue)
				}
			}
		}
	}

}

func TestMarshalOutputs(t *testing.T) {
	d := &dynamodb{Name: "foo", CalculatedName: "foobar"}

	outputs := MarshalOutputs(d)

	if len(outputs) != 1 {
		t.Errorf("marshalled outputs '%d' not equal to 1 '%+v'", len(outputs), outputs)
	}

	expected := map[string]string{
		"CalcName": "foobar",
	}

	for k, v := range expected {
		for _, output := range outputs {
			if *output.OutputKey == k {
				if *output.OutputValue != v {
					t.Errorf("parameter value '%s' doesn't match expected 'foo'", *output.OutputValue)
				}
			}
		}
	}
}

func TestMarshalTypes(t *testing.T) {
	cases := []testCase{
		testCase{
			name:    "TestFourItemNameFoo",
			keyType: parameterKeyName,
			object:  &dynamodb{Name: "foo", CalculatedName: "foobar"},
			length:  4,
			keyName: "Name",
			value:   "foo",
		},
		testCase{
			name:    "TestThreeItemArray",
			keyType: parameterKeyName,
			object:  &dynamodb{ArrayValue: []string{"test", "1", "2"}},
			length:  3,
			keyName: "Array",
			value:   "test, 1, 2",
		},
		testCase{
			name:    "TestFiveItemBody",
			keyType: parameterKeyName,
			object:  &dynamodb{Name: "foo", NestedType: nestedType{Body: "blah"}},
			length:  5,
			keyName: "Body",
			value:   "blah",
		},
		testCase{
			name:    "TestThreeItemBoolean",
			keyType: parameterKeyName,
			object:  &dynamodb{Boolean: true},
			length:  3,
			keyName: "Boolean",
			value:   "true",
		},
		testCase{
			name:    "TestThreeItemInteger",
			keyType: parameterKeyName,
			object:  &dynamodb{Integer: 65},
			length:  3,
			keyName: "Integer",
			value:   "65",
		},
		testCase{
			name:    "TestThreeItemInteger",
			keyType: outputKeyName,
			object:  &dynamodb{Name: "foo", CalculatedName: "foobar"},
			length:  1,
			keyName: "CalcName",
			value:   "foobar",
		},
		testCase{
			name:    "TestFourEmptyName",
			keyType: parameterKeyName,
			object:  &dynamodb{Name: "", Boolean: true},
			length:  3,
			keyName: "Boolean",
			value:   "true",
		},
	}

	for _, tCase := range cases {
		types := make(map[string]string)

		MarshalTypes(types, tCase.object, tCase.keyType)
		if len(types) != tCase.length {
			t.Errorf("%s marshalled response length '%d' not equal to '%d', '%+v'", tCase.name, len(types), tCase.length, types)
		}

		var key string
		var ok bool
		if key, ok = (types)[tCase.keyName]; !ok {
			t.Errorf("%s key '%s' doesn't exist, '%+v'", tCase.name, tCase.keyName, types)
		}

		if key != tCase.value {
			t.Errorf("%s '%+v' value doesn't equal '%s'", tCase.name, key, tCase.value)
		}
	}
}
