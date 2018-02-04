package news.mimir.streamlistener.config.util;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class EnvNotFoundException extends RuntimeException {

    public EnvNotFoundException(String notFoundKey) {
        super(String.format("Env var: %s not found", notFoundKey));
    }

}
