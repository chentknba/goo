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

type HandleFunc func(request interface{})

type Service interface {
    SetA()
    SetB()
}


// concrete
type baseService struct {
}

func (s baseService) SetA() {
    fmt.Println("set_a func")
}

func (s baseService) SetB() {
    fmt.Println("set_b func")
}

func NewBasicService() Service{
    return baseService{}
}

// handlefunc set
type HandleFuncSets struct {
    SetAFunc HandleFunc
    SetBFunc HandleFunc
}

func (h HandleFuncSets) SetA() {
    h.SetAFunc(struct{}{})
}

func (h HandleFuncSets) SetB() {
    h.SetBFunc(struct{}{})
}

// wrap
func MakeSetAHandleFunc(s Service) HandleFunc {
    return func(request interface{}) {
        s.SetA()
    }
}

func MakeSetBHandleFunc(s Service) HandleFunc {
    return func(request interface{}) {
        s.SetB()
    }
}

// service holder
type Holder interface{
    Loop()
}

type ServiceHolder struct {
    handlers HandleFuncSets

    inputc chan []byte
}

func (hd *ServiceHolder) Loop() {
    for r := range(hd.inputc) {
        var request Request
		if err := json.Unmarshal(r, &request); err != nil {
			continue
		}

        switch request.Method {
        case "a":
            hd.handlers.SetA()
        }
    }
}

func NewServiceHolder(hs HandleFuncSets) *ServiceHolder {
    return &ServiceHolder{hs, make(chan []byte)}
}

// client holder
type ClientHolder struct {
    ch chan []byte
}

func (c *ClientHolder) Push(data []byte) {
    c.ch <- data
}

func NewClientHolder(sh *ServiceHolder) *ClientHolder {
    return &ClientHolder{sh.inputc}
}

////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	//Gen(Con())
	//select {}

    var service Service
    service = NewBasicService()

    var setAHandler HandleFunc
    setAHandler = MakeSetAHandleFunc(service)

    var setBHandler HandleFunc
    setBHandler = MakeSetBHandleFunc(service)

    handlesets := HandleFuncSets{
        SetAFunc : setAHandler,
        SetBFunc : setBHandler,
    }

    hd := NewServiceHolder(handlesets)
    cd := NewClientHolder(hd)

    go hd.Loop()

    f := func() {
        tick := time.Tick(1*time.Second)

        for _ = range(tick) {
            request := Request{"a", "123"}
            r, err := json.Marshal(request)
            if err != nil {
                fmt.Println("request json marshal err:", err)
                continue
            }

            cd.Push(r)
        }
    }

    go f()
    go f()

    select{}
}




