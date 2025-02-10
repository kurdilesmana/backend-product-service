package helperModel

type BaseResponseModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

type BaseResponseMobileModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Status  int         `json:"status"`
}

type ErrRespValidationModel struct {
	Message         string      `json:"message"`
	Data            interface{} `json:"data"`
	Error           string      `json:"error"`
	ErrorValidation interface{} `json:"error_validation"`
}
