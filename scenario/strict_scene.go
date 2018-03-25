package scenario

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

func (s *StrictScene) Forward(args ...interface{}) interface{} {
	//receivedCall := &Call{args: args}
	//s.receivedCalls = append(s.receivedCalls, receivedCall)
	//callIndex := len(s.receivedCalls)
	//if callIndex > len(s.expectedCalls) {
	//	s.failures = append(
	//		s.failures,
	//		fmt.Sprintf(
	//			"Unexpected receivedCall %v", receivedCall.String(),
	//		),
	//	)
	//	return
	//}
	//
	//expectedCall := s.expectedCalls[callIndex]
	//callLen := len(receivedCall.args)
	//expectedLen := len(expectedCall.args)
	//biggerLen := int(math.Max(float64(callLen), float64(expectedLen)))
	//if callLen != expectedLen {
	//	s.failures = append(
	//		s.failures,
	//		fmt.Sprintf(
	//			"Arg length does not match (received: %v, expectedCall: %v)",
	//			callLen,
	//			expectedLen,
	//		),
	//	)
	//}
	//
	//for i := 0; i < biggerLen; i++ {
	//	missingReceivedArg, missingExpectedArg := false, false
	//	var receivedArg, expectedArg interface{}
	//	if i < callLen {
	//		receivedArg =receivedCall.args[i]
	//	} else {
	//		missingReceivedArg = true
	//	}
	//
	//	if i < expectedLen {
	//		expectedArg = expectedCall.args[i]
	//	} else {
	//		missingExpectedArg = true
	//	}
	//
	//	areEqual := receivedCall == expectedCall
	//
	//	if missingExpectedArg {
	//		expectedArg = "<Missing Arg>"
	//	}
	//
	//	if missingReceivedArg {
	//		receivedArg = "<Missing Arg>"
	//	}
	//
	//	if missingExpectedArg && missingReceivedArg {
	//		panic("StrictScene::Forward [missingExpectedArg && missingReceivedArg]")
	//	}
	//
	//	if missingExpectedArg || missingReceivedArg || !areEqual {
	//		s.failures =
	//	}
	//}
	return nil
}

func (s *StrictScene) Summarize() []string {
	if len(s.failures) == 0 && len(s.expectedCalls) != len(s.receivedCalls) {
		//i := len()
	}
	return s.failures
}

func NewStrictScene() *StrictScene {
	return &StrictScene{}
}
