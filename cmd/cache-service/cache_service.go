/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// The main package is the entry point for the Cache Service, providing caching functionality.
package main

import (
	"fmt"
	"os"

	"github.com/TencentBlueKing/bk-bscp/cmd/cache-service/app"
	"github.com/TencentBlueKing/bk-bscp/cmd/cache-service/options"
	_ "github.com/TencentBlueKing/bk-bscp/internal/i18n/translations"
	"github.com/TencentBlueKing/bk-bscp/pkg/cc"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
)

func main() {
	cc.InitService(cc.CacheServiceName)

	opts := options.InitOptions()
	if err := app.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "start cache service failed, err: %v", err)
		logs.CloseLogs()
		os.Exit(1)
	}
}
