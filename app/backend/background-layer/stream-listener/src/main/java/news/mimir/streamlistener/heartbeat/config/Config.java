package news.mimir.streamlistener.heartbeat.config;

import news.mimir.streamlistener.config.util.EnvVar;

/**
 * Configuration for heartbeat emission.
 * @author simon.g.lindgren@gmail.com
 */
public class Config {

    public final String HEARTBEAT_FILE;
    public final int HEARTBEAT_INTERVAL_SECONDS;
    public final String DEFAULT_INTERVAL_SECONDS = "15";

    public Config() {
        this.HEARTBEAT_FILE = EnvVar.get("HEARTBEAT_FILE");
        this.HEARTBEAT_INTERVAL_SECONDS = getHeartbeatInterval();
    }

    /**
     * Gets the heartbeat interval from the environment.
     * @return Heartbeat interval in seconds
     */
    private int getHeartbeatInterval() {
        String intervalStr = EnvVar.get("HEARTBEAT_INTERVAL", DEFAULT_INTERVAL_SECONDS);
        return Integer.parseInt(intervalStr);
    }

}
