import java.util.*;
import java.lang.*;
import java.io.*;

public class Main {
    static String mainThreadName;
    static BufferedWriter output;
    public static void printNumber(double x) throws Exception{
        if(mainThreadName == Thread.currentThread().getName()) {
            Exception exception = new Exception();
            throw exception;
        }else {
            output.write(String.valueOf(x));
            output.write(String.valueOf("\n"));
        }
    }

    public static void main(String[] args){
        try {
            mainThreadName = Thread.currentThread().getName();
            String input_path = args[0];
            String output_path = args[1];
            File file=new File(input_path);
            BufferedReader br=new BufferedReader(new InputStreamReader(new FileInputStream(file),"GBK"));
            File file2=new File(output_path);
            output=new BufferedWriter(new OutputStreamWriter(new FileOutputStream(file2),"UTF-8"));
            String string = br.readLine();
            String[] slist = string.split(" ");
            double[] array = new double[slist.length];
            for(int i = 0; i < slist.length; i++) {
                double x = Double.parseDouble(slist[i]);
                array[i] = x;
            }
            Solution sol = new Solution();
            sol.sleepSort(array);
            br.close();
            output.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}