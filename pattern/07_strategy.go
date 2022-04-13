package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
//Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и
//помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
//Паттерн Стратегия предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются, и вынести их в собственные классы, называемые стратегиями.
//
//Вместо того, чтобы изначальный класс сам выполнял тот или иной алгоритм, он будет играть роль контекста,
//ссылаясь на одну из стратегий и делегируя ей выполнение работы.
//Чтобы сменить алгоритм, вам будет достаточно подставить в контекст другой объект-стратегию.
// Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
//
// Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные объекты-поведения (например, отличающиеся балансом скорости и потребления ресурсов).
//
// Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.

//+Горячая замена алгоритмов на лету.
// Изолирует код и данные алгоритмов от остальных классов.
// Уход от наследования к делегированию.
// Реализует принцип открытости/закрытости.

//- Усложняет программу за счёт дополнительных классов.
// Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
// Структура кэша
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

// Сеттер, обеспечивающий выбор стратегии
func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Инетерфейс сратегий
type evictionAlgo interface {
	evict(c *cache)
}

// Стратегии с заглушками вместо реализации

// Стратегия 1
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// Стратегия 2
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

// Стратегия 3
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

func ExampleStrategy() {
	fmt.Println("Strategy example")
	fmt.Println()

	// Инициализация и заполнение кэша со стратегией lfu
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	// Смена стратегии
	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	// Смена стратегии
	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
