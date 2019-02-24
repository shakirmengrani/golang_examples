package InnerPkg

import "testing"

func TestFunc(t *testing.T) {
	di := NewDI(&Config{Name: "hello", Port: 3000})
	if di.Config.Port != 3000 {
		t.Errorf("Expecting port is %d but got %d", 8000, di.Config.Port)
	}

}
