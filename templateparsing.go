package templateparsing

import (
	"bytes"
	"fmt"
	"text/template"
)

// ParseTemplate ...
func ParseTemplate(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return ""
	}
	temp := buf.String()

	return temp
}
