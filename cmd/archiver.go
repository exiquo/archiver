package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func ArchiveDirectory(srcPath string, tgtPath string) error {
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(tgtPath); os.IsNotExist(err) {
		return err
	}

	files, err := listFiles(srcPath)
	if err != nil {
		return err
	}

	outputFile := fmt.Sprintf("%s%d_archive.tar.gz", tgtPath, time.Now().UnixMilli())

	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	err = createArchive(files, out)
	if err != nil {
		return fmt.Errorf("Error creating archive: %s", err)
	}

	return nil
}

func listFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func createArchive(files []string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err := addToArchive(tw, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return nil
	}

	header.Name = filename

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
