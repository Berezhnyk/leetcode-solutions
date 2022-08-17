package detectcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle1(t *testing.T) {
	n1 := &ListNode{
		Val: 3,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 0,
	}
	n4 := &ListNode{
		Val: -4,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	r := detectCycle(n1)
	assert.Equal(t, n2, r)
}

func Test_detectCycle2(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}

	n1.Next = n2
	n2.Next = n1

	r := detectCycle(n1)
	assert.Equal(t, n1, r)
}

func Test_detectCycle3(t *testing.T) {
	n1 := &ListNode{
		Val: 3,
	}

	r := detectCycle(n1)
	assert.Nil(t, r)
}

func Test_detectCycle4(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n1.Next = n2
	n2.Next = n3

	r := detectCycle(n1)
	assert.Nil(t, r)
}

func Test_detectCycle5(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n4 := &ListNode{
		Val: 4,
	}
	n5 := &ListNode{
		Val: 5,
	}
	n6 := &ListNode{
		Val: 6,
	}
	n7 := &ListNode{
		Val: 7,
	}
	n8 := &ListNode{
		Val: 8,
	}
	n9 := &ListNode{
		Val: 9,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8
	n8.Next = n9
	n9.Next = n6

	r := detectCycle(n1)
	assert.Equal(t, n6, r)
}

func Test_detectCycle6(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 1,
	}
	n3 := &ListNode{
		Val: 1,
	}
	n4 := &ListNode{
		Val: 2,
	}
	n5 := &ListNode{
		Val: 2,
	}
	n6 := &ListNode{
		Val: 2,
	}
	n7 := &ListNode{
		Val: 3,
	}
	n8 := &ListNode{
		Val: 3,
	}
	n9 := &ListNode{
		Val: 3,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8
	n8.Next = n9
	n9.Next = n5

	r := detectCycle(n1)
	assert.Equal(t, n5, r)
}

func Benchmark_detectCycle(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			detectCycle(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			})

			detectCycle(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			})

			detectCycle(&ListNode{
				Val: 1,
			})

			n1 := &ListNode{
				Val: 3,
			}
			n2 := &ListNode{
				Val: 2,
			}
			n3 := &ListNode{
				Val: 0,
			}
			n4 := &ListNode{
				Val: -4,
			}

			n1.Next = n2
			n2.Next = n3
			n3.Next = n4
			n4.Next = n2

			detectCycle(n1)
		}
	})
}
