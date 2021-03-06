// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package howto

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestBasicJobSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := basicJobSearch(buf, tc.ProjectID, testCompany.Name, "SWE"); err != nil {
			r.Errorf("basicJobSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("basicJobSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestCategoryFilterSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := categoryFilterSearch(buf, tc.ProjectID, testCompany.Name, []string{"COMPUTER_AND_IT"}); err != nil {
			r.Errorf("categoryFilterSearch: %v", err)
		}
		want := "Jobs:"
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("categoryFilterSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestEmploymentTypesSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := employmentTypesSearch(buf, tc.ProjectID, testCompany.Name, []string{"FULL_TIME"}); err != nil {
			r.Errorf("employmentTypesSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("employmentTypesSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestDateRangeSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := dateRangeSearch(buf, tc.ProjectID, testCompany.Name, "2000-01-01T00:00:00.01Z", "2099-01-01T00:00:00.01Z"); err != nil {
			r.Errorf("dateRangeSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("dateRangeSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestLanguageCodeSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := languageCodeSearch(buf, tc.ProjectID, testCompany.Name, []string{"en-US"}); err != nil {
			r.Errorf("languageCodeSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("languageCodeSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestCompanyDisplayNameSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := companyDisplayNameSearch(buf, tc.ProjectID, testCompany.Name, []string{"Google Sample"}); err != nil {
			r.Errorf("companyDisplayNameSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("companyDisplayNameSearch got %q, want to contain %q", got, want)
		}
	})
}

func TestCompensationSearch(t *testing.T) {
	tc := testutil.SystemTest(t)

	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := compensationSearch(buf, tc.ProjectID, testCompany.Name); err != nil {
			r.Errorf("compensationSearch: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("compensationSearch got %q, want to contain %q", got, want)
		}
	})
}
