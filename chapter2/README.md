# Chapter 2

| Abstact Concept | Actual Term |
|--------|-------------|
| Goroutine| Thread      |
| Channel| Mutex       |

## Channel vs Mutex

### 1. 데이터 소유권을 이전하려 하는가?
만약 데이터를 이전하고 공유하기 위해서라면 channel을 사용하는 것이 맞다

### 2. 구조체 내부 상태를 보호하고자 하는가?
```go
type Counter struct {
	mu sync.Mutex
	value int
}
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
```
`Counter` 구조체의 내용을 보호하기 위한 수단으로는 channel 을 사용하면 안된다. 왜냐하면 외부로 데이터를 공유할 여지가 없기 때문이다.

또한, Increment() 함수에 접근하는 영역 자체가 임계영역으로 판단이된다.

### 3. 여러 부분의 논리를 조정해야 하는가?
사실 이 부분은 1번의 이유와 비슷하다. return value 를 여러 곳에서 사용할 확률이 높다는 의미이므로, 이 경우 channel을 사용하는 것이 맞다

### 4. 성능상의 임계 영역
만약, performance review 결과 특정 영역의 성능이 낮다면, mutex를 사용하자

&rarr; 왜냐하면, channel 은 memory 접근 동기화를 하는 과정이 내부적으로 encoding 되어 있기 때문이다.

> ***메모리접근 동기화**란, 하나의 리소스에 대해 하나의 thread만 접근이 가능하게 제약한다는 의미*

