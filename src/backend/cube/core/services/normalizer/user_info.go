package normalizer

import (
	"crypto/sha256"
	"encoding/hex"
	"neural_storage/cube/core/entities/user"
)

func (n *Normalizer) normalizePwd(pwd string) string {
	res := sha256.New()
	res.Write([]byte(pwd))
	return hex.EncodeToString(res.Sum(nil))
}

func (n *Normalizer) NormalizeUserInfo(info user.Info) (user.Info, error) {
	if info.Pwd() != "" {
		np := n.normalizePwd(info.Pwd())
		info.SetPwd(np)
	}

	return info, nil
}
