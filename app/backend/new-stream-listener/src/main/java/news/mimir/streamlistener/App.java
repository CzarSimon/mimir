package news.mimir.streamlistener;

import news.mimir.streamlistener.client.HttpClient;
import news.mimir.streamlistener.client.TweetHandlerClient;
import news.mimir.streamlistener.handler.Handler;
import news.mimir.streamlistener.handler.HandlerImpl;
import news.mimir.streamlistener.heartbeat.HeartbeatRunner;
import news.mimir.streamlistener.stream.Stream;
import news.mimir.streamlistener.stream.StreamImpl;

import java.util.logging.Logger;

/**
 * Entrypoint for the stream listener
 * @author simon.g.lindgren@gmail.com
 */
public class App {

    private static final Logger log = Logger.getLogger(App.class.getName());

    /**
     * Starts heartbeat emission in a background thread.
     */
    private static void startHeartbeat() {
        HeartbeatRunner runner = new HeartbeatRunner();
        Thread thread = new Thread(runner);
        thread.start();
    }

    public static void main(String... args) {
        log.info("Starting stream listener");
        HttpClient client = new TweetHandlerClient();
        Handler handler = new HandlerImpl(client);
        Stream stream = new StreamImpl(client, handler);
        startHeartbeat();
        stream.start();
    }

}
