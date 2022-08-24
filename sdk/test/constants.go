package test

import (
	"strings"

	"github.com/google/uuid"
)

var (
	TEST_UUID       = "73577357-7357-7357-7357-735773577357"
	CLEAN_TEST_UUID = "73577357735773577357735773577357"
	NULL_UUID       = uuid.Nil.String()
	CLEAN_NULL_UUID = strings.Replace(NULL_UUID, "-", "", -1)
	TEST_PREFIX     = "test"
	TEST_TRACE_ID   = "test_trace_id"
	DEFAULT_PREFIX  = "id"
)
