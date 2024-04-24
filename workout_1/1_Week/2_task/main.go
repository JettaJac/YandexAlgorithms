package main

import (
	"fmt"
)

func triagle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if triagle(a, b, c) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}




// #include <iostream>
// #include <vector>
// #include <algorithm>

// using namespace std;

// int main() {
//     int n;
//     cin >> n;

//     vector<string> events(n);
//     for (int i = 0; i < n; i++) {
//         cin >> events[i];
//     //    cout << events[i] << endl;
//     }
    

//     //sort(events.begin(), events.end());

//     int days = 1;
//     for (int i = 1; i < n; i++) {
//     //    cout << events[i][0] << endl;
//         if (events[i] <= events[i - 1]) {
//             days++;
            
//         }
//     }

//     cout << days << endl;

//     return 0;
// }