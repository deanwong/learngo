package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair struct {
	right  int
	height int
}
type hp []pair

// 实现堆接口 的 五个方法
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) } // 这里的 v 输入的pair变量不能为指针
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline2(buildings [][]int) (ans [][]int) {
	// 1. 建立边界数组 并排序
	n := len(buildings)
	boundaries := make([]int, 0, 2*n) // 易错点  忘记加0  表示申请2n长度 从0开始添加  否则是从2n添加
	for i := 0; i < n; i++ {
		boundaries = append(boundaries, buildings[i][0], buildings[i][1])
	}
	sort.Ints(boundaries)
	// 2. 获取每个边界处最高点，将其存储到 ans 数组中记录
	// 遍历边界数组boundaries，
	// 将建筑左边界小于等于 boundary边界的存入优先队列（这里用堆），采用大顶堆记录最高点
	// 若建筑右边界小于等于 boundary 则弹出
	// 考虑？什么时候存入 ans数组
	// 按理说，每个boundary 及其 高度 都要存入 ans数组，
	// 但要考虑一个特殊情况，当多个边界相同高度时，只需要存储 初始边界和高度
	// （ 因此要判断 当前boundary对应的高度 是否与前一个存储ans的高度 是否相同？ 相同 则略过）
	// 最后还要考虑个边界条件，当ans 最开始为空时，直接存入 boundary 及其对应的边界
	index := 0
	h := hp{}
	for i := 0; i < len(boundaries); i++ {
		// 左边界 小于等于 boundary 的建筑，进入优先队列（大顶堆）
		boundary := boundaries[i]
		for index < n && buildings[index][0] <= boundary {
			p := pair{
				right:  buildings[index][1],
				height: buildings[index][2],
			}
			heap.Push(&h, p)
			index++
		}
		// 优先队列中 右边界小于等于 boundary 的建筑（说明当前建筑已经算过了） 弹出
		for len(h) > 0 && h[0].right <= boundary {
			//h.Pop()  错误用法
			heap.Pop(&h)
		}
		// 计算当前最高高度
		maxh := 0 // 重要  用于建筑间为平地时  记录0高度
		if len(h) > 0 {
			maxh = h[0].height
		}
		// 存入结果数组 ans
		if len(ans) == 0 || maxh != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxh})
		}
	}
	return
}

/*
class Solution {
    public List<List<Integer>> getSkyline(int[][] bs) {
        List<List<Integer>> ans = new ArrayList<>();

        // 预处理所有的点，为了方便排序，对于左端点，令高度为负；对于右端点令高度为正
        List<int[]> ps = new ArrayList<>();
        for (int[] b : bs) {
            int l = b[0], r = b[1], h = b[2];
            ps.add(new int[]{l, -h});
            ps.add(new int[]{r, h});
        }

        // 先按照横坐标进行排序
        // 如果横坐标相同，则按照左端点排序
        // 如果相同的左/右端点，则按照高度进行排序
        Collections.sort(ps, (a, b)->{
            if (a[0] != b[0]) return a[0] - b[0];
            return a[1] - b[1];
        });

        // 大根堆
        PriorityQueue<Integer> q = new PriorityQueue<>((a,b)->b-a);
        int prev = 0;
        q.add(prev);
        for (int[] p : ps) {
            int point = p[0], height = p[1];
            if (height < 0) {
                // 如果是左端点，说明存在一条往右延伸的可记录的边，将高度存入优先队列
                q.add(-height);
            } else {
                // 如果是右端点，说明这条边结束了，将当前高度从队列中移除
                q.remove(height);
            }

            // 取出最高高度，如果当前不与前一矩形“上边”延展而来的那些边重合，则可以被记录
            int cur = q.peek();
            if (cur != prev) {
                List<Integer> list = new ArrayList<>();
                list.add(point);
                list.add(cur);
                ans.add(list);
                prev = cur;
            }
        }
        return ans;
    }
}
*/

type maxheap []int

func (h maxheap) Len() int            { return len(h) }
func (h maxheap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxheap) Push(v interface{}) { *h = append(*h, v.(int)) } // 这里的 v 输入的pair变量不能为指针
func (h *maxheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
// func (h *maxheap) Peek() interface{}  { a := *h; v := a[len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	ps := make([][2]int, 0, 2*n)
	// 预处理所有的点，为了方便排序，对于左端点，令高度为负；对于右端点令高度为正
	for i := 0; i < n; i++ {
		ps = append(ps, [2]int{buildings[i][0], -buildings[i][2]})
		ps = append(ps, [2]int{buildings[i][1], buildings[i][2]})
	}
	fmt.Println(ps)
	// 先按照横坐标进行排序
	// 如果横坐标相同，则按照左端点排序
	// 如果相同的左/右端点，则按照高度进行排序
	sort.Slice(ps, func(i, j int) bool {
		if ps[i][0] != ps[j][0] {
			return ps[i][0] > ps[j][0]
		}
		return ps[i][1] > ps[j][1]
	})
	fmt.Println(ps)
	h := maxheap{}
	pre := 0
	heap.Push(&h, pre)
	for i := 0; i < len(ps); i++ {
		point, height := ps[i][0], ps[i][1]
		if height < 0 {
			// 如果是左端点，说明存在一条往右延伸的可记录的边，将高度存入优先队列
			heap.Push(&h, -height)
		} else {
			// 如果是右端点，说明这条边结束了，将当前高度从队列中移除
			heap.Pop(&h)
		}
		// 取出最高高度，如果当前不与前一矩形“上边”延展而来的那些边重合，则可以被记录
		maxHeight := heap.

	}
	return nil
}

func main() {
	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}})) // [[2 10] [3 15] [7 12] [12 0] [15 10] [20 8] [24 0]]
	fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))                                           // [[0 3] [5 0]]
}
