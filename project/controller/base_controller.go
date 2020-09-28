package controller

import "fmt"

type BaseController struct {
}

func init() {
	fmt.Printf("hello world")
}
func (b *BaseController) Test() {

}
