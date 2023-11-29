package main

import (
	// "problem-solving/linked_list"
	"problem-solving/tools_scripts"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// MazeSolver()

	// if err := tools_scripts.StartWebhook(); err != nil {
	// 	log.Fatalln(err)
	// }
	tools_scripts.BulkRemoveTogaiCustomers()

	// tools_scripts.GetCustomerCards()
	// tools_scripts.GetCustomerPaymentMethods()
	// linked_list.RunAddTwoNumbers()
	// linked_list.RunHasCycle()
	// linked_list.RunTestQueue()
	// linked_list.Insertion()
	// linked_list.ConvertLinkedListToArray()

	// RunProductExceptSelf()
	// RunInsertDltGetRand()
	// RunHIndex()
	// RunMoveZeroes()
	// run_binary_search()
	// RunQuicksort()
	// RunBubbleSort()
	// RunCanConstruct()
	// RunJumpGame2()
	// RunMaxProfit2()
	// RunJumpGame()
	// RunRemoveDuplicates2()
	// RunLengthOfLastWord()
	// RunTextJustification()

	// RunEncodingDecodingStrings()
	// tools_scripts.BulkRemoveTogaiCustomers()
}
