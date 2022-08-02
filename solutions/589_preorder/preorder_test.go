package preorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorder1(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 5, 6, 2, 4}, preorder(&Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}))
}

func Test_preorder2(t *testing.T) {
	assert.EqualValues(t, []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}, preorder(&Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 2,
			},
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 6,
					},
					&Node{
						Val: 7,
						Children: []*Node{
							&Node{
								Val: 11,
								Children: []*Node{
									&Node{
										Val: 14,
									},
								},
							},
						},
					},
				},
			},
			&Node{
				Val: 4,
				Children: []*Node{
					&Node{
						Val: 8,
						Children: []*Node{
							&Node{
								Val: 12,
							},
						},
					},
				},
			},
			&Node{
				Val: 5,
				Children: []*Node{
					&Node{
						Val: 9,
						Children: []*Node{
							&Node{
								Val: 13,
							},
						},
					},
					&Node{
						Val: 10,
					},
				},
			},
		},
	}))
}

func Benchmark_preorder(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			preorder(&Node{
				Val: 1,
				Children: []*Node{
					&Node{
						Val: 2,
					},
					&Node{
						Val: 3,
						Children: []*Node{
							&Node{
								Val: 6,
							},
							&Node{
								Val: 7,
								Children: []*Node{
									&Node{
										Val: 11,
										Children: []*Node{
											&Node{
												Val: 14,
											},
										},
									},
								},
							},
						},
					},
					&Node{
						Val: 4,
						Children: []*Node{
							&Node{
								Val: 8,
								Children: []*Node{
									&Node{
										Val: 12,
									},
								},
							},
						},
					},
					&Node{
						Val: 5,
						Children: []*Node{
							&Node{
								Val: 9,
								Children: []*Node{
									&Node{
										Val: 13,
									},
								},
							},
							&Node{
								Val: 10,
							},
						},
					},
				},
			})
		}

		preorder(&Node{
			Val: 1,
			Children: []*Node{
				&Node{
					Val: 3,
					Children: []*Node{
						&Node{
							Val: 5,
						},
						&Node{
							Val: 6,
						},
					},
				},
				&Node{
					Val: 2,
				},
				&Node{
					Val: 4,
				},
			},
		})
	})
}
