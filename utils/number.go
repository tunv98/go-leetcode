package utils

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func Contains[T comparable](arr []T, val T) bool {
	for _, element := range arr {
		if val == element {
			return true
		}
	}
	return false
}
