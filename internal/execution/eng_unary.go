// Code generated by genlib2. DO NOT EDIT.

package execution

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/maxsei/tensor/internal/storage"
)

func (e E) Neg(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		NegI(a.Ints())
		return nil
	case Int8:
		NegI8(a.Int8s())
		return nil
	case Int16:
		NegI16(a.Int16s())
		return nil
	case Int32:
		NegI32(a.Int32s())
		return nil
	case Int64:
		NegI64(a.Int64s())
		return nil
	case Uint:
		NegU(a.Uints())
		return nil
	case Uint8:
		NegU8(a.Uint8s())
		return nil
	case Uint16:
		NegU16(a.Uint16s())
		return nil
	case Uint32:
		NegU32(a.Uint32s())
		return nil
	case Uint64:
		NegU64(a.Uint64s())
		return nil
	case Float32:
		NegF32(a.Float32s())
		return nil
	case Float64:
		NegF64(a.Float64s())
		return nil
	case Complex64:
		NegC64(a.Complex64s())
		return nil
	case Complex128:
		NegC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Neg", t)
	}
}

func (e E) Inv(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		InvI(a.Ints())
		return nil
	case Int8:
		InvI8(a.Int8s())
		return nil
	case Int16:
		InvI16(a.Int16s())
		return nil
	case Int32:
		InvI32(a.Int32s())
		return nil
	case Int64:
		InvI64(a.Int64s())
		return nil
	case Uint:
		InvU(a.Uints())
		return nil
	case Uint8:
		InvU8(a.Uint8s())
		return nil
	case Uint16:
		InvU16(a.Uint16s())
		return nil
	case Uint32:
		InvU32(a.Uint32s())
		return nil
	case Uint64:
		InvU64(a.Uint64s())
		return nil
	case Float32:
		InvF32(a.Float32s())
		return nil
	case Float64:
		InvF64(a.Float64s())
		return nil
	case Complex64:
		InvC64(a.Complex64s())
		return nil
	case Complex128:
		InvC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Inv", t)
	}
}

func (e E) Square(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		SquareI(a.Ints())
		return nil
	case Int8:
		SquareI8(a.Int8s())
		return nil
	case Int16:
		SquareI16(a.Int16s())
		return nil
	case Int32:
		SquareI32(a.Int32s())
		return nil
	case Int64:
		SquareI64(a.Int64s())
		return nil
	case Uint:
		SquareU(a.Uints())
		return nil
	case Uint8:
		SquareU8(a.Uint8s())
		return nil
	case Uint16:
		SquareU16(a.Uint16s())
		return nil
	case Uint32:
		SquareU32(a.Uint32s())
		return nil
	case Uint64:
		SquareU64(a.Uint64s())
		return nil
	case Float32:
		SquareF32(a.Float32s())
		return nil
	case Float64:
		SquareF64(a.Float64s())
		return nil
	case Complex64:
		SquareC64(a.Complex64s())
		return nil
	case Complex128:
		SquareC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Square", t)
	}
}

func (e E) Cube(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		CubeI(a.Ints())
		return nil
	case Int8:
		CubeI8(a.Int8s())
		return nil
	case Int16:
		CubeI16(a.Int16s())
		return nil
	case Int32:
		CubeI32(a.Int32s())
		return nil
	case Int64:
		CubeI64(a.Int64s())
		return nil
	case Uint:
		CubeU(a.Uints())
		return nil
	case Uint8:
		CubeU8(a.Uint8s())
		return nil
	case Uint16:
		CubeU16(a.Uint16s())
		return nil
	case Uint32:
		CubeU32(a.Uint32s())
		return nil
	case Uint64:
		CubeU64(a.Uint64s())
		return nil
	case Float32:
		CubeF32(a.Float32s())
		return nil
	case Float64:
		CubeF64(a.Float64s())
		return nil
	case Complex64:
		CubeC64(a.Complex64s())
		return nil
	case Complex128:
		CubeC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Cube", t)
	}
}

func (e E) Exp(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		ExpF32(a.Float32s())
		return nil
	case Float64:
		ExpF64(a.Float64s())
		return nil
	case Complex64:
		ExpC64(a.Complex64s())
		return nil
	case Complex128:
		ExpC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Exp", t)
	}
}

func (e E) Tanh(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		TanhF32(a.Float32s())
		return nil
	case Float64:
		TanhF64(a.Float64s())
		return nil
	case Complex64:
		TanhC64(a.Complex64s())
		return nil
	case Complex128:
		TanhC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Tanh", t)
	}
}

func (e E) Log(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		LogF32(a.Float32s())
		return nil
	case Float64:
		LogF64(a.Float64s())
		return nil
	case Complex64:
		LogC64(a.Complex64s())
		return nil
	case Complex128:
		LogC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Log", t)
	}
}

func (e E) Log2(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		Log2F32(a.Float32s())
		return nil
	case Float64:
		Log2F64(a.Float64s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Log2", t)
	}
}

func (e E) Log10(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		Log10F32(a.Float32s())
		return nil
	case Float64:
		Log10F64(a.Float64s())
		return nil
	case Complex64:
		Log10C64(a.Complex64s())
		return nil
	case Complex128:
		Log10C128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Log10", t)
	}
}

func (e E) Sqrt(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		SqrtF32(a.Float32s())
		return nil
	case Float64:
		SqrtF64(a.Float64s())
		return nil
	case Complex64:
		SqrtC64(a.Complex64s())
		return nil
	case Complex128:
		SqrtC128(a.Complex128s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Sqrt", t)
	}
}

func (e E) Cbrt(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		CbrtF32(a.Float32s())
		return nil
	case Float64:
		CbrtF64(a.Float64s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Cbrt", t)
	}
}

func (e E) InvSqrt(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Float32:
		InvSqrtF32(a.Float32s())
		return nil
	case Float64:
		InvSqrtF64(a.Float64s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for InvSqrt", t)
	}
}

func (e E) NegIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return NegIterI(a.Ints(), ait)
	case Int8:
		return NegIterI8(a.Int8s(), ait)
	case Int16:
		return NegIterI16(a.Int16s(), ait)
	case Int32:
		return NegIterI32(a.Int32s(), ait)
	case Int64:
		return NegIterI64(a.Int64s(), ait)
	case Uint:
		return NegIterU(a.Uints(), ait)
	case Uint8:
		return NegIterU8(a.Uint8s(), ait)
	case Uint16:
		return NegIterU16(a.Uint16s(), ait)
	case Uint32:
		return NegIterU32(a.Uint32s(), ait)
	case Uint64:
		return NegIterU64(a.Uint64s(), ait)
	case Float32:
		return NegIterF32(a.Float32s(), ait)
	case Float64:
		return NegIterF64(a.Float64s(), ait)
	case Complex64:
		return NegIterC64(a.Complex64s(), ait)
	case Complex128:
		return NegIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for NegIter", t)
	}
}

func (e E) InvIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return InvIterI(a.Ints(), ait)
	case Int8:
		return InvIterI8(a.Int8s(), ait)
	case Int16:
		return InvIterI16(a.Int16s(), ait)
	case Int32:
		return InvIterI32(a.Int32s(), ait)
	case Int64:
		return InvIterI64(a.Int64s(), ait)
	case Uint:
		return InvIterU(a.Uints(), ait)
	case Uint8:
		return InvIterU8(a.Uint8s(), ait)
	case Uint16:
		return InvIterU16(a.Uint16s(), ait)
	case Uint32:
		return InvIterU32(a.Uint32s(), ait)
	case Uint64:
		return InvIterU64(a.Uint64s(), ait)
	case Float32:
		return InvIterF32(a.Float32s(), ait)
	case Float64:
		return InvIterF64(a.Float64s(), ait)
	case Complex64:
		return InvIterC64(a.Complex64s(), ait)
	case Complex128:
		return InvIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for InvIter", t)
	}
}

func (e E) SquareIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return SquareIterI(a.Ints(), ait)
	case Int8:
		return SquareIterI8(a.Int8s(), ait)
	case Int16:
		return SquareIterI16(a.Int16s(), ait)
	case Int32:
		return SquareIterI32(a.Int32s(), ait)
	case Int64:
		return SquareIterI64(a.Int64s(), ait)
	case Uint:
		return SquareIterU(a.Uints(), ait)
	case Uint8:
		return SquareIterU8(a.Uint8s(), ait)
	case Uint16:
		return SquareIterU16(a.Uint16s(), ait)
	case Uint32:
		return SquareIterU32(a.Uint32s(), ait)
	case Uint64:
		return SquareIterU64(a.Uint64s(), ait)
	case Float32:
		return SquareIterF32(a.Float32s(), ait)
	case Float64:
		return SquareIterF64(a.Float64s(), ait)
	case Complex64:
		return SquareIterC64(a.Complex64s(), ait)
	case Complex128:
		return SquareIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for SquareIter", t)
	}
}

func (e E) CubeIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return CubeIterI(a.Ints(), ait)
	case Int8:
		return CubeIterI8(a.Int8s(), ait)
	case Int16:
		return CubeIterI16(a.Int16s(), ait)
	case Int32:
		return CubeIterI32(a.Int32s(), ait)
	case Int64:
		return CubeIterI64(a.Int64s(), ait)
	case Uint:
		return CubeIterU(a.Uints(), ait)
	case Uint8:
		return CubeIterU8(a.Uint8s(), ait)
	case Uint16:
		return CubeIterU16(a.Uint16s(), ait)
	case Uint32:
		return CubeIterU32(a.Uint32s(), ait)
	case Uint64:
		return CubeIterU64(a.Uint64s(), ait)
	case Float32:
		return CubeIterF32(a.Float32s(), ait)
	case Float64:
		return CubeIterF64(a.Float64s(), ait)
	case Complex64:
		return CubeIterC64(a.Complex64s(), ait)
	case Complex128:
		return CubeIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for CubeIter", t)
	}
}

func (e E) ExpIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return ExpIterF32(a.Float32s(), ait)
	case Float64:
		return ExpIterF64(a.Float64s(), ait)
	case Complex64:
		return ExpIterC64(a.Complex64s(), ait)
	case Complex128:
		return ExpIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for ExpIter", t)
	}
}

func (e E) TanhIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return TanhIterF32(a.Float32s(), ait)
	case Float64:
		return TanhIterF64(a.Float64s(), ait)
	case Complex64:
		return TanhIterC64(a.Complex64s(), ait)
	case Complex128:
		return TanhIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for TanhIter", t)
	}
}

func (e E) LogIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return LogIterF32(a.Float32s(), ait)
	case Float64:
		return LogIterF64(a.Float64s(), ait)
	case Complex64:
		return LogIterC64(a.Complex64s(), ait)
	case Complex128:
		return LogIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for LogIter", t)
	}
}

func (e E) Log2Iter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return Log2IterF32(a.Float32s(), ait)
	case Float64:
		return Log2IterF64(a.Float64s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for Log2Iter", t)
	}
}

func (e E) Log10Iter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return Log10IterF32(a.Float32s(), ait)
	case Float64:
		return Log10IterF64(a.Float64s(), ait)
	case Complex64:
		return Log10IterC64(a.Complex64s(), ait)
	case Complex128:
		return Log10IterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for Log10Iter", t)
	}
}

func (e E) SqrtIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return SqrtIterF32(a.Float32s(), ait)
	case Float64:
		return SqrtIterF64(a.Float64s(), ait)
	case Complex64:
		return SqrtIterC64(a.Complex64s(), ait)
	case Complex128:
		return SqrtIterC128(a.Complex128s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for SqrtIter", t)
	}
}

func (e E) CbrtIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return CbrtIterF32(a.Float32s(), ait)
	case Float64:
		return CbrtIterF64(a.Float64s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for CbrtIter", t)
	}
}

func (e E) InvSqrtIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Float32:
		return InvSqrtIterF32(a.Float32s(), ait)
	case Float64:
		return InvSqrtIterF64(a.Float64s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for InvSqrtIter", t)
	}
}

func (e E) Abs(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		AbsI(a.Ints())
		return nil
	case Int8:
		AbsI8(a.Int8s())
		return nil
	case Int16:
		AbsI16(a.Int16s())
		return nil
	case Int32:
		AbsI32(a.Int32s())
		return nil
	case Int64:
		AbsI64(a.Int64s())
		return nil
	case Float32:
		AbsF32(a.Float32s())
		return nil
	case Float64:
		AbsF64(a.Float64s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Abs", t)
	}
}

func (e E) Sign(t reflect.Type, a *storage.Header) (err error) {
	switch t {
	case Int:
		SignI(a.Ints())
		return nil
	case Int8:
		SignI8(a.Int8s())
		return nil
	case Int16:
		SignI16(a.Int16s())
		return nil
	case Int32:
		SignI32(a.Int32s())
		return nil
	case Int64:
		SignI64(a.Int64s())
		return nil
	case Float32:
		SignF32(a.Float32s())
		return nil
	case Float64:
		SignF64(a.Float64s())
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Sign", t)
	}
}

func (e E) AbsIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return AbsIterI(a.Ints(), ait)
	case Int8:
		return AbsIterI8(a.Int8s(), ait)
	case Int16:
		return AbsIterI16(a.Int16s(), ait)
	case Int32:
		return AbsIterI32(a.Int32s(), ait)
	case Int64:
		return AbsIterI64(a.Int64s(), ait)
	case Float32:
		return AbsIterF32(a.Float32s(), ait)
	case Float64:
		return AbsIterF64(a.Float64s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for AbsIter", t)
	}
}

func (e E) SignIter(t reflect.Type, a *storage.Header, ait Iterator) (err error) {
	switch t {
	case Int:
		return SignIterI(a.Ints(), ait)
	case Int8:
		return SignIterI8(a.Int8s(), ait)
	case Int16:
		return SignIterI16(a.Int16s(), ait)
	case Int32:
		return SignIterI32(a.Int32s(), ait)
	case Int64:
		return SignIterI64(a.Int64s(), ait)
	case Float32:
		return SignIterF32(a.Float32s(), ait)
	case Float64:
		return SignIterF64(a.Float64s(), ait)
	default:
		return errors.Errorf("Unsupported type %v for SignIter", t)
	}
}

func (e E) Clamp(t reflect.Type, a *storage.Header, minVal interface{}, maxVal interface{}) (err error) {
	var ok bool
	switch t {
	case Int:
		var min, max int
		if min, ok = minVal.(int); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampI(a.Ints(), min, max)
		return nil
	case Int8:
		var min, max int8
		if min, ok = minVal.(int8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampI8(a.Int8s(), min, max)
		return nil
	case Int16:
		var min, max int16
		if min, ok = minVal.(int16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampI16(a.Int16s(), min, max)
		return nil
	case Int32:
		var min, max int32
		if min, ok = minVal.(int32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampI32(a.Int32s(), min, max)
		return nil
	case Int64:
		var min, max int64
		if min, ok = minVal.(int64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampI64(a.Int64s(), min, max)
		return nil
	case Uint:
		var min, max uint
		if min, ok = minVal.(uint); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampU(a.Uints(), min, max)
		return nil
	case Uint8:
		var min, max uint8
		if min, ok = minVal.(uint8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampU8(a.Uint8s(), min, max)
		return nil
	case Uint16:
		var min, max uint16
		if min, ok = minVal.(uint16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampU16(a.Uint16s(), min, max)
		return nil
	case Uint32:
		var min, max uint32
		if min, ok = minVal.(uint32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampU32(a.Uint32s(), min, max)
		return nil
	case Uint64:
		var min, max uint64
		if min, ok = minVal.(uint64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampU64(a.Uint64s(), min, max)
		return nil
	case Float32:
		var min, max float32
		if min, ok = minVal.(float32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(float32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampF32(a.Float32s(), min, max)
		return nil
	case Float64:
		var min, max float64
		if min, ok = minVal.(float64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(float64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		ClampF64(a.Float64s(), min, max)
		return nil
	default:
		return errors.Errorf("Unsupported type %v for Clamp", t)
	}
}

func (e E) ClampIter(t reflect.Type, a *storage.Header, ait Iterator, minVal interface{}, maxVal interface{}) (err error) {
	var ok bool
	switch t {
	case Int:
		var min, max int
		if min, ok = minVal.(int); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterI(a.Ints(), ait, min, max)
	case Int8:
		var min, max int8
		if min, ok = minVal.(int8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterI8(a.Int8s(), ait, min, max)
	case Int16:
		var min, max int16
		if min, ok = minVal.(int16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterI16(a.Int16s(), ait, min, max)
	case Int32:
		var min, max int32
		if min, ok = minVal.(int32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterI32(a.Int32s(), ait, min, max)
	case Int64:
		var min, max int64
		if min, ok = minVal.(int64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(int64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterI64(a.Int64s(), ait, min, max)
	case Uint:
		var min, max uint
		if min, ok = minVal.(uint); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterU(a.Uints(), ait, min, max)
	case Uint8:
		var min, max uint8
		if min, ok = minVal.(uint8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint8); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterU8(a.Uint8s(), ait, min, max)
	case Uint16:
		var min, max uint16
		if min, ok = minVal.(uint16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint16); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterU16(a.Uint16s(), ait, min, max)
	case Uint32:
		var min, max uint32
		if min, ok = minVal.(uint32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterU32(a.Uint32s(), ait, min, max)
	case Uint64:
		var min, max uint64
		if min, ok = minVal.(uint64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(uint64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterU64(a.Uint64s(), ait, min, max)
	case Float32:
		var min, max float32
		if min, ok = minVal.(float32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(float32); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterF32(a.Float32s(), ait, min, max)
	case Float64:
		var min, max float64
		if min, ok = minVal.(float64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, min, minVal), "Clamp() min")
		}
		if max, ok = maxVal.(float64); !ok {
			return errors.Wrap(errors.Errorf(typeMismatch, max, maxVal), "Clamp() max")
		}
		return ClampIterF64(a.Float64s(), ait, min, max)
	default:
		return errors.Errorf("Unsupported type %v for Clamp", t)
	}
}
