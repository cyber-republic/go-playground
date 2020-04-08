// Copyright 2011-2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
)

func init() {
	http.HandleFunc("/", editHandler)
}

var editTemplate = template.Must(template.ParseFiles("static/index.html"))

type editData struct {
	Snippet *Snippet
	Share   bool
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	snip := &Snippet{Body: []byte(helloPlayground)}
	if strings.HasPrefix(r.URL.Path, "/p/") {
		if !*flagAllowShare {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		id := r.URL.Path[3:]
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucketSnippets)
			v := b.Get([]byte(id))
			if v != nil {
				snip.Body = v
			}
			return nil
		})
	}

	editTemplate.Execute(w, &editData{snip, *flagAllowShare})
}

const helloPlayground = `package main

import (
	"fmt"

	"github.com/cyber-republic/go-grpc-adenine/elastosadenine"
	"github.com/cyber-republic/go-grpc-adenine/elastosadenine/stubs/health_check"
)

const (
	// This is the shared secret that you need to generate an API key or get an API key
	// This value is the same for both the gRPC client and the gRPC server
	sharedSecretAdenine string = "7XDnFBdHafpPyIC4nrtuJ5EUYVqdEKjW"
)

func main() {
	fmt.Println("Hello, Nucleus Go Playground. This is a simple example of how to interact with Elastos Smartweb Service.")

 	grpcServerHost := "192.168.1.23" // Replace this with the host/IP where Elastos Smartweb Service is running
    grpcServerPort := 8001			// Replace this with the port where Elastos Smartweb Service is running
    production := false				// Set this to True if Elastos Smartweb Service uses TLS/SSL
    
    didToUse := "iHdasfhasdflkHdasfasdfD"

	fmt.Println("--> Checking the health status of gRPC server(Elastos Smartweb Service)")
	healthCheckTest(grpcServerHost, grpcServerPort, production)

	fmt.Println("--> Generate API Key - SHARED_SECRET_ADENINE")
	generateAPIKeyDemo(grpcServerHost, grpcServerPort, production, didToUse)

	fmt.Println("--> Get API Key - SHARED_SECRET_ADENINE")
	getAPIKeyDemo(grpcServerHost, grpcServerPort, production, didToUse)
}

func healthCheckTest(grpcServerHost string, grpcServerPort int, production bool) {
	healthCheck := elastosadenine.NewHealthCheck(grpcServerHost, grpcServerPort, production)
	defer healthCheck.Close()

	response := healthCheck.Check()
	if response.Status != health_check.HealthCheckResponse_SERVING {
		fmt.Println("grpc server is not running properly")
	} else {
		fmt.Println("grpc server is running")
	}
}

func generateAPIKeyDemo(grpcServerHost string, grpcServerPort int, production bool, didToUse string) {
	common := elastosadenine.NewCommon(grpcServerHost, grpcServerPort, production)
	defer common.Close()

	response := common.GenerateAPIRequest(sharedSecretAdenine, didToUse)
	if response.Status {
		fmt.Printf("Output: %s\n", response.Output)
	} else {
		fmt.Printf("Error Message: %s\n", response.StatusMessage)
	}
}

func getAPIKeyDemo(grpcServerHost string, grpcServerPort int, production bool, didToUse string) {
	common := elastosadenine.NewCommon(grpcServerHost, grpcServerPort, production)
	defer common.Close()

	response := common.GetAPIKey(sharedSecretAdenine, didToUse)
	if response.Status {
		fmt.Printf("Output: %s\n", response.Output)
	} else {
		fmt.Printf("Error Message: %s\n", response.StatusMessage)
	}
}
`
