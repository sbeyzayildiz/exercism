package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct{}
type Portuguese struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}
func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func SayHello(name string, germanGreter Greeter) string {
	return "I can speak " + germanGreter.LanguageName() + ": " + germanGreter.Greet(name)
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
