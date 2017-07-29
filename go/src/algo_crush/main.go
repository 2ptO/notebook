package main

import "fmt"

func main() {
	/*
	 * Read N, M
	 * initialize slice of length M, capacity N. Initialize each item in the slice to zero.
	 * Get M operations one by one.
	 * for each operation, add incremeter to the elements between the indices
	 */
	var N, M int64
	fmt.Scanf("%d %d", &N, &M)

	/*
	 * a, b in the queries range between 1 - N, inclusive.
	 * So we need N+1 elements in the data to allow difference sum calculation.
	 */
	data := make([]int64, N+1)

	/*
	 * Using a difference array with prefix sum scan reduces
	 * the complexity from O(MN) to O(M+N). With O(MN) complexity,
	 * this easily times out.
	 */
	for i := int64(0); i < M; i++ {
		var a, b, adder int64
		fmt.Scanf("%d %d %d", &a, &b, &adder)
		data[a] += adder
		if (b + 1) <= N {
			data[b+1] = data[b+1] - adder
		}
	}

	/*
	 * find the maximum through prefix sum scan
	 */
	var max, x int64
	for i := int64(1); i <= N; i++ {
		x = data[i] + x
		if x > max {
			max = x
		}
	}

	fmt.Println(max)
}
