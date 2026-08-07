// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/irs/pkg/file"
)

func FuzzCreateFile(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		fl, err := file.CreateFile([]byte(contents))
		if err != nil || fl == nil {
			return
		}
		_ = fl.Validate()
		_ = fl.Ascii()
		_, _ = fl.TCC()
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("{}")
	f.Add("T")

	roots := []string{
		filepath.Join("..", "testdata"),
		filepath.Join("..", "..", "docs", "examples"),
	}
	for _, root := range roots {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			base := strings.ToLower(filepath.Base(path))
			if ext == ".json" || ext == ".ascii" || strings.Contains(base, "ascii") || ext == ".txt" {
				bs, err := os.ReadFile(path)
				if err != nil || len(bs) > 512*1024 {
					return nil
				}
				f.Add(string(bs))
			}
			return nil
		})
	}
}
