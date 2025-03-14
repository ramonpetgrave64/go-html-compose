package doc

import (
	"io"
)

// contContainerStruct contains content without itself being an element.
// See https://html.spec.whatwg.org/multipage/dom.html#flow-content.
type contContainerStruct struct {
	Children []IContent
}

// ContContainer creates a ContContainer.
func ContContainer(children ...IContent) IContent {
	return &contContainerStruct{
		Children: children,
	}
}

// RenderContent renders the items in the FlowStruct.
func (c *contContainerStruct) RenderConent(wr io.Writer) error {
	for _, elem := range c.Children {
		if err := elem.RenderConent(wr); err != nil {
			return err
		}
	}
	return nil
}
