package dss

import (
	"reflect"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func Test_sumSlice(t *testing.T) {
	type args struct {
		num []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero value",
			args: args{
				num: []float64{0, 0, 0, 0},
			},
			want: 0,
		},
		{
			name: "positive and negative value",
			args: args{
				num: []float64{4, 3, 2, -5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumSlice(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNew2DSlice(t *testing.T) {
	type args struct {
		lRow int
		lCol int
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "3 x 3 slice",
			args: args{
				lCol: 3,
				lRow: 3,
			},
			want: [][]float64{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "3 x 4 slice",
			args: args{
				lCol: 4,
				lRow: 3,
			},
			want: [][]float64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New2DSlice(tt.args.lRow, tt.args.lCol)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newDivisor(t *testing.T) {
	type args struct {
		l       int
		matrixs [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDivisor(tt.args.l, tt.args.matrixs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMooraSpec_normalization(t *testing.T) {
	tests := []struct {
		name    string
		m       MooraSpec
		want    [][]float64
		wantErr bool
	}{
		{
			name: "should pass and return normalization matrixs",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want: [][]float64{
				{0.7427813527082074, 0.3713906763541037, 0.5570860145311556},
				{0.8703882797784892, 0.3481553119113957, 0.3481553119113957},
				{0.7624928516630234, 0.457495710997814, 0.457495710997814},
				{0.4923659639173309, 0.6154574548966636, 0.6154574548966636},
				{0.4242640687119285, 0.565685424949238, 0.7071067811865475},
			},
			wantErr: false,
		},
		{
			name: "should fail cause some field are empty",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail cause some field are empty",
			m: MooraSpec{
				Alternative: []string{},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail cause some field are empty",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{},
				Weight:       []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail cause some field are empty",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail because the length of alternative and assesment is not match",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
					"row 4",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail because the length of assesment column is not match with criteria type or weight",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40},
					{20, 20, 30, 50},
					{30, 20, 30, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
		{
			name: "should fail because the length of assesment column is not match with criteria type or weight",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40},
					{20, 20, 30, 50},
					{30, 20, 30, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			want:    [][]float64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Normalization()
			assert.Equal(t, tt.want, got)
			if tt.wantErr == true {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func TestMooraSpec_optimization(t *testing.T) {
	type args struct {
		normalMatrix [][]float64
	}
	tests := []struct {
		name string
		m    MooraSpec
		args args
		want map[string]float64
	}{
		{
			name: "should pass and return optimization result",
			m: MooraSpec{
				Alternative: []string{
					"row 1",
					"row 2",
					"row 3",
				},
				Assesment: [][]float64{
					{40, 50, 50, 40, 30},
					{20, 20, 30, 50, 40},
					{30, 20, 30, 50, 50},
				},
				CriteriaType: []string{
					"benefit",
					"benefit",
					"benefit",
					"cost",
					"cost",
				},
				Weight: []float64{2.2, 2.1, 2.1, 1.8, 1.8},
			},
			args: args{
				normalMatrix: [][]float64{
					{0.7427813527082074, 0.3713906763541037, 0.5570860145311556},
					{0.8703882797784892, 0.3481553119113957, 0.3481553119113957},
					{0.7624928516630234, 0.457495710997814, 0.457495710997814},
					{0.4923659639173309, 0.6154574548966636, 0.6154574548966636},
					{0.4242640687119285, 0.565685424949238, 0.7071067811865475},
				},
			},
			want: map[string]float64{
				"row 1": 3.4132352932525665,
				"row 2": 0.3828694523657452,
				"row 3": 0.5368407551281025,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Optimization(tt.args.normalMatrix)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetRank(t *testing.T) {
	type args struct {
		list map[string]float64
	}
	tests := []struct {
		name string
		args args
		want RankList
	}{
		{
			name: "t",
			args: args{
				list: map[string]float64{
					"asf":  1.2,
					"adf":  2.2,
					"avcv": 3.5,
				},
			},
			want: []Result{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRank(tt.args.list)
			log.Print(got)
		})
	}
}
