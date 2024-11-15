package util

var (
	TabContent     = []byte(`	`)
	NewlineContent = []byte("\n")
)

func GetTabBytes(tabs int) []byte {
	tabBytes := []byte{}
	i := 0
	for i < tabs {
		tabBytes = append(tabBytes, TabContent...)
		i++
	}
	return tabBytes
}
