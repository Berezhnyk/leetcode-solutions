package levelorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder1(t *testing.T) {
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

func Test_levelOrder2(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, levelOrder(&TreeNode{
		Val: 1,
	}))
}

func Test_levelOrder3(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil))
}

func Test_levelOrder4(t *testing.T) {
	assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9}}, levelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

func Benchmark_levelOrder(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			levelOrder(&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			})
			levelOrder(&TreeNode{
				Val: 1,
			})
			levelOrder(nil)
		}
	})
}
