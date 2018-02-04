package news.mimir.streamlistener.stream;

import news.mimir.streamlistener.client.HttpClient;
import news.mimir.streamlistener.config.Config;
import news.mimir.streamlistener.handler.Handler;
import news.mimir.streamlistener.listener.RateLimitListener;
import news.mimir.streamlistener.listener.StreamListener;
import twitter4j.*;

import java.util.logging.Logger;

/**
 * @see Stream
 * @author simon.g.lindgren@gmail.com
 */
public class StreamImpl implements Stream {

    private static final Logger log = Logger.getLogger(StreamImpl.class.getName());

    private HttpClient client;
    private TwitterStream stream;

    /**
     * Constructor
     * @param client
     * @param handler
     */
    public StreamImpl(HttpClient client, Handler handler) {
        this.client = client;
        stream = new TwitterStreamFactory(Config.TWITTER_CONFIG).getInstance();
        stream.addListener(new StreamListener(handler));
        stream.addRateLimitStatusListener(new RateLimitListener());
    }

    /**
     * @see Stream#start()
     */
    @Override
    public void start() {
        try {
            String[] tickers = client.getFilter();
            logTrackedTickers(tickers);
            stream.filter(tickers);
        } catch (Exception e) {
            stop();
        }
    }

    /**
     * @see Stream#stop()
     */
    @Override
    public void stop() {
        stream.shutdown();
    }

    /**
     * Logs the tickers tracked by the stream.
     * @param tickers
     */
    private void logTrackedTickers(String[] tickers) {
        log.info("Tracking tickers: " + String.join(", ", tickers));
    }

}
