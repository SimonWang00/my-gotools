package main

/**
 * @Author: SimonWang00
 * @Description: es7接口使用，翻页、排序、字段过滤，搜索引擎
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/12/26 14:36
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)


var (
	URL = "http://127.0.0.1:9200/"
	INDEX = "goods_price"
	TYPE = "_doc"
	LogFilePath = "./es/log"
)

// 商品价格查询的数据模型结构体
type GoodsPrice struct {
	Commodity_id         string `json:"commodity_id"`
	Shop_id              string `json:"shop_id"`
	Seller_id            string `json:"seller_id"`
	Brand_id             string `json:"brand_id"`
	Category_id          string `json:"category_id"`
	Commodity_name       string `json:"commodity_name"`
	Commodity_tag        string `json:"commodity_tag"`
	Shop_name            string `json:"shop_name"`
	Price                string `json:"price"`
	Price_tag            string `json:"price_tag"`
	Commodity_image_link string `json:"commodity_image_link"`
	Shop_link            string `json:"shop_link"`
	Detail_link          string `json:"detail_link"`
	Review_link          string `json:"review_link"`
	Review_number        string `json:"review_number"`
	Sales_volume         string `json:"sales_volume"`
	Create_date          string `json:"create_date"`
	Update_date          string `json:"update_date"`
	Is_dis               string `json:"is_dis"`
	Discount             string `json:"discount"`
	Brand_name           string `json:"brand_name"`
	Site_id              string `json:"site_id"`
	Is_sell              string `json:"is_sell"`
	Category             string `json:"category"`
	Discount_link        string `json:"discount_link"`
}

//查询结果汇总
type EsData struct {
	Items []interface{} `json:"items"`
}

//组装返回请求
type EsResult struct {
	Status string `json:"status"`
	Total  int64  `json:"total"`
	Data   EsData `json:"data"`
}


// InitES7 初始化ES7参数,初始化,获取es信息
func InitES7()  {
	//es 连接
	EsClient, _ := ESClient()
	info, code, err := EsClient.Ping(URL).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esversion, err := EsClient.ElasticsearchVersion(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

// ESClient创建es客户端
func ESClient() (client *elastic.Client, err error) {
	file := LogFilePath
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766) // 应该判断error，此处简略
	cfg := []elastic.ClientOptionFunc{
		elastic.SetURL(URL),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
		elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)),
	}
	client, err = elastic.NewClient(cfg...)
	return client, err
}

// 批量写es
func BulkAdd(docs []GoodsPrice) {
	client, _ := ESClient()
	bulkRequest := client.Bulk()
	defer client.Stop()
	for _, doc := range docs {
		Commodity_id := doc.Commodity_id
		esRequest := elastic.NewBulkIndexRequest().
			Index(INDEX).
			Type(TYPE).
			Id(Commodity_id).
			Doc(doc)
		bulkRequest = bulkRequest.Add(esRequest)
	}
	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		fmt.Printf("批量写发生错误！err:", err)
	}
}

// isExists查询id是否存在
func isExists(id string) bool {
	client, _ := ESClient()
	defer client.Stop()
	exist, _ := client.Exists().Index(INDEX).Type(TYPE).Id(id).Do(context.Background())
	if !exist {
		//weblog.Info("ID may be incorrect! " + id)
		log.Println("ID may be incorrect! ", id)
		return false
	}
	return true
}

// Get 根据id查询
func (doc *GoodsPrice) Get(id string) {
	client, _ := ESClient()
	defer client.Stop()
	if !isExists(id) {
		return
	}
	esResponse, err := client.Get().Index(INDEX).Id(id).Do(context.Background())
	if err != nil {
		// Handle Error
		return
	}
	json.Unmarshal(esResponse.Source, &doc)
}

//Update 更新文档
func Update(updateField *map[string]interface{}, id string) {
	client, _ := ESClient()
	defer client.Stop()
	if !isExists(id) {
		return
	}
	_, err := client.Update().Index(INDEX).Id(id).Doc(updateField).Do(context.Background())
	if err != nil {
		//Handle Error
	}
}

//Delete 删除文档
func Delete(id string) {
	client, _ := ESClient()
	defer client.Stop()
	if !isExists(id) {
		return
	}
	_, err := client.Delete().Index(INDEX).Id(id).Do(context.Background())
	if err != nil {
		//Handle Error
	}
}

//搜索条件类
type SearchPriceOption struct {
	Brand      string `json:"brand_name"`  //品牌名称
	Limit      int    `json:"limit"`       //每页数量
	Page       int    `json:"offset"`      //页码
	Key        string `json:"key"`         //用户输入搜索关键词
	SiteId     string `json:"site_id"`     //电商平台
	PriceOrder string `json:"price_order"` //价格排序:0-低到高，1-高到低
	SalesOrder string `json:"sales_order"` //销量排序:1-高到低
}

//Search搜索商品价格
//（搜索是ES非常引以为傲的功能，以下示例是相对复杂的查询条件，需将查询结果排序，按页返回）
//SearchPrice 条件搜索，根据用户输入商品短文本搜索返回商品
func (criteria *SearchPriceOption) SearchPrice() (ET EsResult, err error) {
	client, _ := ESClient()
	defer client.Stop()
	//短语搜索 commodity_name字段中 含有关键短文本,采用的是ik分词器
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("commodity_name.IKS", criteria.Key))
	if criteria.PriceOrder != ""{
		// price 不能为空
		boolQ.MustNot(elastic.NewMatchQuery("price",""))
	}
	if criteria.SalesOrder != ""{
		// sales_volume 不能为空
		boolQ.MustNot(elastic.NewMatchQuery("sales_volume",""))
	}
	if criteria.SiteId != "" {
		boolQ.Must(elastic.NewMatchQuery("site_id", criteria.SiteId))
	}
	if criteria.Brand != "" {
		boolQ.Must(elastic.NewMatchQuery("brand_name", criteria.Brand))
	}
	//搜索字段显示高亮
	hl := elastic.NewHighlight()
	hl = hl.Fields(elastic.NewHighlighterField("commodity_name"))
	hl.HighlightFilter(true)
	hl.RequireFieldMatch(true)
	hl = hl.PreTags("<div style='color:red'>").PostTags("</div>")
	flag := NewStrategy(criteria)
	var esResponse *elastic.SearchResult // 接收返回结果
	switch flag {
	case 1:
		esResponse, err = client.Search().
			Index(INDEX).
			Highlight(hl).
			Query(boolQ).
			Sort("price", true).
			From(criteria.Page * criteria.Limit).
			Size(criteria.Limit). //(page) * size
			Pretty(true).
			Do(context.Background())
	case 2:
		esResponse, err = client.Search().
			Index(INDEX).
			Highlight(hl).
			Query(boolQ).
			Sort("price", false).
			From(criteria.Page * criteria.Limit).
			Size(criteria.Limit).
			Pretty(true).
			Do(context.Background())
	case 3:
		esResponse, err = client.Search().
			Index(INDEX).
			Highlight(hl).
			Query(boolQ).
			Sort("sales_volume", true).
			From(criteria.Page * criteria.Limit).
			Size(criteria.Limit).
			Pretty(true).
			Do(context.Background())
	case 4:
		esResponse, err = client.Search().
			Index(INDEX).
			Highlight(hl).
			Query(boolQ).
			Sort("sales_volume", false).
			From(criteria.Page * criteria.Limit).
			Size(criteria.Limit).
			Pretty(true).
			Do(context.Background())
	default:
		esResponse, err = client.Search().
			Index(INDEX).
			Highlight(hl).
			Query(boolQ).
			From(criteria.Page * criteria.Limit).
			Size(criteria.Limit).
			Pretty(true).
			Do(context.Background())
	}
	if err != nil {
		loginf := fmt.Sprintf("查询elasticsearch异常!， elasticsearch,error:", err)
		log.Println(loginf)
		//fmt.Printf("elasticsearch,error:", err)
		panic("查询elasticsearch异常!")
	}
	ET = packagePriceResponse(esResponse)
	return ET, err
}

//NewStrategy设计模式中的策略模式
//翻页、页码、指定店铺、价格排序、销量排序五个条件
//支持的典型搜索场景：
//1.关键词 + 价格升序
//2.关键词 + 价格降序
//3.关键词 + 销量升序
//4.关键词 + 销量降序
//5.默认关键词
func NewStrategy(criteria *SearchPriceOption) (flag int) {
	switch criteria.PriceOrder {
	case "0":
		return 1
	case "1":
		return 2
	}
	switch criteria.SalesOrder {
	case "0":
		return 3
	case "1":
		return 4
	}
	return 5
}


// packagePriceResponse封装返回结果
func packagePriceResponse(esResponse *elastic.SearchResult) (ET EsResult) {
	var docs []interface{}
	for _, value := range esResponse.Hits.Hits {
		var doc *GoodsPrice
		err := json.Unmarshal(value.Source, &doc)
		if err != nil {
			loginf := fmt.Sprintf("elasticsearch查询结果转换错误！,error:", err)
			log.Println(loginf)
		}
		docs = append(docs, *doc)
	}
	ERT := EsResult{
		Status: "sucess",
		Total:  esResponse.Hits.TotalHits.Value,
		Data:   EsData{Items: docs},
	}
	//fmt.Println("ERT:",ERT)
	return ERT
}


func main() {
	opt := &SearchPriceOption{
		Brand:      "华为",
		Limit:      10,
		Page:       1,
		Key:        "P50",
		SiteId:     "京东",
		PriceOrder: "1",
		SalesOrder: "",
	}
	price, err := opt.SearchPrice()
	if err != nil{
		panic(err)
	}
	fmt.Println(price)
}
