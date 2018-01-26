package news.mimir.streamlistener.client;

import twitter4j.TweetEntity;
import twitter4j.api.TweetsResources;

import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class MockClient implements HttpClient {

    private static final Logger log = Logger.getLogger(MockClient.class.getName());

    /**
     * (non-Javadoc)
     * @see HttpClient#send(String)
     */
    public void send(String message) {
        log.info("Sending message: " + message);
    }

    /**
     * (non-Javadoc)
     * @see HttpClient#getFilter()
     */
    public String[] getFilter() {
        return new String[]{
                "$AAPL", "$AMZN", "$FB"
        };
    }

}
