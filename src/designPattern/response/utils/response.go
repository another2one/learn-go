package utils

type ResponseInter interface {
	Check(request *Request)
}

type LoginCheck struct{}

func (lc LoginCheck) Check(request *Request) {
	if request.Name != "lizhi" {
		panic("用户名错误")
	} else {
		request.LoginStatus = true
	}
}

type NameLenCheck struct{}

func (nlc NameLenCheck) Check(request *Request) {
	if len(request.Name) > 10 {
		panic("用户名不能大于10个字符")
	}
}
