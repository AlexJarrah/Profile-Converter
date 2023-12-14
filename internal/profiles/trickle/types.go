package trickle

type Profile struct {
	Keyword      string `csv:"Keyword"`
	Size         string `csv:"SIZE/QUANTITY"`
	FirstName    string `csv:"FIRST NAME"`
	LastName     string `csv:"LAST NAME"`
	Email        string `csv:"EMAIL"`
	PhoneNumber  string `csv:"PHONE NUMBER"`
	Address1     string `csv:"ADDRESS 1"`
	Address2     string `csv:"ADDRESS 2"`
	State        string `csv:"STATE"`
	City         string `csv:"CITY"`
	Zip          string `csv:"ZIP"`
	Country      string `csv:"COUNTRY"`
	CCNumber     string `csv:"CC NUMBER"`
	Month        string `csv:"MONTH"`
	Year         string `csv:"YEAR"`
	CVC          string `csv:"CVC"`
	TaskQuantity string `csv:"Task Quantity"`
	RetryDelay   string `csv:"Retry Delay"`
	MonitorDelay string `csv:"Monitor Delay"`
	Site         string `csv:"Site"`
	Mode         string `csv:"Mode"`
	_2CaptchaKEY string `csv:"2CaptchaKEY"`
	SHIPPINGRATE string `csv:"SHIPPINGRATE"`
}
