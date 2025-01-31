// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type ResetAllAction struct {
	actionutils.ParentAction
}

func (this *ResetAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo("忽略所有证书的OCSP状态")

	_, err := this.RPC().SSLCertRPC().ResetAllSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetAllSSLCertsWithOCSPErrorRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
