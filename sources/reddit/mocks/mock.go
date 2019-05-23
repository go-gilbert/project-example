// Code generated by MockGen. DO NOT EDIT.
// Source: source.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	geddit "github.com/jzelinskie/geddit"
	reflect "reflect"
)

// RedditMock is a mock of redditReader interface
type RedditMock struct {
	ctrl     *gomock.Controller
	recorder *RedditMockMockRecorder
}

// RedditMockMockRecorder is the mock recorder for RedditMock
type RedditMockMockRecorder struct {
	mock *RedditMock
}

// NewRedditMock creates a new mock instance
func NewRedditMock(ctrl *gomock.Controller) *RedditMock {
	mock := &RedditMock{ctrl: ctrl}
	mock.recorder = &RedditMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *RedditMock) EXPECT() *RedditMockMockRecorder {
	return m.recorder
}

// SubredditSubmissions mocks base method
func (m *RedditMock) SubredditSubmissions(arg0 string, arg1 geddit.PopularitySort, arg2 geddit.ListingOptions) ([]*geddit.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubredditSubmissions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*geddit.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditSubmissions indicates an expected call of SubredditSubmissions
func (mr *RedditMockMockRecorder) SubredditSubmissions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditSubmissions", reflect.TypeOf((*RedditMock)(nil).SubredditSubmissions), arg0, arg1, arg2)
}
