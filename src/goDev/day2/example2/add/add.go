package add

var Name string
var Age int

// init 方法会在 文件全局变量申明之后执行
// 在 main 方法执行之前执行
func init() {
	Name = "alex"
	Age = 23
}
