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

```go
type value struct {
	mu sync.Mutex
	value int
}

var wg sync.WaitGroup
printSum := func(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()
	
	time.Sleep(2*time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()
	
	fmt.Printf("sum=%v\n", v1.value, v2.value)
}

var a, b value
wg.Add(2)
go printSum(&a, &b)
go printSum(&b, &a)
wg.Wait()
```

1. 상호 배제
   1. `printSum(&a, &b)` 와 `printSum(&b, &a)` 는 서로 다른 각기의 프로세스이고 a와 b에 대해 배타적 권리를 가진다
2. 대기 조건
   1. `printSum(&a, &b)` 의 경우에는 a 리소스를 확보함과 동시에 b 리소스를 기다리고 있다
3. 비선점
   1. `printSum(&a, &b)` 는 a 리소스를 공유할 수 있는 여지를 주지 않는다
4. 순환 대기
   1. `printSum(&a, &b)` 와 `printSum(&b, &a)` 는 서로 비선점이 되기를 기다리고 있다

## Starvation

&rarr; Mainly resolved by the critical area

```go
var wg sync.WaitGroup
var sharedLock sync.Mutex
const runtime = 1*time.Second

greedyWorker := func() {
	defer wg.Done()
	
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(3*time.Nanosecond)
		sharedLock.Unlock()
		count ++
        }
	
	fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}

politeWorker := func() {
    defer wg.Done()
    
    var count int
    for begin := time.Now(); time.Since(begin) <= runtime; {
        sharedLock.Lock()
        time.Sleep(1*time.Nanosecond)
        sharedLock.Unlock()

        sharedLock.Lock()
        time.Sleep(1*time.Nanosecond)
        sharedLock.Unlock()
		
        sharedLock.Lock()
        time.Sleep(1*time.Nanosecond)
        sharedLock.Unlock()
        count ++
    }
    
    fmt.Printf("Polite worker was able to execute %v work loops\n", count)
}
```
Greedyworker 는 PoliteWorker 가 자원을 차지할 여지 조차를 없애 버린다 &rarr; 피해를 유발하는 것으로 생각하면 된다
