/*
 * Copyright (c) 2020 Rollbar, Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package client

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

// TestClientNoToken checks that a warning message is logged when a
// RollbarAPIClient is initialized without an API token.
func (s *Suite) TestClientNoToken() {
	var buf bytes.Buffer
	log.Logger = log.Logger.Output(&buf)
	NewClient(DefaultBaseURL, "") // Valid, but probably not what you want, thus warn
	bs := buf.String()
	s.NotZero(bs)
	s.Contains(bs, "warn")
	s.Contains(bs, "Rollbar API token not set")
}

// TestClientNoBaseURL checks that an error is logged when a RollbarAPIClient is
// initialized without an API base URL.
func (s *Suite) TestClientNoBaseURL() {
	var buf bytes.Buffer
	multiWriter := io.MultiWriter(os.Stderr, &buf)
	log.Logger = log.Logger.Output(multiWriter)
	NewClient("", "placeholder") // Invalid base URL
	bs := buf.String()
	s.NotZero(bs)
	s.Contains(bs, "error")
	s.Contains(bs, "Rollbar API base URL not set")
}
