package check

import (
	"fitan/gen/internal/parser"
	"fmt"
	"strconv"
	"strings"
)

//GetTestParamInTmpl return param list
func (m *InterfaceMethod) GetTestParamInTmpl() string {
	return testParamToString(m.Params)
}

// GetTestResultParamInTmpl return result list
func (m *InterfaceMethod) GetTestResultParamInTmpl() string {
	var res []string
	for i := range m.Result {
		tmplString := fmt.Sprintf("res%d", i+1)
		res = append(res, tmplString)
	}
	return strings.Join(res, ",")
}

// testParamToString param list to string used in tmpl
func testParamToString(params []parser.Param) string {
	var res []string
	for i, param := range params {
		tmplString := fmt.Sprintf("tt.Input.Args[%d].(%s)", i, param.Type)
		res = append(res, tmplString)
	}
	return strings.Join(res, ",")
}

// GetAssertInTmpl assert in diy test
func (m *InterfaceMethod) GetAssertInTmpl() string {
	var res []string
	for i := range m.Result {
		tmplString := fmt.Sprintf("assert(t, %v, res%d, tt.Expectation.Ret[%d])", strconv.Quote(m.MethodName), i+1, i)
		res = append(res, tmplString)
	}
	return strings.Join(res, "\n")
}
