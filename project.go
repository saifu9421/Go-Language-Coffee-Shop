package main

import (
	"fmt";
	// "strconv"
	// "net"
)
type Coffee struct{
   name string;
    price int;	 
};
	 
  func main(){
	var arr []int;
	var sum int  = 0;
	coffeeList := map[int]Coffee{
		1:{
		 name: "Hot Coffee",
		 price: 200,
		},
		 2:{
			 name: "Cool Coffee",
			 price: 350,
		 },
		 3:{
			 name: "Black Coffee",
			  price: 400,
		 },
		};

		parcess := make(map[int]Coffee);
		// countParcessCoffee := make(map[int]string);

	 fmt.Println("#-------------------Self Service Coffee Shop------------------#");
	
	 for true{
		  var n int;
		     fmt.Println("Enter the value 0 vew items");
			 fmt.Println("Enter the value -1 exit");
		
		   fmt.Scan(&n);
		     if n == -1{
				 break;
			 }else if n == 0 { 
				    for quantity,value := range coffeeList{
						   fmt.Println(quantity,"",value.name,"=",value.price);
					};
					
					//    var tmp  int = 1;
	                    var userInput int;
					    fmt.Println("#------- BUY COFFEE --------#");
					  fmt.Scan(&userInput); 
					      for  key:= range coffeeList{
							    if key == userInput{
								   parcess[userInput] = coffeeList[key];
								   arr = append(arr, key);
								//    countParcessCoffee[sum] = parcess[userInput].name;
								sum++;
								}
							};
			 }
  };

        var x []int;
   for i:= 0;i<len(arr);i++{
	         var add int = 1;
	      for j:=i+1;j<len(arr);j++{
			 if arr[i] == arr[j]{
				 add++;
				 x = append(x, add);
			 };
		  };
   };

    var totalAmount int;
	 for _,cof := range parcess{
		 totalAmount += cof.price; 
	 };

fmt.Println("Item Name        Price        Quantity        Amount");
fmt.Println("-------------------------------------------------------","\n",);
fmt.Println("                                   ",sum);
   for  _,coffee := range parcess{
	    fmt.Println(coffee.name,"      ",coffee.price,);
   };
fmt.Println("----------------------------------------------",totalAmount);
//   fmt.Println(x);
// fmt.Println(countParcessCoffee);
 fmt.Println(x);
  }
