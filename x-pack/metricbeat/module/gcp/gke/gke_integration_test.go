// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration && gcp

package gke

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/elastic/beats/v7/x-pack/metricbeat/module/gcp/metrics"
)

func TestFetch(t *testing.T) {
	config := metrics.GetConfigForTest(t, "gke")
	fmt.Printf("%+v\n", config)

	metricSet := mbtest.NewReportingMetricSetV2WithContext(t, config)
	events, errs := mbtest.ReportingFetchV2WithContext(metricSet)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}

	assert.NotEmpty(t, events)
	mbtest.TestMetricsetFieldsDocumented(t, metricSet, events)
}
