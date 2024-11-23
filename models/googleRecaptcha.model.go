package models

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challengeTs"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"errorСodes"`
}

type RecaptchaResponseV3 struct {
	Success     bool     `json:"success"`
	Score       float64  `json:"score"`
	Action      string   `json:"action"`
	ChallengeTS string   `json:"challengeTs"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"errorCodes"`
}
