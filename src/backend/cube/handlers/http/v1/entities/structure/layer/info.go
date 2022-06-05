package layer

type Info struct {
	ID             int `json:"id"`
	LimitFunc      string `json:"limit_func"`
	ActivationFunc string `json:"activation_func"`
}
