# Chapter 1

## Race Condition
특정 순서대로 진행되어야 하는 logic이 있지만 그대로 실행되지 못하는 경우

### Data Race
특정 데이터를 동시에 접근하려고 하는 시도를 뜻함

```go
var data int
go func() {
	data++
}()
time.Sleep(1*time.Second)
if data == 0 {
	fmt.Printf("the value is %d", data)
}
```
goroutine이 끝날 때까지 기다릴 "수"는 있지만 data race 조건을 완전히 해결하지 못한다

```go
var memoryAccess sync.Mutex
var value int
go func() {
	memoryAccess.Lock()
	value++
	memoryAccess.Unlock()
}()

memoryAccess.Lock()
if value == 0 {
	fmt.Printf("the value is %v\n", value)
} else {
    fmt.Printf("the value is %v\n", value)
}
memoryAccess.Unlock()
```
**Data Race** 는 해결할 수 있지만, **Race Condition** 자체는 해결하지 못함
&rarr; 왜냐하면 접근의 순서를 강제하지는 못하기 때문

## Dead Lock
### Requirements
1. 상호 배제 (Mutual Exclusion)
- 동시에 실행되는 프로세스가 임의의 리소스에 대해 배타적 권리를 가진다

2. 대기 조건 (Wait For Condition)
- 하나의 리소스를 보유함과 동시에 다른 리소스를 기다리고 있음

3. 비선점 (Non-Preemption)
- 동시에 실행되고 있는 프로세스 중 특정 리소스를 점유하고 있는 프로세스는 다른 프로세스가 점유할 수 있는 여지를 주지 않음

4. 순환 대기(Circular Wait)
- 동시에 실행되는 프로세스가 서로를 기다림