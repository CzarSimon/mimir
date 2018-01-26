package news.mimir.streamlistener.client;

/**
 * @author simon.g.lindgren@gmail.com
 */
public interface HttpClient {

    /**
     * Sends a supplied message.
     * @param message - Message to send
     */
    void send(String message);

    /**
     * Gets filter terms used to track a stream
     * @return
     */
    String[] getFilter();

}
