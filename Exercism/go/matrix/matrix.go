package matrix

import (
    "fmt"
    "strings"
	"strconv"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
    rows := strings.Split(s, "\n")
    m := make(Matrix, len(rows))
    for i,row := range rows {
        cols := strings.Fields(row)
        if i > 0 {
            if len(cols) != len(strings.Fields(rows[i-1])) {
                return nil, fmt.Errorf("")
            }
        }
        m[i] = make([]int, len(cols))
        for j,col := range cols {
            var err error
            m[i][j], err = strconv.Atoi(col)
            if err != nil {
                return nil, fmt.Errorf("")
            }
        }
    }
	return &m, nil
}

func (m Matrix) Cols() [][]int {
    result := make(Matrix, len(m[0]))
    for i := range result {
        result[i] = make([]int, len(m))
    }
    for i:=0; i<len(m); i++ {
        for j:=0; j<len(m[0]); j++ {
            result[j][i] = m[i][j]
        }
    }
	return result
}

func (m Matrix) Rows() [][]int {
    result := make(Matrix, len(m))
    for i := range result {
        result[i] = make([]int, len(m[0]))
    }
    for i:=0; i<len(m); i++ {
        for j:=0; j<len(m[0]); j++ {
            result[i][j] = m[i][j]
        }
    }
	return result
}

func (m *Matrix) Set(row, column, value int) bool {
    if row < 0 || column < 0 || row >= len(*m) || column >= len((*m)[0]) {
        return false
    }
	(*m)[row][column] = value
    return true
}