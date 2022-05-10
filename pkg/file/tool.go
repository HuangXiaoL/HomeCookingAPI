package file

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	// ListFilePrefix 列表文件的前缀
	ListFilePrefix = "  "
	// DownloadResources 下载的资源地址
	DownloadResources = "./resources/"
	// DownloadZip 下载的资源压缩包存放地址
	DownloadZip = "./zip/"
)

// Init 初始化文件夹
func Init() {
	PrepareFolders(DownloadResources)
	PrepareFolders(DownloadZip)
}

// ReadNovelFileName 文件名称读取
func ReadNovelFileName(path string) map[string]int {
	arrayFileName := map[string]int{}
	pathSeparator := string(os.PathSeparator)
	level := 1
	fileName := listAllFileByName(level, pathSeparator, path)
	for _, v := range fileName {
		arrayFileName[v] = 1
	}
	return arrayFileName
}

// listAllFileByName 文件列表
func listAllFileByName(level int, pathSeparator, fileDir string) map[int]string {
	var (
		num      = 1                    //计数器
		fileName = make(map[int]string) //文件名称

	)
	files, _ := ioutil.ReadDir(fileDir)
	tmpPrefix := ""
	for i := 1; i < level; i++ {
		tmpPrefix = tmpPrefix + ListFilePrefix
	}
	for _, o := range files {
		if o.IsDir() {
			listAllFileByName(level+1, pathSeparator, fileDir+o.Name())
		} else {
			fileName[num] = tmpPrefix + o.Name()
		}
		num++
	}

	return fileName
}

// PrepareFolders 准备存储文件夹
func PrepareFolders(address string) error {
	if !IsExist(address) {
		err := os.MkdirAll(address, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//IsExist 判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

// ZipFileInfo 压缩文件信息参数
type ZipFileInfo struct {
	ZipName   string         //压缩文件名称
	Parameter []ZipParameter //压缩参数数组
}

// ZipParameter 压缩文件需要的参数
type ZipParameter struct {
	FileAddress string //源文件地址
	Name        string // 处理后的文件名称
}

// ZipFile 压缩文件
func (z *ZipFileInfo) ZipFile() error {
	file, err := os.Create(DownloadZip + z.ZipName + ".zip")
	if err != nil {
		return err
	}
	defer file.Close()

	zipwriter := zip.NewWriter(file)
	defer zipwriter.Close()

	for _, f := range z.Parameter {
		iowriter, err := zipwriter.Create(f.Name)
		if err != nil {
			if os.IsPermission(err) {

				return errors.New("权限不足: " + err.Error())
			}

			return fmt.Errorf("Create file %s error: %s ", f.Name, err.Error())
		}

		content, err := ioutil.ReadFile(DownloadResources + z.ZipName + "/" + f.Name)
		if err != nil {
			content = []byte("")
		}
		iowriter.Write(content)
	}

	// 创建空目录
	zipwriter.Create("/name/dir/")
	return nil
}

// DeleteFile 删除文件
func DeleteFile(folderAddress, fileName string) {
	os.Remove(folderAddress + "/" + fileName)
	files, _ := ListDir(folderAddress, "")
	if folderAddress == DownloadZip {
		return
	}
	if len(files) == 0 {
		os.Remove(folderAddress)
	}
}

// ListDir 获取指定路径下的所有文件，只搜索当前路径，不进入下一级目录，可匹配后缀过滤（suffix为空则不过滤
func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix) //匹配后缀

	for _, _file := range _dir {
		if _file.IsDir() {
			continue //忽略目录
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//文件后缀匹配
			files = append(files, path.Join(dir, _file.Name()))
		}
	}

	return files, nil
}
