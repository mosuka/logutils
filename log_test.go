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
	"os"
	"strings"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	logLevel := "DEBUG"
	logFilename := "log_test.log"
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

	logger := NewLogger(logLevel, logFilename, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	msg := "[DEBUG] debug message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog := <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[INFO] info message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[WARN] warning message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[ERR] error message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}
}

func TestInfo(t *testing.T) {
	logLevel := "INFO"
	logFilename := "log_test.log"
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

	logger := NewLogger(logLevel, logFilename, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	//msg := "[DEBUG] debug message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	msg := "[INFO] info message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog := <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[WARN] warning message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[ERR] error message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}
}

func TestWarn(t *testing.T) {
	logLevel := "WARN"
	logFilename := "log_test.log"
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

	logger := NewLogger(logLevel, logFilename, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	//msg := "[DEBUG] debug message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	//msg := "[INFO] info message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	msg := "[WARN] warning message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog := <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}

	msg = "[ERR] error message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog = <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}
}

func TestErr(t *testing.T) {
	logLevel := "ERR"
	logFilename := "log_test.log"
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

	logger := NewLogger(logLevel, logFilename, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	//msg := "[DEBUG] debug message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	//msg := "[INFO] info message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	//msg := "[WARN] warning message"
	//logger.Printf(msg)
	//time.Sleep(5*time.Millisecond)
	//lastLog := <-ch
	//
	//if !strings.HasSuffix(*lastLog, msg) {
	//	t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	//}

	msg := "[ERR] error message"
	logger.Printf(msg)
	time.Sleep(5 * time.Millisecond)
	lastLog := <-ch

	if !strings.HasSuffix(*lastLog, msg) {
		t.Errorf("expected content to see %v, saw %v", msg, *lastLog)
	}
}
