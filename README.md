## 介绍
dbobj 提供一个操作数据库的工具包，通过直接写sql的方式操作数据库。目前支持mysql数据库

### 获取方法：

> go get github.com/hcdast/dbobj

## 使用方法:

### mysql
1. 请设置环境变量.HBIGDATA_HOME.这个变量中创建目录conf.然后将dbobj中的app.conf复制到conf中.

### 接入项目

```shell
    mkdir $HBIGDATA_HOME/conf
    cp app.conf $HBIGDATA_HOME/conf
```

在指定的配置文件目录中创建配置文件,配置文件名称指定为:app.conf,在文件中输入下面信息:

#### mysql配置文件

```shell
    DB_NAME=
    DB_USERNAME=
    DB_PASSWORD=
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_CHARSET=
    Driver_Name=
    MaxOpenConns=
    MaxIdleConns=
    Location=
```

### 单元测试
在项目根目录下执行以下命令：
```shell
    go test
```