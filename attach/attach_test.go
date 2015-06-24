package attach

import (
	"github.com/tokuhirom/go-hsperfdata/hsperfdata"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	repo, err := hsperfdata.New()
	if err != nil {
		t.Error(err)
	}

	files, err := repo.GetFiles()
	if files == nil || len(files) == 0 {
		t.Skipf("No java process")
	}

	if err != nil {
		t.Error(err)
	}

	file := files[0]
	pid := file.GetPid()
	sock, err := New(pid)
	if err != nil {
		t.Error(err)
	}
	defer sock.Close()
	err = sock.Execute("threaddump")
	if err != nil {
		t.Error(err)
	}

	threaddump, err := sock.ReadString()
	if err != nil {
		t.Error(err)
	}
	if strings.Index(threaddump, "JNI global references") < 0 {
		t.Errorf("Cannot get threaddump")
	}
}
