/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"github.com/apache/incubator-devlake/models/common"
	"time"
)

type ZentaoStoryRes struct {
	ID          uint64        `json:"id"`
	Vision      string        `json:"vision"`
	Parent      uint64        `json:"parent"`
	Product     uint64        `json:"product"`
	Branch      int           `json:"branch"`
	Module      int           `json:"module"`
	Plan        string        `json:"plan"`
	Source      string        `json:"source"`
	SourceNote  string        `json:"sourceNote"`
	FromBug     int           `json:"fromBug"`
	Feedback    int           `json:"feedback"`
	Title       string        `json:"title"`
	Keywords    string        `json:"keywords"`
	Type        string        `json:"type"`
	Category    string        `json:"category"`
	Pri         int           `json:"pri"`
	Estimate    int           `json:"estimate"`
	Status      string        `json:"status"`
	SubStatus   string        `json:"subStatus"`
	Color       string        `json:"color"`
	Stage       string        `json:"stage"`
	Mailto      []interface{} `json:"mailto"`
	Lib         int           `json:"lib"`
	FromStory   uint64        `json:"fromStory"`
	FromVersion int           `json:"fromVersion"`
	OpenedBy    struct {
		ID       uint64 `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"openedBy"`
	OpenedDate *time.Time `json:"openedDate"`
	AssignedTo struct {
		ID       uint64 `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"assignedTo"`
	AssignedDate *time.Time `json:"assignedDate"`
	ApprovedDate string     `json:"approvedDate"`
	LastEditedBy struct {
		ID       uint64 `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"lastEditedBy"`
	LastEditedDate *time.Time `json:"lastEditedDate"`
	ChangedBy      string     `json:"changedBy"`
	ChangedDate    string     `json:"changedDate"`
	ReviewedBy     struct {
		ID       uint64 `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"reviewedBy"`
	ReviewedDate *time.Time `json:"reviewedDate"`
	ClosedBy     struct {
		ID       uint64 `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"closedBy"`
	ClosedDate       *time.Time `json:"closedDate"`
	ClosedReason     string     `json:"closedReason"`
	ActivatedDate    string     `json:"activatedDate"`
	ToBug            int        `json:"toBug"`
	ChildStories     string     `json:"childStories"`
	LinkStories      string     `json:"linkStories"`
	LinkRequirements string     `json:"linkRequirements"`
	DuplicateStory   uint64     `json:"duplicateStory"`
	Version          int        `json:"version"`
	StoryChanged     string     `json:"storyChanged"`
	FeedbackBy       string     `json:"feedbackBy"`
	NotifyEmail      string     `json:"notifyEmail"`
	URChanged        string     `json:"URChanged"`
	Deleted          bool       `json:"deleted"`
	PriOrder         string     `json:"priOrder"`
	PlanTitle        string     `json:"planTitle"`
	ProductStatus    string     `json:"productStatus"`
}

type ZentaoStory struct {
	common.NoPKModel
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT  NOT NULL"`
	ID           uint64 `json:"id" gorm:"primaryKey;type:BIGINT  NOT NULL" `
	Product      uint64 `json:"product"`
	Branch       int    `json:"branch"`
	Version      int    `json:"version"`
	OrderIn      int    `json:"order"`
	Vision       string `json:"vision"`
	Parent       uint64 `json:"parent"`
	Module       int    `json:"module"`
	Plan         string `json:"plan"`
	Source       string `json:"source"`
	SourceNote   string `json:"sourceNote"`
	FromBug      int    `json:"fromBug"`
	Feedback     int    `json:"feedback"`
	Title        string `json:"title"`
	Keywords     string `json:"keywords"`
	Type         string `json:"type"`
	Category     string `json:"category"`
	Pri          int    `json:"pri"`
	Estimate     int    `json:"estimate"`
	Status       string `json:"status"`
	SubStatus    string `json:"subStatus"`
	Color        string `json:"color"`
	Stage        string `json:"stage"`
	//Mailto           []interface{} `json:"mailto"`
	Lib              int    `json:"lib"`
	FromStory        uint64 `json:"fromStory"`
	FromVersion      int    `json:"fromVersion"`
	OpenedById       uint64
	OpenedByName     string
	OpenedDate       *time.Time `json:"openedDate"`
	AssignedToId     uint64
	AssignedToName   string
	AssignedDate     *time.Time `json:"assignedDate"`
	ApprovedDate     string     `json:"approvedDate"`
	LastEditedId     uint64
	LastEditedDate   *time.Time `json:"lastEditedDate"`
	ChangedDate      string     `json:"changedDate"`
	ReviewedById     uint64     `json:"reviewedBy"`
	ReviewedDate     *time.Time `json:"reviewedDate"`
	ClosedId         uint64
	ClosedDate       *time.Time `json:"closedDate"`
	ClosedReason     string     `json:"closedReason"`
	ActivatedDate    string     `json:"activatedDate"`
	ToBug            int        `json:"toBug"`
	ChildStories     string     `json:"childStories"`
	LinkStories      string     `json:"linkStories"`
	LinkRequirements string     `json:"linkRequirements"`
	DuplicateStory   uint64     `json:"duplicateStory"`
	StoryChanged     string     `json:"storyChanged"`
	FeedbackBy       string     `json:"feedbackBy"`
	NotifyEmail      string     `json:"notifyEmail"`
	URChanged        string     `json:"URChanged"`
	Deleted          bool       `json:"deleted"`
	PriOrder         string     `json:"priOrder"`
	PlanTitle        string     `json:"planTitle"`
}

func (ZentaoStory) TableName() string {
	return "_tool_zentao_stories"
}