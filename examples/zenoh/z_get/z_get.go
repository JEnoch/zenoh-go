/*
 * Copyright (c) 2017, 2020 ADLINK Technology Inc.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License 2.0 which is available at
 * http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
 * which is available at https://www.apache.org/licenses/LICENSE-2.0.
 *
 * SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
 *
 * Contributors:
 *   ADLINK zenoh team, <zenoh@adlink-labs.tech>
 */

package main

import (
	"fmt"
	"os"

	"github.com/eclipse-zenoh/zenoh-go"
)

func main() {
	selector := "/demo/example/**"
	if len(os.Args) > 1 {
		selector = os.Args[1]
	}

	var locator *string
	if len(os.Args) > 2 {
		locator = &os.Args[2]
	}

	s, err := zenoh.NewSelector(selector)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Login to Zenoh...")
	y, err := zenoh.Login(locator, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Use Workspace on '/'")
	root, _ := zenoh.NewPath("/")
	w := y.Workspace(root)

	fmt.Println("Get from " + s.ToString())
	for _, pv := range w.Get(s) {
		fmt.Println("  " + pv.Path().ToString() + " : " + pv.Value().ToString())
	}

	err = y.Logout()
	if err != nil {
		panic(err.Error())
	}

}
