package bidder

import "testing"

func TestReadConfig(t *testing.T) {
	f, e := ReadConfig("tmp/config.json")
	if e != nil {
		t.Error(e)
	}
	if len(f.Port) == 0 {
		t.Error("Port param is empty")
	}
}
