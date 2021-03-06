// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Tamir Duberstein (tamird@gmail.com)

package grpcutil

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"github.com/cockroachdb/cockroach/util/log"
)

func init() {
	grpclog.SetLogger(&logger{})
}

// NB: This interface is implemented by a pointer because using a value causes
// a synthetic stack frame to be inserted on calls to the interface methods.
// Specifically, we get a stack frame that appears as "<autogenerated>", which
// is not useful in logs.
//
// Also NB: we pass a depth of 2 here because all logging calls originate from
// the logging adapter file in grpc, which is an additional stack frame away
// from the actual logging site.
var _ grpclog.Logger = (*logger)(nil)

type logger struct{}

func (*logger) Fatal(args ...interface{}) {
	log.FatalfDepth(context.TODO(), 2, "", args...)
}

func (*logger) Fatalf(format string, args ...interface{}) {
	log.FatalfDepth(context.TODO(), 2, format, args...)
}

func (*logger) Fatalln(args ...interface{}) {
	log.FatalfDepth(context.TODO(), 2, "", args...)
}

func (*logger) Print(args ...interface{}) {
	log.InfofDepth(context.TODO(), 2, "", args...)
}

func (*logger) Printf(format string, args ...interface{}) {
	log.InfofDepth(context.TODO(), 2, format, args...)
}

func (*logger) Println(args ...interface{}) {
	log.InfofDepth(context.TODO(), 2, "", args...)
}
