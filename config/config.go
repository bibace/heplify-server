package config

var Setting HeplifyServer

type HeplifyServer struct {
	Network        string `default:"udp"`
	Protobuf       bool   `default:"false"`
	HEPAddr        string `default:"0.0.0.0:9060"`
	HEPWorkers     int    `default:"100"`
	MQName         string `default:""`
	MQAddr         string `default:""`
	MQTopic        string `default:""`
	PromAddr       string `default:""`
	PromTargetIP   string `default:""`
	PromTargetName string `default:""`
	RTPAgent       bool   `default:"false"`
	DBShema        string `default:"homer5"`
	DBDriver       string `default:"mysql"`
	DBAddr         string `default:"localhost:3306"`
	DBUser         string `default:"root"`
	DBPass         string `default:""`
	DBDataTable    string `default:"homer_data"`
	DBConfTable    string `default:"homer_configuration"`
	DBTableSpace   string `default:""`
	DBBulk         int    `default:"200"`
	DBTimer        int    `default:"2"`
	DBRotate       bool   `default:"true"`
	DBPartLog      string `default:"6h"`
	DBPartSip      string `default:"2h"`
	DBPartQos      string `default:"12h"`
	DBDropDays     int    `default:"0"`
	DBDropOnStart  bool   `default:"false"`
	Dedup          bool   `default:"false"`
	AlegID         string `default:"x-cid"`
	LogDbg         string `default:""`
	LogLvl         string `default:"info"`
	LogStd         bool   `default:"false"`
	Config         string `default:"./heplify-server.toml"`
	Version        bool   `default:"false"`
}

func NewConfig() *HeplifyServer {
	return &HeplifyServer{
		Network:        "udp",
		Protobuf:       false,
		HEPAddr:        "0.0.0.0:9060",
		HEPWorkers:     100,
		MQName:         "",
		MQAddr:         "",
		MQTopic:        "",
		PromAddr:       "",
		PromTargetIP:   "",
		PromTargetName: "",
		RTPAgent:       false,
		DBShema:        "homer5",
		DBDriver:       "mysql",
		DBAddr:         "localhost:3306",
		DBUser:         "root",
		DBPass:         "",
		DBDataTable:    "homer_data",
		DBConfTable:    "homer_configuration",
		DBTableSpace:   "",
		DBBulk:         200,
		DBTimer:        2,
		DBRotate:       true,
		DBPartLog:      "6h",
		DBPartSip:      "2h",
		DBPartQos:      "12h",
		DBDropDays:     0,
		DBDropOnStart:  false,
		Dedup:          false,
		AlegID:         "x-cid",
		LogDbg:         "",
		LogLvl:         "info",
		LogStd:         false,
		Config:         "./heplify-server.toml",
		Version:        false,
	}
}

func Get() *HeplifyServer {
	return NewConfig()
}
