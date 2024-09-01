public class Test {
    public static void main(String[] args) throws InterruptedException {
        long start = System.currentTimeMillis();
        Thread[] ts = new Thread[1000];
        for (int i = 0; i < 1000; i++) {

            ts[i] = new Thread(() -> count());
        }
        for (int i = 0; i < 1000; i++) {

            ts[i].join();
        }
        
        long end = (System.currentTimeMillis() - start);
        System.out.println("total time :" + end);
    }

    private static void count() {
        long res = 0;
        for (long i = 0; i < 100000000; i++) {
            res += 1;
        }
    }
}
