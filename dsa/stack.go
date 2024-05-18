package dsa


type Stack[T any] []T


func (s *Stack[T]) Pop() T{
    arr := *s
    val := arr[0]
    *s = arr[1 : len(arr)]
    return val
}

func (s *Stack[T]) Peek() T {
    arr := *s
    return arr[0]
}

func (s *Stack[T]) Push(x T) {
    *s = append(*s, x)
}
