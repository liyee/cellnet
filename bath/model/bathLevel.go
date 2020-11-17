package model

import (
	"strconv"

	"github.com/davyxu/cellnet/bath/comm"
)

type BathLevelFunc interface {
	GetBathLevel() map[string]string
	GetSilgleBathLevel(string) int
}

type ItemLevelWaitInfo struct {
	Info map[string]string
}

func (waitInfo ItemLevelWaitInfo) GetBathLevel() map[string]string {
	waitParam := []string{"bathLevel:1", "rec_w_max", "chr_w_max", "bap_w_max", "spy1sau_w_max"}
	waitInfo.Info = comm.GetHash(waitParam)
	return waitInfo.Info
}

func (waitInfo ItemLevelWaitInfo) GetSilgleBathLevel(name string) int {
	waitParam := []string{"bathLevel:1", "rec_w_max", "chr_w_max", "bap_w_max", "spy1sau_w_max"}
	waitInfo.Info = comm.GetHash(waitParam)

	val, _ := strconv.Atoi(waitInfo.Info[name])
	return val
}
