package comfymill

import "github.com/davidroman0O/comfylite3"

func NewDatabase(opts ...comfylite3.ComfyOption) (*comfylite3.ComfyDB, error) {
	return comfylite3.Comfy(opts...)
}
