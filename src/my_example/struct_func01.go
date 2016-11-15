/*  project: my_example
    author:diogo
    time: 2016/11/15-19:27
*/

package main

import (
	"fmt"
)

func main() {
	p := Person{2, "张三"}

	p.test(1)
	/** out
	Id: 2 Name 张三
	x= 1
	*/

	//把 p.test 方法重新定义
	var f1 func(int) = p.test
	f1(2) //==> p.test(2);
	/** out
	Id: 2 Name: 张三
	x= 2
	*/

	//直接调用 test();
	Person.test(p, 3)
	/**out
	Id: 2 Name: 张三
	x= 3
	*/


	//另类转换
	var f2 func(Person, int) = Person.test
	f2(p, 4)
	/**out
	Id: 2 Name: 张三
	x= 4

	 */


}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name:", this.Name)
	fmt.Println("x=", x)
}
