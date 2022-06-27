package main

import "fmt"

// 实际场景下 go 的使用
// BaseCurlInter 实现基本增删改查
// user 为用户表，提供es 和 mysql 实现 其中 updateById 需要鉴权重写
// main 里面为 service 调用 只需动态修改 userInter 实现就可切换数据库查询
// TODO: 乐外卖场景下还有哪些问题需要模拟实现
// 		1. 打印机多种实现 每种的支持方法还不同

type BaseCurlInter interface {
	updateById(id int64, info map[string]string) string
	deleteById(id int64) string
}

// mysql 基超
type BaseCurl struct{}

func (b *BaseCurl) updateById(id int64, info map[string]string) string {
	return "base updateById"
}

func (b *BaseCurl) deleteById(id int64) string {
	return "base deleteById"
}

// es 基超
type BaseEsCurl struct{}

func (b *BaseEsCurl) updateById(id int64, info map[string]string) string {
	return "base es updateById"
}

func (b *BaseEsCurl) deleteById(id int64) string {
	return "base es deleteById"
}

// user
type UserInter interface {
	BaseCurlInter
}

type User struct {
	BaseCurl
}

// updateById user需要判断权限 这里覆写公共方法
func (u *User) updateById(id int64, info map[string]string) string {
	return "user updateById"
}

// user es
type UserEs struct {
	BaseEsCurl
}

// updateById user需要判断权限 这里覆写公共方法 es
func (u *UserEs) updateById(id int64, info map[string]string) string {
	return "user es updateById"
}

func main() {
	var userInter UserInter

	// 只需要修改 UserInter 实现即可实现 es he mysql 切换
	user := &User{}
	// user := &UserEs{}
	userInter = user
	s := userInter.updateById(1, map[string]string{"name": "lizhi"})
	fmt.Println(s)
	s = userInter.deleteById(1)
	fmt.Println(s)
}
