package dss

import (
	"errors"
	"math"
	"sort"
	"strings"
)

type MooraSpec struct {
	Alternative  []string    `json:"alternative"`
	Assesment    [][]float64 `json:"assesment"`
	CriteriaType []string    `json:"criteria_type"`
	Weight       []float64   `json:"weight"`
}

type Result struct {
	Alternative string
	Score       float64
}

type RankList []Result

// implement sort interface
func (r RankList) Len() int           { return len(r) }
func (r RankList) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RankList) Less(i, j int) bool { return r[i].Score > r[j].Score }

var (
	ErrEmptyField  = errors.New("field is empty")
	ErrAlternative = errors.New("alternative and assesment input is not match")
	ErrAssesment   = errors.New("assesment and criteriatype or weight is not match")
)

func (m MooraSpec) Normalization() ([][]float64, error) {
	if len(m.Alternative) == 0 || len(m.Assesment) == 0 || len(m.CriteriaType) == 0 || len(m.Weight) == 0 {
		return [][]float64{}, ErrEmptyField
	}
	if len(m.Alternative) != len(m.Assesment) {
		return [][]float64{}, ErrAlternative
	}
	if len(m.Assesment[0]) != len(m.CriteriaType) || len(m.Assesment[0]) != len(m.Weight) {
		return [][]float64{}, ErrAssesment
	}

	matrix := New2DSlice(len(m.CriteriaType), len(m.Assesment))
	var index int
	for _, val := range m.Assesment {
		for i, v := range val {
			matrix[i][index] = v
		}
		index++
	}

	divisor := newDivisor(len(m.CriteriaType), matrix)
	for i, val := range matrix {
		for j, v := range val {
			matrix[i][j] = v / divisor[i]
		}
	}

	return matrix, nil
}
func (m MooraSpec) Optimization(normalMatrix [][]float64) RankList {
	matrix := New2DSlice(len(m.Assesment), len(m.CriteriaType))
	for i, val := range normalMatrix {
		for j, v := range val {
			matrix[j][i] = v
		}
	}

	r := make(map[string]float64)
	for idx, val := range matrix {
		max := make([]float64, len(m.CriteriaType))
		min := make([]float64, len(m.CriteriaType))
		for i, v := range val {
			if strings.ToLower(m.CriteriaType[i]) == "benefit" {
				max[i] = v * m.Weight[i]
			}
			if strings.ToLower(m.CriteriaType[i]) == "cost" {
				min[i] = v * m.Weight[i]
			}
		}
		tmp := sumSlice(max) - sumSlice(min)
		r[m.Alternative[idx]] = tmp
	}
	res := getRank(r)
	return res
}

func getRank(list map[string]float64) RankList {
	res := make(RankList, len(list))
	var idx int
	for k, v := range list {
		res[idx] = Result{k, v}
		idx++
	}
	sort.Sort(res)
	return res
}

func sumSlice(num []float64) float64 {
	var res float64
	for _, v := range num {
		res += v
	}
	return res
}

func New2DSlice(lRow, lCol int) [][]float64 {
	col := make([][]float64, lRow)
	for i := range col {
		col[i] = make([]float64, lCol)
	}
	return col
}

func newDivisor(l int, matrixs [][]float64) []float64 {
	div := make([]float64, l)
	var idx int
	for _, val := range matrixs {
		var tmp float64
		for _, v := range val {
			tmp += v * v
		}
		div[idx] = math.Sqrt(tmp)
		idx++
	}
	return div
}
