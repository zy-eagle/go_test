package main

var eventByName = make(map[string][]func(interface{}))

/**
 * 注册事件.
 * @param name 事件名
 * @param callback 事件回调函数
 */
func RegistryEvent(name string, callback func(interface{})) {
	funcList := eventByName[name]
	funcList = append(funcList, callback)
	eventByName[name] = funcList
}

/**
 * 调用事件.
 * @param name 事件名
 * @param param 事件调用函数参数
 */
func CallEvent(name string, param interface{}) {
	funcList := eventByName[name]
	for _, callback := range funcList {
		callback(param)
	}
}
