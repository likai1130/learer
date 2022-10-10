package main

func main()  {
	var fragment Fragment = GetPodAction{}

	fragment.Exec()
}


type Fragment interface {
	Exec() error
}
type GetPodAction struct {}
func (g GetPodAction) Exec() error {
	return nil
}
