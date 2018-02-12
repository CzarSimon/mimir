package news.mimir.streamlistener.heartbeat;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class HeartbeatRunner implements Runnable {

    private FileHeartbeat heartbeat;

    public HeartbeatRunner() {
        this.heartbeat = new FileHeartbeatImpl();
    }

    @Override
    public void run() {
        heartbeat.run();
    }

}
