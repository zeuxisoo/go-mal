package setting

var (
    Address string
    Port    int
)

func NewSetting() {
    Address = "127.0.0.1"
    Port    = 25
}
