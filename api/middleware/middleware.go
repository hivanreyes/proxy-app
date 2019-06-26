package middleware

import (
	"bufio"
	"fmt"
	iris "github.com/kataras/iris"
	"log"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

// Queue
type Queue struct {
	Domain string
	Weigth int
	Priority int
}

// Que declaration
var Que []string

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
	
	for scanner.Scan() {
		if scanner.Text() == "" {
			tmp = &Queue{}
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
				data = append(data, tmp)
			}
		}
	}
	return data;
}

func ProxyMiddleware(c iris.Context){
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("FROM HEADER", domain)
	for _, row := range repo.Read() {
		fmt.Println("FROM SOURCE", row.Domain)

		//  ALGORITHM HERE...
		//  USE QUE

	}
	Que = append(Que, domain)

	c.Next()
}