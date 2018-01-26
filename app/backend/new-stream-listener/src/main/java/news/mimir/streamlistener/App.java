package news.mimir.streamlistener;

import news.mimir.streamlistener.client.HttpClient;
import news.mimir.streamlistener.client.TweetHandlerClient;
import news.mimir.streamlistener.handler.Handler;
import news.mimir.streamlistener.handler.HandlerImpl;
import news.mimir.streamlistener.stream.Stream;
import news.mimir.streamlistener.stream.StreamImpl;

import java.util.logging.Logger;

/**
 * Entrypoint for the stream listener
 * @author simon.g.lindgren@gmail.com
 */
public class App {

    private static final Logger log = Logger.getLogger(App.class.getName());

    public static void main(String... args) {
        log.info("Starting stream listener");
        HttpClient client = new TweetHandlerClient();
        Handler handler = new HandlerImpl(client);
        Stream stream = new StreamImpl(client, handler);
        stream.start();
    }

}
