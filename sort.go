//source: https://medium.com/vendasta/golang-the-time-complexity-of-append-2177dcfb6bad, https://yourbasic.org/golang/compare-slices/, https://yourbasic.org/golang/quicksort-optimizations/

package main

//Lånat från yourbasic.org, jämför slices T(n) är O(n)
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {

			return false
		}
	}
	return true
}

//sorterar vektorerna i en 2-partition, inspirerat av https://yourbasic.org/golang/quicksort-optimizations/, T(n) är O(n)
func sort(vek []int) (newlist []int) {
	low, high := 0, len(vek)
	for i := 0; i < high; {
		/* Invariant:
		v[:low] < 0
		v[i:] >= 0*/
		if v := vek[i]; v < 0 {
			vek[i] = vek	[low]
			vek[low] = v
			low++
			i++
		} else {
			vek[i] = vek[high-1]
			vek[high-1] = v
			high--
		}
	}
	return vek
}

//testfunktion som testar 4 olika fall och ger panik om sorteringen inte är korrekt
func test() {
	//testar basfallet
	vek := []int{1, -1}
	testvek := []int{-1, 1}

	if !equal(sort(vek), testvek) {
		panic("Fel!")
	}

	//testar fallet 0, -1, 1, för att testa listor med bara ett element
	vek = []int{0}
	testvek = []int{0}
	if !equal(sort(vek), testvek) {
		panic("Fel!")
	}

	vek = []int{-1}
	testvek = []int{-1}
	if !equal(sort(vek), testvek) {
		panic("Fel!")
	}

	vek = []int{1}
	testvek = []int{1}
	if !equal(sort(vek), testvek) {
		panic("Fel!")
	}
}

func main() {
	//vek := []int{-1, 5, -2, -4, 2, 3, 6, -8, -13, 0, 0, -24, -1}
	//fmt.Println(sort(vek))
	test()
}
