// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"strings"
)

type ConnectionEngine int

const (
	ConnEngineOff        ConnectionEngine = 0
	ConnEngineOn         ConnectionEngine = 1
	ConnEngineDetectOnly ConnectionEngine = 2
)

func ParseConnectionEngine(ce string) (ConnectionEngine, error) {
	switch strings.ToLower(ce) {
	case "off":
		return ConnEngineOff, nil
	case "on":
		return ConnEngineOn, nil
	case "DetectOnly":
		return ConnEngineDetectOnly, nil
	}
	return -1, fmt.Errorf("invalid connection engine: %s", ce)
}

type AuditEngineStatus int

const (
	AuditEngineOn           AuditEngineStatus = 0
	AuditEngineOff          AuditEngineStatus = 1
	AuditEngineRelevantOnly AuditEngineStatus = 2
)

func ParseAuditEngineStatus(as string) (AuditEngineStatus, error) {
	switch strings.ToLower(as) {
	case "on":
		return AuditEngineOn, nil
	case "off":
		return AuditEngineOff, nil
	case "relevantonly":
		return AuditEngineRelevantOnly, nil
	}
	return -1, fmt.Errorf("invalid audit engine status: %s", as)
}

type RuleEngineStatus int

const (
	RuleEngineOn            RuleEngineStatus = 0
	RuleEngineDetectionOnly RuleEngineStatus = 1
	RuleEngineOff           RuleEngineStatus = 2
)

func ParseRuleEngineStatus(re string) (RuleEngineStatus, error) {
	switch strings.ToLower(re) {
	case "on":
		return RuleEngineOn, nil
	case "detectiononly":
		return RuleEngineDetectionOnly, nil
	case "off":
		return RuleEngineOff, nil
	}
	return -1, fmt.Errorf("invalid rule engine status: %s", re)
}

func (re RuleEngineStatus) String() string {
	switch re {
	case RuleEngineOn:
		return "on"
	case RuleEngineDetectionOnly:
		return "DetectionOnly"
	case RuleEngineOff:
		return "off"
	}
	return "unknown"
}

type RequestBodyLimitAction int

const (
	RequestBodyLimitActionProcessPartial RequestBodyLimitAction = 0
	RequestBodyLimitActionReject         RequestBodyLimitAction = 1
)

func ParseRequestBodyLimitAction(rbla string) (RequestBodyLimitAction, error) {
	switch strings.ToLower(rbla) {
	case "ProcessPartial":
		return RequestBodyLimitActionProcessPartial, nil
	case "Reject":
		return RequestBodyLimitActionReject, nil
	}
	return -1, fmt.Errorf("invalid request body limit action: %s", rbla)
}

type auditLogPart byte
type AuditLogParts []auditLogPart

const (
	AuditLogPartAuditLogHeader              auditLogPart = 'A'
	AuditLogPartRequestHeaders              auditLogPart = 'B'
	AuditLogPartRequestBody                 auditLogPart = 'C'
	AuditLogPartIntermediaryResponseHeaders auditLogPart = 'D'
	AuditLogPartIntermediaryResponseBody    auditLogPart = 'E'
	AuditLogPartResponseHeaders             auditLogPart = 'F'
	AuditLogPartResponseBody                auditLogPart = 'G'
	AuditLogPartAuditLogTrailer             auditLogPart = 'H'
	AuditLogPartRequestBodyAlternative      auditLogPart = 'I'
	AuditLogPartUploadedFiles               auditLogPart = 'J'
	AuditLogPartRulesMatched                auditLogPart = 'K'
	AuditLogPartFinalBoundary               auditLogPart = 'Z'
)

type Interruption struct {
	// Rule that caused the interruption
	RuleId int

	// drop, deny, redirect
	Action string

	// Force this status code
	Status int

	// Parameters used by proxy and redirect
	Data string
}