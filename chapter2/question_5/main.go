package main

import (
	"fmt"
	"math"
	"strconv"
)

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
func main() {
	r := Rectangle{width: 10.43, height: 5.5}
	r.Area()
	r.Perimeter()

	c := Circle{radius: 5.5}
	c.Area()
	c.Perimeter()

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() {
	result := r.width * r.height
	fmt.Println("Rectangle's Area is " + strconv.FormatFloat(result, 'f', 2, 64))
}

func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle's Perimeter is " + strconv.FormatFloat(2*(r.width+r.height), 'f', 2, 64))
}

type Circle struct {
	radius float64
}

func (r Circle) Area() {
	fmt.Println("Circle's Area is " + strconv.FormatFloat(math.Pi*r.radius*r.radius, 'f', 2, 64))
}

func (r Circle) Perimeter() {
	fmt.Println("Circle's Perimeter is " + strconv.FormatFloat(2*math.Pi*r.radius, 'f', 2, 64))
}
