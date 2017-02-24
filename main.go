package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////
type Request struct {
	Method string `json:method`
	Param  string `json:param`
}

type Response struct {
	Code int   `json:code`
	Err  error `json:error`
}

func Con() chan []byte {
	ch := make(chan []byte)

	go func() {
		for r := range ch {
			var request Request
			if err := json.Unmarshal(r, &request); err != nil {
				continue
			}

			switch request.Method {
			case "get":
				fmt.Println("get method, param:", request.Param)
			case "set":
				fmt.Println("set method, param:", request.Param)
			default:
				fmt.Println("unknown method")
			}
		}

	}()

	return ch
}

func Gen(ch chan []byte) {
	go func() {
		tick := time.Tick(1 * time.Second)

		for _ = range tick {
			request := Request{"get", "pppppppp"}
			r, err := json.Marshal(request)
			if err != nil {
				fmt.Println("marshal request err:", err)
				continue
			}

			ch <- r
		}
	}()
}

////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	//Gen(Con())
	//select {}

}




