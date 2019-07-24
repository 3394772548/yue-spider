package main

import (
	"./docs"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
)

var (
	h         bool
	port      string
	accessKey string
	proxyUrl  string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&port, "port", "8080", "set http port")
	flag.StringVar(&accessKey, "accessKey", "", "set accessKey")
	flag.StringVar(&proxyUrl, "proxy", "", "set proxy")
}

// @title  yue-spider API
// @version 1.0
// @description 淘宝App 闲鱼App 相关Api 调用，仅供测试使用。 点击接口展开可以查看详情，点击"Try it out" 填写参数即可进行接口测试 更多交流 3394772548@qq.com
// @contact.name API Support
// @contact.url
// @contact.email 3394772548@qq.com
// @host
// @BasePath /
func main() {
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	docs.SwaggerInfo.Title = "yue-spider"
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Static("/docs", "./docs")
	url := ginSwagger.URL("/docs/swagger.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//宝贝基本信息
	r.GET("/mtop.taobao.detail.getdetail", MtopTaobaoDetailGetDetail)
	//宝贝详情
	r.GET("/mtop.taobao.detail.getdesc", MtopTaobaoDetailGetdesc)
	//宝贝评价
	r.GET("/mtop.taobao.rate.detaillist.get", MtopTaobaoRateDetaillistGet)
	//淘宝搜索
	r.GET("/mtop.taobao.wsearch.appsearch", MtopTaobaoWsearchAppsearch)
	//淘宝店铺搜索
	//r.GET("/mtop.taobao.wsearch.appsearch.shop",MtopTaobaoWsearchAppsearchShop)
	//淘宝买家秀
	r.GET("/mtop.taobao.ugc.tql.facade", MtopTaobaoUgcTqlFacade)
	//闲鱼搜索
	r.GET("/mtop.taobao.idle.main.item.search", MtopTaobaoIdleMainItemSearch)
	//闲鱼宝贝基本信息
	r.GET("/mtop.taobao.idle.item.detail", MtopTaobaoIdleItemDetail)
	//闲鱼评论
	//
	//淘口令解析
	r.GET("/mtop.taobao.sharepassword.querypassword", MtopTaobaoSharepasswordQuerypassword)
	log.Println("爬虫系统已启动.API接口文档地址: http://localhost:" + port + "/swagger/index.html")
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
