package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)
func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("网页读取成功")
			return
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
func work(start, end int) {
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Printf("正在读取第%d页\n", i)
		result, err := httpGet(url)
		if err != nil {
			fmt.Println("读取网页出现错误：", err)
			return
		}
		file, err2 := os.Create("第" + strconv.Itoa(i) + "页.html")
		if err2 != nil {
			fmt.Println("Create err:", err)
			continue
		}
		file.WriteString(result)
		file.Close()
	}

}

func main() {
	var start, end int
	fmt.Print("请输入要开始爬取的页数：")
	fmt.Scan(&start)
	fmt.Print("请输入要结束的爬取页数：")
	fmt.Scan(&end)
	work(start, end)
}
