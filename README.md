# 公共组件
## 认证模块
### 基本认证
组件支持的基础认证方式：

- [x] 大于(Gt)
- [x] 大于等于(Ge)
- [x] 小于(Lt)
- [x] 小于等于(Le)
- [x] 不等于(Ne)
- [x] 等于(Eq)
- [x] 非空(NotEmpty)
- [x] 日期格式(IsDateType)
- [x] 密码格式(IsPassword)

### 使用方式
#### 示例
```golang
func init() {
	RegisterVerifies([]Func{
		{
			Name:        "lt",
			Description: "长度或值不在合法范围",
			CallParam:   Lt,
		},
		{
			Name:        "gt",
			Description: "长度或值不在合法范围",
			CallParam:   Gt,
		},
		{
			Name:        "ge",
			Description: "长度或值不在合法范围",
			CallParam:   Ge,
		},
		{
			Name:        "le",
			Description: "长度或值不在合法范围",
			CallParam:   Le,
		},
		{
			Name:        "eq",
			Description: "长度或值不在合法范围",
			CallParam:   Eq,
		},
		{
			Name:        "ne",
			Description: "长度或值不在合法范围",
			CallParam:   Ne,
		},
		{
			Name:        "password",
			Description: "密码格式不正确",
			Call:        IsPassword,
		},
		{
			Name:        "mobile",
			Description: "手机号码格式认证失败",
			Call:        IsMobilePhone,
		},
		{
			Name:        "notEmpty",
			Description: "字段不能为空",
			Call:        NotEmpty,
		},
		{
			Name:        "date",
			Description: "日期格式不准确",
			CallParam:   IsDateType,
		},
	})
}
type Student struct {
	Name       string   `verify:"notEmpty"`
	Age        int      `verify:"gt(5),le(8)"`
	Class      string   `verify:"eq(7)"`
	CreateTime string   `verify:"date(2006-01-02|2006/01/02)"`
	UpdateTime string   `verify:"date(2006-01-02|2006/01/02)"`
	Book       []string `verify:"gt(0)"`
	Password   string   `verify:"password"`
	Mobile     string   `verify:"mobile"`
}

func TestVerify(t *testing.T) {
	err := Verify(Student{
		Name:       "name",
		Age:        8,
		Class:      "testnam",
		CreateTime: "2018/05/05",
		UpdateTime: "2018-05-05",
		Book:       []string{"book"},
		Password:   "pasab145",
		Mobile:     "18010058148",
		Email:      "597410004@qq.com",
	})
}
```
#### 说明
1. 必须在`init()`方法中注册需要使用的认证方式
2. 注册结构体说明: 
    - `name`对应`tag`
    - `description`表示认证错误后提示信息
    - `CallParam`表示该认证方式需要参数
    - `Call`表示该认证方式不需要参数
3. 框架分为参数调用和非参数调用两种方式,参数使用`()`表示,例:`gt(2)`
4. 框架支持多种认证方式同时使用,比如:
    1. 需要一个数字大于3小于20,则`tag`中用`,`分隔开,例:`gt(3),lt(20)`;
    2. 需要多个日期格式认证使用`|`分隔开,例:`date(2006-01-02|2006/01/02)`
#### 自定义认证方式
1. 定义认证方式
2. 注册该认证方式
3. 在字段中添加 `tag`
4. 样例:
```golang
// 定义一个认证email的方法
func email(v interface{}) bool {
	value := v.(reflect.Value)
	if ok, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", value.String()); !ok {
		return false
	}
	return true
}
// 注册认证方法
RegisterVerify(Func{
		Name:        "email",
		Description: "邮箱格式不正确",
		Call:        email,
	})
// 使用
type Email struct{
    Email string `verify:"email"`
}
```
## excel导入和导出
## 模拟数据
## 更多...