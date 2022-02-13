# This package log and modify data for [http.ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter)

Bu paketi http.ResponseWriter'a yazılan dataları debug edebilmek ve değiştirmek için kullanabilirsiniz. 

## Install
	go get github.com/s4l1h/responselogger

## Usage [Echo](https://echo.labstack.com/) FrameWork 

	e := echo.New()

	e.Use(responselogger.EchoMiddleware(func(b []byte) []byte{
		log.Print(string(b))
		return b
	}))


## Usage [http.Handler](https://golang.org/pkg/net/http/#Handler)

	http.HandleFunc("/path", responselogger.Middleware(handler,func(b []byte) []byte{
		log.Print(string(b))
		return b
	}))

## Complex Example Log and modify data

	func changeAllAdultWords(b []byte) []byte {
		data := string(b)
		log.Print(data)                                 // print log
		data = strings.Replace(data, "porn", "***", -1) // modify data (replace "porn" to "***")
		return []byte(data)                             // return new data for writer
	}
	e := echo.New()
	e.Use(responselogger.EchoMiddleware(changeAllAdultWords))