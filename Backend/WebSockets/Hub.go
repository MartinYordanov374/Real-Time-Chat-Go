package WebSockets


func NewHub() *Hub{
	return &Hub{
		ActiveClients: make(map[*Client]bool),
	}
}

func (TargetHub *Hub) RegisterClient(Client *Client){
	TargetHub.Mutex.Lock()
	defer TargetHub.Mutex.Unlock()
	TargetHub.ActiveClients[Client] = true
}

func (TargetHub *Hub) UnregisterClient(Client *Client){
	TargetHub.Mutex.Lock()
	defer TargetHub.Mutex.Unlock()

	delete(TargetHub.ActiveClients, Client)
	close(Client.SendChannel)
}
