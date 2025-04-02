// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelfiber instruments the github.com/gofiber/fiber package.
//
// Currently there are two ways the code can be instrumented. One is
// instrumenting the routing of a received message (the Middleware function)
// and instrumenting the response generation through template evaluation (the
// HTML function).
package otelfiber // import "go.opentelemetry.io/contrib/instrumentation/github.com/gofiber/fiber/otelfiber"
