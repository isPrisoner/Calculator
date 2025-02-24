package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// 实现栈
type Stack struct {
	elements []interface{}
}

// 创建一个新的栈
func NewStack() *Stack {
	return &Stack{}
}

// 判断栈是否为空
func (s *Stack) empty() bool {
	return len(s.elements) == 0
}

// 把数据添加到栈中
func (s *Stack) push(x interface{}) {
	s.elements = append(s.elements, x)
}

// 弹出栈顶元素
func (s *Stack) pop() (interface{}, error) {
	if s.empty() {
		return nil, errors.New("栈为空")
	}
	ret := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return ret, nil
}

// 查看栈顶元素
func (s *Stack) top() (interface{}, error) {
	if s.empty() {
		return nil, errors.New("栈为空")
	}
	return s.elements[len(s.elements)-1], nil
}

// 实现队列
type Queue struct {
	elements []interface{}
}

// 创建一个新的队列
func NewQueue() *Queue {
	return &Queue{}
}

// 判断队列是否为空
func (q *Queue) empty() bool {
	return len(q.elements) == 0
}

// 把数据添加到队列中
func (q *Queue) push(x interface{}) {
	q.elements = append(q.elements, x)
}

// 弹出列队头元素
func (q *Queue) pop() (interface{}, error) {
	if q.empty() {
		return nil, errors.New("队列为空")
	}
	if len(q.elements) == 1 {
		ret := q.elements[0]
		r := ret.(string)
		println(r)
		q.elements = q.elements[0:0]
		return ret, nil
	} else {
		ret := q.elements[0]
		q.elements = q.elements[1 : len(q.elements)-1]
		return ret, nil
	}
}

// 把接口类型的数字转化为整型
func InterToNum(i interface{}) float64 {
	str := i.(string)
	ret, _ := strconv.ParseFloat(str, 10)
	return ret
}

// 把需要计算的式子转化为后缀表达式
func Transform(S *Stack, Q *Queue, input string) error {
	temp := ""
	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case "+":
			//如果temp中有数字，就把其放入队列
			if temp != "" {
				Q.push(temp)
				temp = ""
			}
			if S.empty() { //如果栈为空，直接入栈
				S.push(string(input[i]))
			} else { // 如果栈不为空
				m, _ := S.top()
				if m.(string) == "(" { //前一个是左括号直接入栈
					S.push(string(input[i]))
				} else { //否则全出
					for {
						t, _ := S.pop()
						Q.push(t.(string))
						a, _ := S.top()
						if S.empty() || a.(string) == "(" { //直到栈为空或者碰到左括号为止
							break
						}
					}
					S.push(string(input[i]))
				}
			}
		case "-":
			//如果temp中有数字，就把其放入队列
			if temp != "" {
				Q.push(temp)
				temp = ""
			}
			if S.empty() { //如果栈为空，直接入栈
				S.push(string(input[i]))
			} else { // 如果栈不为空
				m, _ := S.top()
				if m.(string) == "(" { //前一个是左括号直接入栈
					S.push(string(input[i]))
				} else { //否则全出
					for {
						t, _ := S.pop()
						Q.push(t.(string))
						a, _ := S.top()
						if S.empty() || a.(string) == "(" { //直到栈为空或者碰到左括号为止
							break
						}
					}
					S.push(string(input[i]))
				}
			}
		case "*":
			//如果temp中有数字，就把其放入队列
			if temp != "" {
				Q.push(temp)
				temp = ""
			}
			if S.empty() { //如果栈为空直接入栈
				S.push(string(input[i]))
			} else { //反之，查看栈顶元素
				t, _ := S.top()
				//如果栈顶为加减号或左括号，直接入栈
				if t.(string) == "+" || t.(string) == "-" || t.(string) == "(" {
					S.push(string(input[i]))
				} else {
					for {
						j, _ := S.pop()
						Q.push(j.(string))
						a, _ := S.top()
						//直到栈为空或者碰到左括号为止
						if S.empty() || a.(string) == "(" || a.(string) == "+" || a.(string) == "-" {
							break
						}
					}
					S.push(string(input[i]))
				}
			}
		case "/":
			//如果temp中有数字，就把其放入队列
			if temp != "" {
				Q.push(temp)
				temp = ""
			}
			//如果栈为空直接入栈
			if S.empty() {
				S.push(string(input[i]))
			} else { //反之，查看栈顶元素
				t, _ := S.top()
				//如果栈顶为加减号或左括号，直接入栈
				if t.(string) == "+" || t.(string) == "-" || t.(string) == "(" {
					S.push(string(input[i]))
				} else {
					for {
						j, _ := S.pop()
						Q.push(j.(string))
						a, _ := S.top()
						//直到栈为空或者碰到左括号、加号、减号为止
						if S.empty() || a.(string) == "(" || a.(string) == "+" || a.(string) == "-" {
							break
						}
					}
					S.push(string(input[i]))
				}
			}
		case "(":
			//如果是左括号，则直接放入栈中
			S.push(string(input[i]))
		case ")":
			//如果temp中有数字，就把其放入队列
			if temp != "" {
				Q.push(temp)
				temp = ""
			}
			for {
				//把栈顶元素弹出
				j, _ := S.pop()
				Q.push(j.(string))
				a, _ := S.top()
				if a.(string) == "(" { //直到碰到左括号为止,然后带走左括号
					_, _ = S.pop()
					break
				}
			}
		default:
			if '0' <= input[i] && input[i] <= '9' {
				temp += string(input[i])
			} else {
				return errors.New("输入错误！")
			}
		}
	}
	//把最后一个数字放入队列
	if temp != "" {
		Q.push(temp)
	}
	//若栈还有运算符就出栈
	for {
		if S.empty() {
			break
		}
		t, _ := S.pop()
		Q.push(t.(string))
	}
	return nil
}

// 按照后缀表达式进行运算
func Calculate(S *Stack, Q *Queue) float64 {
	for i := 0; i < len(Q.elements); i++ {
		t := Q.elements[i].(string)
		switch t {
		case "+":
			interNum1, _ := S.pop()
			interNum2, _ := S.pop()
			num1 := InterToNum(interNum1)
			num2 := InterToNum(interNum2)
			ret := num2 + num1
			ret1 := strconv.FormatFloat(ret, 'f', 10, 64)
			S.push(ret1)
		case "-":
			interNum1, _ := S.pop()
			interNum2, _ := S.pop()
			num1 := InterToNum(interNum1)
			num2 := InterToNum(interNum2)
			ret := num2 - num1
			ret1 := strconv.FormatFloat(ret, 'f', 10, 64)
			S.push(ret1)
		case "*":
			interNum1, _ := S.pop()
			interNum2, _ := S.pop()
			num1 := InterToNum(interNum1)
			num2 := InterToNum(interNum2)
			ret := num2 * num1
			ret1 := strconv.FormatFloat(ret, 'f', 10, 64)
			S.push(ret1)
		case "/":
			interNum1, _ := S.pop()
			interNum2, _ := S.pop()
			num1 := InterToNum(interNum1)
			num2 := InterToNum(interNum2)
			ret := num2 / num1
			ret1 := strconv.FormatFloat(ret, 'f', 10, 64)
			S.push(ret1)
		default:
			S.push(t)
		}
	}
	i, _ := S.pop()
	return InterToNum(i)
}

// 判断输入的字符串是否仅含有0-9，一元运算符，左右小括号和小写字母n
func isValidExpression(s string) bool {
	re := regexp.MustCompile(`^[0-9+\-*/n()]+$`)
	return re.MatchString(s)
}

func main() {
	for {
		fmt.Println("如果输入n，则直接退出程序。")
		fmt.Printf("请输入：")
		var input string
		_, err01 := fmt.Scan(&input) //获取输入内容
		if err01 != nil {
			fmt.Println("输入错误，请重新输入！！！")
		}
		if isValidExpression(input) == false {
			fmt.Println("输入错误！！！")
			break
		}
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(string); ok {
					fmt.Println("捕获到错误:", err)
				} else {
					fmt.Println("当前格式不能进行运算")
				}
			}
		}()
		if input == "n" {
			break
		}
		S := NewStack()
		Q := NewQueue()
		err := Transform(S, Q, input)
		if err != nil {
			fmt.Println(err)
		}
		ret := Calculate(S, Q)
		fmt.Println("结果为: ", ret)
	}
}
