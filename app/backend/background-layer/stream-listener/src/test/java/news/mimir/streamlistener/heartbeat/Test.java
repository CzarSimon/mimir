package news.mimir.streamlistener.heartbeat;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class Test {

    private static void startHeartbeat() {
        Thread thread = new Thread(new HeartbeatRunner());
        thread.start();
    }

    public static void main(String... args) {
        startHeartbeat();
        while (true) {
            System.out.println("hello");
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                System.out.println(e.getMessage());
            }
        }
    }

}
