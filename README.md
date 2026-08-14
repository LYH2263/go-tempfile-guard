# 临时文件守卫

CLI → Atomic Writer → Tempfile Guard；失败时应保留错误。

## 测试

```bash
set GOTOOLCHAIN=local
go test ./... -count=1
```
