// The MIT License (MIT)
//
// Copyright (c) 2015 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package goprops

import (
	"fmt"
	"testing"
)

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatal(fmt.Sprintln("Expected:", expected, ", got:", actual))
	}
}

func TestRead(t *testing.T) {
	props, err := ReadFile("test.properties")
	if err != nil {
		t.Fatal(err)
	}
	for key, value := range props {
		fmt.Println(key, value)
	}
	assertEquals(t, 4, len(props))
	assertEquals(t, "value1", props["key1"])
	assertEquals(t, "value2", props["key2"])
	assertEquals(t, "value3", props["key3"])
	assertEquals(t, "value4", props["key4"])
}

func TestReadInvalidSyntax(t *testing.T) {
	props, err := ReadFile("invalid.properties")
	if err == nil {
		t.Fatal("ReadFile should fail")
	}
	if props != nil {
		t.Fatal("props should be nil")
	}
}

func TestWrite(t *testing.T) {
	props := Properties{}
	props["key1"] = "value1"
	props["key2"] = "value2"
	props["key3"] = "value3"

	err := WriteFile("testwrite.properties", props)
	if err != nil {
		t.Fatal(err)
	}
}
