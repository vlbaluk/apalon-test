package cached

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/vlbaluk/apalon-test/services"
	"github.com/vlbaluk/apalon-test/services/interfaces"
	. "math/big"
)

var cachePattern = "%s:%s:%s"

type MathService struct {
	Cache   *cache.Cache
	Service *services.MathService
}

func (ctrl MathService) CachedCalculate(x *Float, y *Float, keyHead string, fn interfaces.Calculate) (string, bool) {
	key := fmt.Sprintf(cachePattern, keyHead, x.String(), y.String())
	cached, exist := ctrl.Cache.Get(key)
	if exist {
		return cached.(string), true
	}

	res := fn(x, y)
	ctrl.Cache.Set(key, res, 0)
	return res, false
}
