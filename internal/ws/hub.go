package ws
type Room struct {
	Id string
	Clients map[string]*Client
}
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
}
func NewHub ()*Hub{
	return &Hub{
		Rooms: make(map[string]*Room),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
	}
}
func (h *Hub)Run(){
	for{
		select{
			
		}
	}

}