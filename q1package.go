package pdcq1


var InputString = "Hello World"


func PrintOutput(name string) string {
	
    switch day:=name;
 day{ 
       case "Pakistan": 
       fmt.Println("Covid19 Cases in Pakistan are 1,526 ") 
       break;
       case "South Korea": 
       fmt.Println("Covid19 Cases in South Korea are 9,583 ") 
       break;
       case "France": 
       fmt.Println("Covid19 Cases in France are 37,575") 
       break;
       case "Message": 
       fmt.Println("Stay Home and Wash your hands at regular intervals") 
       break;
       default:  
       fmt.Println("Exiting Now.") 
   } 
}
