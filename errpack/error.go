package errpack

//进行错误处理

import (
	"errors"
	"runtime"
	"time"

	"github.com/238Studio/child-nodes-assist/util/uuid"
)

// CustomError 是自定义的错误类型，包括错误消息和堆栈信息。
type CustomError struct {
	ErrorLevel   int    //错误等级
	ErrorModule  int    //错误模块
	ErrorMessage string //错误信息
	Stack        string //调用堆栈

	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`     //创建时间
	UUID      string    `gorm:"primaryKey;type:CHAR(36)"` //UUID

	//FIXME:并发安全？
	ChildErrsUUID  string //孩子错误的UUID
	ParentsErrUUID string //父亲错误UUID
}

// NewError 创建一个新的 CustomError。
// 传入：错误等级 错误类型 错误信息
// 传出：CustomError
func NewError(errorLevel int, errorModule int, err error) *CustomError {
	// 判断错误是否是nil
	if err == nil {
		//错误为空时直接返回空值
		return nil
	}

	// 获取堆栈信息
	stack := getStackTrace()
	return &CustomError{
		ErrorLevel:   errorLevel,
		ErrorModule:  errorModule,
		ErrorMessage: err.Error(),
		Stack:        stack,
		UUID:         uuid.GenerateUUID(),
	}
}

// Error 实现了错误接口的 Error 方法。
// 传入：CustomError
// 传出：错误信息
func (e *CustomError) Error() string {
	return e.ErrorMessage + "\n" + e.Stack
}

// getStackTrace 获取当前堆栈信息。
// 传入：无
// 传出：堆栈信息
func getStackTrace() string {
	buf := make([]byte, 4096)
	runtime.Stack(buf, false)
	stackTrace := string(buf)
	return stackTrace
}

// ErrToStruct 将错误转换为结构体
// 传入：错误信息
// 传出：错误结构体
func ErrToStruct(err error) *CustomError {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr
	}
	return nil
}

// ExtractStackTrace 从错误消息中提取堆栈信息。
// 传入：错误信息
// 传出：堆栈信息
func ExtractStackTrace(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.Stack
	}
	return ""
}

// ExtractErrorMessage 从错误消息中提取错误消息。
// 传入：错误信息
// 传出：错误消息
func ExtractErrorMessage(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorMessage
	}
	return err.Error()
}

// ExtractErrorModule 从错误消息中提取错误模块。
// 传入：错误信息
// 传出：错误类型
func ExtractErrorModule(err error) int {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorModule
	}
	return -1
}

// ExtractErrorLevel 从错误消息中提取错误等级。
// 传入：错误信息
// 传出：错误等级
func ExtractErrorLevel(err error) int {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorLevel
	}
	return 0
}

// ExtractErrorUUID 从错误消息中提取错误UUID。
// 传入：错误信息
// 传出：错误UUID
func ExtractErrorUUID(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.UUID
	}
	return ""
}

//func ChangeErrorLevel(err error, errorLevel int) error {
//}
