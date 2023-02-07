//source: https://gist.github.com/ihkN/3c796e613f179ae223bd283775d904f8

//En modifierad, förkortad version av quicksort används för att hitta typvärdet

package main

//Hittar maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findmax(vek []int) int {
	temp := vek[0]
	for _, v := range vek {
		if temp < v {
			temp = v
		}
	}
	return temp
}

//Hittar maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findmin(vek []int) int {
	temp := vek[0]
	for _, v := range vek {
		if temp > v {
			temp = v
		}
	}
	if temp >= 0 {
		temp = 0
	}
	return temp
}

//Hittar indexet för maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findindex(vek []int) int {
	temp := vek[0]
	index := 0
	for i, v := range vek {
		if temp < v {
			temp = v
			index = i
		}
	}
	return index
}

//gör en tom vektor med längd [max], T(n) är O(n)
func makerange(max, min int) []int {
	vek := make([]int, max+1-min)
	for i := range vek {
		vek[i] = 0 //fyller vektorn med nollor
	}
	return vek
}

//Räknar ut typvärdet, tar in en slice och returnerar en int, T(n) är O(n)
func count(vek []int) int {
	if len(vek) == 0 {
		panic("Listan är tom!")
	}
	max := findmax(vek)
	min := findmin(vek)
	counter := makerange(max, min)

	for _, v := range vek {
		counter[v-min] += 1
	}

	mode := findindex(counter)
	mode += min

	return mode
}

//testfunktion
func tcount() {
	//basfall
	vek := []int{1, 1, 2}
	if count(vek) != 1 {
		panic("Typvärdet ska vara 1 men det erhållna värdet är " + string(count(vek)))
	}

	//testa om negativa tal fungerar
	vek = []int{-1, 2, -1}
	if count(vek) != -1 {
		panic("Typvärdet ska vara 1 men det erhållna värdet är " + string(count(vek)))
	}

	//testar två lika vanliga värden
	vek = []int{-1, 1, -1, 1}
	if count(vek) != -1 {
		panic("Typvärdet ska vara 1 men det erhållna värdet är " + string(count(vek)))
	}
}

func main() {
	tcount()
}
