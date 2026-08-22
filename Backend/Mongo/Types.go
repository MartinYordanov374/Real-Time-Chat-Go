package MongoConfig

type RequestStatus string

const (
	RequestPending  RequestStatus = "pending"
	RequestAccepted RequestStatus = "accepted"
	RequestRejected RequestStatus = "rejected"
	RequestError RequestStatus = "Request wasn't found"
)
