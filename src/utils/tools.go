package utils

import (
	"fmt"
	"sort"
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
	"os"
)

func Test()  {
	fmt.Print("222222")
	fmt.Print("aaa")
	fmt.Print("1111")

}

func Sourt(a...int) <-chan int {
	out:=make(chan int)
	go func() {
		for _  , v:=range a{
			out  <-  v
		}
		close(out)
	}()
	return out
}

func InMes(in  <- chan int) <-chan int{
	chanints1 := make(chan int)
	go func() {
		var sot [] int
		for v:= range in {
			sot=append(sot,v)
		}
        sort.Ints(sot)
		for _,v:=range sot{
			chanints1<-v
		}
		close(chanints1)
	}()
	return chanints1
}

//定义一个类型

type Spider struct {
	url string
	header map[string]string
}
//定义一个爬虫get方法

//定义 Spider get的方法
func (keyword Spider) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)
}

func Parse()  {
	header := map[string]string{
		"Host": "movie.douban.com",
		"Connection": "keep-alive",
		"Cache-Control": "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer": "https://movie.douban.com/top250",
	}

	//创建excel文件
	f, err := os.Create("C:/haha3.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//写入标题
	f.WriteString("电影名称"+"\t"+"评分"+"\t"+"评价人数"+"\t"+"\r\n")

	//循环每页解析并把结果写入excel
	for i:=0;i<10;i++{
		fmt.Println("正在抓取第"+strconv.Itoa(i)+"页......")
		url:="https://movie.douban.com/top250?start="+strconv.Itoa(i*25)+"&filter="
		spider := &Spider{url, header}
		html := spider.get_html_header()

		//评价人数
		pattern2:=`<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(pattern2)
		find_txt2 := rp2.FindAllStringSubmatch(html,-1)

		//评分
		pattern3:=`property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		find_txt3 := rp3.FindAllStringSubmatch(html,-1)

		//电影名称
		pattern4:=`img alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		find_txt4 := rp4.FindAllStringSubmatch(html,-1)

		// 写入UTF-8 BOM
		f.WriteString("\xEF\xBB\xBF")
		//  打印全部数据和写入excel文件
		for i:=0;i<len(find_txt2);i++{
			f.WriteString(find_txt4[i][1]+"\t"+find_txt3[i][1]+"\t"+find_txt2[i][1]+"\t"+"\r\n")

		}
	}

}



//文件struct

type FiledDetail struct {
     name string   //文件名
	 typeString string //文件 类型
	 pathString string  //文件存放路径
}

//检测文件是否存在 true 不存在  false 存在
func CheckStatus(fileName string)  bool {
	status:=true
	if info, e := os.Stat(fileName);e!=nil{
		fmt.Println(info)
		status=false
	}
	return status
}