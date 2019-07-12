# gotools
1. gotools是一个带角色、权限、用户管理的Go脚手架
2. gotools可一键生成开箱即用的源码程序，可一键生成简单的增删改查前后端代码
2. 前端使用Vue、ElementUI、vue-element-admin等开源技术
3. 后端使用gin、xormplus、casbin、gf、jwt等开源技术

# 权限层级机构
1. 比如我的系统层级结构是：平台->银行->服务商->商家，我们暂且叫这四个层级为机构类型，如下图：
![组织架构](https://github.com/gopark001/gotools/blob/master/1562936179826.jpg)
2. 业务要求每个机构类型下的机构都能自由定义角色、权限、用户管理
例如：
  平台下管理20个银行，这20个银行都有自己的角色、权限、用户管理
  每个银行管理100个服务商，这100个服务商都有自己的角色、权限、用户管理，以此类推到商家
3. 

## 使用gotools创建项目
* 1. git clone https://github.com/gopark001/gotools
* 2. cd gotools/configs
* 3. 修改mysql.toml，将mysql配置信息改为你的数据库信息，数据库名称可随意
* 4. 修改casbin.toml，将mysql配置信息改为你的数据库信息，且数据库名称必须为casbin
* 5. 修改org_type.json
* 3. go build gotools.go
* 4. gotools -newProject hello

