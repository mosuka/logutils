// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logutils

import (
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	accesslog "github.com/mash/go-accesslog"
)

func TestApacheCombined(t *testing.T) {
	logFilename := "http_log_test.log"
	logMaxSize := 500
	logMaxBackups := 3
	logMaxAge := 30
	logCompress := false

	_, err := os.Stat(logFilename)
	if !os.IsNotExist(err) {
		os.Remove(logFilename)
	}

	tailer, err := NewTailer(logFilename)
	defer func() {
		err := tailer.Stop()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()
	if err != nil {
		t.Errorf("%v", err)
	}
	ch := tailer.Tail()

	logger := NewHTTPLogger(logFilename, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	apacheCompinedLogger := NewApacheCombinedLogger(logger)

	header := http.Header{}
	header.Add("Referer", "/index.html")
	header.Add("User-Agent", "test-client")
	record := accesslog.LogRecord{
		Size:          10,
		RequestHeader: header,
		Ip:            "127.0.0.1",
		Username:      "testuser",
		Time:          time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		Method:        "GET",
		Uri:           "/test",
		Protocol:      "HTTP/1.0",
		Status:        200,
	}

	apacheCompinedLogger.Log(record)
	time.Sleep(5 * time.Millisecond)
	lastLog := <-ch

	if !strings.HasSuffix(*lastLog, `127.0.0.1 - testuser [31/Dec/2014 12:13:24 +0000] "GET /test HTTP/1.0" 200 10 "/index.html" "test-client"`) {
		t.Errorf("expected content to see %v, saw %v", `127.0.0.1 - testuser [31/Dec/2014 12:13:24 +0000] "GET /test HTTP/1.0" 200 10 "/index.html" "test-client"`, *lastLog)
	}
}
