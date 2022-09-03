package LeetCode

/**
 * 146. LRU Cache
 * 描述：
 * 难度：Medium
 * 类型：
 */
type LRUCache struct {
	cacheList []*Node
	cacheMap  map[int]int
	Cap       int
}

type Node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheList: make([]*Node, 0),
		cacheMap:  make(map[int]int),
		Cap:       capacity,
	}
}

func (lc *LRUCache) Get(key int) int {
	if index, ok := lc.cacheMap[key]; ok {
		// 记录节点，避免丢失，重新插入头部
		node := lc.cacheList[index]
		// 将当前访问节点在map中的下标修改为0，并累加其余key在map中下标
		lc.cacheMap[key] = 0
		for k, v := range lc.cacheMap {
			if k != key {
				lc.cacheMap[k] = v + 1
			}
		}
		return node.val
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if index, ok := lc.cacheMap[key]; ok {
		// key已存在，先更新节点值，再移动位置
		lc.cacheList[index].val = value
		lc.cacheMap[key] = 0
		for k, v := range lc.cacheMap {
			if k != key {
				lc.cacheMap[k] = v + 1
			}
		}
		return
	} else {
		// key不存在，直接从头部插入新节点，并记录key与节点映射关系
		lc.cacheList = append([]*Node{&Node{key: key, val: value}}, lc.cacheList...)
		lc.cacheMap[key] = 0
		for k, v := range lc.cacheMap {
			if k != key {
				lc.cacheMap[k] = v + 1
			}
		}
	}
	// 超过cache容量则从尾部删除节点
	if len(lc.cacheMap) > lc.Cap {
		tail := lc.cacheList[len(lc.cacheList)-1]
		lc.cacheList = lc.cacheList[:len(lc.cacheList)-1]
		delete(lc.cacheMap, tail.key)
	}
}
