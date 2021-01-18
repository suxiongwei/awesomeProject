package main

import (
	"fmt"
	"time"
	"os"
)

// 有init()函数则会先执行init()函数
func init()  {
	fmt.Println("Init")
}

func test()  {
	/**
	使用 := 赋值操作符可以高效地创建一个新的变量，变量类型将由编译器自动推断。
	这是声明变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
	 */
	s := 1
	fmt.Println(s)

	var a int = 10
	var b int = 20

	fmt.Println("多重赋值")
	// 左值和右值按照从左到右的顺序赋值
	b, a = a, b
	fmt.Println(a)
	fmt.Println(b)
}

func GetData() (int, int) {
	return 10,20
}

// 常量组中如果不指定类型和初始值，则与上一行非空常量的值相同
const (
	consta = 10
	constb
	constc
)

const (
	iotaa = iota
	iotab
	iotac
)

func test1()  {
	var a int8 = 4
	var ptr = &a

	fmt.Printf("a的值为 %d\n", a)
	fmt.Printf("& ptr的值为 %d\n", ptr)
	fmt.Printf("* ptr的值为 %d\n", *ptr)

}

func testIf()  {
	if  a := 10; a > 10{
		fmt.Println("大于十")
	}else if a == 10{
		fmt.Println("等于十")
	}else {
		fmt.Println("小于十")
	}
}

func testFor() {
	// 常规写法
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// 省略条件表达式
	for i := 1; ; i++ {
		// 跳出循环
		if i > 10 {
			break
		}
		fmt.Printf("%d ", i)
	}

	// for关键字后只有1个条件表达式，效果类似其他编程语言中的while循环
	i := 0
	for i <= 10 {
		// 跳出循环
		if i > 10 {
			break
		}
		fmt.Printf("%d ", i)
		i ++
	}

	// for关键字后无表达式，效果与其他编程语言的for(;;) {}一致，此时for执行无限循环。
	j := 0
	for  {
		// 跳出循环
		if j > 11 {
			break
		}
		fmt.Printf("%d ", j)
		j++
	}

	/**
	for ... range
	for循环的range格式对string、slice、array、map、channel等进行迭代循环。
	array、slice、string返回索引和值；
	map返回键和值；
	channel只返回通道内的值
	 */
	str := "123ABCabc苏"
	for i, value := range str {
		fmt.Printf("第 %d 的ASCLL值=%d, 字符是%c \n", i, value, value)
	}
}

/**
接受不定数量的参数
 */
func testFun(arg...int8)  {
	
}

/**
测试函数
 */
// 声明一个函数类型
type processFunc func(int) bool

// 判断元素是否为偶数
func isEven(integer int)  bool{
	if integer % 2 ==0 {
		return true
	}
	return false
}

// 判断元素是否为奇数
func isOdd(integer int)  bool{
	if integer % 2 == 0 {
		return false
	}
	return true
}

// 根据函数来处理切片，根据元素奇偶数分组，返回新的切片
func filter(slice []int, f processFunc) []int{
	var result []int
	for _,value := range slice{
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func visit(list []float64, f func(float64))  {
	for _,value := range list{
		f(value)
	}
}

func testArray()  {
	// 如果忽略 [] 中的数字，不设置数组长度，Go语言会根据元素的个数来设置数组的长度。可以忽略声明中数组的长度并将其替换为“…”。编译器会自动计算长度。
	array := [...]int{1,2,3,4}
	fmt.Println(array)
	fmt.Printf("数组长度 %d\n", len(array))

	// 遍历数组方式一
	for i := 0; i < len(array); i++{
		fmt.Println(array[i])
	}

	fmt.Println()

	// 遍历数组方式二
	for _,value := range array{
		fmt.Println(value)
	}
}

func testSlice()  {
	arr0 := [...]string{"a","b","c","d","e","f","g","h","i","j","k"}
	fmt.Println("cap(arr0)", cap(arr0), arr0)

	// 截取数组形成切片,arr中从下标startIndex到endIndex-1下的元素创建为一个新的切片（前闭后开），长度为endIndex-startIndex。
	s01 := arr0[2:8]
	fmt.Printf("%T \n", s01)
	fmt.Println("cap(s01)", cap(arr0), s01)
}

func testMap()  {
	// 短变量声明初始化时间
	rating := map[string]float64{"c":5, "Go":4.5, "Python":4.5, "C++":3}
	fmt.Println(rating)

	// 创建map后赋值
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"

	// 遍历Map
	for k,v := range countryMap{
		fmt.Println("国家：", k, "首都", v)
	}

	// 只展示value
	for _,v := range countryMap{
		fmt.Println( "首都", v)
	}

	// 只展示key
	for k := range countryMap{
		fmt.Println( "国家", k)
	}

	// 判断元素是否存在
	value,ok := countryMap["China"]
	fmt.Println("China是否存在", value, ok)

	value1,ok := countryMap["English"]
	fmt.Println("China是否存在", value1, ok)

	delete(countryMap, "China")
	fmt.Println(countryMap)

	// Go语言没有为map提供清空所有元素的函数，清空map的唯一办法是重新make一个新的map。不用担心垃圾回收的效率，Go语言的垃圾回收比写一个清空函数更高
	countryMap = make(map[string]string)
	fmt.Println(countryMap)
}
// 结构体
type Teacher struct {
	name string
	age int8
	sex byte
}

// 结构体的匿名字段
type User struct {
	string
	int8
}

// 结构体嵌套
type Address struct {
	province, city string
}

type Person struct {
	name string
	age int8
	address *Address
}

// 采用匿名字段模拟继承关系
type Student struct {
	Person
	schoolName string
}

type Employee struct {
	name,currency string
	salary float64
}

// 方法
func (e Employee)printSalary()  {
	fmt.Printf("员工姓名：%s, 薪资：%s%.2f \n", e.name, e.currency, e.salary)
}

// 函数
func printSalary(e Employee) {
	fmt.Printf("员工姓名：%s, 薪资：%s%.2f \n", e.name, e.currency, e.salary)
}

// 接口
type Phone interface {
	call()
}

type AndroidPhone struct {
}

type Iphone struct {
}

func (a AndroidPhone) call()  {
	fmt.Println("我是安卓手机 我信号牛皮")
}

func (a Iphone) call()  {
	fmt.Println("我是苹果手机 我信号垃圾")
}

// 测试异常
func testError()  {
	res, err := Divide(100, 0)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}
}

func Divide(dividee float64, divider float64)(float64, error)  {
	if divider == 0{
		//return 0, errors.New("出错，除数不可以为0")
		return 0, MyError{time.Now(), "出错，除数不可以为0"}
	}else {
		return dividee /divider, nil
	}
}

// 测试自定义异常 定义结构体 表示自定义的异常
type MyError struct {
	when time.Time
	what string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v:%v", e.when, e.what)
}

// 测试defer
func testDefer()  {
	defer funA()
	funB()
	defer funC()
	fmt.Println("testDefer执行结束")
}

func funA()  {
	fmt.Println("这是funA")
}
func funB()  {
	fmt.Println("这是funB")
}
func funC()  {
	fmt.Println("这是funC")
}

// 测试defer参数,延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行。
func testDeferParameter()  {
	a := 1
	b := 2
	// 执行延迟函数的时候 a,b 分别未1，2 而不是实际调用的时候复制未11，12
	defer printAdd(a, b, true)

	a = 11
	b = 12
	printAdd(a, b, false)
}

func printAdd(a, b int, flag bool)  {
	if flag {
		fmt.Println("延迟执行函数，参数a，b分别为%d,%d", a, b)
	}else {
		fmt.Println("未延迟执行函数，参数a，b分别为%d,%d", a, b)
	}
}

func testPanic()  {
	fmt.Println("Hello, World!")
	panic("发生致命错误")
	fmt.Println("over")
}

func testRecover()  {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复啦，获取recover的返回值:", msg)
		}
	}()
	fmt.Println("testRecover")

	for i := 0; i < 10; i++{
		fmt.Println("当前i值", i)

		if i == 5 {
			panic("testRecover 恐慌了")
		}
	}
}


/**
只有package名称为main的包可以包含main()函数。
func main()是程序入口。所有Go函数以关键字func开头，每一个可执行程序都必须包含main()函数，通常是程序启动后第一个执行的函数
*/
func main1()  {
	//fmt.Println("Hello, World!")
	//test()

	// 匿名变量
	//a,_ := GetData()
	//_,b := GetData()
	//
	//c := string(a)
	//fmt.Printf("GetData a 的值为 %d\n", a)
	//fmt.Println(b)
	//fmt.Println(c)
	//
	//var b1 bool = false
	//fmt.Println(b1)
	//
	//fmt.Println("枚举")
	//fmt.Println(consta, constb, constc)
	//
	//fmt.Println("iota")
	//fmt.Println(iotaa, iotab, iotac)

	//test1()
	//testIf()
	//testFor()

	//slice := []int{1,2,3,4,5,6,7}
	//fmt.Println("slice = ", slice)
	//odd := filter(slice, isOdd)
	//fmt.Println("odd = ", odd)
	//even := filter(slice, isEven)
	//fmt.Println("even = ", even)

	// 匿名函数
	//func(data int){
	//	fmt.Println("hello:", data)
	//}(100)
	//
	//// 匿名函数使用方式二
	//f := func(data string) {
	//	fmt.Println(data)
	//}
	//f("欢迎学习go语言")
	//
	//// 匿名函数用作回调函数,第二个参数传入了匿名函数，后面的大括号实现了匿名函数的功能
	//// 调用函数，对每个数进行求平方根操作
	//arr := []float64{1,9,16,25,30}
	//visit(arr, func(v float64) {
	//	v = math.Pow(v, 2)
	//	fmt.Printf("%.0f \n", v)
	//})

	// 空指针
	//var p *int
	//if nil == p {
	//	fmt.Println("值为nil")
	//}

	//testArray()
	//testSlice()
	//testMap()

	// 实例化结构体1
	//var t1 Teacher
	//fmt.Println(t1)
	//
	//t1.name = "Steven"
	//t1.age = 11
	//t1.sex = 1
	//fmt.Println(t1)
	//
	//// 实例化结构体2
	//t2 := Teacher{}
	//t2.name = "Steven1"
	//t2.age = 12
	//t2.sex = 2
	//fmt.Println(t2)
	//
	//// 实例化结构体3
	//t3 := Teacher{
	//	name: "harden",
	//	age: 32,
	//	sex: 1,
	//}
	//
	//fmt.Println(t3)
	//
	//// 实例化结构体4
	//t4 := Teacher{ "iven",32,1}
	//fmt.Println(t4)


	// 匿名结构体 这个相比Java就简洁的多，在定义接口返回值时，可以不用去定义一个类
	//addr := struct {
	//	province, city string
	//}{"陕西", "西安"}
	//fmt.Println(addr)
	//
	//// 结构体的匿名字段
	//user := User{"苏雄伟", 26}
	//fmt.Println(user)

	// 测试结构体嵌套
	//person := Person{}
	//person.name = "蔡虚空"
	//person.age = 11
	//
	//address := Address{city: "徐州", province: "江苏"}
	//person.address = &address
	//fmt.Println(person)
	//fmt.Println(person.address.province)

	// 测试继承
	//student := Student{}
	//student.Person = person
	//student.schoolName = "CUMT"
	//fmt.Println(student)

	// 测试函数与方法
	//empl := Employee{"Giao", "$", 100}
	//printSalary(empl)
	//empl.printSalary()

	// 测试接口
	//var phone Phone
	//phone = new(AndroidPhone)
	//phone.call()

	// 测试异常
	//testError()

	/**
	测试defer,程序输出：
	这是funB
	testDefer执行结束
	这是funC
	这是funA
	*/
	//testDefer()

	// 测试defer参数
	//testDeferParameter()

	// 测试panic
	//testPanic()

	//testRecover()

	// 创建目录
	fileName := "/Users/smzdm/Desktop/test"
	err := os.Mkdir(fileName, os.ModePerm)
	if err != nil{
		fmt.Println("err:", err.Error())
	}else {
		fmt.Println("创建目录成功")
	}
}

