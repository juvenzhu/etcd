// Copyright 2015 CoreOS, Inc.
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

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	etcdPath := flag.String("etcd-path", filepath.Join(os.Getenv("GOPATH"), "bin/etcd"), "the path to etcd binary")
	port := flag.String("port", ":9027", "port to serve agent server")
	flag.Parse()

	a, err := newAgent(*etcdPath)
	if err != nil {
		log.Fatal(err)
	}
	a.serveRPC(*port)

	var done chan struct{}
	<-done
}
