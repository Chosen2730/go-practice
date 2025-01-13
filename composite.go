package main

import "fmt"

type Composite struct {
	name      string
	maths     string
	chemistry string
	biology   string
	physics   string
	english   string
	utme      int32
}

func (c *Composite) updateName(name string) {
	c.name = name
	fmt.Println("Updated username: ", c.name)
}

func (c *Composite) updateSubjectScore(key string, value string) {
	switch key {
	case "maths":
		c.maths = value
	case "chemistry":
		c.chemistry = value
	case "biology":
		c.biology = value
	case "physics":
		c.physics = value
	case "english":
		c.english = value
	default:
		{
			fmt.Println("Invalid subject entered, accepted subjects are maths, chemistry, biology, physics and english")
			c.updateSubjectScore(key, value)
		}
	}
	fmt.Printf("Updated composite: %v: %v\n", key, value)

}

func createComposite() Composite {
	student := Composite{
		name:      "",
		maths:     "",
		chemistry: "",
		biology:   "",
		physics:   "",
		english:   "",
		utme:      0,
	}
	fmt.Println("Student created successfully")
	return student
}
