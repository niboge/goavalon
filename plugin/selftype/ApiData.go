package selftype

type Object map[string]interface{}

// 构建model查询
type ModelCond struct{
	Where string
	Bind interface{}
	Column string
	Order string
	Group string
}