package scenario

import "fmt"

type StrictScene struct {
	expectedCalls []*Call
	receivedCalls []*Call
	failures      []string
}

func (s *StrictScene) Call(args ...interface{}) (c *Call) {
	c = &Call{args: args}
	s.expectedCalls = append(s.expectedCalls, c)
	return
}

func (s *StrictScene) Forward(args ...interface{}) (result interface{}) {
	receivedCall := &Call{args: args}
	s.receivedCalls = append(s.receivedCalls, receivedCall)
	if len(s.receivedCalls) > len(s.expectedCalls) {
		s.fail("Unexpected: %v", receivedCall)
		return
	}

	expectedCall := s.expectedCalls[len(s.receivedCalls)-1]
	if !receivedCall.InputEqual(expectedCall) {
		leftArity, _ := expectedCall.Arity()
		rightArity, _ := receivedCall.Arity()
		s.fail(
			"Calls differ; Expected[%v-ary]: %v; Received[%v-ary]: %v",
			leftArity,
			expectedCall,
			rightArity,
			receivedCall,
		)
		return
	}
	return nil
}

func (s *StrictScene) Summarize() []string {
	if len(s.failures) == 0 && len(s.expectedCalls) != len(s.receivedCalls) {
		//i := len()
	}
	return s.failures
}

func (s *StrictScene) fail(format string, args ...interface{}) {
	s.failures = append(s.failures, fmt.Sprintf(format, args...))
}

func NewStrictScene() *StrictScene {
	return &StrictScene{}
}
