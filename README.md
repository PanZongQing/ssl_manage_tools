# 使用方式

1. linux系统创建一个带管理员权限的用户，配置nopasswd去掉sudo密码交互
2. nginx配置文件标准化，统一路径和使用域名作为ssl证书的文件名

### 目录结构

configpath: 本地存放远程服务器拉去下来的nginx配置文件
core: 核心功能，目录检查和ssh方法
server: 后台服务
server/api: 接口方法，包括编辑文档和上传下载功能
templates: 前端模板
uploadier: 上传文件保存路径



