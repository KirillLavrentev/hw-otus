package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Проверка на логику выталкивания элементов из-за размера очереди (например: n = 3, добавили 4 элемента - 1й из кэша вытолкнулся);
		c := NewCache(2)

		wasInCache := c.Set("a", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("b", 2)
		require.False(t, wasInCache)

		wasInCache = c.Set("c", 3)
		require.False(t, wasInCache)

		val, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, 2, val)

		val, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, val)

		wasInCache = c.Set("d", 4)
		require.False(t, wasInCache)

		val, ok = c.Get("b")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, val)

		val, ok = c.Get("d")
		require.True(t, ok)
		require.Equal(t, 4, val)

		val, ok = c.Get("d")
		require.True(t, ok)
		require.Equal(t, 4, val)

		val, ok = c.Get("d")
		require.True(t, ok)
		require.Equal(t, 4, val)

		val, ok = c.Get("b")
		require.False(t, ok)
		require.Nil(t, val)
	})

	// t.Run("purge old elements", func(t *testing.T) {
	// 	// Проверка на логику выталкивания давно используемых элементов (например: n = 3, добавили 3 элемента, обратились несколько раз
	// 	// к разным элементам: изменили значение, получили значение и пр. - добавили 4й элемент, из первой тройки вытолкнется тот элемент,
	// 	// что был затронут наиболее давно).
	// 	c := NewCache(3)

	// 	wasInCache := c.Set("a", 11)
	// 	require.False(t, wasInCache)

	// 	wasInCache = c.Set("b", 22)
	// 	require.False(t, wasInCache)

	// 	wasInCache = c.Set("c", 33)
	// 	require.False(t, wasInCache)

	// 	val, ok := c.Get("a")
	// 	require.True(t, ok)
	// 	require.Equal(t, 11, val)

	// 	val, ok = c.Get("c")
	// 	require.True(t, ok)
	// 	require.Equal(t, 33, val)

	// 	wasInCache = c.Set("a", 1)
	// 	require.True(t, wasInCache)

	// 	wasInCache = c.Set("b", 2)
	// 	require.True(t, wasInCache)

	// 	wasInCache = c.Set("f", 1234)
	// 	require.False(t, wasInCache)

	// 	val, ok = c.Get("c")
	// 	require.False(t, ok)
	// 	require.Nil(t, val)

	// 	// wasInCache = c.Set("b", 55)
	// 	// require.False(t, wasInCache)

	// 	// val, ok = c.Get("q")
	// 	// require.False(t, ok)
	// 	// require.Nil(t, val)

	// })
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
