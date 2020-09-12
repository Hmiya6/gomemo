package main

import (
	"archive/zip"
	"io/ioutil"
	"os"
)

func main() {
	// 1. create zip file
	zipFile, err := os.Create("<filename1>")
	if err != nil {
		/* handle error */
	}
	defer zipFile.Close()

	// 2. set a zip.Writer from zipFile
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 3. read a file
	data, err := ioutil.ReadFile("<filename2>")
	if err != nil {
		/* handle error */
	}

	// 4. make file to writer in .zip
	f, err := zipWriter.Create("filepath in .zip")
	if err != nil {
		/* handle error */
	}

	// 5. write the data in file in .zip
	_, err = f.Write(data)
	if err != nil {
		/* handle error */
	}
}
