package basicMatrix_test

import (
	"basicMatrix"
	"fmt"
	"testing"
)

func TestPutGet(t *testing.T) {
	matrix := basicMatrix.NewMatrix(4, 3)
	matrix.Put(0, 0, 5.0)

	if matrix.Get(0, 0) != 5.0 {
		t.Error("Values not equal")
	}
}

func TestMultipliedByScalar(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(2, 2)
	mat1.Put(0, 0, 1)
	mat1.Put(0, 1, 2)
	mat1.Put(1, 0, 3)
	mat1.Put(1, 1, 4)

	result := mat1.MultipliedByScalar(2.0)
	if result.Get(0, 0) != 2.0 {
		t.Error("Value not equal")
	}
	if result.Get(0, 1) != 4.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 0) != 6.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 1) != 8.0 {
		t.Error("Value not equal")
	}
}

func TestAdd(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(2, 2)
	mat1.Put(0, 0, 1)
	mat1.Put(0, 1, 2)
	mat1.Put(1, 0, 3)
	mat1.Put(1, 1, 4)

	mat2 := basicMatrix.NewMatrix(2, 2)
	mat2.Put(0, 0, 5)
	mat2.Put(0, 1, 7)
	mat2.Put(1, 0, 11)
	mat2.Put(1, 1, 13)

	result := mat1.Add(mat2)

	if result.Get(0, 0) != 6.0 {
		t.Error("Value not equal")
	}
	if result.Get(0, 1) != 9.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 0) != 14.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 1) != 17.0 {
		t.Error("Value not equal")
	}
}

func TestMultipledBy(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(2, 2)
	mat1.Put(0, 0, 1)
	mat1.Put(0, 1, 2)
	mat1.Put(1, 0, 3)
	mat1.Put(1, 1, 4)

	mat2 := basicMatrix.NewMatrix(2, 2)
	mat2.Put(0, 0, 5)
	mat2.Put(0, 1, 6)
	mat2.Put(1, 0, 7)
	mat2.Put(1, 1, 8)

	result := mat1.MultipliedBy(mat2)
	if result.Get(0, 0) != 19.0 {
		t.Error("Value not equal")
	}
	if result.Get(0, 1) != 22.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 0) != 43.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 1) != 50.0 {
		t.Error("Value not equal")
	}

}

func TestTranspose(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(3, 2)
	mat1.Put(0, 0, 1)
	mat1.Put(0, 1, 2)
	mat1.Put(1, 0, 3)
	mat1.Put(1, 1, 4)
	mat1.Put(2, 0, 5)
	mat1.Put(2, 1, 6)

	result := mat1.Transpose()
	if result.Get(0, 0) != 1.0 {
		t.Error("Value not equal")
	}
	if result.Get(0, 1) != 3.0 {
		t.Error("Value not equal")
	}
	if result.Get(0, 2) != 5.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 0) != 2.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 1) != 4.0 {
		t.Error("Value not equal")
	}
	if result.Get(1, 2) != 6.0 {
		t.Error("Value not equal")
	}

}

func TestMatrixExcludingRowAndCol(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(3, 3)
	mat1.Put(0, 0, 1)
	mat1.Put(0, 1, 2)
	mat1.Put(0, 2, 3)
	mat1.Put(1, 0, 4)
	mat1.Put(1, 1, 5)
	mat1.Put(1, 2, 6)
	mat1.Put(2, 0, 7)
	mat1.Put(2, 1, 8)
	mat1.Put(2, 2, 9)

	var result *basicMatrix.Matrix
	result = mat1.MatrixExcludingRowAndCol(0, 0)

	if result.Get(0, 0) != 5 {
		t.Error("Value not equal")
	}
	if result.Get(0, 1) != 6 {
		t.Error("Value not equal")
	}
	if result.Get(1, 0) != 8 {
		t.Error("Value not equal")
	}
	if result.Get(1, 1) != 9 {
		t.Error("Value not equal")
	}
}

func TestDeterminant(t *testing.T) {
	mat1 := basicMatrix.NewMatrix(3, 3)
	mat1.Put(0, 0, 6)
	mat1.Put(0, 1, 1)
	mat1.Put(0, 2, 1)
	mat1.Put(1, 0, 4)
	mat1.Put(1, 1, -2)
	mat1.Put(1, 2, 5)
	mat1.Put(2, 0, 2)
	mat1.Put(2, 1, 8)
	mat1.Put(2, 2, 7)

	// found this example on the internets
	if mat1.Determinant() != -306 {
		t.Error("Values not equal", mat1.Determinant())
	}
}

func TestInverse(t *testing.T) {
	/*
		mat1 := basicMatrix.NewMatrix(2, 2)
		mat1.Put(0, 0, 4)
		mat1.Put(0, 1, 7)
		mat1.Put(1, 0, 2)
		mat1.Put(1, 1, 6)
	*/
	mat1 := basicMatrix.NewMatrix(5, 5)
	mat1.Put(0, 0, 6)
	mat1.Put(0, 1, 1)
	mat1.Put(0, 2, 1)
	mat1.Put(0, 3, 1)
	mat1.Put(0, 4, 1)

	mat1.Put(1, 0, 4)
	mat1.Put(1, 1, 2)
	mat1.Put(1, 2, 5)
	mat1.Put(1, 3, 1)
	mat1.Put(1, 4, 1)

	mat1.Put(2, 0, 2)
	mat1.Put(2, 1, 8)
	mat1.Put(2, 2, 7)
	mat1.Put(2, 3, 1)
	mat1.Put(2, 4, 1)

	mat1.Put(3, 0, 1)
	mat1.Put(3, 1, 2)
	mat1.Put(3, 2, 3)
	mat1.Put(3, 3, 27)
	mat1.Put(3, 4, 1)

	mat1.Put(4, 0, 1)
	mat1.Put(4, 1, 1)
	mat1.Put(4, 2, 1)
	mat1.Put(4, 3, 1)
	mat1.Put(4, 4, 1)

	// for any matrix with even dimensions, every odd INDEX per row has wrong sign

	mat1.PrettyPrint()

	fmt.Printf("determinant of mat: %f\n", mat1.Determinant())
	fmt.Printf("\n\n")

	result, err := mat1.Inverse()
	if err == nil {
		result.PrettyPrint()
	}
}
