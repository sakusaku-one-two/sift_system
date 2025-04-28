package util




type (
	
	ForEachCallBack[T any] func(value *T)
	MapCallBack[T any] func(value *T) *T
	KeyIdentityCallBack[T any,K comparable] func(value *T) K

	ListI[T any,K comparable] interface {
		Len() int
		Append(target *T)
		Delete(target *T) bool
		Exists(target *T) bool
		ExistsByKey(target *T,callBack KeyIdentityCallBack[T,K]) bool
		ForEach(callBack ForEachCallBack[T])
		Map(callBack MapCallBack[T]) ListI[T,K]
	}

	List[T any,K comparable] struct {
		Values []*T
	} 

)


func NewList[T any,K comparable]() *List[T,K] {
	return new(List[T,K])
} 


func (l *List[T,K]) Len() int {
	return len(l.Values)
}

func (l *List[T,K]) Append(target *T) {
	l.Values = append(l.Values,target)
}

func (l *List[T,K]) Delete(target *T)bool {
	if l.Len() == 0 {
		return false
	}
	
	target_idx := -1
	for idx,value_ptr := range l.Values {
		if target == value_ptr {
			target_idx = idx
			break
		}
	}
	
	if target_idx == -1 {
		return false
	}

	l.Values = append(l.Values[:target_idx],l.Values[target_idx+1:]...)
	return true
} 



func (l *List[T,K]) Exists(target *T) bool {

	for _,value_ptr := range l.Values {
		if target == value_ptr {
			return true
		}
	}

	return false
}

func (l *List[T,K]) ExistsByKey(target *T,callBack KeyIdentityCallBack[T,K]) bool {
	keyValue := callBack(target)
	for _, value_ptr := range l.Values {
		if callBack(value_ptr) == keyValue {
			return true
		}
	}
	return false
}


func (l *List[T,K]) ForEach(callBack ForEachCallBack[T]) {
	for _,ptr_of_value := range l.Values {
		callBack(ptr_of_value)
	}
}

func (l *List[T,K]) Map(callBack MapCallBack[T]) *List[T,K] {
	new_list := NewList[T,K]()
	for _, ptr_of_struct := range l.Values {
		result := callBack(ptr_of_struct)
		if result != nil {
			new_list.Append(result)
		}
	}

	return new_list
}