/**
interal 代码包让一些程序实体仅仅能被当前模块中的其它代码引用（模块级私有)）。
interal包下申明的公开实体仅能被该代码包的直接父包及其子包 中的代码使用
*/
package internal

import (
	"io"
	"fmt"
)

func Hello(w io.Writer,name string){
	fmt.Fprintf(w,"hello,%s!\n",name);
}
