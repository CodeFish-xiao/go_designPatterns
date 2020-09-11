package proxy

import "testing"

func TestProxy(t *testing.T) {
	gaocuilan := &Gaocuilan{99}
	sun := &Sunwukong{gaocuilan}
	sun.call()
}
