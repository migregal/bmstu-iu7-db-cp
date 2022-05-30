package layer

func (i *Info) ID() int {
	return i.id
}

func (i *Info) LimitFunc() string {
	return i.limitFunc
}

func (i *Info) SetLimitFunc(limitFunc string) {
	i.limitFunc = limitFunc
}

func (i *Info) ActivationFunc() string {
	return i.activationFunc
}

func (i *Info) SetActivationFunc(activationFunc string) {
	i.activationFunc = activationFunc
}
