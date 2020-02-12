package models

type LynkDriverAvailableAfterOffline struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type   string `json:"type"`
			Fields []struct {
				Type     string `json:"type"`
				Optional bool   `json:"optional"`
				Field    string `json:"field"`
				Name     string `json:"name,omitempty"`
				Version  int    `json:"version,omitempty"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
	} `json:"schema"`
	Payload struct {
		Before  struct {
			DriverID     string `json:"DriverId"`
			TimeInMins   int    `json:"TimeInMins"`
			TripAssigned int    `json:"TripAssigned"`
			TripID       string `json:"TripID"`
			CreatedDate  int64  `json:"CreatedDate"`
		}`json:"before,omitempty"`
		After  struct {
			DriverID     string `json:"DriverId"`
			TimeInMins   int    `json:"TimeInMins"`
			TripAssigned int    `json:"TripAssigned"`
			TripID       string `json:"TripID"`
			CreatedDate  int64  `json:"CreatedDate"`
		} `json:"after,omitempty"`
		Source struct {
			Version       string      `json:"version"`
			Connector     string      `json:"connector"`
			Name          string      `json:"name"`
			TsMs          int64       `json:"ts_ms"`
			Snapshot      string      `json:"snapshot"`
			Db            string      `json:"db"`
			Schema        string      `json:"schema"`
			Table         string      `json:"table"`
			ChangeLsn     interface{} `json:"change_lsn"`
			CommitLsn     string      `json:"commit_lsn"`
			EventSerialNo interface{} `json:"event_serial_no"`
		} `json:"source"`
		Op   string `json:"op"`
		TsMs int64  `json:"ts_ms"`
	} `json:"payload"`
} 
