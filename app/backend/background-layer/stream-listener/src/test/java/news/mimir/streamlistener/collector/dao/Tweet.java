package news.mimir.streamlistener.collector.dao;

import java.sql.SQLException;

/**
 * Data access interface for tweets.
 * @author simon.g.lindgren@gmail.com
 */
public interface Tweet {

    /**
     * Inserts a tweets in a connected database.
     * @param tweet
     * @throws SQLException
     */
    void insert(String tweet) throws SQLException;

}
