/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"net/http"
	"net/url"
	"os"
)

func parseQueryString(vals url.Values) string {
	var rtn = ""
	var first = true
	for key, value := range vals {
		if first == true {
			first = false
			rtn += "?" + key + "=" + value[0]
		} else {
			rtn += "&" + key + "=" + value[0]
		}
	}
	return rtn
}

func getCacheHost() string {
	var rtn = ""
	if os.Getenv("CACHE_HOST") != "" {
		rtn = os.Getenv("CACHE_HOST")
	} else {
		rtn = "http://localhost:3010"
	}
	return rtn
}

func buildHeaders(pr *http.Request, sr *http.Request) {
	h := pr.Header
	for hn, v := range h {
		//fmt.Print("header: ")
		//fmt.Print(hn)
		//fmt.Print(" value: ")
		//fmt.Println(v[0])
		sr.Header.Set(hn, v[0])
	}
}
