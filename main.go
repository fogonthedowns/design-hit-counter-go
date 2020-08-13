package main

const fiveMinutes = 5 * 60

type HitCounter struct {
	secondsBuckets map[int]int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		secondsBuckets: make(map[int]int),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.secondsBuckets[timestamp]++
	for seconds, _ := range this.secondsBuckets {
		if seconds < timestamp-fiveMinutes {
			delete(this.secondsBuckets, seconds)
		}
	}

}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	var total int
	for seconds, count := range this.secondsBuckets {
		if seconds > timestamp-fiveMinutes {
			total += count
		} else {
			delete(this.secondsBuckets, seconds)
		}
	}
	return total
}
