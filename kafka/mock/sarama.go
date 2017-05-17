// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/Shopify/sarama (interfaces: Client,Consumer,PartitionConsumer)

package mock

import (
	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Brokers() []*sarama.Broker {
	ret := _m.ctrl.Call(_m, "Brokers")
	ret0, _ := ret[0].([]*sarama.Broker)
	return ret0
}

func (_mr *_MockClientRecorder) Brokers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Brokers")
}

func (_m *MockClient) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockClient) Closed() bool {
	ret := _m.ctrl.Call(_m, "Closed")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockClientRecorder) Closed() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Closed")
}

func (_m *MockClient) Config() *sarama.Config {
	ret := _m.ctrl.Call(_m, "Config")
	ret0, _ := ret[0].(*sarama.Config)
	return ret0
}

func (_mr *_MockClientRecorder) Config() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Config")
}

func (_m *MockClient) Coordinator(_param0 string) (*sarama.Broker, error) {
	ret := _m.ctrl.Call(_m, "Coordinator", _param0)
	ret0, _ := ret[0].(*sarama.Broker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Coordinator(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Coordinator", arg0)
}

func (_m *MockClient) GetOffset(_param0 string, _param1 int32, _param2 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "GetOffset", _param0, _param1, _param2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetOffset(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOffset", arg0, arg1, arg2)
}

func (_m *MockClient) InSyncReplicas(_param0 string, _param1 int32) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "InSyncReplicas", _param0, _param1)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) InSyncReplicas(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InSyncReplicas", arg0, arg1)
}

func (_m *MockClient) Leader(_param0 string, _param1 int32) (*sarama.Broker, error) {
	ret := _m.ctrl.Call(_m, "Leader", _param0, _param1)
	ret0, _ := ret[0].(*sarama.Broker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Leader(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Leader", arg0, arg1)
}

func (_m *MockClient) Partitions(_param0 string) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "Partitions", _param0)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Partitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Partitions", arg0)
}

func (_m *MockClient) RefreshCoordinator(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RefreshCoordinator", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) RefreshCoordinator(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RefreshCoordinator", arg0)
}

func (_m *MockClient) RefreshMetadata(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RefreshMetadata", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) RefreshMetadata(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RefreshMetadata", arg0...)
}

func (_m *MockClient) Replicas(_param0 string, _param1 int32) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "Replicas", _param0, _param1)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Replicas(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Replicas", arg0, arg1)
}

func (_m *MockClient) Topics() ([]string, error) {
	ret := _m.ctrl.Call(_m, "Topics")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Topics() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Topics")
}

func (_m *MockClient) WritablePartitions(_param0 string) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "WritablePartitions", _param0)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) WritablePartitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WritablePartitions", arg0)
}

// Mock of Consumer interface
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *_MockConsumerRecorder
}

// Recorder for MockConsumer (not exported)
type _MockConsumerRecorder struct {
	mock *MockConsumer
}

func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &_MockConsumerRecorder{mock}
	return mock
}

func (_m *MockConsumer) EXPECT() *_MockConsumerRecorder {
	return _m.recorder
}

func (_m *MockConsumer) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConsumerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockConsumer) ConsumePartition(_param0 string, _param1 int32, _param2 int64) (sarama.PartitionConsumer, error) {
	ret := _m.ctrl.Call(_m, "ConsumePartition", _param0, _param1, _param2)
	ret0, _ := ret[0].(sarama.PartitionConsumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConsumerRecorder) ConsumePartition(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConsumePartition", arg0, arg1, arg2)
}

func (_m *MockConsumer) HighWaterMarks() map[string]map[int32]int64 {
	ret := _m.ctrl.Call(_m, "HighWaterMarks")
	ret0, _ := ret[0].(map[string]map[int32]int64)
	return ret0
}

func (_mr *_MockConsumerRecorder) HighWaterMarks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HighWaterMarks")
}

func (_m *MockConsumer) Partitions(_param0 string) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "Partitions", _param0)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConsumerRecorder) Partitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Partitions", arg0)
}

func (_m *MockConsumer) Topics() ([]string, error) {
	ret := _m.ctrl.Call(_m, "Topics")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConsumerRecorder) Topics() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Topics")
}

// Mock of PartitionConsumer interface
type MockPartitionConsumer struct {
	ctrl     *gomock.Controller
	recorder *_MockPartitionConsumerRecorder
}

// Recorder for MockPartitionConsumer (not exported)
type _MockPartitionConsumerRecorder struct {
	mock *MockPartitionConsumer
}

func NewMockPartitionConsumer(ctrl *gomock.Controller) *MockPartitionConsumer {
	mock := &MockPartitionConsumer{ctrl: ctrl}
	mock.recorder = &_MockPartitionConsumerRecorder{mock}
	return mock
}

func (_m *MockPartitionConsumer) EXPECT() *_MockPartitionConsumerRecorder {
	return _m.recorder
}

func (_m *MockPartitionConsumer) AsyncClose() {
	_m.ctrl.Call(_m, "AsyncClose")
}

func (_mr *_MockPartitionConsumerRecorder) AsyncClose() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AsyncClose")
}

func (_m *MockPartitionConsumer) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPartitionConsumerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockPartitionConsumer) Errors() <-chan *sarama.ConsumerError {
	ret := _m.ctrl.Call(_m, "Errors")
	ret0, _ := ret[0].(<-chan *sarama.ConsumerError)
	return ret0
}

func (_mr *_MockPartitionConsumerRecorder) Errors() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Errors")
}

func (_m *MockPartitionConsumer) HighWaterMarkOffset() int64 {
	ret := _m.ctrl.Call(_m, "HighWaterMarkOffset")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockPartitionConsumerRecorder) HighWaterMarkOffset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HighWaterMarkOffset")
}

func (_m *MockPartitionConsumer) Messages() <-chan *sarama.ConsumerMessage {
	ret := _m.ctrl.Call(_m, "Messages")
	ret0, _ := ret[0].(<-chan *sarama.ConsumerMessage)
	return ret0
}

func (_mr *_MockPartitionConsumerRecorder) Messages() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Messages")
}