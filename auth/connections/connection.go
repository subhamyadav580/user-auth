package connections

type DBConnection struct {
	Connect error
	Ping    error
	Close   error
}
