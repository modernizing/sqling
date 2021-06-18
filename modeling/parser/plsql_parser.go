// MIT License
//
//Copyright (c) 2017 Li Xin
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PlSql(source string) {
	codeFiles := make([]string, 0)
	filepath.Walk(source, func(path string, fi os.FileInfo, err error) error {
		codeFiles = append(codeFiles, path)
		return nil
	})

	allP := parseAllPkg(codeFiles)
	fmt.Println(allP.Procedures)

	pTree, pTables := allP.Print("")

	fmt.Println(pTables, pTree)
}


func parseAllPkg(codeFiles []string) *AllProcedure {
	allP := NewAllProcedure()
	pkg := ""
	doFiles(codeFiles, func() {
		pkg = ""
	}, func(line, codeFileName string) {
		line = strings.ToUpper(line)

		if strings.Contains(line, "PKG_") && strings.Contains(line, "CREATE") {
			doCreatePkg(line, func(s string) {
				pkg = s
			})
		}

		if strings.Contains(line, "PROCEDURE") && strings.Contains(line, "P_") {
			doCreateProcedure(line, func(s string) {
				allP.Add(pkg, s)
			})
		}
	})

	procedure := ""
	isComments := false

	doFiles(codeFiles, func() {
		pkg = ""
		procedure = ""
		isComments = false
	}, func(line, codeFileName string) {
		line = strings.ToUpper(line)
		line = strings.Trim(line, " ")
		if strings.HasPrefix(line, "/*") {
			isComments = true
		}

		if strings.HasSuffix(line, "*/") || strings.HasSuffix(line, "*/;") {
			isComments = false
		}

		if isComments {
			return
		}

		if strings.HasPrefix(line, "--") {
			return
		}

		if strings.Contains(line, "PKG_") && strings.Contains(line, "CREATE") {
			doCreatePkg(line, func(s string) {
				pkg = s
			})
		}

		if strings.Contains(line, "PROCEDURE") && strings.Contains(line, "P_") {
			doCreateProcedure(line, func(s string) {
				procedure = s
			})
		}

		if strings.Contains(line, "PKG_") && strings.Contains(line, "(") {
			doPkgLine(line, emptyFilter, func(p string, sp string) {
				allP.AddCall(pkg, procedure, p, sp)
			})
		}

		if strings.Contains(line, "P_") && strings.Contains(line, "(") {
			doCreateProcedure(line, func(s string) {
				allP.AddCall(pkg, procedure, pkg, s)
			})
		}

		if strings.Contains(line, " T_") || strings.Contains(line, ",T_") {
			doSplit(line, tableSplit, func(table string) {
				if strings.HasPrefix(table, "T_") && !IsChineseChar(table) && !strings.Contains(table, ";") && !strings.Contains(table, "、") {
					isWrite := strings.Contains(line, "INSERT ") || strings.Contains(line, "UPDATE ") || strings.Contains(line, "DELETE ")
					allP.AddTable(pkg, procedure, table, isWrite)
				}
			})
		}
	})

	return allP
}


func doFiles(fileNames []string, fileCallback func(), callback func(string, string)) {
	for _, fileName := range fileNames {
		file, _ := os.Open(fileName)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		fileCallback()
		for scanner.Scan() {
			line := scanner.Text()
			callback(line, fileName)
		}

		file.Close()
	}
}

var emptyFilter = func(line string) bool {
	return true
}

var pkgSplit = func(r rune) bool {
	return r == ' ' || r == '(' || r == ',' || r == '\'' || r == '"' || r == ')'
}
var tableSplit = func(r rune) bool {
	return r == ' ' || r == ',' || r == '.' || r == '"' || r == ':' || r == '(' || r == ')' || r == '）' || r == '%' || r == '!' || r == '\''
}

func isComment(first string) bool {
	return strings.HasPrefix(first, "/*") || strings.HasPrefix(first, "*") || strings.HasPrefix(first, "//")
}

func doSplit(line string, f func(rune) bool, callback func(string)) {
	tmp := strings.FieldsFunc(line, f)
	for _, s := range tmp {
		callback(s)
	}
}

func doPkg(s string, pkgFilter func(line string) bool, callback func(string, string)) {
	pkg := ""
	sp := ""
	if strings.HasPrefix(s, "PKG_") && pkgFilter(strings.Split(s, ".")[0]) {
		s = strings.Replace(s, "\"", "", -1)
		tmp := strings.Split(s, ".")
		pkg = tmp[0]
		if len(tmp) > 1 && strings.HasPrefix(tmp[1], "P_") {
			sp = tmp[1]
			callback(pkg, sp)
		}
	}
}

func doPkgLine(line string, pkgFilter func(line string) bool, callback func(string, string)) {
	doSplit(line, pkgSplit, func(s string) {
		doPkg(s, pkgFilter, callback)
	})
}

func doTableLine(line string, tableFilter func(line string) bool, callback func(string)) {
	doSplit(line, tableSplit, func(s string) {
		if strings.HasPrefix(s, "T_") && tableFilter(s) && !IsChineseChar(s) {
			callback(s)
		}
	})
}

func doCreatePkg(line string, callback func(string)) {
	tmp := strings.FieldsFunc(line, pkgSplit)
	for _, key := range tmp {
		if strings.HasPrefix(key, "PKG_") {
			callback(key)
		}
	}
}
func doCreateProcedure(line string, callback func(string)) {
	tmp := strings.FieldsFunc(line, pkgSplit)
	for _, key := range tmp {
		if strings.HasPrefix(key, "P_") {
			callback(key)
		}
	}
}
