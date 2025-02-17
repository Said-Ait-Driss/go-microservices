package response


type DataResponse struct {
	Data 	 interface{} 	`json:"data,omitempty"`
	Message  string     	`json:"messsage,omitempty"`
	Status   string      	`json:"status,omitempty"`
}
