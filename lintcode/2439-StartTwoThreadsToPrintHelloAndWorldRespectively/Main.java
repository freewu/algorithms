import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) {

        try {
            String inputPath = args[0];
            String outputPath = args[1];

            PrintStream ps = new PrintStream(outputPath);
            System.setOut(ps);

            Solution solution = new Solution();

            new Thread(() -> {
            try {
                solution.printHello();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                solution.printWorld();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }).start();
            
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}