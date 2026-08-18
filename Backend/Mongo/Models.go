package MongoConfig

type User struct {
	Username		string		`json:Username`
	HashedPassword	string		`json:HashedPassword`
	Email			string		`json:Email`
}
