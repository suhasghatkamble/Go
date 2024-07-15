package main
import "fmt"

//Author structure
type author struct{
	name string
	branch string
	particles int 
	salary int 
}

//Methods with a receiver
func (a author) show(){
	fmt.Println("Author's Name:",a.name)
	fmt.Println("Author's Branch:",a.branch)
	fmt.Println("Author's Particles:",a.particles)
	fmt.Println("Author's Salary:",a.salary)
}

type Mutatable struct{
	a int 
	b int 
}

func (m *Mutatable) StayTheSame(){
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate(){
	m.a = 10
	m.b = 20
}

func main(){
	res := author{
		name:"Sona",
		branch:"CSE",
		particles: 203,
		salary: 34000,
	}

	res.show()

	m := &Mutatable{9,0}
	fmt.Println(m)

	m.StayTheSame()
	fmt.Println(m)

	m.Mutate()
	fmt.Println(m)
}