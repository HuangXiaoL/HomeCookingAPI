package recipe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/HuangXiaoL/HomeCookingAPI/pkg/file"
)

func Init() {
	// 赋值为空
	StorageRecipeInfoArr = []*StorageRecipeInfo{}
	for i := 1; i <= 405; i++ {
		getRecipeInfo(i)
	}
	s,_:=json.Marshal(StorageRecipeInfoArr)
	err := ioutil.WriteFile(file.DownloadResources+"/infoDate.txt", s, 0777)
	if err!=nil {
		log.Println("初始化基本数据失败")
	}

}

// InitData 初始化数据
func InitData()  {
	StorageRecipeInfoArr = []*StorageRecipeInfo{}
	resData,err:=ioutil.ReadFile(file.DownloadResources+"/infoDate.txt")
	if err!=nil {
		log.Println(err)
		return
	}

	if err:=json.Unmarshal(resData,&StorageRecipeInfoArr);err!=nil{
		log.Println(err)
		return
	}
	log.Println("初始化本地数据成功")
}

// ReturnDish 返回个菜品
func ReturnDish() *StorageRecipeInfo {
	min := 10
	max := len(StorageRecipeInfoArr)
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return StorageRecipeInfoArr[randNum]
}

// getRecipeInfo 获取菜谱信息
func getRecipeInfo(num int) error {
	client := &http.Client{}
	fmt.Println(url + strconv.Itoa(num) + "/")
	reqest, err := http.NewRequest("GET", url+strconv.Itoa(num)+"/", nil) //建立一个请求
	if err != nil {
		return err
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	q := reqest.URL.Query()

	reqest.URL.RawQuery = q.Encode()
	response, err := client.Do(reqest) //提交
	if err != nil {
		return err
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	doc.Find(`div.s_list > ul > li`).EachWithBreak(func(i int, sel *goquery.Selection) bool {
		name := sel.Find("div.ins > p.name.kw > a").Text()
		href, exists := sel.Find("div.ins > p.name.kw > a").Attr("href")
		if !exists {
			err = fmt.Errorf("第%d个href不存在", i)
			return exists
		}
		imageHref, exists := sel.Find("a > img").Attr("src")
		if !exists {
			err = fmt.Errorf("第%d个href不存在", i)
			return exists
		}
		s := &StorageRecipeInfo{}
		s.Name = name
		s.ImageAddress = imageHref
		s.Link = href
		StorageRecipeInfoArr = append(StorageRecipeInfoArr, s)
		return exists
	})
	return nil
}
