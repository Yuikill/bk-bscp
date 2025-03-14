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

package service

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/TencentBlueKing/bk-bscp/internal/runtime/shutdown"
	"github.com/TencentBlueKing/bk-bscp/pkg/cc"
	"github.com/TencentBlueKing/bk-bscp/pkg/criteria/errf"
	"github.com/TencentBlueKing/bk-bscp/pkg/kit"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
	"github.com/TencentBlueKing/bk-bscp/pkg/rest"
)

// HealthyHandler livenessProbe 健康检查
func (p *proxy) HealthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// ReadyHandler ReadinessProbe 健康检查
func (p *proxy) ReadyHandler(w http.ResponseWriter, r *http.Request) {
	p.Healthz(w, r)
}

// Healthz service health check.
func (p *proxy) Healthz(w http.ResponseWriter, r *http.Request) {
	if shutdown.IsShuttingDown() {
		logs.Errorf("service healthz check failed, current service is shutting down")
		w.WriteHeader(http.StatusServiceUnavailable)
		rest.WriteResp(w, rest.NewBaseResp(errf.UnHealth, "current service is shutting down"))
		return
	}

	if err := p.state.Healthz(); err != nil {
		logs.Errorf("etcd healthz check failed, err: %v", err)
		rest.WriteResp(w, rest.NewBaseResp(errf.UnHealth, "etcd healthz error, "+err.Error()))
		return
	}

	rest.WriteResp(w, rest.NewBaseResp(errf.OK, "healthy"))
}

// LogoutHandler return redirect url
func (p *proxy) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	data := p.authorizer.LogOut(r)
	render.Render(w, r, rest.OKRender(data))
}

// UserInfoHandler 鉴权后的用户信息接口
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	k := kit.MustGetKit(r.Context())

	user := kit.User{Username: k.User}
	render.Render(w, r, rest.OKRender(user))
}

// FeatureFlags feature flags
type FeatureFlags struct {
	// BizView 是否开启业务体验
	BizView bool `json:"BIZ_VIEW"`
	// ResourceLimit 业务资源限制
	ResourceLimit cc.ResourceLimit `json:"RESOURCE_LIMIT"`
	// TrpcGoPlugin trpc go plugin
	TrpcGoPlugin TrpcGoPlugin `json:"TRPC_GO_PLUGIN"`
}

// TrpcGoPlugin trpc go plugin
type TrpcGoPlugin struct {
	Enable           bool   `json:"enable"`
	ModuleDomain     string `json:"module_domain"`
	BscpModuleDomain string `json:"bscp_module_domain"`
}

// FeatureFlagsHandler 特性开关接口
func FeatureFlagsHandler(w http.ResponseWriter, r *http.Request) {
	featureFlags := FeatureFlags{}

	biz := r.URL.Query().Get("biz")
	// set biz_view feature flag
	bizViewConf := cc.ApiServer().FeatureFlags.BizView
	featureFlags.BizView = *bizViewConf.Default
	if enable, ok := bizViewConf.Spec[biz]; ok && enable != nil {
		featureFlags.BizView = *enable
	}
	// set biz resource limit
	resourceLimitConf := cc.ApiServer().FeatureFlags.ResourceLimit
	featureFlags.ResourceLimit = resourceLimitConf.Default

	if resource, ok := resourceLimitConf.Spec[biz]; ok {
		if resource.MaxFileSize != 0 {
			featureFlags.ResourceLimit.MaxFileSize = resource.MaxFileSize
		}
		if resource.AppConfigCnt != 0 {
			featureFlags.ResourceLimit.AppConfigCnt = resource.AppConfigCnt
		}
		if resource.TmplSetTmplCnt != 0 {
			featureFlags.ResourceLimit.TmplSetTmplCnt = resource.TmplSetTmplCnt
		}
		if resource.MaxUploadContentLength != 0 {
			featureFlags.ResourceLimit.MaxUploadContentLength = resource.MaxUploadContentLength
		}
		// NOCC:golint/todo(忽略)
		// nolint TODO：其他资源限制
	}

	featureFlags.TrpcGoPlugin = TrpcGoPlugin(cc.ApiServer().FeatureFlags.TrpcGoPlugin)

	render.Render(w, r, rest.OKRender(featureFlags))
}
