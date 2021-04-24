1\. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？  
> 应该Wrap这个error，职责分离，dao层只负责执行逻辑，由调用方判断错误及逻辑处理，代码见 week02/main.go。
