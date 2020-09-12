# archive/zip の使い方
Docment: https://golang.org/pkg/archive/zip/   
archive/tar でも基本的に同じ.

## 1. 基本
### 1.1. zip する
1. .zip の `file.File` をつくる
2. `zip.Writer`をつくる
3. .zip へ入れるファイルを読み込む
4. 書き込み先のファイルをつくる
5. 読み込んだファイルのデータを .zip 内の書き込み先ファイルへ書き込む
```go
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
```

### 1.2 unzip する
1. `zip.OpenReader()` で zip 内を読み取る
2. `zip.Reader` の `File` からファイルを読み込む
```go
// 1. open reader
r, err := zip.OpenReader("<zip filename>")
if err != nil {
    /* handle error */
}
defer r.Close()

// 2. iterate through the files in the archive
for _, f := range r.File {
    /* do something */
}
```


---
## 2. 応用
あるディレクトリのすべてのファイルをまとめて zip 化する.

`buildArchive()` 処理の流れ 
1. `retreiveFiles()` ですべてのファイルパスを得る (要所は `filepath.Walk()` )
2. `compressFiles()` でそれらのファイルを .zip ファイルへ書き込む  
    1. `addFiles()` で複数のファイルを .zip へ追加している

```go

var (
    // target directory path
    targetDir string
    
    // zip file's name
    outputName string
)

// buildArchive() wraps all functions to generate archive.
func buildArchive() error {
    files, err := retreiveFiles()
    if err != nil {
        return err
    }
    if compressFiles(files) != nil {
        return err
    }
    return nil
}


// retreiveFiles() obtain file list to archive.
func retreiveFiles() ([]string, error) {
    files := []string{}

    // walk in targetDir recursively
    err := filepath.Walk(targetDir, visitDirs(&files))
    if err != nil {
        newErrStr := fmt.Sprintf("Error while visiting directories: %s", err.Error())
        err = errors.New(newErrStr)
        return nil, err
    }

    return files, nil
}


// filepath.WalkFunc wrapper
func visitDirs(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		// skip "directory" files.
		if info.IsDir() {
			return nil
        }
        
		// skip myself
		if filepath.Base(path) == outputName {
			return nil
		}

		// add file to files list
		*files = append(*files, path)
		return err
	}
}


// compressFiles() makes zip file from file list
func compressFiles(files []string) error {
	// set output file
	outFile, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// set zip writer
	zipWriter := zip.NewWriter(outFile)

	// add files into zip
	err = addFiles(zipWriter, files)
	if err != nil {
		return err
	}

	defer zipWriter.Close()
	return nil
}


// addFiles() adds files from file list into a zip file
func addFiles(zipWriter *zip.Writer, files []string) error {

	// do the following against all file
	for _, file := range files {

		// read file data
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		// remove superfluos directory structure
		relativePath := strings.TrimPrefix(file, targetDir)

		// add file in zip file
		f, err := zipWriter.Create(relativePath)
		if err != nil {
			return err
		}

		// write content into file in .zip
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

```


---
## 参考
List the files in a folder with Go (英語)  
https://flaviocopes.com/go-list-files/

