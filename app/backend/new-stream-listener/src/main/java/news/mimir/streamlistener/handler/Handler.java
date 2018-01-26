package news.mimir.streamlistener.handler;

/**
 * @author simon.g.lindgren@gmail.com
 */
public interface Handler {

    /**
     * Handles a given tweet
     * @param tweet
     */
    void handleTweet(String tweet);

}
