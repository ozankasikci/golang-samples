// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
	dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
)

func TestTemplateSamples(t *testing.T) {
	testutil.SystemTest(t)
	buf := new(bytes.Buffer)
	fullID := "projects/" + projectID + "/inspectTemplates/golang-samples-test-template"
	// Delete template before trying to create it since the test uses the same name every time.
	listInspectTemplates(buf, client, projectID)
	got := buf.String()
	if strings.Contains(got, fullID) {
		buf.Reset()
		deleteInspectTemplate(buf, client, fullID)
		if got := buf.String(); !strings.Contains(got, "Successfully deleted inspect template") {
			t.Fatalf("failed to delete template")
		}
	}
	buf.Reset()
	createInspectTemplate(buf, client, projectID, dlppb.Likelihood_POSSIBLE, 0, "golang-samples-test-template", "Test Template", "Template for testing", nil)
	got = buf.String()
	if !strings.Contains(got, "Successfully created inspect template") {
		t.Fatalf("failed to createInspectTemplate: %s", got)
	}
	buf.Reset()
	listInspectTemplates(buf, client, projectID)
	got = buf.String()
	if !strings.Contains(got, fullID) {
		t.Fatalf("failed to list newly created template (%s): %q", fullID, got)
	}
	buf.Reset()
	deleteInspectTemplate(buf, client, fullID)
	if got := buf.String(); !strings.Contains(got, "Successfully deleted inspect template") {
		t.Fatalf("failed to delete template")
	}
}
