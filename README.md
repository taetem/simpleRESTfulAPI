# A simple API

Used for a basic intro to RESTful API servers in Golang as part of the Women Who Go Utah Meetup. This API features a mock database in the form of a slice of Person objects, and contains no middleware, security, or Swagger documentation. It is a very basic starting point to begin learning more about these topics.

The code used here is heavily based on the code created by Nic Raboy, and is sourced from his article [Create A Simple RESTful API With Golang](https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/). All credit for the endpoints created herein goes to him. Please check out his article for more insight on how to use this code!

To run this mock api, clone the repository into your gopath src folder (ie /home/username/go/src), and from within the cloned repository, run the command `go run main.go`.

If you run into errors running the cloned repository, most likely your issue is that Gorilla Mux is not in your gopath. This can be resolved with either the command `go get github.com/gorilla/mux` or the command `go get -d ./...`, which tells go to recursively download all the 3rd party dependencies required by your code. Make sure to `go build` or `go install` after getting these dependencies!

You can use a tool like [Postman](www.getpostman.com), [Insomnia](insomnia.rest), or [cURL](curl.haxx.se) to test all the endpoints in this code. The GET endpoints can be tested through a web browser - for example, the GetAttendees endpoint can be viewed at localhost:12345/attendees

For more information on how to create HTTP endpoints in Go, check out [Go by Example](https://gobyexample.com/), which covers how to create [HTTP servers](https://gobyexample.com/http-servers) and [clients](https://gobyexample.com/http-clients) without a 3rd party package like Gorilla Mux. You may also find their [JSON](https://gobyexample.com/json) examples useful. 

Other Useful Resources/Topics of Study:

[Using Postman to troubleshoot RESTful APIS](https://www.thepolyglotdeveloper.com/2015/01/using-postman-troubleshoot-restful-api-requests/)
[Go Walkthrough - JSON](https://medium.com/go-walkthrough/go-walkthrough-encoding-json-package-9681d1d37a8f) - This briefly covers the difference between JSON marshalling and JSON encoding/decoding. It is general practice to use an encoder/decoder when dealing with HTTP requests, as in this mock API.
[Encoding/JSON package](https://golang.org/pkg/encoding/json/) - Helpful documentation on the encoding/json golang package, along with examples on how to use it. Focus on the Decoder and Encoder to understand the mock API given here.
[Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)- Covers the improvements to the basic Go HTTP package that mux provides, as well as how to use mux to create your own APIs.
[What is REST?](https://www.codecademy.com/articles/what-is-rest)
[What is CRUD?](https://www.codecademy.com/articles/what-is-crud)


Good luck, and welcome to the exciting world of Web APIs (Application Programming Interfaces)!