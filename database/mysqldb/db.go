/*
 Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2016 Ken Williamson
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
package mysqldb

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//ConnectDb connect to db
func ConnectDb(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//ConnectionTest get a row. Passing in tx allows for transactions
func ConnectionTest(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(ConnectionTestQuery, args...)
	return rowPtr
}

//InsertClient insert
func InsertClient(args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, InsertClientQuery, args...)
	return success, insID
}

//UpdateClient updates a row. Passing in tx allows for transactions
func UpdateClient(args ...interface{}) bool {
	success := crud.Update(nil, UpdateClientQuery, args...)
	return success
}

//GetClient get a row. Passing in tx allows for transactions
func GetClient(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(ClientGetQuery, args...)
	return rowPtr
}

//GetClientList get a row. Passing in tx allows for transactions
func GetClientList(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(ClientGetListQuery, args...)
	return rowsPtr
}

//DeleteClient delete
func DeleteClient(args ...interface{}) bool {
	success := crud.Delete(nil, ClientDeleteQuery, args...)
	return success
}

//InsertRestRoute insert
func InsertRestRoute(args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, InsertRestRouteQuery, args...)
	return success, insID
}

//UpdateRestRoute updates a row. Passing in tx allows for transactions
func UpdateRestRoute(args ...interface{}) bool {
	success := crud.Update(nil, UpdateRestRouteQuery, args...)
	return success
}

//GetRestRoute get a row. Passing in tx allows for transactions
func GetRestRoute(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(RestRouteGetQuery, args...)
	return rowPtr
}

//GetRestRouteList get a row. Passing in tx allows for transactions
func GetRestRouteList(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(RestRouteGetListQuery, args...)
	return rowsPtr
}

//DeleteRestRoute delete
func DeleteRestRoute(args ...interface{}) bool {
	success := crud.Delete(nil, RestRouteDeleteQuery, args...)
	return success
}

//InsertRouteURL insert
func InsertRouteURL(args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, InsertRouteURLQuery, args...)
	return success, insID
}

//UpdateRouteURL updates a row. Passing in tx allows for transactions
func UpdateRouteURL(args ...interface{}) bool {
	success := crud.Update(nil, UpdateRouteURLQuery, args...)
	return success
}

//ActivateRouteURL updates a row. Passing in tx allows for transactions
func ActivateRouteURL(args ...interface{}) bool {
	success := crud.Update(nil, ActivateRouteURLQuery, args...)
	return success
}

//DeactivateOtherRouteURLs updates a row. Passing in tx allows for transactions
func DeactivateOtherRouteURLs(args ...interface{}) bool {
	success := crud.Update(nil, DeactivateOtherRouteURLsQuery, args...)
	return success
}

//GetRouteURL get a row. Passing in tx allows for transactions
func GetRouteURL(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(RouteURLGetQuery, args...)
	return rowPtr
}

//GetRouteURLList get a row. Passing in tx allows for transactions
func GetRouteURLList(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(RouteURLGetListQuery, args...)
	return rowsPtr
}

//GetRouteNameURLList get a row. Passing in tx allows for transactions
func GetRouteNameURLList(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(RouteNameURLGetListQuery, args...)
	return rowsPtr
}

//DeleteRouteURL delete
func DeleteRouteURL(args ...interface{}) bool {
	success := crud.Delete(nil, RouteURLDeleteQuery, args...)
	return success
}

//GetRouteURLs get a row. Passing in tx allows for transactions
func GetRouteURLs(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(GetRouteURLsQuery, args...)
	return rowsPtr
}

//CloseDb close connection to db
func CloseDb() bool {
	res := crud.Close()
	return res
}
