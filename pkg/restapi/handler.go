package restapi

import "github.com/emicklei/go-restful/v3"

func Hello(request *restful.Request, response *restful.Response) {
	payload := map[string]string{"payload": "Hello World"}
	err := response.WriteEntity(payload)
	if err != nil {
		return
	}
}
