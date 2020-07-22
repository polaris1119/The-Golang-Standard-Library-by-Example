# 《Go语言标准库》The Golang Standard Library by Example #

Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。

标准库基于最新版本Go。注：目前 Go 标准库文档并没有标识某个 API 基于哪个版本的 Go，将来会加上这部分 [issue](https://github.com/golang/go/issues/5778)。

讲解中涉及到特定操作系统时，针对的都是 Linux/amd64。Go 中相关系统调用在 Linux 下，对于同一个系统调用，如果有对应的 `at` 版本，使用的都是 `at` 版本，如 `open` 系统调用使用都是 `openat`。更多信息参考 [Go语言中文网博客中关于系统调用的文章](http://blog.studygolang.com)。

## 交流 ##

欢迎大家加入QQ群：192706294 《Go语言实现与标准库》交流群

Go语言构建的 Go语言中文网：[http://studygolang.com](http://studygolang.com)

关注作者公众号，加微信好友、进微信交流群。

![](polarisxu-qrcode-small.jpg)

## 阅读 ##

为了更方便阅读，Go语言中文网搭建了阅读平台，可以更友好的在线阅读。

[Go语言中文网——Go语言标准库](http://books.studygolang.com/The-Golang-Standard-Library-by-Example)

## 捐赠 ##

如果您觉得本书对您有帮助，通过微信或支付宝捐赠作者，金额随意！

**由于无法从支付方获取支付者信息，请在支付的留言备注功能中附上 Go语言中文网账户的昵称等信息，以便我们记录！**

## 目录 ##
[目录](directory.md)
## 贡献者 ##

[hikerell](https://github.com/hikerell)

## 反馈 ##

由于本人能力有限，书中难免有写的不对之处，且目前所写内容没有经过校正。如果阅读过程中有任何疑问或觉得不对之处，欢迎提出，谢谢！

## 版权声明 ##

本书所有内容遵循 [CC-BY-SA 3.0协议（署名-相同方式共享）](http://zh.wikipedia.org/wiki/Wikipedia:CC-by-sa-3.0%E5%8D%8F%E8%AE%AE%E6%96%87%E6%9C%AC)

1. 常见误解

2. 常用手法

3. 如何理解，使用

4. 为什么接口如此组织

5. 和其它语言对比优缺点
