//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:09
//==================================
package jorm

type jMap map[string]interface{}

//设置参数
func (j jMap) Set(key string, value interface{}) {
	j[key] = value
}

//获取参数
func (j jMap) Get(key string) interface{} {
	if j == nil {
		return ""
	}
	jv := j[key]
	return jv
}

//删除参数
func (j jMap) Remove(key string) {
	delete(j, key)
}
