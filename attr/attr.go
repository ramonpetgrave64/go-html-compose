package attr

import (
	"fmt"
	"html"
	"io"
)

type AttributeStruct struct {
	Name  string
	Value string
}

func (attr *AttributeStruct) Render(wr io.Writer) {
	wr.Write([]byte(fmt.Sprintf(`%s="%s"`, attr.Name, html.EscapeString(attr.Value))))
}

func Attr(name string, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}
