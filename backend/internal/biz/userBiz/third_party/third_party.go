package third_party

import (
	"github.com/google/wire"
	"github.com/wzyjerry/auth/internal/biz/userBiz/third_party/github"
	"github.com/wzyjerry/auth/internal/biz/userBiz/third_party/microsoft"
)

var ProviderSet = wire.NewSet(github.New, microsoft.New)
