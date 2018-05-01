## Go爬虫
> 处理和分发爬取珍爱网用户信息的数据任务,并提供关键字收索以及简单前端页面数据展示

- 基于 `Go` 语言标准库开发网络爬虫,不使用现有的爬虫库/框架
- 使用 `Docker` 容器运行 `ElasticSearch` 作为数据存储以及全文收索引擎
- 使用 `Go` 语言标准模板库实现http数据展示部分
- 爬虫框架由单任务版过渡到并发版(多个 `Goroutine}` ),直至分布式爬取数据

## Go爬虫单机版
> 处理和分发爬取珍爱网用户信息的数据任务
### 单任务版爬虫架构

<img src = "http://orj2jcr7i.bkt.clouddn.com/%E5%8D%95%E4%BB%BB%E5%8A%A1%E7%89%88%E7%88%AC%E8%99%AB%E6%9E%B6%E6%9E%84.png" alt="单任务版爬虫架构">

> - 输入请求，通过Engine去解析每个请求，并分发每个请求。
> - 对于一个总请求，通过正则解析出每个城市和每个城市的url,在城市列表解析其中加入城市解析函数，最后加入请求队列中


### 爬虫总体算法

#### 解析器<Parser>

- 输入：`utf-8` 编码的文本
- 输出：`Request{URL，对应Parser}` 列表，`Item` 列表

<img src = "http://orj2jcr7i.bkt.clouddn.com/Parser.png" alt="爬虫总体算法框架">


### 并发版爬虫框架

<img src = "http://orj2jcr7i.bkt.clouddn.com/%E5%B9%B6%E5%8F%91%E7%89%88%E7%88%AC%E8%99%AB%E6%9E%B6%E6%9E%84.png" alt="单任务版爬虫架构">


#### Scheduler实现I 

> **简单调度器**
> - 所有 `Worker` 共用一个输入
> - 实现效率过慢，等待耗时操作

<img src = "https://on-img.com/chart_image/59a84c4ce4b02082b1db046a.png" alt="简单调度器">

#### Scheduler实现II

> **并发调度器**
> - 并发分发Request
> - 控制力弱, 分发出去的goroutine,就收不回来了；并且所有Worker都在抢同一个channel的东西，也没办法控制
> - 限制了负载均衡
> - 调度器使用worker channel来存储请求，开启多个worker groutine去获取请求，在分发任务的engine中等待结果
<img src = "https://on-img.com/chart_image/5ab717c9e4b0a248b0e1bff4.png" alt="并发调度器">

