package server

// apiStartJobReq expected API payload for `/start`
type APIStartJobReq struct {
	Command  string `json:"command"`
	Path     string `json:"path"`
	WorkerID string `json:"worker_id"`
}

// apiStartJobRes returned API payload for `/start`
type APIStartJobRes struct {
	JobID string `json:"job_id"`
}

// apiStopJobReq expected API payload for `/stop`
type APIStopJobReq struct {
	JobID    string `json:"job_id"`
	WorkerID string `json:"worker_id"`
}

// apiStopJobRes returned API payload for `/stop`
type APIStopJobRes struct {
	Success bool `json:"success"`
}

// apiQueryJobReq expected API payload for `/query`
type APIQueryJobReq struct {
	JobID    string `json:"job_id"`
	WorkerID string `json:"worker_id"`
}

// apiQueryJobRes returned API payload for `/query`
type APIQueryJobRes struct {
	Done      bool   `json:"done"`
	Error     bool   `json:"error"`
	ErrorText string `json:"error_text"`
}

// apiError is used as a generic api response error
type APIError struct {
	Error string `json:"error"`
}