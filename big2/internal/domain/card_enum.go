// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package domain

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// CompareResultSmaller is a CompareResult of type Smaller.
	CompareResultSmaller CompareResult = iota
	// CompareResultBigger is a CompareResult of type Bigger.
	CompareResultBigger
	// CompareResultEqual is a CompareResult of type Equal.
	CompareResultEqual
)

var ErrInvalidCompareResult = errors.New("not a valid CompareResult")

const _CompareResultName = "SmallerBiggerEqual"

var _CompareResultMap = map[CompareResult]string{
	CompareResultSmaller: _CompareResultName[0:7],
	CompareResultBigger:  _CompareResultName[7:13],
	CompareResultEqual:   _CompareResultName[13:18],
}

// String implements the Stringer interface.
func (x CompareResult) String() string {
	if str, ok := _CompareResultMap[x]; ok {
		return str
	}
	return fmt.Sprintf("CompareResult(%d)", x)
}

var _CompareResultValue = map[string]CompareResult{
	_CompareResultName[0:7]:                    CompareResultSmaller,
	strings.ToLower(_CompareResultName[0:7]):   CompareResultSmaller,
	_CompareResultName[7:13]:                   CompareResultBigger,
	strings.ToLower(_CompareResultName[7:13]):  CompareResultBigger,
	_CompareResultName[13:18]:                  CompareResultEqual,
	strings.ToLower(_CompareResultName[13:18]): CompareResultEqual,
}

// ParseCompareResult attempts to convert a string to a CompareResult.
func ParseCompareResult(name string) (CompareResult, error) {
	if x, ok := _CompareResultValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _CompareResultValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return CompareResult(0), fmt.Errorf("%s is %w", name, ErrInvalidCompareResult)
}
