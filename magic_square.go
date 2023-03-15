package main

import "fmt"

//"fmt"

type Matrix [][]int32

func is_magic_square(v []int32) bool {
	var a Matrix
	// Convert vector into 3 X 3 matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = v[3*i+j]
		}
	}

	var s int32 = 0
	for j := 0; j < 3; j++ {
		s += a[0][j]
	}

	var tmp int32
	// Checking if each row sum is same
	for i := 1; i <= 2; i++ {
		tmp = 0
		for j := 0; j < 3; j++ {
			tmp += a[i][j]
			if tmp != s {
				return false
			}
		}
	}

	// Checking if each column sum is same
	for j := 0; j < 3; j++ {
		tmp = 0
		for i := 0; i < 3; i++ {
			tmp += a[i][j]
			if tmp != s {
				return false
			}
		}
	}

	// Checking if diagonal 1 sum is same
	tmp = 0
	for i := 0; i < 3; i++ {
		tmp += a[i][i]
		if tmp != s {
			return false
		}

	}

	// Checking if diagonal 2 sum is same
	tmp = 0
	for i := 0; i < 3; i++ {
		tmp += a[2-i][i]
		if tmp != s {
			return false
		}
		return true
	}

	return false
}

func find_magic_squares(X Matrix) {
	var v []int32
	var i int32
	// Initialing the vector
	for i = 0; i < 9; i++ {
		v[i] = i + 1
	}

	/*
	   for ;;{
	     if is_magic_square(v){

	     }
	*/

}

func diff(a []int32, b []int32) int32 {
	var res int32 = 0
	for i := 0; i < 9; i++ {
		res += Abs(a[i] - b[i])
	}
	return res
}

func wrapper(a []int32) {
	//	res := math.MaxUint32
	//var m_squares Matrix
	//	fmt.Println(find_magic_squares(m_squares))

	/*
	  for (int i = 0; i < magic_squares.size(); ++i) {

	        // Finding the difference with each magic square
	        // and assigning the minimum value.
	        res = min(res, diff(v, magic_squares[i]));
	    }
	    return res;
	*/
}

// Abs returns the absolute value of x.
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var v = []int32{4, 9, 2, 1, 5, 7, 8, 1, 5}
	//wrapper(v)
	fmt.Println(is_magic_square(v))

}
