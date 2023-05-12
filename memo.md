## hint1: un-exported fields

``` go
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}
```

## hint2: asserting method

``` go
> Equal asserts that two objects are equal.
//
>    assert.Equal(t, 123, 123)
//
> Pointer variable equality is determined based on the equality of the
> referenced values (as opposed to the memory addresses). Function equality
> cannot be determined and will always fail.
func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if !ObjectsAreEqual(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true
}
```

in the comparing method `ObjectsAreEqual`, it uses `reflect.DeepEqual`.

``` go
> ObjectsAreEqual determines if two objects are considered equal.
//
> This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
```

## reflect.DeepEqual

> > DeepEqual reports whether x and y are “deeply equal,” defined as follows.
> Two values of identical type are deeply equal if one of the following cases applies.
> Values of distinct types are never deeply equal.
>
> Array values are deeply equal when their corresponding elements are deeply equal.
>
> Struct values are deeply equal if their corresponding fields,
> both exported and unexported, are deeply equal.
>
> Func values are deeply equal if both are nil; otherwise they are not deeply equal.
>
> Interface values are deeply equal if they hold deeply equal concrete values.
>
> Map values are deeply equal when all of the following are true:
> they are both nil or both non-nil, they have the same length,
> and either they are the same map object or their corresponding keys
> (matched using Go equality) map to deeply equal values.
>
> Pointer values are deeply equal if they are equal using Go's == operator
> or if they point to deeply equal values.
>
> Slice values are deeply equal when all of the following are true:
> they are both nil or both non-nil, they have the same length,
> and either they point to the same initial entry of the same underlying array
> (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
> Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
> are not deeply equal.
>
> Other values - numbers, bools, strings, and channels - are deeply equal
> if they are equal using Go's == operator.
>
> In general DeepEqual is a recursive relaxation of Go's == operator.
> However, this idea is impossible to implement without some inconsistency.
> Specifically, it is possible for a value to be unequal to itself,
> either because it is of func type (uncomparable in general)
> or because it is a floating-point NaN value (not equal to itself in floating-point comparison),
> or because it is an array, struct, or interface containing
> such a value.
> On the other hand, pointer values are always equal to themselves,
> even if they point at or contain such problematic values,
> because they compare equal using Go's == operator, and that
> is a sufficient condition to be deeply equal, regardless of content.
> DeepEqual has been defined so that the same short-cut applies
> to slices and maps: if x and y are the same slice or the same map,
> they are deeply equal regardless of content.
>
> As DeepEqual traverses the data values it may find a cycle. The
> second and subsequent times that DeepEqual compares two pointer
> values that have been compared before, it treats the values as
> equal rather than examining the values to which they point.
> This ensures that DeepEqual terminates.

'both exported and unexported, are deeply equal for struct.

## MessageState in proto

``` go
// MessageState is a data structure that is nested as the first field in a
// concrete message. It provides a way to implement the ProtoReflect method
// in an allocation-free way without needing to have a shadow Go type generated
// for every message type. This technique only works using unsafe.
//
// Example generated code:
//
//	type M struct {
//		state protoimpl.MessageState
//
//		Field1 int32
//		Field2 string
//		Field3 *BarMessage
//		...
//	}
//
//	func (m *M) ProtoReflect() protoreflect.Message {
//		mi := &file_fizz_buzz_proto_msgInfos[5]
//		if protoimpl.UnsafeEnabled && m != nil {
//			ms := protoimpl.X.MessageStateOf(Pointer(m))
//			if ms.LoadMessageInfo() == nil {
//				ms.StoreMessageInfo(mi)
//			}
//			return ms
//		}
//		return mi.MessageOf(m)
//	}
//
// The MessageState type holds a *MessageInfo, which must be atomically set to
// the message info associated with a given message instance.
// By unsafely converting a *M into a *MessageState, the MessageState object
// has access to all the information needed to implement protobuf reflection.
// It has access to the message info as its first field, and a pointer to the
// MessageState is identical to a pointer to the concrete message value.
//
// Requirements:
//   - The type M must implement protoreflect.ProtoMessage.
//   - The address of m must not be nil.
//   - The address of m and the address of m.state must be equal,
//     even though they are different Go types.
type MessageState struct {
	pragma.NoUnkeyedLiterals
	pragma.DoNotCompare
	pragma.DoNotCopy

	atomicMessageInfo *MessageInfo
}
```
