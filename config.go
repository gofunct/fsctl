package fsctl

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (c *Fs) ReadAllJsonAndYaml() error {
	if err := filepath.Walk(os.Getenv("PWD"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}
		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == "json" {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			if err := c.ReadConfig(bytes.NewBuffer(b)); err != nil {
				panic(err)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Fs) Render(s string) string {
	if strings.Contains(s, "{{") {
		t, err := template.New("").Funcs(sprig.GenericFuncMap()).Parse(s)
		if err != nil {
			c.Exit(1, errFmt, err, "failed to render string")
		}
		buf := bytes.NewBuffer(nil)
		if err := t.Execute(buf, c.AllSettings()); err != nil {
			c.Exit(1, errFmt, err, "failed to render string")
		}
		return buf.String()
	}
	return s
}

func (c *Fs) Sync() {
	for _, e := range os.Environ() {
		sp := strings.Split(e, "=")
		c.SetDefault(strings.ToLower(sp[0]), sp[1])
	}
	for k, v := range c.AllSettings() {
		val, ok := v.(string)
		if ok {
			if err := os.Setenv(k, val); err != nil {
				c.Exit(1, errFmt, err, "failed to bind config to env variable")
			}
		}
	}
}
