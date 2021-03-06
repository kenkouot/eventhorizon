// Copyright (c) 2014 - The Event Horizon authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"os"
	"testing"

	"github.com/looplab/eventhorizon/publisher/testutil"
)

func TestEventPublisher(t *testing.T) {
	// Local Redis testing with Docker
	url := os.Getenv("REDIS_HOST")

	if url == "" {
		// Default to localhost
		url = "localhost:6379"
	}

	publisher1, err := NewEventPublisher("test", url, "")
	if err != nil {
		t.Fatal("there should be no error:", err)
	}
	defer publisher1.Close()

	// Another bus to test the observer.
	publisher2, err := NewEventPublisher("test", url, "")
	if err != nil {
		t.Fatal("there should be no error:", err)
	}
	defer publisher2.Close()

	// Wait for subscriptions to be ready.
	<-publisher1.ready
	<-publisher2.ready

	testutil.EventPublisherCommonTests(t, publisher1, publisher2)
}
