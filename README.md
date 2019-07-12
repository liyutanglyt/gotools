# gotools
1. gotools是一个带角色、权限、用户管理的Go脚手架
2. gotools可一键生成开箱即用的源码程序，可一键生成简单的增删改查前后端代码
2. 前端使用Vue、ElementUI、vue-element-admin等开源技术
3. 后端使用gin、xormplus、casbin、gf、jwt等开源技术

# 权限层级机构
1. 比如我的系统层级结构是：平台->银行->服务商->商家，我们暂且叫这四个层级为机构类型，如下图：


## 使用gotools创建项目
* 1. git clone http://git.ae-pay.com/wangwei/gotools
* 2. cd gotools/configs
* 3. 修改mysql.toml，将mysql配置信息改为你的数据库信息，数据库名称可随意
* 4. 修改casbin.toml，将mysql配置信息改为你的数据库信息，且数据库名称必须为casbin
* 5. 修改org_type.json
* 3. go build gotools.go
* 4. gotools -newProject hello

