package news.mimir.streamlistener.client;

import news.mimir.streamlistener.collector.Config;
import news.mimir.streamlistener.collector.dao.Tweet;
import news.mimir.streamlistener.collector.dao.TweetImpl;

import java.sql.SQLException;
import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class CollectionClient extends TweetHandlerClient {

    private static final Logger log =
            Logger.getLogger(CollectionClient.class.getName());

    private Tweet tweetDao;

    public CollectionClient() {
        super();
        Config config = new Config();
        this.tweetDao = new TweetImpl(config.connect());
    }

    @Override
    public void send(String message) {
        log.info("Storing tweet: " + message);
        try {
            tweetDao.insert(message);
        } catch (SQLException e) {
            log.severe(e.getMessage());
            System.exit(1);
        }
    }

}
