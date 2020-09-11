package facade

import "testing"

func TestFacade(t *testing.T) {
	startBtn := &StartBtn{} //开机
	startBtn.start()
}
