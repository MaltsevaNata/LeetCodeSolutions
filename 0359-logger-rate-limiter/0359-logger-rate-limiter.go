type Logger struct {
    LastMsgs map[string]int
}


func Constructor() Logger {
    return Logger{LastMsgs: make(map[string]int)}
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if ts, ok := this.LastMsgs[message]; ok {
        if timestamp - ts >= 10 {
            this.LastMsgs[message] = timestamp
            return true
        }
        return false
    }
    this.LastMsgs[message] = timestamp
    return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */