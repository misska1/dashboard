// Copyright 2017 The Kubernetes Dashboard Authors.
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

package service

import (
	"log"

	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	"github.com/kubernetes/dashboard/src/app/backend/resource/event"
	client "k8s.io/client-go/kubernetes"
)

// GetServiceEvents returns model events for a service with the given name in the given
// namespace
func GetServiceEvents(client client.Interface, dsQuery *dataselect.DataSelectQuery, namespace, serviceName string) (
	*common.EventList, error) {

	serviceEvents, err := event.GetEvents(client, namespace, serviceName)

	if err != nil {
		return nil, err
	}
	if !event.IsTypeFilled(serviceEvents) {
		serviceEvents = event.FillEventsType(serviceEvents)
	}

	events := event.CreateEventList(serviceEvents, dsQuery)
	log.Printf("Found %d events related to %s service in %s namespace",
		len(events.Events), serviceName, namespace)

	return &events, nil
}