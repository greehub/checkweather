package main

const (
	apiUrl                = "https://api.openweathermap.org/data/2.5/weather"
	MsgMissingParameters  = "missing lat or lon parameter"
	MsgApiError           = "error occurred while fetching data from Open weather api"
	MsgResponseAttributes = "missing attributes in response"
	MsgErrorLoadingEnv    = "Error loading .env file"
)
