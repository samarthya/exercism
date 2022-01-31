package twobucket

import "errors"

type Bucket struct {
	cap, fill int
}

var (
	ErrorFull          = errors.New("bucket is full")
	ErrorEmpty         = errors.New("bucket is empty")
	ErrorInvalid       = errors.New(" invalid bucket")
	ErrorInvalidAmount = errors.New(" invalid amount")
)

func (b *Bucket) Capacity() int {
	return b.cap
}

func (b *Bucket) FillItUp() {
	b.fill = b.cap
}

func (b *Bucket) EmptyIt() {
	b.fill = 0
}

func (b *Bucket) CurrentFill() int {
	return b.fill
}

func (b *Bucket) IsFull() bool {
	if b.cap > 0 {
		return b.cap == b.fill
	}
	return false
}

func (b *Bucket) IsEmpty() bool {
	return b.fill == 0
}

func (b *Bucket) Add(c int) (int, error) {
	if !b.IsFull() {

		rem := b.cap - b.fill

		if rem >= c {
			b.fill += c
			return 0, nil
		} else if c > rem {
			c -= rem
			b.fill += rem
			return c, nil
		}
	}
	return c, ErrorFull
}

func (b *Bucket) Reduce(c int) (int, error) {
	if !b.IsEmpty() {
		if c <= b.fill {
			b.fill -= c
			return b.fill, nil
		} else {
			c -= b.fill
			b.fill = 0
			return c, nil
		}
	}
	return 0, ErrorEmpty
}

func (b *Bucket) Success(v int) bool {
	return b.fill == v
}

func (b *Bucket) PourTillCapacity(b2 *Bucket) (bool, error) {
	if b2.IsFull() {
		return false, ErrorFull
	}

	capacityRemaining := b2.Capacity() - b2.CurrentFill()

	if capacityRemaining >= b.CurrentFill() {
		b2.Add(b.CurrentFill())
		b.Reduce(b.CurrentFill())
	} else if capacityRemaining < b.CurrentFill() {
		b.Reduce(capacityRemaining)
		b2.FillItUp()
	}

	return true, nil
}

func returnBucket(v int) string {
	switch v {
	case 1:
		return ONE
	default:
		return TWO
	}
}

const (
	ONE = "one"
	TWO = "two"
)

// Determine how many actions are required to measure an exact number of liters
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	/**
	There are only 3 possible actions:
	1. Pouring one bucket into the other bucket until either: a) the first bucket is empty b) the second bucket is full
	2. Emptying a bucket and doing nothing to the other.
	3. Filling a bucket and doing nothing to the other.
	After an action, you may not arrive at a state where the starting bucket is empty and the other bucket is full.
	*/
	var totalSteps int

	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 {
		return ONE, totalSteps, 0, ErrorInvalid
	}

	if goalAmount <= 0 {
		return ONE, totalSteps, 0, ErrorInvalidAmount
	}

	var bucketMap = map[string]*Bucket{
		ONE: &Bucket{sizeBucketOne, 0},
		TWO: &Bucket{sizeBucketTwo, 0},
	}

	otherBucket := TWO

	if startBucket == TWO {
		otherBucket = ONE
	} else if startBucket != ONE {
		return ONE, totalSteps, 0, ErrorInvalid
	}

	for {
		switch {

		case bucketMap[otherBucket].Success(goalAmount):
			if ok, o := checkConditionMet(bucketMap[startBucket], bucketMap[otherBucket], goalAmount); ok {
				return otherBucket, totalSteps, o, nil
			}

		case bucketMap[startBucket].Success(goalAmount):
			if ok, o := checkConditionMet(bucketMap[startBucket], bucketMap[otherBucket], goalAmount); ok {
				return startBucket, totalSteps, o, nil
			}

		case bucketMap[startBucket].IsEmpty():
			//Step 1 Fill the start bucket
			totalSteps++

			// Always start  by filling the first one
			err := Action3(bucketMap[startBucket], bucketMap[otherBucket])

			if err != nil {
				return ONE, totalSteps, 0, err
			}

		case bucketMap[otherBucket].Capacity() == goalAmount:
			totalSteps++
			bucketMap[otherBucket].FillItUp()

		case bucketMap[otherBucket].IsFull():

			//Step 1 Fill the start bucket
			totalSteps++
			Action2(bucketMap[otherBucket], bucketMap[startBucket])

		default:
			totalSteps++
			if _, err := Action1(bucketMap[startBucket], bucketMap[otherBucket]); err != nil {
				return ONE, totalSteps, 0, err
			}
		}
	}
}

func Action1(b1, b2 *Bucket) (bool, error) {
	return b1.PourTillCapacity(b2)
}

func Action2(b1, b2 *Bucket) {
	b1.EmptyIt()
	return
}

// Fill It Up
func Action3(b1, b2 *Bucket) error {
	if b1.IsFull() {
		return ErrorFull
	}

	b1.FillItUp()
	return nil
}

func checkConditionMet(b1, b2 *Bucket, g int) (bool, int) {
	if b1.Success(g) {
		return true, b2.CurrentFill()
	}

	if b2.Success(g) {
		return true, b1.CurrentFill()
	}

	return false, 0
}
