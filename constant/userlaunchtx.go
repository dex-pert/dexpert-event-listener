package constant

const (
    SwapTypeSwap   = 0
    SwapTypeSniper = 1
    SwapTypeLaunch = 99
)

func WithSwapName(swapType int64) string {
    switch swapType {
    case 0:
        return "swap"
    case 1:
        return "sniper"
    case 99:
        return "token creation"
    default:
        return ""
    }
}
