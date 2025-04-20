package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/prometheus/common/log"
	"user/common"
	"user/domain/model"
	"user/domain/service"
	handler "user/handle"
	user "user/proto"
)

func main() {

	//	配置中心
	consulConfig,err := common.GetConsulConfig("127.0.0.1",8500,"/mycro/config")
	if err != nil {
		log.Error(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8001"),
		micro.Registry(consulRegistry),
	)

	mysqlInfo := common.GetMysqlConfig(consulConfig,"mysql")
	//创建数据库连接
	db,err :=gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)


	//初始化服务
	srv.Init()

	//创建服务实例
	userDataService := service.NewUserDataService(model.NewUserRepository(db))
	//注册Handler
	err = user.RegisterUserServiceHandler(srv.Server(),&handler.User{UserDataService:userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
