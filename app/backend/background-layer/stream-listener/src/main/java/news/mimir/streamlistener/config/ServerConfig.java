package news.mimir.streamlistener.config;

import news.mimir.streamlistener.config.util.EnvVar;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class ServerConfig {

    public final String POST_TWEET_ROUTE;
    public final String GET_TICKERS_ROUTE;

    /**
     * Constructor
     */
    public ServerConfig() {
        String baseURL = getBaseURL();
        this.POST_TWEET_ROUTE = baseURL + "/api/tweet";
        this.GET_TICKERS_ROUTE = baseURL + "/api/tweet/tickers";
    }

    /**
     * Gets the base url for the tweet handler
     * @return base url
     */
    private String getBaseURL() {
        String protocol = EnvVar.get("TWEET_HANDLER_PROTOCOL", "http");
        String host = EnvVar.get("TWEET_HANDLER_HOST");
        String port = EnvVar.get("TWEET_HANDLER_PORT");
        return String.format("%s://%s:%s", protocol, host, port);
    }

}
