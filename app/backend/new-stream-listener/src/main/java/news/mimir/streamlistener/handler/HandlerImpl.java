package news.mimir.streamlistener.handler;

import news.mimir.streamlistener.client.HttpClient;

import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class HandlerImpl implements Handler {

    private static final Logger log =
            Logger.getLogger(HandlerImpl.class.getName());

    /**
     * Client to send incoming messages
     */
    private HttpClient client;

    /**
     * Constructor
     * @param client
     */
    public HandlerImpl(HttpClient client) {
        this.client = client;
    }

    /**
     * (non-Javadoc)
     * @see Handler#handleTweet(String)
     */
    public void handleTweet(String tweet) {
        if (!isEmpty(tweet)) {
            client.send(tweet);
        } else {
            log.info("Empty tweet");
        }
    }

    /**
     * Checks if a tweet is empty
     * @param tweet
     * @return
     */
    private boolean isEmpty(String tweet) {
        return (tweet.equals("") || tweet == null);
    }
}
