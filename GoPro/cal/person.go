package cal

type Person struct {
	name string
	age int
}

func (p *Person) SetName(name string){
	p.name = name
}

func (p *Person) SetAge(age int){
	p.age = age
}

func (p *Person) Getname() string{
	return p.name
}

func (p *Person) Getage() string{
	return p.name
}