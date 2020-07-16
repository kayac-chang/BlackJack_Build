package ulg168utils

var Conf *ULG168Conf

var (
	CMDc2sLogin             int
	CMDc2sWalletTransfer    int
	CMDc2sLogout            int
	CMDs2cLoginAck          int
	CMDs2cWalletTransferAck int
	CMDs2cLogoutAck         int
	CMDs2cGameNotFinished   int
	CMDs2cMemberInfo        int
)

type ULG168Conf struct {
	APIHost       string
	MaintainAPI   string
	MaintainToken string
	GameID        string
	GameType      string
	ENV           string
	CMDPrefix     int
}

func InitConf(c *ULG168Conf) {
	Conf = c
	CMDc2sLogin = Conf.CMDPrefix + 150
	CMDc2sWalletTransfer = Conf.CMDPrefix + 151
	CMDc2sLogout = Conf.CMDPrefix + 152

	CMDs2cLoginAck = Conf.CMDPrefix + 50
	CMDs2cWalletTransferAck = Conf.CMDPrefix + 51
	CMDs2cLogoutAck = Conf.CMDPrefix + 52
	CMDs2cGameNotFinished = Conf.CMDPrefix + 53
	CMDs2cMemberInfo = Conf.CMDPrefix + 95
}
