package news.mimir.streamlistener.stream;

/**
 * Interface for starting and controlling a twitter stream
 * @author simon.g.lindgren@gmail.com
 */
public interface Stream {

    /**
     * Starts a stream
     */
    void start();

    /**
     * Stops a stream
     */
    void stop();

}
