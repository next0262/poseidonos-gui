/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Intel Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
 
package dagent

import (
	"pnconnector/src/routers/m9k/model"
	"pnconnector/src/setting"
	"bytes"
	"dagent/src/routers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	X_REQUEST_ID        = "X-request-Id"
	X_REQUEST_ID_VALUE  = "25154800-f145-42b9-8954-c50bada0ab5a"
	ACCEPT              = "Accept"
	ACCEPT_VALUE        = "application/json"
	AUTHORIZATION       = "Authorization"
	AUTHORIZATION_VALUE = "Samsung Rocks"
	TIMESTAMP           = "ts"
	ONE_GB              = 1073741824
	POST                = "POST"
	CREATE_VOLUME       = "/m9k/api/ibofos/v1/volume"
	MOUNT_VOLUME        = "/m9k/api/ibofos/v1/volume/mount"
)

var reqBody *model.VolumeParam = &model.VolumeParam{
	Name:        "Test_Volume",
	NewName:     "",
	Size:        ONE_GB,
	Maxiops:     0,
	Maxbw:       0,
	NameSuffix:  0,
	TotalCount:  3,
	StopOnError: true,
	MountAll:    true,
}

func TestCreateMultiVolWrongParams(t *testing.T) {
	setting.LoadConfig()
	gin.SetMode(gin.TestMode)
	fmt.Println("START")
	bytesRepresentation, err := json.Marshal(*reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytesRepresentation)
	fmt.Println(string(bytesRepresentation))
	request, _ := http.NewRequest(POST, CREATE_VOLUME, bytes.NewBuffer(bytesRepresentation))
	request.Header.Add(X_REQUEST_ID, X_REQUEST_ID_VALUE)
	request.Header.Add(ACCEPT, ACCEPT_VALUE)
	request.Header.Add(AUTHORIZATION, AUTHORIZATION_VALUE)
	request.Header.Add(TIMESTAMP, string(time.Now().Unix()))
	response := httptest.NewRecorder()
	routers.InitRouter().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "400 response is expected")
	fmt.Println("SUCCESS ", response.Body)
}

func TestCreateMultiVolAllTrue(t *testing.T) {
	setting.LoadConfig()
	gin.SetMode(gin.TestMode)
	fmt.Println("START")
	bytesRepresentation, err := json.Marshal(*reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytesRepresentation)
	var jsonStr = []byte(`{"param":{
        	"name":"Test_Volume",
        	"newname":"",
       		"size":1024,
        	"maxiops":0,
        	"maxbw":0,
       		"namesuffix":0,
        	"totalcount":3,
        	"stoponerror":true,
        	"mountall":true
		}
	}`)
	fmt.Println(string(bytesRepresentation))
	fmt.Println("Parameters", string(jsonStr), "\nMarshalled JSON", string(bytesRepresentation))
	request, _ := http.NewRequest(POST, CREATE_VOLUME, bytes.NewBuffer(jsonStr))
	request.Header.Add(X_REQUEST_ID, X_REQUEST_ID_VALUE)
	request.Header.Add(ACCEPT, ACCEPT_VALUE)
	request.Header.Add(AUTHORIZATION, AUTHORIZATION_VALUE)
	request.Header.Add(TIMESTAMP, string(time.Now().Unix()))
	response := httptest.NewRecorder()
	routers.InitRouter().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "202 response is expected")
	fmt.Println("SUCCESS ", response.Body)
}

func TestCreateMultiVolAllFalseServiceUnavailable(t *testing.T) {
	setting.LoadConfig()
	gin.SetMode(gin.TestMode)
	fmt.Println("START")
	bytesRepresentation, err := json.Marshal(*reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytesRepresentation)
	var jsonStr = []byte(`{"param":{
                "Name":"Test_Volume",
                "NewName":"",
                "Size":1024,
                "Maxiops":0,
                "Maxbw":0,
                "NameSuffix":0,
                "TotalCount":3,
                "StopOnError":false,
                "MountAll":false
                }
        }`)
	fmt.Println(string(bytesRepresentation))
	fmt.Println("Parameters", string(jsonStr), "\nMarshalled JSON", string(bytesRepresentation))
	request, _ := http.NewRequest(POST, CREATE_VOLUME, bytes.NewBuffer(jsonStr))
	request.Header.Add(X_REQUEST_ID, X_REQUEST_ID_VALUE)
	request.Header.Add(ACCEPT, ACCEPT_VALUE)
	request.Header.Add(AUTHORIZATION, AUTHORIZATION_VALUE)
	request.Header.Add(TIMESTAMP, string(time.Now().Unix()))
	response := httptest.NewRecorder()
	routers.InitRouter().ServeHTTP(response, request)
	assert.Equal(t, 503, response.Code, "503 response is expected")
	fmt.Println("SUCCESS", response.Body)
}

func TestMountMultiVolRequest(t *testing.T) {
	time.Sleep(7 * time.Second)
	setting.LoadConfig()
	gin.SetMode(gin.TestMode)
	fmt.Println("START")
	bytesRepresentation, err := json.Marshal(*reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytesRepresentation)
	var jsonStr = []byte(`{"param":{
                "Name":"Test_Volume",
                "NewName":"",
                "Size":1024,
                "Maxiops":0,
                "Maxbw":0,
                "NameSuffix":0,
                "TotalCount":3,
                "StopOnError":false
                }
        }`)
	fmt.Println(string(bytesRepresentation))
	fmt.Println("Parameters", string(jsonStr), "\nMarshalled JSON", string(jsonStr))
	request, _ := http.NewRequest(POST, MOUNT_VOLUME, bytes.NewBuffer(jsonStr))
	request.Header.Add(X_REQUEST_ID, X_REQUEST_ID_VALUE)
	request.Header.Add(ACCEPT, ACCEPT_VALUE)
	request.Header.Add(AUTHORIZATION, AUTHORIZATION_VALUE)
	request.Header.Add(TIMESTAMP, string(time.Now().Unix()))
	response := httptest.NewRecorder()
	routers.InitRouter().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "202 response is expected")
	fmt.Println("SUCCESS", response.Body)
}
