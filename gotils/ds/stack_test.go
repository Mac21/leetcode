package ds

import (
	"testing"
)

func TestPushOne(t *testing.T) {
    st := Stack[int]{}
    st = st.Push(7)

    if st.Len() != 1 {
        t.Fatalf("Expected stack len to be 1 got %d", st.Len())
    }
}

func TestPop(t *testing.T) {
    st := Stack[int]{}
    st = st.Push(7).Push(8)

    st, top, _ := st.Pop()
    if st.Len() != 1 && top != 8 {
        t.Fatalf("Expected Pop 8 from top of stack got %d", top)
    }

    st, top, _ = st.Pop()
    if st.Len() != 0 && top != 7 {
        t.Fatalf("Expected Pop 7 from top of stack got %d", top)
    }

    _, top, empty := st.Pop()
    if !empty {
        t.Fatalf("Expected third pop to be empty got %d", top)
    }
}
