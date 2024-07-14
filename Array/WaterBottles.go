package Array

//There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles from the market with one full water bottle.
//
//The operation of drinking a full water bottle turns it into an empty bottle.
//
//Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.

//Example 1:
//Input: numBottles = 9, numExchange = 3
//Output: 13
//Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
//Number of water bottles you can drink: 9 + 3 + 1 = 13.
//Example 2:

//Input: numBottles = 15, numExchange = 4
//Output: 19
//Explanation: You can exchange 4 empty bottles to get 1 full water bottle.
//Number of water bottles you can drink: 15 + 3 + 1 = 19.

//Approach
//First drink all bottles and then now numBottles is number of empty bottles.
//While number of empty bottles >= exchange requirements:
//Change all possible and drink all.
//Add unchanged bottles to new exchanged bottles.

func numWaterBottles(numBottles int, numExchange int) int {
	sum := numBottles
	for numBottles >= numExchange {
		sum += numBottles / numExchange
		numBottles = (numBottles/numExchange + numBottles%numExchange)
	}
	return sum
}
