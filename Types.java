// public class Types {
//     public static void main(String a[]){
//         String name1 = "Emma";
//         String name2 = "Bolu";
//         char btw = 'A';
//         int b = 3;
//         int c = 17;
//         int d = 16;
//         int e = 11;
//         System.out.println(name1 + " " + "is" + " " + "sitting" + " " + "close" + " " + "to" + " " + name2);
//          System.out.println(btw);
//           System.out.println(b + c);
//            System.out.println(d - e);
//     }
    
// }
    
// write a program that allows users to make payment of goods and services, as well as printing the receipts as follows: 

// public class Types {

//     public static void main(String[] args) {
//         String n_o_g1 ="toothpaste";
//         String n_o_g2 ="air freshener";
//         String n_o_g3 ="toothbrush";

//         int quantity_of_paste = 5;
//         int quantity_of_air_freshener = 5;
//         int quantity_of_brush = 5;

//         int price_of_paste = 500;
//         int price_of_air_freshner = 200;
//         int price_of_brush = 300;

//         System.out.println("Total amount paid for toothpaste is " + "#" + quantity_of_paste *price_of_paste);
//         System.out.println("Total amount paid for toothbrush is " + "#" + quantity_of_brush *price_of_brush);
//         System.out.println("Total amount paid for air freshener is " + "#" + quantity_of_air_freshener *price_of_air_freshner);
        
//     }
// }


// public class Types {
//     public static void main(String[] args) {
//         char sign = '$';
//         String goods_name = "laptop";
//         int quantity = 3;
//         int price = 100;
//         int total = quantity*price;
//         System.out.println("Goods" + " " + "Qty" + " " + "Price" + " " + "Total" + "\n" + goods_name + " " + quantity + " " + sign + price + " " + sign + total);
//         // System.out.printf("| %-10s | %-8s | %4s |%n", "CATEGORY", "NAME", "BITS");
//         System.out.printf("%-10s %-8s %-8s %-8s", "Goods", "Qty", "Price", "Total" );
       

//     }

    
// }

// goods, qty, price, total. ASS- Difference between integer and big integer, float and double- indicate it in byte.
// BigInteger is a class in Java that can store values bigger than what primitive data types can store

import java.util.Scanner; // program uses class Scanner
public class Types {
// // main method begins execution of Java application
    public static void main(String[] args) {
// // create a Scanner to obtain input from the command window
//     Scanner input = new Scanner(System.in);
//     System.out.print("Enter first integer: ");
//     int number1 = input.nextInt(); // read first number from user
//     System.out.print("Enter second integer: ");
//     int number2 = input.nextInt(); // read second number from user
//     int sum = number1 + number2; // add numbers, then store total in sum
//     System.out.printf("Sum is %d%n", sum);
//     }
// }

Scanner input = new Scanner(System.in);

System.out.print("How old are you?: ");
int age = input.nextLine();

System.out.print("What is your name?: ");

input.nextLine(); // this will remove the hanging return character
String name = input.nextLine();
    }
}
