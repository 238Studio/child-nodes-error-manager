# 错误常量

## 错误等级

| 错误等级                    | 值 | 描述                            |
|-------------------------|---|-------------------------------|
| `NotException`          | 0 | 没有错误                          |
| `TrivialException`      | 1 | 轻微错误。直接传入数据库。                 |
| `CommonException`       | 2 | 一般错误。需要进行特殊处理。                |
| `FatalException`        | 3 | 严重错误。被`recover`或者入参检测所捕获的错误。  |
| `ExtremeFatalException` | 4 | 极其严重的错误。此类错误一般无法捕获，会直接导致程序崩溃。 |

## 错误模块定

| 错误类型       | 值 | 描述    |
|------------|---|-------|
| `Assist`   | 0 | 辅助包   |
| `Config`   | 1 | 配置包   |
| `Database` | 2 | 数据库   |
| `Network`  | 3 | 网络    |
| `Device`   | 4 | 下位机   |
| `File`     | 5 | 文件管理器 |
| `Script`   | 6 | 脚本    |