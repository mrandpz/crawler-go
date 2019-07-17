
1、单任务爬虫到并发版爬虫  
2、并发版本爬虫   
3、数据存储和展示   
4、分布式爬虫  

ElasticSearch

index/type/id

database/table/字段

### 基本使用方法

index/type/_search 找出所有内容  
index/type/_search?q=golang 找出所有内容

用post 可以省略id，会自动帮我们生成id

ES用的v6版本，一直下不下来，正确的下载方式：  
go get github.com/olivere/elastic  
git checkout -b release-branch.v6 origin/release-branch.v6  

消息传递的方法
REST 
RPC 
中间件  


分布式爬虫

运行 persist 下的itemsaver.go
运行 woker下面的 woker.go
运行 main.go 