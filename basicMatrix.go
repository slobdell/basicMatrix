package basicMatrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	data    []float64
	rows    int
	columns int
}

func NewMatrix(rows, columns int) *Matrix {
	return &Matrix{
		data:    make([]float64, rows*columns, rows*columns),
		rows:    rows,
		columns: columns,
	}
}

func NewIdentityMatrix(rows, columns int) *Matrix {
	if rows != columns {
		panic("Identity matrix should be square")
	}
	matrix := NewMatrix(rows, columns)
	c := 0
	for r := 0; r < rows; r++ {
		matrix.Put(r, c, 1)
		c++
	}
	return matrix
}

func (m *Matrix) rowColToOffset(row, column int) int {
	offset := m.columns * row
	offset += column
	return offset
}
func (m *Matrix) Put(row, column int, value float64) {
	m.data[m.rowColToOffset(row, column)] = value
}

func (m *Matrix) Get(row, column int) float64 {
	return m.data[m.rowColToOffset(row, column)]
}

func (m *Matrix) getRow(row int) []float64 {
	startIndex := m.rowColToOffset(row, 0)
	endIndex := startIndex + m.columns
	return m.data[startIndex:endIndex]
}

func (m *Matrix) getCol(col int) []float64 {
	returnValues := make([]float64, m.rows, m.rows)
	for r := 0; r < m.rows; r++ {
		returnValues[r] = m.getRow(r)[col]
	}
	return returnValues
}

func (m *Matrix) MultipliedByScalar(value float64) *Matrix {
	newData := make([]float64, len(m.data), len(m.data))
	for i := 0; i < len(m.data); i++ {
		newData[i] = m.data[i] * value
	}
	return &Matrix{
		data:    newData,
		rows:    m.rows,
		columns: m.columns,
	}
}

func (m *Matrix) MultipliedBy(otherMatrix *Matrix) *Matrix {
	if m.columns != otherMatrix.rows {
		panic("Matrices are not compatible for multiplication")
	}
	resultantMatrix := NewMatrix(m.rows, otherMatrix.columns)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < otherMatrix.columns; c++ {
			row := m.getRow(r)
			col := otherMatrix.getCol(c)
			value := 0.0
			for i := 0; i < len(row); i++ {
				value += row[i] * col[i]
			}
			resultantMatrix.Put(r, c, value)
		}
	}
	return resultantMatrix
}

func (m *Matrix) Transpose() *Matrix {
	transpose := NewMatrix(m.columns, m.rows)

	for r := 0; r < m.rows; r++ {
		row := m.getRow(r)

		transposeCol := r
		for i := 0; i < len(row); i++ {
			transpose.Put(i, transposeCol, row[i])
		}
	}
	return transpose
}

func (m *Matrix) Add(otherMatrix *Matrix) *Matrix {
	if m.rows != otherMatrix.rows || m.columns != otherMatrix.columns {
		panic("Cannot add matrices of different dimensions")
	}
	newData := make([]float64, len(m.data), len(m.data))
	for i := 0; i < len(m.data); i++ {
		newData[i] = m.data[i] + otherMatrix.data[i]
	}
	return &Matrix{
		data:    newData,
		rows:    m.rows,
		columns: m.columns,
	}
}

func (m *Matrix) Determinant() float64 {
	if m.rows != m.columns {
		panic("Cannot take the determinant of a non-square matrix")
	}
	if m.rows == 1 {
		return m.Get(0, 0)
	} else if m.rows == 2 {
		return m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	} else {
		sign := 1
		total := 0.0
		r := 0
		for c := 0; c < m.columns; c++ {
			value := m.Get(r, c)
			value *= m.MatrixExcludingRowAndCol(r, c).Determinant()
			value *= float64(sign)
			sign *= -1
			total += value
		}
		return total
	}
}

func (m *Matrix) MatrixExcludingRowAndCol(omitRow, omitCol int) *Matrix {
	newSize := (m.rows - 1) * (m.columns - 1)
	newData := make([]float64, newSize, newSize)

	newIndex := 0
	for i := 0; i < len(m.data); i++ {
		if (i / m.columns) == omitRow {
			continue
		} else if (i % m.columns) == omitCol {
			continue
		}
		newData[newIndex] = m.data[i]
		newIndex++
	}
	return &Matrix{
		data:    newData,
		rows:    m.rows - 1,
		columns: m.columns - 1,
	}
}

func (m *Matrix) Inverse() (*Matrix, error) {
	if m.rows != m.columns {
		panic("Cannot take inverse of a non-square matrix")
	}
	if m.Determinant() == 0 {
		return nil, errors.New("Cannot take inverse of matrix")
	}
	newMatrix := NewMatrix(m.rows, m.columns)
	for r := 0; r < m.rows; r++ {
		sign := 1
		if r%2 == 1 {
			sign = -1
		}
		for c := 0; c < m.columns; c++ {
			newMatrix.Put(
				r,
				c,
				float64(sign)*m.MatrixExcludingRowAndCol(r, c).Determinant(),
			)

			sign *= -1
		}
	}
	diagSwapped := newMatrix.Transpose()
	return diagSwapped.MultipliedByScalar(1.0 / m.Determinant()), nil
}

func (m *Matrix) PrettyPrint() {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.columns; c++ {
			fmt.Printf("%.6f  ", m.Get(r, c))
		}
		fmt.Printf("\n")
	}
}
