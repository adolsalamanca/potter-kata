package checkout

import (
	"testing"
)

func Test_purchaste(t *testing.T) {
	tests := []struct {
		name  string
		books []Book
		want  float64
	}{
		{
			name: "should return no discount for one book that doesn't much harry potter",
			books: []Book{
				{
					Title: "The automatic millionaire",
					Price: 8.0,
				},
			},
			want: 8.0,
		},
		{
			name: "should return 5 percent discount for two different harry potter books",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},

				{
					Title: HarryPotterSecond,
					Price: 8.0,
				},
			},
			want: 15.2,
		},
		{
			name: "should not do discount buying the same two harry potter books",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},

				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
			},
			want: 16,
		},
		{
			name: "should do 5 percent discount buying 2 different harry potter books, and no discount to the remaining book",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
				{
					Title: HarryPotterSecond,
					Price: 8.0,
				},

				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
			},
			want: 23.2,
		},
		{
			name: "should do 10 percent discount buying 3 different harry potter books",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
				{
					Title: HarryPotterSecond,
					Price: 8.0,
				},

				{
					Title: HarryPotterThird,
					Price: 8.0,
				},
			},
			want: 21.60,
		},
		{
			name: "should do 20 percent discount buying 4 different harry potter books",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
				{
					Title: HarryPotterSecond,
					Price: 8.0,
				},

				{
					Title: HarryPotterThird,
					Price: 8.0,
				},
				{
					Title: HarryPotterFourth,
					Price: 8.0,
				},
			},
			want: 25.60,
		},
		{
			name: "should do 25 percent discount buying 5 different harry potter books",
			books: []Book{
				{
					Title: HarryPotterFirst,
					Price: 8.0,
				},
				{
					Title: HarryPotterSecond,
					Price: 8.0,
				},

				{
					Title: HarryPotterThird,
					Price: 8.0,
				},
				{
					Title: HarryPotterFourth,
					Price: 8.0,
				},
				{
					Title: HarryPotterFifth,
					Price: 8.0,
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purchase(tt.books); got != tt.want {
				t.Errorf("purchaste() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
