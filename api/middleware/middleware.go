package middleware

import (
	"bufio"
	iris "github.com/kataras/iris"
	"log"
	"os"
	"path/filepath"
	"strings"
	"strconv"
	"container/heap"
)

// Queue
type Queue struct {
	Domain string
	Weigth int
	Priority int
	index int
}

// Que declaration
var Que []string
var readFromFile bool = false

type Repository interface {
	Read() []*Queue
}

func(q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var data []*Queue
	tmp := &Queue{}
	index :=0
	
	for scanner.Scan() {
		if scanner.Text() == "" {
			tmp = &Queue{}
			index++
			continue
		}
		lineArray := strings.Split(scanner.Text(), ":")
		if len(lineArray) == 1 {
			tmp.Domain = lineArray[0]
		}
		if len(lineArray) > 1 {
			val, _ := strconv.Atoi(lineArray[1])
			if lineArray[0] == "weight" {
				tmp.Weigth = val;
			}else if lineArray[0] == "priority"{
				tmp.Priority = val;
				tmp.index = index
				data = append(data, tmp)
			}
		}
	}
	return data;
}

type PriorityQueue []*Queue

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	var sum1 = pq[i].Priority + pq[i].Weigth
	var sum2 = pq[j].Priority + pq[j].Weigth
	return sum1 < sum2
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Queue)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Queue, domain string, weigth int, priority int) {
	item.Domain = domain
	item.Weigth = weigth
	item.Priority = priority
	heap.Fix(pq, item.index)
}

var pq = make(PriorityQueue, 0)

func ProxyMiddleware(c iris.Context){
	Que = Que[:0]
	domain := c.GetHeader("domain")
	priority := c.GetHeader("priority")
	weigth := c.GetHeader("weigth")
	if len(domain) == 0 || len(priority) == 0 || len(weigth) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}

	heap.Init(&pq)

	if(readFromFile == false) {
		var repo Repository
		repo = &Queue{}
		for _, row := range repo.Read() {
			heap.Push(&pq, row)
			pq.update(row, row.Domain, row.Weigth, row.Priority)
		}
		readFromFile = true
	}

	priorityInt, _ := strconv.Atoi(priority)
	weightInt, _ := strconv.Atoi(weigth)
	newEntry := &Queue{}
	newEntry.Domain = domain
	newEntry.Weigth = weightInt
	newEntry.Priority = priorityInt

	heap.Push(&pq, newEntry)
	pq.update(newEntry, newEntry.Domain, newEntry.Weigth, newEntry.Priority)

	for i := pq.Len(); i > 0; i-- {
		Que = append(Que, pq[i-1].Domain)
	}
	c.Next()
}