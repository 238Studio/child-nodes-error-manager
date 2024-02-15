package errpack

// Exceptions 错误等级的定义
const (
	NotException          = iota // 无错误
	TrivialException             // 轻微错误
	CommonException              // 一般错误
	FatalException               // 严重错误
	ExtremeFatalException        // 极其严重错误
)

// ErrorModule 错误模块定义
const (
	Assist   = iota // 辅助包
	Config          // 配置
	Database        // 数据库
	Network         // 网络
	Device          // 下位机
	File            //文件
	Script          // 脚本
	Python          //python脚本运行错误
)
