//  ThingsConstruction, a code generator for WoT-based models
//  Copyright (C) 2017,2018  @aschmidt75
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as published
//  by the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//  This program is dual-licensed. For commercial licensing options, please
//  contact the author(s).
//

//
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Counter interface {
	Increment(category string, detail string)
}

type InMemoryCounter struct {
}

type RedisCounter struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewInMemoryCounter() *InMemoryCounter {
	c := &InMemoryCounter{}
	return c
}

func (imc InMemoryCounter) Increment(category string, detail string) {
	//
}

func NewRedisCounter(redisHost string, redisPort int) *RedisCounter {
	c := &RedisCounter{
		Host: redisHost,
		Port: redisPort,
	}
	return c
}

func (rc RedisCounter) Increment(category string, detail string) {
	go func(category string, detail string) {
		c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", rc.Host, rc.Port))
		if err != nil {
			Debug.Printf("Unable to open connection to redis: %s", err)
			return
		}
		defer c.Close()

		cmd := fmt.Sprintf("%s:%s", category, detail)
		_, err = c.Do("INCR", cmd)
		if err != nil {
			Debug.Printf("Error incrementing value: %s", err)
			return
		}

	}(category, detail)
}
