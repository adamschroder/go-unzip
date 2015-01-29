package unZipper

import (
	"archive/zip"
	// "bufio"
	// "fmt"
	"io"
	// "log"
	"os"
	"path/filepath"
	"strings"
)

func UnZip(src, dest string) error {

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		// fmt.Println(f.Name)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(fpath, 0755)
			if err != nil {
				return err
			}
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			if len(f.Name) < 150 {

				err = os.MkdirAll(fdir, 0755)
				if err != nil {
					return err
				}

				f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}

				defer f.Close()

				_, err = io.Copy(f, rc)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
