/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2016 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiversions

import (
	"github.com/rackspace/gophercloud"
	"net/http"
)

// Get will retrieve the image type with the provided ID. To extract the image
// type from the result, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient) GetResult {
	var res GetResult
	reqOpts := gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK, http.StatusMultipleChoices},
	}
	_, res.Err = client.Get(getURL(client), &res.Body, &reqOpts)
	return res
}
