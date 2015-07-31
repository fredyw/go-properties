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
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile(`^([^=]+)=(.*)$`)
)

// Properties contains the list of properties
type Properties map[string]string

// Read reads properties from a reader.
func Read(in io.Reader) (Properties, error) {
	scanner := bufio.NewScanner(in)
	props := Properties{}
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// Skip comment
		if len(line) == 0 || line[0] == '#' {
			lineNumber++
			continue
		}
		if groups := regex.FindStringSubmatch(line); groups != nil {
			key, value := groups[1], groups[2]
			props[strings.TrimSpace(key)] = strings.TrimSpace(value)
		} else {
			return nil, errors.New(fmt.Sprintf("Invalid syntax at line %d", lineNumber))
		}
		lineNumber++
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return props, nil
}

// ReadFile reads properties from a file.
func ReadFile(fileName string) (Properties, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Read(file)
}

// Write writes properties into a writer
func Write(out io.Writer, props Properties) {
	for key, value := range props {
		fmt.Fprintln(out, key, "=", value)
	}
	fmt.Fprintln(out)
}

// WriteFile writes properties into a file
func WriteFile(fileName string, props Properties) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	Write(file, props)
	return nil
}
