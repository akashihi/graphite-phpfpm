/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"strings"
)

/*
  {
    "pool":"www",
    "process manager":"static",
    "start time":1443721960,
    "start since":73451,
    "accepted conn":1470924,
    "listen queue":0,
    "max listen queue":1,
    "listen queue len":128,
    "idle processes":509,
    "active processes":3,
    "total processes":512,
    "max active processes":283,
    "max children reached":0,
    "slow requests":0
  }
*/
type Status struct {
	Accept      string
	Queue       string
	MaxQueue    string
	Idle        string
	Active      string
	Total       string
	MaxActive   string
	MaxChildren string
	Slow        string
}

func parse(page string) Status {
	var result = Status{}

	var rawData interface{}
	d := json.NewDecoder(strings.NewReader(page))
	d.UseNumber()
	err := d.Decode(&rawData)
	if err != nil {
		log.Error("Can't parse status data: %v", err)
		return result
	}

	data := rawData.(map[string]interface{})

	result.Accept = string(data["accepted conn"].(json.Number))
	result.Queue = string(data["listen queue"].(json.Number))
	result.MaxQueue = string(data["max listen queue"].(json.Number))
	result.Idle = string(data["idle processes"].(json.Number))
	result.Active = string(data["active processes"].(json.Number))
	result.Total = string(data["total processes"].(json.Number))
	result.MaxActive = string(data["max active processes"].(json.Number))
	result.MaxChildren = string(data["max children reached"].(json.Number))
	result.Slow = string(data["slow requests"].(json.Number))

	return result
}
