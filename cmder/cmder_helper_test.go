package cmder

import (
	"testing"

	"github.com/tickstep/cloudpan189-go/internal/config"
)

// TestShouldRelogin 验证：只有"无缓存用户(nil 或 UID==0)"才需要密码重登；
// 已有缓存用户(UID!=0，含 token 过期返回的 UID==1 stub)直接复用缓存的
// sessionKey/sessionSecret，不再每条命令都做密码登录（限流根因）。
func TestShouldRelogin(t *testing.T) {
	cases := []struct {
		name string
		user *config.PanUser
		want bool
	}{
		{"nil 用户 → 需重登", nil, true},
		{"UID==0(未匹配的空 stub) → 需重登", &config.PanUser{UID: 0}, true},
		{"UID==1(token 过期 stub, 仍有缓存 AppToken) → 复用缓存, 不重登", &config.PanUser{UID: 1}, false},
		{"UID==123(正常登录用户) → 复用缓存, 不重登", &config.PanUser{UID: 123}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ShouldRelogin(c.user); got != c.want {
				t.Errorf("ShouldRelogin(%+v) = %v, want %v", c.user, got, c.want)
			}
		})
	}
}
