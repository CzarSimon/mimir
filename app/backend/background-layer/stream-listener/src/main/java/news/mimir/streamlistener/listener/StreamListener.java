package news.mimir.streamlistener.listener;

import news.mimir.streamlistener.handler.Handler;
import twitter4j.RawStreamListener;

import java.util.logging.Logger;

/**
 * Listener for a twitter stream with raw unformated messages.
 * @author simon.g.lindgren@gmail.com
 */
public class StreamListener implements RawStreamListener {

    private static final Logger log =
            Logger.getLogger(StreamListener.class.getName());

    private Handler handler;

    /**
     * Constructor
     * @param handler
     */
    public StreamListener(Handler handler) {
        this.handler = handler;
    }

    /**
     * Handles an incoming status message.
     * @param s
     */
    public void onMessage(String s) {
        handler.handleTweet(s);
    }

    /**
     * Handles exception by logging the message.
     * @param e
     */
    public void onException(Exception e) {
        log.severe(e.getMessage());
    }

}
