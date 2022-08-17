package isvalidbst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST1(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

func Test_isValidBST2(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

func Test_isValidBST3(t *testing.T) {
	assert.Equal(t, true, isValidBST(&TreeNode{
		Val: 1,
	}))
}

func Test_isValidBST4(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
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

func Test_isValidBST5(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}))
}

func Test_isValidBST6(t *testing.T) {
	assert.Equal(t, true, isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

func Test_isValidBST7(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}))
}

func Test_isValidBST8(t *testing.T) {
	assert.Equal(t, true, isValidBST(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -1,
		},
	}))
}

// [5,4,6,null,null,3,7]
func Test_isValidBST9(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

// [1,1]
func Test_isValidBST10(t *testing.T) {
	assert.Equal(t, false, isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}))
}

func Benchmark_isValidBST(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isValidBST(&TreeNode{
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
			isValidBST(&TreeNode{
				Val: 1,
			})
			isValidBST(nil)
		}
	})
}
