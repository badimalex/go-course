package homework09

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (c Employee) Age() int {
	return c.age
}

func (c Customer) Age() int {
	return c.age
}

type User interface {
	Age() int
}

func MaxAge(people ...User) int {
	max := 0
	for _, user := range people {
		if user.Age() > max {
			max = user.Age()
		}
	}
	return max
}
