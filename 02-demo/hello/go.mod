module pogusanqian.com/hello

go 1.16

require (
	pogusanqian.com/greetings v0.0.0-00010101000000-000000000000
)

// 使用go mod eidt引入的本地包greetings
replace pogusanqian.com/greetings => ../greetings
