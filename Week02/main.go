/**
* @Author:wjp
* @Date:2020/12/2 1:13 下午
 */
package main


import "fmt"

func main()  {
	us,err:= UserServer(1)
	if err!=nil{
		return
	}
	fmt.Println(us)
}