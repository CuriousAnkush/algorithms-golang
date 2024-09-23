package stockprice

import (
	"math"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want StockPrice
	}{
		{
			name: "T1",
			want: StockPrice{
				auditRecord:  make(map[int]*Price),
				maxTimeStamp: math.MinInt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockPrice_Current(t *testing.T) {
	tests := []struct {
		name    string
		want    []int
		updates [][]int
		max     []int
		min     []int
	}{
		{
			name: "T1",
			want: []int{10, 5, 5},
			updates: [][]int{
				{1, 10},
				{2, 5},
				{1, 3},
			},
			max: []int{10, 10, 5},
			min: []int{10, 5, 3},
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for i := range tt.updates {
				updateRecord := tt.updates[i]
				this.Update(updateRecord[0], updateRecord[1])
				if got := this.Current(); got != tt.want[i] {
					t.Errorf("Current() = %v, want %v", got, tt.want)
				}
				maxGot := this.Maximum()
				minGot := this.Minimum()

				if maxGot != tt.max[i] {
					t.Errorf("Maximum() = %v, want %v", maxGot, tt.max[i])
				}

				if minGot != tt.min[i] {
					t.Errorf("Minimum() = %v, want %v", minGot, tt.min[i])
				}
			}

		})
	}
}
