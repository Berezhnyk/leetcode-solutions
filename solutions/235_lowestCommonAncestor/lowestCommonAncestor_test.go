package lowestcommonancestor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor1(t *testing.T) {
	result := lowestCommonAncestor(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}, &TreeNode{Val: 2}, &TreeNode{Val: 8})
	assert.Equal(t, 6, result.Val)
}

func Test_lowestCommonAncestor2(t *testing.T) {
	result := lowestCommonAncestor(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}, &TreeNode{Val: 2}, &TreeNode{Val: 4})
	assert.Equal(t, 2, result.Val)
}

func Test_lowestCommonAncestor3(t *testing.T) {
	result := lowestCommonAncestor(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}, &TreeNode{Val: 2}, &TreeNode{Val: 1})
	assert.Equal(t, 2, result.Val)
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lowestCommonAncestor(&TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			}, &TreeNode{Val: 2}, &TreeNode{Val: 8})
			lowestCommonAncestor(&TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			}, &TreeNode{Val: 2}, &TreeNode{Val: 4})
		}
	})
}
