package unZipper

import (
	"archive/zip"
	// "bufio"
	// "fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnZip(src, dest, ignoreDir string) error {

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

		if lastIndex := strings.LastIndex(fpath, ignoreDir); lastIndex == -1 {
			if f.FileInfo().IsDir() {
				err = os.MkdirAll(fpath, 0755)
				if err != nil {
					log.Fatal(err)
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
						log.Fatal(err)
						return err
					}

					fi, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
					if err != nil {
						log.Fatal(err)
						return err
					}

					defer fi.Close()

					_, err = io.Copy(fi, rc)
					if err != nil {
						log.Fatal(err)
						return err
					}
				}
			}
		}
	}

	return nil
}
