package fscmd

import (
	"github.com/gofunct/fsctl"
	"github.com/spf13/cobra"
)

type FsCmd struct {
	*fsctl.Fs
	*cobra.Command
}

func NewFsCmd(name, usg string) *FsCmd {
	return &FsCmd{
		Fs: fsctl.NewFs(),
		Command: &cobra.Command{
			Use:   name,
			Short: usg,
		},
	}
}
