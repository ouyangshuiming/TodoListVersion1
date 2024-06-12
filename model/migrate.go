package model

//通过go代码在数据库中建表

func migration() {
	//AutoMigrate()接收一个结构体对象的地址作为参数
	//然后根据这个对象中有x个属性就建一张有x个字段的表
	//如果表不存在就会创建表，如果表已经存在而且表结构要
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Task{})

	//表关联关系
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")

}
