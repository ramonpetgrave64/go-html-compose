package doc

import (
	"io"
)

type IContent interface {
	RenderConent(wr io.Writer) (err error)
}

// ContContainerStruct contains content without itself being an element.
// See https://html.spec.whatwg.org/multipage/dom.html#flow-content.
type ContContainerStruct struct {
	Children []IContent
}

// ContContainer creates a ContContainer.
func ContContainer(children ...IContent) *ContContainerStruct {
	return &ContContainerStruct{
		Children: children,
	}
}

// RenderContent renders the items in the FlowStruct.
func (c *ContContainerStruct) RenderConent(wr io.Writer) error {
	for _, elem := range c.Children {
		if err := elem.RenderConent(wr); err != nil {
			return err
		}
	}
	return nil
}
