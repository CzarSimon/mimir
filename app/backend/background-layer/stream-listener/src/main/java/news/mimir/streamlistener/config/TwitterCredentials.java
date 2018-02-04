package news.mimir.streamlistener.config;

import news.mimir.streamlistener.config.util.EnvVar;

import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class TwitterCredentials {

    private static final Logger log =
            Logger.getLogger(TwitterCredentials.class.getName());

    protected final String CONSUMER_KEY;
    protected final String CONSUMER_SECRET;
    protected final String ACCESS_TOKEN;
    protected final String ACCESS_TOKEN_SECRET;

    /**
     * Constructor
     */
    public TwitterCredentials() {
        this.CONSUMER_KEY = EnvVar.get("TWITTER_CONSUMER_KEY");
        this.CONSUMER_SECRET = EnvVar.get("TWITTER_CONSUMER_SECRET");
        this.ACCESS_TOKEN = EnvVar.get("TWITTER_ACCESS_TOKEN");
        this.ACCESS_TOKEN_SECRET = EnvVar.get("TWITTER_ACCESS_TOKEN_SECRET");
    }

}
