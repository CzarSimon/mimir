package news.mimir.streamlistener;

import news.mimir.streamlistener.client.CollectionClient;
import news.mimir.streamlistener.client.HttpClient;
import news.mimir.streamlistener.handler.Handler;
import news.mimir.streamlistener.handler.HandlerImpl;
import news.mimir.streamlistener.stream.Stream;
import news.mimir.streamlistener.stream.StreamImpl;

import java.util.logging.Logger;

/**
 * Tracks an identical stream but diverts collected tweets into a store for future testing.
 * @author simon.g.lindgren@gmail.com
 */
public class TweetCollector {

    private static final Logger log = Logger.getLogger(TweetCollector.class.getName());

    public static void main(String... args) {
        log.info("Starting collection");
        HttpClient client = new CollectionClient();
        Handler handler = new HandlerImpl(client);
        Stream stream = new StreamImpl(client, handler);
        stream.start();
    }

}
