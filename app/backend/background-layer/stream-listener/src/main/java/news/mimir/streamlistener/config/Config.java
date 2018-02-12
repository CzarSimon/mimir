package news.mimir.streamlistener.config;

import twitter4j.conf.Configuration;
import twitter4j.conf.ConfigurationBuilder;

/**
 * @author simon.g.lindgren@gmail.com
 */
public final class Config {

    public static final Configuration TWITTER_CONFIG = setupTwitterConfig();
    public static final ServerConfig SERVER_CONFIG = new ServerConfig();

    /**
     * Private constructor to prevent instantiation.
     */
    private Config() {}

    /**
     * Creates twitter configuration by reading environment variables
     * @return Twitter configuration
     */
    private static Configuration setupTwitterConfig() {
        TwitterCredentials credentials = new TwitterCredentials();
        ConfigurationBuilder builder = new ConfigurationBuilder();
        builder.setOAuthConsumerKey(credentials.CONSUMER_KEY)
                .setOAuthConsumerSecret(credentials.CONSUMER_SECRET)
                .setOAuthAccessToken(credentials.ACCESS_TOKEN)
                .setOAuthAccessTokenSecret(credentials.ACCESS_TOKEN_SECRET);
        return builder.build();
    }

}
