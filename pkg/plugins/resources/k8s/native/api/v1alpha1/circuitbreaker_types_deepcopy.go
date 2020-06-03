package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (cb *CircuitBreaker) DeepCopyInto(out *CircuitBreaker) {
	*out = *cb
	out.TypeMeta = cb.TypeMeta
	cb.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = runtime.DeepCopyJSON(cb.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPermission.
func (cb *CircuitBreaker) DeepCopy() *CircuitBreaker {
	if cb == nil {
		return nil
	}
	out := new(CircuitBreaker)
	cb.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (cb *CircuitBreaker) DeepCopyObject() runtime.Object {
	if c := cb.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (l *CircuitBreakerList) DeepCopyInto(out *CircuitBreakerList) {
	*out = *l
	out.TypeMeta = l.TypeMeta
	out.ListMeta = l.ListMeta
	if l.Items != nil {
		in, out := &l.Items, &out.Items
		*out = make([]CircuitBreaker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPermissionList.
func (l *CircuitBreakerList) DeepCopy() *CircuitBreakerList {
	if l == nil {
		return nil
	}
	out := new(CircuitBreakerList)
	l.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (l *CircuitBreakerList) DeepCopyObject() runtime.Object {
	if c := l.DeepCopy(); c != nil {
		return c
	}
	return nil
}