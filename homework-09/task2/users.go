package homework09

type Employee struct {
	Age int
}

type Customer struct {
	Age int
}

func OldestPerson(people ...any) any {
	max := 0
	var oldest any

	for _, user := range people {
		switch u := user.(type) {
		case Employee:
			if u.Age > max {
				max = u.Age
				oldest = u
			}
		case Customer:
			if u.Age > max {
				max = u.Age
				oldest = u
			}
		}
	}

	return oldest
}
