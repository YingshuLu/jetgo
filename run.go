package jetgo

import "github.com/yingshulu/jetgo/vm"

// Run source code with data input
func Run(source string, data interface{}) (interface{}, error) {
	var fact interface{}
	if data != nil {
		fact = vm.Fact(data)
	}
	r, err := vm.Compile(source)
	if err != nil {
		return nil, err
	}
	v := vm.New()
	return v.Run(r, fact)
}
