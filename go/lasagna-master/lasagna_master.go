package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, preTime int) int {
	layersCount := len(layers)
	if preTime == 0 {
		return layersCount * 2
	}
	return layersCount * preTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

	sauce := 0.2
	noodle := 50

	noodleCount := 0
	sauceCount := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodleCount++
		}
		if layer == "sauce" {
			sauceCount++
		}
	}

	return noodleCount * noodle, float64(sauceCount) * sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {

	person := float64(scale) / 2.0
	newQuantities := make([]float64, len(quantities))

	for index, quantity := range quantities {
		newQuantities[index] = quantity * person
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
