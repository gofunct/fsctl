package fscmd

import (
	"github.com/gofunct/fsctl"
	"github.com/spf13/cobra"
)

type FsCmd struct {
	*fsctl.Fs
	*cobra.Command
}

func NewFsCmd(name, usg string, cfgPath string) *FsCmd {
	return &FsCmd{
		Fs: fsctl.NewFs(cfgPath),
		Command: &cobra.Command{
			Use:   name,
			Short: usg,
		},
	}
}
