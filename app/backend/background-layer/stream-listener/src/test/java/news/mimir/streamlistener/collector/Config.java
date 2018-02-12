package news.mimir.streamlistener.collector;

import news.mimir.streamlistener.config.util.EnvVar;
import news.mimir.streamlistener.stream.Stream;

import java.io.File;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class Config {

    private static final Logger log = Logger.getLogger(Config.class.getName());

    private final String DB_PATH = EnvVar.get("TEST_RESOURCE_PATH") + "tweet.db";
    private final String URL = "jdbc:sqlite:" + DB_PATH;
    private final String SCHEMA = "CREATE TABLE IF NOT EXISTS TWEET(SEQ_NO INTEGER PRIMARY KEY, TWEET TEXT)";

    /**
     * Connects to sqlite database.
     * @return
     */
    public Connection connect() {
        Connection conn = null;
        boolean dbPresent = isDBPresent();
        try {
            conn = DriverManager.getConnection(URL);
            if (!dbPresent)
                installSchema(conn);
        } catch (SQLException e) {
            log.info(e.getMessage());
        }
        return conn;
    }

    /**
     * Installs schema to supplied database if present.
     * @param conn
     */
    private void installSchema(Connection conn) {
        try {
            PreparedStatement stmt = conn.prepareStatement(SCHEMA);
            stmt.execute();
        } catch (SQLException e) {
            log.severe(e.getMessage());
            throw new RuntimeException(e.getMessage());
        }
    }

    /**
     * Checks if DB is present prior to connection.
     * @return boolean result.
     */
    private boolean isDBPresent() {
        File dbFile = new File(DB_PATH);
        return dbFile.exists();
    }

}
