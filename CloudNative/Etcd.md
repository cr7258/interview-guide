## 1.简述 Etcd 及其特点？
答：Etcd 是 CoreOS 团队发起的开源项目，是一个管理配置信息和服务发现（Service Discovery）的项目，它的目标是构建一个高可用的分布式键值对（key-value） 数据库，基于 Go 语言实现。特点：
-	简单：支持 REST 风格的 HTTP + JSON API。
-	安全：支持 HTTPS 方式访问。
-	快速：支持并发 1k/s 的写操作。
-	可靠：支持分布式结构，基于 Raft 的一致性算法，Raft 是一套通过选举主节点来实现分布式系统一致性的算法。