package model

type Create struct {
	TableName string
	FileName  string
}

type Init struct {
	Driver   string
	Host     string
	Port     int
	DAtabase string
	User     string
	Password string
}

type UpDown struct {
	TableName string
}
