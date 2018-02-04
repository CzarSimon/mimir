package news.mimir.streamlistener.heartbeat;

/**
 * Interface for emiting a heartbeat (sign of life) to a file
 * that can be monitored by an liveness probe.
 * @author simon.g.lindgren@gmail.com
 */
public interface FileHeartbeat {

    /**
     * Continously runs heartbeat emission.
     */
    void run();

    /**
     * Emits a heartbeat to a file.
     */
    void emitToFile();

}
