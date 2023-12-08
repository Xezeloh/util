package function

import (
	"reflect"
)

type (
	// ConsumeFunc 函数类型消费一个泛型值
	ConsumeFunc[T any] func(T)
	// SupplyFunc 函数类型提供一个泛型值
	SupplyFunc[T any] func() T
)

// ConsumeChain 内含一个 ConsumeFunc 类型成员，用于将 ConsumeFunc 调用封装为链式调用形式。
type ConsumeChain[T any] struct {
	consumeFunc ConsumeFunc[T]
}

// NewConsumeChain 将一个 ConsumeFunc 类型的值封装为 ConsumeChain 类型并返回。
func NewConsumeChain[T any](f ConsumeFunc[T]) *ConsumeChain[T] {
	return &ConsumeChain[T]{f}
}

// Consume 消费一个泛型值并返回一个指向 ConsumeChain 本身的指针。
func (c *ConsumeChain[T]) Consume(consumed T) *ConsumeChain[T] {
	c.consumeFunc(consumed)
	return c
}

func NewFirstNonZeroValueConsumeFunc[T any]() ConsumeFunc[SupplyFunc[T]] {
	obtainedNonZeroValue := false
	return func(s SupplyFunc[T]) {
		if !obtainedNonZeroValue {
			v := reflect.ValueOf(s())
			obtainedNonZeroValue = v.IsValid() && !v.IsZero()
		}
	}
}

func NewFirstNonZeroValueConsumeChain[T any]() *ConsumeChain[SupplyFunc[T]] {
	return NewConsumeChain[SupplyFunc[T]](NewFirstNonZeroValueConsumeFunc[T]())
}
