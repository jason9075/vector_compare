package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func vector_similarity(v1, v2 *mat.VecDense) float64 {
	dot := mat.Dot(v1, v2)
	norm := mat.Norm(v1, 2) * mat.Norm(v2, 2)
	sim := dot / norm
	sim = (sim + 1) / 2
	return sim
}
func random_vector(dim int) *mat.VecDense {
	v := make([]float64, dim)
	for i := 0; i < dim; i++ {
		v[i] = rand.Float64()
	}
	return mat.NewVecDense(dim, v)

}
func main() {
	dimension := 128
	number_of_vector := 20000

	all_members := map[int]*mat.VecDense{}
	for i := 0; i < number_of_vector; i++ {
		all_members[i] = random_vector(dimension)
	}
	compare_vector := random_vector(dimension)

	//compare start
	start := time.Now()
	for i := 0; i < number_of_vector; i++ {
		vector_similarity(all_members[i], compare_vector)
		// fmt.Println(sim)
	}
	fmt.Println("compare time: ", time.Since(start))
}
