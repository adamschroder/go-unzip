package unZipper

import (
  "io"
  "os"
  "log"
  "strings"
  "archive/zip"
  "path/filepath"
)

func UnZip (src, dest string) error {

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
    if f.FileInfo().IsDir() {
      os.MkdirAll(fpath, 0755)
    } else {
      var fdir string
      if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
        fdir = fpath[:lastIndex]
      }

      err = os.MkdirAll(fdir, 0755)
      if err != nil {
        log.Fatal(err)
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

  return nil
}