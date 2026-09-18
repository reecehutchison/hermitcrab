package http

func SayHi() string {
	return "Hi"
}

// 'GET / HTTP/1.1
// Host: 127.0.0.1:8080
// User-Agent: curl/8.7.1
// Accept: */*
//                                <-- (this is a blank line!)
// '

type Message struct {
	Method string
	Version string
	Host string
	UserAgent string
	Accept string
} 

func parseMessage(message string) Message {
	messageMap = make(map[string]string)
	

}