//最後成功版
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	//"io/ioutil"
	"embed"

	"github.com/amenzhinsky/go-memexec"	
	//"github.com/akavel/rsrc"  //用來換logo
)

//go:embed aachip8051emulator.zip
var zipData embed.FS


func main() {

data, err := zipData.ReadFile("aachip8051emulator.zip")
	if err != nil {
		fmt.Println("無法讀取內嵌資源文件:", err)
		return
	}



//	zipFilePath :=  "C:\\aachipx52ide\\aachip 8051emulator1.zip" // 您的 ZIP 文件路徑

	// 讀取 ZIP 文件
//	zipData, err := ioutil.ReadFile(zipFilePath)
//	if err != nil {
//		fmt.Println("無法讀取 ZIP 文件:", err)
//		return
//	}

	// 將 ZIP 文件數據解壓縮到記憶體中
	//memZipData, err := unzipToMemory(zipData)
	  memZipData, err := unzipToMemory(data)
	if err != nil {
		fmt.Println("無法解壓縮 ZIP 文件到記憶體:", err)
		return
	}

	// 使用 memexec.New 創建一個 memexec.Exec 實例
	exe, err := memexec.New(memZipData)
	if err != nil {
		fmt.Println("無法創建 memexec.Exec:", err)
		return
	}
	defer exe.Close()

	// 添加您的命令行參數
	argv := []string{} // 假設您的命令行參數存在 argv 中

	// 使用 exe.Command 創建一個 exec.Cmd 實例
	cmd := exe.Command(argv...)

	// 執行命令並獲取輸出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("執行命令時出錯:", err)
		return
	}

	// 輸出執行結果
	fmt.Println(string(output))	
	
}

func unzipToMemory(zipData []byte) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		_, err = io.Copy(buf, rc)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

