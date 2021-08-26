package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "io"
    "strings"
)
//匹配的关键字
var x string = "jdbc"
var x1 string = "user"
var filename = "./ConfigFind.txt"
var f1 *os.File
var err1 error

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil, nil, err
    }

    PthSep := string(os.PathSeparator)
    //suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

    for _, fi := range dir {
        if fi.IsDir() { // 目录, 递归遍历
            dirs = append(dirs, dirPth+PthSep+fi.Name())
            GetFilesAndDirs(dirPth + PthSep + fi.Name())
        } else {
            // 过滤指定格式
            ok := strings.HasSuffix(fi.Name(), ".img")
            if ok {
                files = append(files, dirPth+PthSep+fi.Name())
            }
        }
    }

    return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
    var dirs []string
    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil, err
    }

    PthSep := string(os.PathSeparator)
    //suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

    for _, fi := range dir {
        if fi.IsDir() { // 目录, 递归遍历
            dirs = append(dirs, dirPth+PthSep+fi.Name())
            GetAllFiles(dirPth + PthSep + fi.Name())
        } else {
            // 过滤指定格式
            ok := strings.HasSuffix(fi.Name(), ".xml")
            ok1 := strings.HasSuffix(fi.Name(), ".profile")
            ok2 := strings.HasSuffix(fi.Name(), ".config")
            if ok || ok1 || ok2 {
                files = append(files, dirPth+PthSep+fi.Name())
            }

        }
    }

    // 读取子目录下文件
    for _, table := range dirs {
        temp, _ := GetAllFiles(table)
        for _, temp1 := range temp {
            files = append(files, temp1)
        }
    }

    return files, nil
}

//从获取的文件中查找特定值 jdbc 不区分大小写
func findfile(f string) {
	content, err := ioutil.ReadFile("./" + f)
	if err != nil {
		fmt.Println("读取文件失败，错误:", err)
		return
	}
	a := string(content)
	if strings.Contains((strings.ToLower(a)), x) || strings.Contains((strings.ToLower(a)), x1){
		fmt.Println("该文件可能为数据库配置文件:" + f)
		if checkFileIsExist(filename) { //如果文件存在
        f1, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
        
	    } else {
	        f1, err1 = os.Create(filename) //创建文件
	        
	    }
	    check(err1)
	    io.WriteString(f1, f + "\r\n") //写入文件(字符串)
	    
	    
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

func main() {
    files, dirs, _ := GetFilesAndDirs("./")

    for _, dir := range dirs {
        fmt.Printf("当前文件夹为[%s]\n", dir)
    }
    fmt.Printf("=======================================\n")
    for _, table := range dirs {
        temp, _, _ := GetFilesAndDirs(table)
        for _, temp1 := range temp {
            files = append(files, temp1)
        }
    }

    for _, table1 := range files {
        
        findfile(table1)
    }


    fmt.Printf("=======================================\n")
    xfiles, _ := GetAllFiles("./")
    for _, file := range xfiles {
     
        findfile(file)
    }
}
