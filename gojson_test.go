package gojson

import (
	"bytes"
	"reflect"
	"testing"
)

// 测试 Marshal 函数
func TestMarshal(t *testing.T) {
	testStruct := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}

	b, err := Marshal(testStruct)
	if err != nil {
		t.Errorf("Marshal error: %v", err)
	}

	expected := `{"Name":"Alice","Age":30}`
	if string(b) != expected {
		t.Errorf("Expected %s, got %s", expected, b)
	}
}

// 测试 Unmarshal 函数
func TestUnmarshal(t *testing.T) {
	jsonData := []byte(`{"Name":"Bob","Age":25}`)
	var testStruct struct {
		Name string
		Age  int
	}

	err := Unmarshal(jsonData, &testStruct)
	if err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	if testStruct.Name != "Bob" || testStruct.Age != 25 {
		t.Errorf("Unmarshaled data is incorrect: %+v", testStruct)
	}
}

// 测试 NewDecoder 函数
func TestNewDecoder(t *testing.T) {
	jsonData := `{"Name":"Charlie","Age":40}`
	decoder := NewDecoder(bytes.NewBufferString(jsonData))
	var testStruct struct {
		Name string
		Age  int
	}

	err := decoder.Decode(&testStruct)
	if err != nil {
		t.Errorf("Decode error: %v", err)
	}

	if testStruct.Name != "Charlie" || testStruct.Age != 40 {
		t.Errorf("Decoded data is incorrect: %+v", testStruct)
	}
}

// 测试 NewEncoder 函数
func TestNewEncoder(t *testing.T) {
	testStruct := struct {
		Name string
		Age  int
	}{
		Name: "Dave",
		Age:  50,
	}

	var buf bytes.Buffer
	encoder := NewEncoder(&buf)
	err := encoder.Encode(testStruct)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}

	expected := `{"Name":"Dave","Age":50}`
	if buf.String() != expected {
		t.Errorf("Expected %s, got %s", expected, buf.String())
	}
}

// 测试 Encode 函数
func TestEncode(t *testing.T) {
	testStruct := struct {
		Name string
		Age  int
	}{
		Name: "Eve",
		Age:  28,
	}

	jsonStr, err := Encode(testStruct)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}

	expected := `{"Name":"Eve","Age":28}`
	if jsonStr != expected {
		t.Errorf("Expected %s, got %s", expected, jsonStr)
	}
}

// 测试 MarshalToString 函数
func TestMarshalToString(t *testing.T) {
	testStruct := struct {
		Name string
		Age  int
	}{
		Name: "Frank",
		Age:  35,
	}

	jsonStr, err := MarshalToString(testStruct)
	if err != nil {
		t.Errorf("MarshalToString error: %v", err)
	}

	expected := `{"Name":"Frank","Age":35}`
	if jsonStr != expected {
		t.Errorf("Expected %s, got %s", expected, jsonStr)
	}
}

// 测试 JsonDecode 函数
func TestJsonDecode(t *testing.T) {
	jsonStr := `{"Name":"Grace","Age":42}`
	expected := map[string]interface{}{
		"Name": "Grace",
		"Age":  float64(42),
	}

	result, err := JsonDecode(jsonStr)
	if err != nil {
		t.Errorf("JsonDecode error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试 JsonDecodeNoError 函数
func TestJsonDecodeNoError(t *testing.T) {
	jsonStr := `{"Name":"Harry","Age":47}`
	expected := map[string]interface{}{
		"Name": "Harry",
		"Age":  float64(47),
	}

	result := JsonDecodeNoError(jsonStr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试 JsonEncode 函数
func TestJsonEncode(t *testing.T) {
	data := map[string]interface{}{
		"Name": "Isaac",
		"Age":  52,
	}

	jsonStr, err := JsonEncode(data)
	if err != nil {
		t.Errorf("JsonEncode error: %v", err)
	}

	expected := `{"Age":52,"Name":"Isaac"}`
	if jsonStr != expected {
		t.Errorf("Expected %s, got %s", expected, jsonStr)
	}
}

// 测试 JsonEncodeNoError 函数
func TestJsonEncodeNoError(t *testing.T) {
	data := map[string]interface{}{
		"Name": "Jane",
		"Age":  57,
	}

	jsonStr := JsonEncodeNoError(data)
	expected := `{"Age":57,"Name":"Jane"}`
	if jsonStr != expected {
		t.Errorf("Expected %s, got %s", expected, jsonStr)
	}
}

// 测试 JsonDecodesNoError 函数
func TestJsonDecodesNoError(t *testing.T) {
	jsonStr := `["Kathy", "Leo"]`
	expected := []string{"Kathy", "Leo"}

	result := JsonDecodesNoError(jsonStr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试 ParseQueryString 函数
func TestParseQueryString(t *testing.T) {
	queryStr := "name=Mike&age=62"
	expected := map[string]interface{}{
		"name": "Mike",
		"age":  "62",
	}

	result := ParseQueryString(queryStr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试 IsValidJSON 函数
func TestIsValidJSON(t *testing.T) {
	validJSON := `{"key": "value"}`
	invalidJSON := `{"key": "value"`

	if !IsValidJSON(validJSON) {
		t.Errorf("IsValidJSON should return true for valid JSON")
	}

	if IsValidJSON(invalidJSON) {
		t.Errorf("IsValidJSON should return false for invalid JSON")
	}
}
