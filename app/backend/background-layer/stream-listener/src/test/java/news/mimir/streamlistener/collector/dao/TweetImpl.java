package news.mimir.streamlistener.collector.dao;

import java.sql.Connection;
import java.sql.SQLException;
import java.sql.PreparedStatement;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class TweetImpl implements Tweet {

    private Connection conn;
    private final String INSERT_SQL;

    /**
     * Constructor
     * @param conn
     */
    public TweetImpl(Connection conn) {
        this.conn = conn;
        this.INSERT_SQL = "INSERT INTO TWEET(TWEET) VALUES(?)";
    }

    @Override
    public void insert(String tweet) throws SQLException {
        PreparedStatement stmt = conn.prepareStatement(INSERT_SQL);
        stmt.setString(1, tweet);
        stmt.execute();
    }

}
