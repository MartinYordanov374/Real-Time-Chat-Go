package WebSockets


func NewHub() *Hub{
	return &Hub{
		ActiveClients: make(map[*Client]bool),
	}
}

func RegisterClient(TargetHub *Hub, Client *Client){
	TargetHub.Mutex.Lock()
	defer TargetHub.Mutex.Unlock()

	TargetHub.ActiveClients[Client] = true
}

func UnregisterClient(TargetHub *Hub, Client *Client){
	TargetHub.Mutex.Lock()
	defer TargetHub.Mutex.Unlock()

	delete(TargetHub.ActiveClients, Client)
	close(Client.SendChannel)
}
