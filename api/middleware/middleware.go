package middleware

import (
	"fmt"
	iris "github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain string
	Weigth int
	Priority int
}

type Repository interface {
	Read() []*Queue
}

func(q *Queue) Read() []*Queue {
	return MockQueue()
}

//MockeQueue should mock an Array of Queues
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain: "alpha",
			Weigth: 5,
			Priority: 5,
		},
		{
			Domain: "omega",
			Weigth: 1,
			Priority: 5,
		},
		{
			Domain: "beta",
			Weigth: 5,
			Priority: 1,
		},
	}
}

// Que declaration
var Que []*Queue

// InitQueue should return array of data
func InitQueue(){
	Que = append(Que, &Queue{})
}

func ProxyMiddleware(c iris.Context){
	//domain, c.GetHeader("domain")
	var repo Repository
	repo = &Queue{}
	fmt.Println(repo.Read())
	c.Next()
}