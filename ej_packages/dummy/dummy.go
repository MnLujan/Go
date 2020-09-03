package dummy

var Test string

func init() {
	Test = ":)"
}

func dummy() string {
	return "Esta no la podes llamar"
}

func Dummy() string {
	return "Desde el paquete *dummy*"
}
