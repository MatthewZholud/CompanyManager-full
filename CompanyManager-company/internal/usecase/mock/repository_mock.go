package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase.Repository -o ./repository_mock_test.go

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements Repository
type RepositoryMock struct {
	t minimock.Tester

	funcCreate          func(e *company.Company) (s1 string, err error)
	inspectFuncCreate   func(e *company.Company)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mRepositoryMockCreate

	funcDelete          func(id int64) (s1 string, err error)
	inspectFuncDelete   func(id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mRepositoryMockDelete

	funcGet          func(id int64) (cp1 *company.Company, err error)
	inspectFuncGet   func(id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mRepositoryMockGet

	funcUpdate          func(e *company.Company) (s1 string, err error)
	inspectFuncUpdate   func(e *company.Company)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mRepositoryMockUpdate
}

// NewRepositoryMock returns a mock for Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*RepositoryMockCreateParams{}

	m.DeleteMock = mRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*RepositoryMockDeleteParams{}

	m.GetMock = mRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*RepositoryMockGetParams{}

	m.UpdateMock = mRepositoryMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*RepositoryMockUpdateParams{}

	return m
}

type mRepositoryMockCreate struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockCreateExpectation
	expectations       []*RepositoryMockCreateExpectation

	callArgs []*RepositoryMockCreateParams
	mutex    sync.RWMutex
}

// RepositoryMockCreateExpectation specifies expectation struct of the Repository.Create
type RepositoryMockCreateExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockCreateParams
	results *RepositoryMockCreateResults
	Counter uint64
}

// RepositoryMockCreateParams contains parameters of the Repository.Create
type RepositoryMockCreateParams struct {
	e *company.Company
}

// RepositoryMockCreateResults contains results of the Repository.Create
type RepositoryMockCreateResults struct {
	s1  string
	err error
}

// Expect sets up expected params for Repository.Create
func (mmCreate *mRepositoryMockCreate) Expect(e *company.Company) *mRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &RepositoryMockCreateParams{e}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the Repository.Create
func (mmCreate *mRepositoryMockCreate) Inspect(f func(e *company.Company)) *mRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by Repository.Create
func (mmCreate *mRepositoryMockCreate) Return(s1 string, err error) *RepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &RepositoryMockCreateResults{s1, err}
	return mmCreate.mock
}

//Set uses given function f to mock the Repository.Create method
func (mmCreate *mRepositoryMockCreate) Set(f func(e *company.Company) (s1 string, err error)) *RepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the Repository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the Repository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the Repository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mRepositoryMockCreate) When(e *company.Company) *RepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	expectation := &RepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &RepositoryMockCreateParams{e},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up Repository.Create return parameters for the expectation previously defined by the When method
func (e *RepositoryMockCreateExpectation) Then(s1 string, err error) *RepositoryMock {
	e.results = &RepositoryMockCreateResults{s1, err}
	return e.mock
}

// Create implements Repository
func (mmCreate *RepositoryMock) Create(e *company.Company) (s1 string, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(e)
	}

	mm_params := &RepositoryMockCreateParams{e}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := RepositoryMockCreateParams{e}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("RepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the RepositoryMock.Create")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(e)
	}
	mmCreate.t.Fatalf("Unexpected call to RepositoryMock.Create. %v", e)
	return
}

// CreateAfterCounter returns a count of finished RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mRepositoryMockCreate) Calls() []*RepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*RepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *RepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Create")
	}
}

type mRepositoryMockDelete struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockDeleteExpectation
	expectations       []*RepositoryMockDeleteExpectation

	callArgs []*RepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// RepositoryMockDeleteExpectation specifies expectation struct of the Repository.Delete
type RepositoryMockDeleteExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockDeleteParams
	results *RepositoryMockDeleteResults
	Counter uint64
}

// RepositoryMockDeleteParams contains parameters of the Repository.Delete
type RepositoryMockDeleteParams struct {
	id int64
}

// RepositoryMockDeleteResults contains results of the Repository.Delete
type RepositoryMockDeleteResults struct {
	s1  string
	err error
}

// Expect sets up expected params for Repository.Delete
func (mmDelete *mRepositoryMockDelete) Expect(id int64) *mRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &RepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &RepositoryMockDeleteParams{id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Repository.Delete
func (mmDelete *mRepositoryMockDelete) Inspect(f func(id int64)) *mRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Repository.Delete
func (mmDelete *mRepositoryMockDelete) Return(s1 string, err error) *RepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &RepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &RepositoryMockDeleteResults{s1, err}
	return mmDelete.mock
}

//Set uses given function f to mock the Repository.Delete method
func (mmDelete *mRepositoryMockDelete) Set(f func(id int64) (s1 string, err error)) *RepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Repository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Repository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Repository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mRepositoryMockDelete) When(id int64) *RepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepositoryMock.Delete mock is already set by Set")
	}

	expectation := &RepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &RepositoryMockDeleteParams{id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Repository.Delete return parameters for the expectation previously defined by the When method
func (e *RepositoryMockDeleteExpectation) Then(s1 string, err error) *RepositoryMock {
	e.results = &RepositoryMockDeleteResults{s1, err}
	return e.mock
}

// Delete implements Repository
func (mmDelete *RepositoryMock) Delete(id int64) (s1 string, err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(id)
	}

	mm_params := &RepositoryMockDeleteParams{id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := RepositoryMockDeleteParams{id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("RepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the RepositoryMock.Delete")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(id)
	}
	mmDelete.t.Fatalf("Unexpected call to RepositoryMock.Delete. %v", id)
	return
}

// DeleteAfterCounter returns a count of finished RepositoryMock.Delete invocations
func (mmDelete *RepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of RepositoryMock.Delete invocations
func (mmDelete *RepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mRepositoryMockDelete) Calls() []*RepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*RepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *RepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Delete")
	}
}

type mRepositoryMockGet struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetExpectation
	expectations       []*RepositoryMockGetExpectation

	callArgs []*RepositoryMockGetParams
	mutex    sync.RWMutex
}

// RepositoryMockGetExpectation specifies expectation struct of the Repository.Get
type RepositoryMockGetExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetParams
	results *RepositoryMockGetResults
	Counter uint64
}

// RepositoryMockGetParams contains parameters of the Repository.Get
type RepositoryMockGetParams struct {
	id int64
}

// RepositoryMockGetResults contains results of the Repository.Get
type RepositoryMockGetResults struct {
	cp1 *company.Company
	err error
}

// Expect sets up expected params for Repository.Get
func (mmGet *mRepositoryMockGet) Expect(id int64) *mRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &RepositoryMockGetParams{id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Repository.Get
func (mmGet *mRepositoryMockGet) Inspect(f func(id int64)) *mRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Repository.Get
func (mmGet *mRepositoryMockGet) Return(cp1 *company.Company, err error) *RepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &RepositoryMockGetResults{cp1, err}
	return mmGet.mock
}

//Set uses given function f to mock the Repository.Get method
func (mmGet *mRepositoryMockGet) Set(f func(id int64) (cp1 *company.Company, err error)) *RepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Repository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Repository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Repository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mRepositoryMockGet) When(id int64) *RepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RepositoryMock.Get mock is already set by Set")
	}

	expectation := &RepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &RepositoryMockGetParams{id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Repository.Get return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetExpectation) Then(cp1 *company.Company, err error) *RepositoryMock {
	e.results = &RepositoryMockGetResults{cp1, err}
	return e.mock
}

// Get implements Repository
func (mmGet *RepositoryMock) Get(id int64) (cp1 *company.Company, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(id)
	}

	mm_params := &RepositoryMockGetParams{id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := RepositoryMockGetParams{id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("RepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the RepositoryMock.Get")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(id)
	}
	mmGet.t.Fatalf("Unexpected call to RepositoryMock.Get. %v", id)
	return
}

// GetAfterCounter returns a count of finished RepositoryMock.Get invocations
func (mmGet *RepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of RepositoryMock.Get invocations
func (mmGet *RepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mRepositoryMockGet) Calls() []*RepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*RepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Get")
	}
}

type mRepositoryMockUpdate struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockUpdateExpectation
	expectations       []*RepositoryMockUpdateExpectation

	callArgs []*RepositoryMockUpdateParams
	mutex    sync.RWMutex
}

// RepositoryMockUpdateExpectation specifies expectation struct of the Repository.Update
type RepositoryMockUpdateExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockUpdateParams
	results *RepositoryMockUpdateResults
	Counter uint64
}

// RepositoryMockUpdateParams contains parameters of the Repository.Update
type RepositoryMockUpdateParams struct {
	e *company.Company
}

// RepositoryMockUpdateResults contains results of the Repository.Update
type RepositoryMockUpdateResults struct {
	s1  string
	err error
}

// Expect sets up expected params for Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Expect(e *company.Company) *mRepositoryMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepositoryMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &RepositoryMockUpdateParams{e}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Inspect(f func(e *company.Company)) *mRepositoryMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by Repository.Update
func (mmUpdate *mRepositoryMockUpdate) Return(s1 string, err error) *RepositoryMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepositoryMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &RepositoryMockUpdateResults{s1, err}
	return mmUpdate.mock
}

//Set uses given function f to mock the Repository.Update method
func (mmUpdate *mRepositoryMockUpdate) Set(f func(e *company.Company) (s1 string, err error)) *RepositoryMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the Repository.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the Repository.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the Repository.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mRepositoryMockUpdate) When(e *company.Company) *RepositoryMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepositoryMock.Update mock is already set by Set")
	}

	expectation := &RepositoryMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &RepositoryMockUpdateParams{e},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up Repository.Update return parameters for the expectation previously defined by the When method
func (e *RepositoryMockUpdateExpectation) Then(s1 string, err error) *RepositoryMock {
	e.results = &RepositoryMockUpdateResults{s1, err}
	return e.mock
}

// Update implements Repository
func (mmUpdate *RepositoryMock) Update(e *company.Company) (s1 string, err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(e)
	}

	mm_params := &RepositoryMockUpdateParams{e}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := RepositoryMockUpdateParams{e}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("RepositoryMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the RepositoryMock.Update")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(e)
	}
	mmUpdate.t.Fatalf("Unexpected call to RepositoryMock.Update. %v", e)
	return
}

// UpdateAfterCounter returns a count of finished RepositoryMock.Update invocations
func (mmUpdate *RepositoryMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of RepositoryMock.Update invocations
func (mmUpdate *RepositoryMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mRepositoryMockUpdate) Calls() []*RepositoryMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*RepositoryMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *RepositoryMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Update")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockUpdateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockUpdateDone()
}