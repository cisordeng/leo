# ** leo **
> 用户系统，包括管理员，普通用户的注册，登录业务

## Get
`cd $GOPATH/src`
`git clone git@github.com:cisordeng/leo.git`

## Environment
### database

1. 在mysql中创建`leo`数据库: `CREATE DATABASE leo DEFAULT CHARSET UTF8MB4;`；
2. 将`leo`数据库授权给`leo`用户：`GRANT ALL ON leo.* TO 'leo'@localhost IDENTIFIED BY 's:66668888';`；
3. 项目目录下执行 `go run main.go syncdb -v`


## How use this ?

`bee run`

## Document
```

```