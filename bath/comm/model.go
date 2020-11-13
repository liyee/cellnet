package comm

type ItemLevelInfo struct {
	rec map[string]string
	chr map[string]string
	bap map[string]string
	sau map[string]string
	spy map[string]string
}

type Info struct {
	UserInfo  map[string]string
	ItemsInfo ItemLevelInfo
}

//1:customer:1 2011111 1604978396^10^-1^0^0^0^0
