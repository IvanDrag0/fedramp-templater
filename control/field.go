package control

import (
	"github.com/IvanDrag0/fedramp-templater/common/source"
)

// field is the very simple representation of information.
type field struct {
	source source.Source
	text   string
}
