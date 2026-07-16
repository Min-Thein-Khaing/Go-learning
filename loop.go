package main

import "fmt"

func loop() {
	// -------------------------------------------------------------
	fmt.Println("=== 1. Standard C-style For Loop ===")
	// Variable စမှတ်၊ အဆုံးသတ် သတ်မှတ်ချက်နဲ့ တိုးမယ့်နှုန်း ပါဝင်တဲ့ ပုံမှန် loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Standard Loop Round: %d\n", i)
	}

	// -------------------------------------------------------------
	fmt.Println("\n=== 2. Condition-Only Loop (While Loop ပုံစံ) ===")
	// Condition တစ်ခုတည်းပဲ သုံးပြီး ပတ်တဲ့ loop
	count := 1
	for count <= 3 {
		fmt.Printf("Condition Loop Count: %d\n", count)
		count++ // variable တိုးပေးဖို့ မမေ့ပါနဲ့
	}

	// -------------------------------------------------------------
	fmt.Println("\n=== 3. Infinite Loop (With Break) ===")
	// အဆုံးမရှိ ပတ်နေမယ့် loop ကို သတ်မှတ်ချက်ပြည့်ရင် break နဲ့ ထွက်တာ
	num := 1
	for {
		fmt.Printf("Infinite Loop Round: %d\n", num)
		if num >= 3 {
			fmt.Println("Breaking out of the infinite loop...")
			break // loop ထဲက ထွက်မယ်
		}
		num++
	}

	// -------------------------------------------------------------
	fmt.Println("\n=== 4. For-Range Loop ===")
	// Data အစုအဝေးတွေကို index ကော value ပါ တစ်ခါတည်း ထုတ်ကြည့်တာ
	skills := []string{"Go", "Node.js", "React"}

	for index, value := range skills {
		fmt.Printf("Number %d: Skill name is %s\n", index+1, value)
	}
}