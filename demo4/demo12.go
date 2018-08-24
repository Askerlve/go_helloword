package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	//类型检查
	//注意:{}，除了可以代表空代码块之外，还可以表示不包含任何内容的数据结构/数据类型
	//方式1
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}

	fmt.Printf("The element is %q. (container type: %T)\n",
		container[1], container)

	//方式2
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return

	}
	return
}
