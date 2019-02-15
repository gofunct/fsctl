package fsctl

import (
	"testing"
)

func TestNewFs(t *testing.T) {
	fs := NewFs("https://github.com/gofunct/fsctl/blob/master/testing//config.yaml")
	if fs.GetString("name")  != "Coleman Word" {
		if err := fs.ListTmpFiles(); err != nil {
			t.Fatal(err, "failed to walk temp dir")
		}
		t.Fatal("failed to get name from remote config", fs.Get("name"))
	}
}
