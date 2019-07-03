package main

import (
	"fmt"
	"testing"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	//handlers "wizeline.github.com/hivanreyes/proxy-app/api/handlers"
	//utils "wizeline.github.com/hivanreyes/proxy-app/api/utils"
	//server "wizeline.github.com/hivanreyes/proxy-app/api/server"
	"sync"
)

type Response struct {
	Status int `json:"status,omitempty"`
	Response string `json:"result,omitempty"`
}

func init(){
	wg := &sync.WaitGroup{}
	go func(wg *sync.WaitGroup){
		main()
		//utils.LoadEnv()
		//app := server.SetUp()
		//handlers.HandlerRedirection(app)
		//server.RunServer(app)
	}(wg)
	wg.Wait()
}

func TestAlgorithmn(t *testing.T){
	cases := []struct {
		// Attrs
		Domain string
		Weigth string
		Priority string 
		Output string
	}{
		{Domain: "test1", Weigth: "1", Priority: "1", Output: `["alpha","beta","omega","test1"]`},
		{Domain: "test2", Weigth: "5", Priority: "4", Output: `["test2","alpha","beta","omega","test1"]`},
		{Domain: "test4", Weigth: "5", Priority: "", Output: "error"},
	}

	valuesToCompare := &Response{}
	client := http.Client{}

	for _, singleCase := range cases{
		req, err := http.NewRequest("GET", "http://localhost:8000/ping", nil)
		req.Header.Add("domain", singleCase.Domain)
		req.Header.Add("weigth", singleCase.Weigth)
		req.Header.Add("priority", singleCase.Priority)
		response, err := client.Do(req)

		bytes, err := ioutil.ReadAll(response.Body)

		err = json.Unmarshal(bytes, valuesToCompare)

		fmt.Println(err, "Error json unmarshal")

		//fmt.Println(valuesToCompare.Status)
		//fmt.Println(valuesToCompare.Response)

		assert.Nil(t, err)

		assert.Equal(t, singleCase.Output, valuesToCompare.Response);
	}

}