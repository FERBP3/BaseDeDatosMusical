package modelos

type Rola struct{
	Interprete string
	Titulo string
	Album string
	FechaGrabacion string
	Genero string
	Track string
	Path string
}

type Person struct {
	StageName string
	RealName string
	DateBirth string
	DateDeath string
}

type Group struct {
	Name string
	StartDate string
	EndDate string
	Members []*Person
}